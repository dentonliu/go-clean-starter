package main

import (
	"fmt"
	"net/http"
	"os"
	"path"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/dentonliu/go-clean-starter/internal/car"
	"github.com/dentonliu/go-clean-starter/internal/config"
	"github.com/dentonliu/go-clean-starter/internal/entity"
	"github.com/dentonliu/go-clean-starter/internal/middleware"
	util "github.com/dentonliu/go-clean-starter/internal/util"
)

func main() {
	c, err := config.Load("../configs", "./configs")
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(mysql.Open(c.DSN), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// database initialization
	db.Migrator().AutoMigrate(&entity.Car{})

	r := gin.Default()
	r.Use(cors.Default())

	buildApiRoutes(r, c, db)
	buildWebRoutes(r, c, db)

	r.Run(fmt.Sprintf(":%s", c.ServerPort))
}

// buildApiRoutes registers api routes
func buildApiRoutes(r *gin.Engine, config *config.Config, db *gorm.DB) {
	rg := r.Group("/api")

	rg.GET("/ping", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "pong")
	})

	authHandler := middleware.JWTAuth(config.JWTVerificationKey, middleware.JWTOptions{
		SigningMethod: config.JWTSigningMethod,
	})

	// register car apis
	car.ServeApi(
		rg,
		authHandler,
		car.NewUsecase(car.NewDBRepo(db)),
	)
}

// buildWebRoutes register web routes
func buildWebRoutes(r *gin.Engine, config *config.Config, db *gorm.DB) {
	root, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}

	r.Static("/static", path.Join(root, "/../web/static"))
	r.HTMLRender = util.CreateRenderer(path.Join(root, "/../web/template"))

	rg := r.Group("/w")

	rg.GET("/ping", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "pong")
	})

	car.ServeWeb(rg, car.NewUsecase(car.NewDBRepo(db)))
}
