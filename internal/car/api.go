package car

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/dentonliu/go-clean-starter/internal/entity"
	"github.com/dentonliu/go-clean-starter/internal/util"
)

type apiResource struct {
	usecase *Usecase
}

func ServeApi(rg *gin.RouterGroup, authHandler gin.HandlerFunc, usecase *Usecase) {
	r := &apiResource{usecase}

	rg.GET("/cars/:id", r.get)
	rg.GET("/cars", r.query)

	authRouter := rg.Group("/")
	authRouter.Use(authHandler)
	authRouter.POST("/cars", r.create)
	authRouter.PUT("/cars/:id", r.update)
	authRouter.PATCH("/cars/:id", r.patch)
	authRouter.DELETE("/cars/:id", r.delete)
}

func (r *apiResource) get(ctx *gin.Context) {
	paramID := ctx.Param("id")
	id, err := strconv.ParseUint(paramID, 10, 64)

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	model, err := r.usecase.Get(uint(id))
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if model.ID == 0 {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	ctx.JSON(http.StatusOK, Car{model})
}

func (r *apiResource) query(ctx *gin.Context) {
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

	ctx.JSON(http.StatusOK, paginatedList)
}

func (r *apiResource) create(ctx *gin.Context) {
	var form CarCreate

	if err := ctx.ShouldBindJSON(&form); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	model := entity.Car{
		Brand: form.Brand,
		Color: form.Color,
		Seats: form.Seats,
	}

	err := r.usecase.Create(&model)

	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, Car{model})
}

func (r *apiResource) update(ctx *gin.Context) {
	paramID := ctx.Param("id")
	id, err := strconv.ParseUint(paramID, 10, 64)

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var form CarUpdate

	if err = ctx.ShouldBindJSON(&form); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	model, err := r.usecase.Get(uint(id))
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	model.Brand = form.Brand
	model.Color = form.Color
	model.Seats = form.Seats
	date, err := time.Parse("2006-01-02 15:04:05.000", form.CreatedAt)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	model.CreatedAt = date
	date, err = time.Parse("2006-01-02 15:04:05.000", form.UpdatedAt)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	model.UpdatedAt = date
	if form.DeletedAt != nil {
		date, err = time.Parse("2006-01-02 15:04:05.000", *form.DeletedAt)
		if err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}
		model.DeletedAt = gorm.DeletedAt{
			Time:  date,
			Valid: true,
		}
	}

	err = r.usecase.Update(&model)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, Car{model})
}

func (r *apiResource) patch(ctx *gin.Context) {
	paramID := ctx.Param("id")
	id, err := strconv.ParseUint(paramID, 10, 64)

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var form CarPatch

	if err = ctx.ShouldBindJSON(&form); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	model, err := r.usecase.Get(uint(id))
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	attrs := []string{}
	if form.Brand != nil {
		model.Brand = *form.Brand
		attrs = append(attrs, "Brand")
	}

	if form.Color != nil {
		model.Color = *form.Color
		attrs = append(attrs, "Color")
	}

	if form.Seats != nil {
		model.Seats = *form.Seats
		attrs = append(attrs, "Seats")
	}

	err = r.usecase.Update(&model, attrs...)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, Car{model})
}

func (r *apiResource) delete(ctx *gin.Context) {
	paramID := ctx.Param("id")
	id, err := strconv.ParseUint(paramID, 10, 64)

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	model, err := r.usecase.Delete(uint(id))
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, Car{model})
}
