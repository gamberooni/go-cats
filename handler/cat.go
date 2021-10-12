package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gamberooni/go-cats/model"
	"github.com/gamberooni/go-cats/store"
	"github.com/labstack/echo/v4"
)

// cat handler is a wrapper around the catstore
type CatHandler struct {
	catStore store.CatStore
}

// create cat handler instance to handle requests
func NewCatHandler(cs store.CatStore) *CatHandler {
	return &CatHandler{
		catStore: cs,
	}
}

func (ch *CatHandler) DeleteCatById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))   // convert id from string to int
	err := ch.catStore.DeleteCatById((id)) // invoke the underlying catstore method to delete cat by id
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	log.Printf("Deleted cat with ID: %v", id)
	return c.JSON(http.StatusOK, err)
}

func (ch *CatHandler) UpdateCatById(c echo.Context) error {
	cat := model.Cat{}
	err := c.Bind(&cat)
	if err != nil {
		log.Printf("Failed processing UpdateCat request: %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	id, _ := strconv.Atoi(c.Param("id"))
	result, updateError := ch.catStore.UpdateCatById(id, &cat)
	if updateError != nil {
		return c.JSON(http.StatusInternalServerError, updateError)
	}
	log.Printf("Updated cat with ID: %v", id)

	return c.JSON(http.StatusOK, result)
}

func (ch *CatHandler) GetCatById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	cat, getError := ch.catStore.GetCatById(id)
	if getError != nil {
		return c.JSON(http.StatusInternalServerError, getError)
	}
	log.Printf("Updated cat with ID: %v", id)

	return c.JSON(http.StatusOK, cat)
}

func (ch *CatHandler) GetAllCats(c echo.Context) error {
	cats, err := ch.catStore.GetAllCats()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, cats)
}

func (ch *CatHandler) AddCat(c echo.Context) error {
	cat := model.Cat{}
	err := c.Bind(&cat)
	if err != nil {
		log.Printf("Failed processing AddCat request: %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	createError := ch.catStore.AddCat(&cat)
	if createError != nil {
		return c.JSON(http.StatusInternalServerError, createError)
	}
	log.Printf("Added new cat: %#v", cat)

	return c.JSON(http.StatusOK, cat)
}
