package car

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"

	"github.com/dentonliu/go-clean-starter/internal/entity"
	"github.com/dentonliu/go-clean-starter/internal/util"
)

func TestDBRepo(t *testing.T) {
	db := util.MockDB()
	err := initTable(db)

	if !assert.Nil(t, err) {
		return
	}

	repo := NewDBRepo(db)

	{
		// Create
		model := entity.Car{
			Brand: "BYD Yangwang U8",
			Color: "red",
			Seats: 5,
		}
		err := repo.Create(&model)

		if assert.Nil(t, err) {
			assert.Equal(t, uint(1), model.ID)
		}
	}

	{
		// Update
		model, err := repo.Get(1)
		if assert.Nil(t, err) {
			model.Color = "black"
			err = repo.Update(&model, "Color")
			assert.Nil(t, err)

			model.Seats = 6
			err = repo.Update(&model)
			assert.Nil(t, err)
		}
	}

	{
		// Save
		model := entity.Car{
			Brand: "Li Auto L9",
			Color: "silver",
			Seats: 5,
		}

		err := repo.Save(&model)

		if assert.Nil(t, err) {
			assert.Equal(t, uint(2), model.ID)
		}
	}

	{
		// Query
		models, err := repo.Query(0, 3)
		if assert.Nil(t, err) {
			assert.Equal(t, 2, len(models))
		}
	}

	{
		// Delete
		model, err := repo.Get(2)
		if assert.Nil(t, err) {
			err = repo.Delete(&model)
			assert.Nil(t, err)
		}
	}

	{
		// Count
		c, err := repo.Count()
		if assert.Nil(t, err) {
			assert.Equal(t, 1, c)
		}
	}

	{
		// Get
		model, err := repo.Get(2)
		if assert.Nil(t, err) {
			assert.Zero(t, model.ID)
		}
	}
}

func initTable(db *gorm.DB) error {
	if (db.Migrator().HasTable(&entity.Car{})) {
		err := db.Migrator().DropTable(&entity.Car{})
		if err != nil {
			return err
		}
	}

	return db.Migrator().CreateTable(&entity.Car{})
}
