package car

import (
	"gorm.io/gorm"

	"github.com/dentonliu/go-clean-starter/internal/entity"
)

func NewDBRepo(db *gorm.DB) Repo {
	return &dbRepo{db}
}

type dbRepo struct {
	db *gorm.DB
}

func (r *dbRepo) Get(id uint) (entity.Car, error) {
	var model entity.Car
	res := r.db.Limit(1).Find(&model, id)

	return model, res.Error
}

func (r *dbRepo) Create(model *entity.Car) error {
	res := r.db.Create(model)
	return res.Error
}

func (r *dbRepo) Update(model *entity.Car, attrs ...string) error {
	tx := r.db.Model(model)
	if len(attrs) <= 0 {
		tx.Omit("ID")
	} else {
		tx.Select(attrs)
	}

	res := tx.Updates(*model)
	return res.Error
}

func (r *dbRepo) Save(model *entity.Car) error {
	res := r.db.Save(model)
	return res.Error
}

func (r *dbRepo) Delete(model *entity.Car) error {
	res := r.db.Delete(model)
	return res.Error
}

func (r *dbRepo) Count() (int, error) {
	var count int64
	res := r.db.Model(&entity.Car{}).Count(&count)
	return int(count), res.Error
}

func (r *dbRepo) Query(offset, limit int) ([]entity.Car, error) {
	models := []entity.Car{}
	res := r.db.Offset(offset).Limit(limit).Find(&models)
	return models, res.Error
}
