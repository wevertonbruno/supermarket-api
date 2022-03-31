package main

import (
	"github.com/labstack/echo"
	"github.com/wevertonbruno/supermarket-api-go/adapters/persistence"
	"github.com/wevertonbruno/supermarket-api-go/adapters/web"
	"github.com/wevertonbruno/supermarket-api-go/core/services"
)

func main() {

	server := echo.New()

	prodRepository := persistence.NewProductRepository()
	prodService := services.NewProductService(prodRepository)
	prodController := web.NewProductController(prodService)

	// Routes Product
	productRouter := server.Group("/product")
	productRouter.GET("", prodController.FindAll)
	productRouter.GET("/:id", prodController.FindById)
	productRouter.POST("", prodController.FindAll)
	productRouter.PUT("/:id", prodController.Update)
	productRouter.DELETE("/:id", prodController.Delete)

	server.Logger.Printf("Starting server at http://localhost:8080")
	server.Logger.Fatal(server.Start(":8080"))
}
