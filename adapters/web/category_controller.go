package web

import (
	"github.com/labstack/echo"
	"github.com/wevertonbruno/supermarket-api-go/core/ports"
	"github.com/wevertonbruno/supermarket-api-go/core/ports/DTO"
	"net/http"
)

type CategoryController struct {
	service ports.CategoryServicePort
}

func NewCategoryController(service ports.CategoryServicePort) *CategoryController {
	return &CategoryController{
		service: service,
	}
}

func (p CategoryController) Save(e echo.Context) error {
	var input DTO.InputCategoryDTO

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

func (p CategoryController) Update(e echo.Context) error {
	var input DTO.InputCategoryDTO

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

func (p CategoryController) FindAll(e echo.Context) error {
	response := p.service.FindAll()

	return e.JSON(http.StatusOK, response)
}

func (p CategoryController) FindById(e echo.Context) error {
	ID := e.Param("id")

	response, err := p.service.FindById(ID)
	if err != nil {
		return e.JSON(http.StatusBadRequest, err)
	}

	return e.JSON(http.StatusOK, response)
}

func (p CategoryController) Delete(e echo.Context) error {
	ID := e.Param("id")

	err := p.service.Delete(ID)
	if err != nil {
		return err
	}

	return e.JSON(http.StatusOK, nil)
}
