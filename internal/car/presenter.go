package car

import "github.com/dentonliu/go-clean-starter/internal/entity"

type Car struct {
	entity.Car
}

type CarCreate struct {
	Brand string `json:"brand" form:"brand"`
	Color string `json:"color" form:"color"`
	Seats uint8  `json:"seats" form:"seats"`
}

type CarUpdate struct {
	ID        uint    `json:"id" form:"id"`
	Brand     string  `json:"brand" form:"brand"`
	Color     string  `json:"color" form:"color"`
	Seats     uint8   `json:"seats" form:"seats"`
	CreatedAt string  `json:"created_at" form:"created_at"`
	UpdatedAt string  `json:"updated_at" form:"updated_at"`
	DeletedAt *string `json:"deleted_at" form:"deleted_at"`
}

type CarPatch struct {
	Brand *string `json:"brand" form:"brand"`
	Color *string `json:"color" form:"color"`
	Seats *uint8  `json:"seats" form:"seats"`
}
