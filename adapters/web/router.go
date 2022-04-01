package web

import (
	"database/sql"
	"github.com/labstack/echo"
	"github.com/wevertonbruno/supermarket-api-go/adapters/persistence"
	"github.com/wevertonbruno/supermarket-api-go/core/services"
	"log"
)

func UseRouter(router *echo.Echo) *echo.Echo {
	database, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		log.Fatal(err)
	}

	persistence.CreateDb(database)

	main := router.Group("api/v1")
	{
		products := main.Group("/products")
		{
			prodRepository := persistence.NewProductRepository(database)
			prodService := services.NewProductService(prodRepository)
			prodController := NewProductController(prodService)

			products.GET("", prodController.FindAll)
			products.GET("/:id", prodController.FindById)
			products.POST("", prodController.Save)
			products.PUT("/:id", prodController.Update)
			products.DELETE("/:id", prodController.Delete)
		}

		categories := main.Group("/categories")
		{
			catRepository := persistence.NewCategoryRepository(database)
			catService := services.NewCategoryService(catRepository)
			catController := NewCategoryController(catService)

			categories.GET("", catController.FindAll)
			categories.GET("/:id", catController.FindById)
			categories.POST("", catController.Save)
			categories.PUT("/:id", catController.Update)
			categories.DELETE("/:id", catController.Delete)

		}
	}

	return router
}
