package repository

import (
	"app/db"
	"app/model"
	"fmt"

	"github.com/gocraft/dbr"
)

type (
	ICalorie interface {
		ByUserID(id int64, calorieType int64) (*model.Calories, error)
		AllCaloriesByUserIDAndType(id int64, calorieType int64) (*model.Status, error)
		Create(m *model.CreateCalorie) error
	}

	Calorie struct {
		session *dbr.Session
	}
)

func NewCalorie() ICalorie {
	return &Calorie{
		session: db.GetSession(),
	}
}

func (r *Calorie) ByUserID(id int64, calorieType int64) (*model.Calories, error) {
	m := &model.Calories{}
	_, err := r.session.Select("*").From("calories").
		Where("user_id = ?", id).
		Where("calorie_type = ?", calorieType).
		OrderAsc("created_at").
		Load(m)
	if err != nil {
		return nil, fmt.Errorf("fetch error :%v", err)
	}
	return m, nil
}

func (r *Calorie) AllCaloriesByUserIDAndType(id int64, calorieType int64) (*model.Status, error) {
	m := &model.Status{}
	t, err := r.session.Select("sum(value) as value").From("calories").
		Where("user_id = ?", id).
		Where("calorie_type = ?", calorieType).
		Rows()
	for t.Next() {
		err = t.Scan(&m.Value)
		if err != nil {
			return nil, fmt.Errorf("fetch error :%v", err)
		}
	}
	if err != nil {
		return nil, fmt.Errorf("fetch error :%v", err)
	}
	return m, nil
}

func (r *Calorie) Create(m *model.CreateCalorie) error {
	_, err := r.session.InsertInto("calories").
		Columns("user_id", "calorie_type", "value").
		Record(m).
		Exec()
	if err != nil {
		return err
	}

	return nil
}
