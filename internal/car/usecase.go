package car

import (
	"github.com/dentonliu/go-clean-starter/internal/entity"
)

type Repo interface {
	Get(id uint) (entity.Car, error)
	Create(model *entity.Car) error
	Update(model *entity.Car, columns ...string) error
	Save(model *entity.Car) error
	Delete(model *entity.Car) error
	Count() (int, error)
	Query(offset, limit int) ([]entity.Car, error)
}

type Usecase struct {
	repo Repo
}

func NewUsecase(repo Repo) *Usecase {
	return &Usecase{repo}
}

func (u *Usecase) Get(id uint) (entity.Car, error) {
	return u.repo.Get(id)
}

func (u *Usecase) Create(model *entity.Car) error {
	return u.repo.Create(model)
}

func (u *Usecase) Update(model *entity.Car, attrs ...string) error {
	return u.repo.Update(model, attrs...)
}

func (u *Usecase) Save(model *entity.Car) error {
	return u.repo.Save(model)
}

func (u *Usecase) Delete(id uint) (entity.Car, error) {
	model, err := u.Get(id)
	if err != nil {
		return model, err
	}

	return model, u.repo.Delete(&model)
}

func (u *Usecase) Count() (int, error) {
	return u.repo.Count()
}

func (u *Usecase) Query(offset, limit int) ([]entity.Car, error) {
	return u.repo.Query(offset, limit)
}
