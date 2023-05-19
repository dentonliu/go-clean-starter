package car

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/dentonliu/go-clean-starter/internal/entity"
)

type mockRepo struct {
	models []entity.Car
}

func TestUsecase(t *testing.T) {
	repo := &mockRepo{
		models: []entity.Car{},
	}

	usecase := NewUsecase(repo)

	{
		// Create
		model := entity.Car{
			Brand: "BYD Yangwang U8",
			Color: "red",
			Seats: 5,
		}
		err := usecase.Create(&model)

		if assert.Nil(t, err) {
			assert.Equal(t, uint(1), model.ID)
		}
	}

	{
		// Update
		model, err := usecase.Get(1)
		if assert.Nil(t, err) {
			model.Color = "black"
			err = usecase.Update(&model, "Color")
			assert.Nil(t, err)
		}
	}

	{
		// Save
		model := entity.Car{
			Brand: "Li Auto L9",
			Color: "Silver",
			Seats: 5,
		}

		err := usecase.Save(&model)

		if assert.Nil(t, err) {
			assert.Equal(t, uint(2), model.ID)
		}
	}

	{
		// Query
		models, err := usecase.Query(0, 3)
		if assert.Nil(t, err) {
			assert.Equal(t, 2, len(models))
		}
	}

	{
		// Delete
		model, err := usecase.Delete(2)
		if assert.Nil(t, err) {
			assert.Equal(t, "Li Auto L9", model.Brand)
		}
	}

	{
		// Count
		c, err := usecase.Count()
		if assert.Nil(t, err) {
			assert.Equal(t, 1, c)
		}
	}

	{
		// Get
		model, err := usecase.Get(2)
		if assert.Nil(t, err) {
			assert.Zero(t, model.ID)
		}
	}
}

func (r *mockRepo) Get(id uint) (entity.Car, error) {
	var model entity.Car

	for i, m := range r.models {
		if m.ID == id {
			model = r.models[i]
			break
		}
	}

	return model, nil
}

func (r *mockRepo) Create(model *entity.Car) error {
	model.ID = uint(len(r.models) + 1)
	model.CreatedAt = time.Now()
	model.UpdatedAt = time.Now()
	r.models = append(r.models, *model)
	return nil
}

func (r *mockRepo) Update(model *entity.Car, columns ...string) error {
	for i, m := range r.models {
		if m.ID == model.ID {
			model.UpdatedAt = time.Now()
			r.models[i] = *model
			break
		}
	}

	return nil
}

func (r *mockRepo) Save(model *entity.Car) error {
	if model.ID == 0 {
		return r.Create(model)
	}

	return r.Update(model)
}

func (r *mockRepo) Delete(model *entity.Car) error {
	for i, m := range r.models {
		if m.ID == model.ID {
			r.models = append(r.models[:i], r.models[i+1:]...)
			break
		}
	}
	return nil
}

func (r *mockRepo) Count() (int, error) {
	return len(r.models), nil
}

func (r *mockRepo) Query(offset, limit int) ([]entity.Car, error) {
	if limit <= 0 || offset >= len(r.models) {
		return []entity.Car{}, nil
	}

	if (offset + limit) >= len(r.models) {
		return r.models[offset:], nil
	}

	return r.models[offset:limit], nil
}
