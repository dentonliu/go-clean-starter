package car

import (
	"net/http"
	"testing"

	"github.com/dentonliu/go-clean-starter/internal/entity"
	"github.com/dentonliu/go-clean-starter/internal/util"
)

func TestApiResource(t *testing.T) {
	r := util.MockApiRouter()
	rg := r.Group("/api")

	repo := &mockRepo{
		models: []entity.Car{},
	}

	ServeApi(rg, util.MockJWTAuth(), NewUsecase(repo))

	util.RunApiTests(t, r, []util.TestCase{
		// create failed without data
		{
			http.MethodPost,
			"/api/cars",
			"",
			http.StatusBadRequest,
			"",
		},
		// create successfully
		{
			http.MethodPost,
			"/api/cars",
			`{"brand":"BYD Yangwang U8","color":"red","seats":5}`,
			http.StatusOK,
			"BYD Yangwang U8",
		},
		// patch failed with wrong id
		{
			http.MethodPatch,
			"/api/cars/null",
			"",
			http.StatusBadRequest,
			"",
		},
		// patch failed without data
		{
			http.MethodPatch,
			"/api/cars/1",
			"",
			http.StatusBadRequest,
			"",
		},
		// patch successfully
		{
			http.MethodPatch,
			"/api/cars/1",
			`{"brand":"BYD Yangwang U8","color":"black","seats":5}`,
			http.StatusOK,
			"black",
		},
		// put failed with wrong id
		{
			http.MethodPut,
			"/api/cars/null",
			"",
			http.StatusBadRequest,
			"",
		},
		// put failed without data
		{
			http.MethodPut,
			"/api/cars/1",
			"",
			http.StatusBadRequest,
			"",
		},
		// put successfully
		{
			http.MethodPut,
			"/api/cars/1",
			`{"id":1,"brand":"BYD Yangwang U8","color":"red","seats":5,"created_at":"2023-05-18 08:05:38.270","updated_at":"2023-05-18 08:05:38.270"}`,
			http.StatusOK,
			"red",
		},
		// get failed with wrong id
		{
			http.MethodGet,
			"/api/cars/null",
			"",
			http.StatusBadRequest,
			"",
		},
		// get failed with non-existing id
		{
			http.MethodGet,
			"/api/cars/999",
			"",
			http.StatusNotFound,
			"",
		},
		// get successfully
		{
			http.MethodGet,
			"/api/cars/1",
			"",
			http.StatusOK,
			`"id":1`,
		},
		// query successfully
		{
			http.MethodGet,
			"/api/cars?per_page=20&page=1",
			"",
			http.StatusOK,
			`"total_count":1`,
		},
		// delete failed with wrong id
		{
			http.MethodDelete,
			"/api/cars/null",
			"",
			http.StatusBadRequest,
			"",
		},
		// delete successfully
		{
			http.MethodDelete,
			"/api/cars/1",
			"",
			http.StatusOK,
			"",
		},
	})
}
