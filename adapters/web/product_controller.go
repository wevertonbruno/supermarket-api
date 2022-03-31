package web

import (
	"github.com/labstack/echo"
	"github.com/wevertonbruno/supermarket-api-go/core/ports"
	"github.com/wevertonbruno/supermarket-api-go/core/ports/DTO"
	"net/http"
)

type ProductController struct {
	service ports.ProductServicePort
}

func NewProductController(service ports.ProductServicePort) *ProductController {
	return &ProductController{
		service: service,
	}
}

func (p ProductController) Save(e echo.Context) error {
	var input DTO.InputProductDTO

	err := e.Bind(&input)
	if err != nil {
		return err
	}

	err = p.service.Save(input)
	if err != nil {
		return err
	}

	return e.JSON(http.StatusCreated, nil)
}

func (p ProductController) Update(e echo.Context) error {
	var input DTO.InputProductDTO

	ID := e.Param("id")

	err := e.Bind(&input)
	if err != nil {
		return err
	}

	err = p.service.Update(ID, input)
	if err != nil {
		return err
	}

	return e.JSON(http.StatusOK, nil)
}

func (p ProductController) FindAll(e echo.Context) error {
	response := p.service.FindAll()

	return e.JSON(http.StatusOK, response)
}

func (p ProductController) FindById(e echo.Context) error {
	ID := e.Param("id")

	response, err := p.service.FindById(ID)
	if err != nil {
		return err
	}

	return e.JSON(http.StatusOK, response)
}

func (p ProductController) Delete(e echo.Context) error {
	ID := e.Param("id")

	err := p.service.Delete(ID)
	if err != nil {
		return err
	}

	return e.JSON(http.StatusOK, nil)
}
