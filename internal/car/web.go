package car

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/dentonliu/go-clean-starter/internal/util"
)

type webResource struct {
	usecase *Usecase
}

func ServeWeb(rg *gin.RouterGroup, usecase *Usecase) {
	r := &webResource{usecase}

	rg.GET("/car/index", r.index)
}

func (r *webResource) index(ctx *gin.Context) {
	total, err := r.usecase.Count()
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	paginatedList := util.GetPaginatedListFromRequest(ctx, total)

	models, err := r.usecase.Query(paginatedList.Offset(), paginatedList.Limit())
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	cars := []Car{}
	for _, model := range models {
		cars = append(cars, Car{model})
	}

	paginatedList.Items = cars

	ctx.HTML(http.StatusOK, "car.index", gin.H{
		"title":     "Cars",
		"paginated": paginatedList,
		"items":     cars,
	})
}
