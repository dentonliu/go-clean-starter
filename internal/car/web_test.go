package car

import (
	"net/http"
	"testing"

	"github.com/dentonliu/go-clean-starter/internal/entity"
	"github.com/dentonliu/go-clean-starter/internal/util"
)

func TestWebResource(t *testing.T) {
	r := util.MockWebRouter()
	rg := r.Group("/w")

	repo := &mockRepo{
		models: []entity.Car{
			{
				Brand: "BYD Yangwang U8",
				Color: "red",
				Seats: 5,
			},
			{
				Brand: "Li Auto L9",
				Color: "silver",
				Seats: 5,
			},
		},
	}

	ServeWeb(rg, NewUsecase(repo))

	util.RunWebTests(t, r, []util.TestCase{
		// car index
		{
			http.MethodGet,
			"/w/car/index",
			"",
			http.StatusOK,
			"BYD Yangwang U8",
		},
	})
}
