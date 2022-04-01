package main

import (
	"database/sql"
	"github.com/labstack/echo"
	_ "github.com/mattn/go-sqlite3"
	"github.com/wevertonbruno/supermarket-api-go/adapters/persistence"
	"github.com/wevertonbruno/supermarket-api-go/adapters/web"
	"github.com/wevertonbruno/supermarket-api-go/core/services"
)

func main() {

	server := echo.New()
	database, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		server.Logger.Fatal(err)
	}

	persistence.CreateDb(database)

	catRepository := persistence.NewCategoryRepository(database)
	catService := services.NewCategoryService(catRepository)
	catController := web.NewCategoryController(catService)

	prodRepository := persistence.NewProductRepository(database)
	prodService := services.NewProductService(prodRepository)
	prodController := web.NewProductController(prodService)

	// Routes Category
	categoryRouter := server.Group("/category")
	categoryRouter.GET("", catController.FindAll)
	categoryRouter.GET("/:id", catController.FindById)
	categoryRouter.POST("", catController.Save)
	categoryRouter.PUT("/:id", catController.Update)
	categoryRouter.DELETE("/:id", catController.Delete)

	// Routes Product
	productRouter := server.Group("/product")
	productRouter.GET("", prodController.FindAll)
	productRouter.GET("/:id", prodController.FindById)
	productRouter.POST("", prodController.Save)
	productRouter.PUT("/:id", prodController.Update)
	productRouter.DELETE("/:id", prodController.Delete)

	server.Logger.Printf("Starting server at http://localhost:8080")
	server.Logger.Fatal(server.Start(":8080"))
}
