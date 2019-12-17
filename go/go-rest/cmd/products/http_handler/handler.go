package http_handler

import (
	"github.com/labstack/echo"
	"net/http"
	"strconv"

	"go-tdd/internal/products"
)

type Handler struct {
	service products.Service
}

func NewServerHandler(e *echo.Echo, service products.Service) *Handler {
	h := &Handler{
		service: service,
	}
	h.RegistryURI(e)
	return h
}

func (h *Handler) RegistryURI(e *echo.Echo) {
	e.POST("/products", h.Create)
	e.GET("/products", h.FindAll)
	e.GET("/products/:id", h.FindById)
	e.DELETE("/products/:id", h.Delete)
	e.PUT("/products/:id", h.Update)
}

func (h *Handler) Create(c echo.Context) error {
	return nil
}
func (h *Handler) FindAll(c echo.Context) error {
	return nil
}
func (h *Handler) FindById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	product, err := h.service.FindById(int32(id))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, product)
}
func (h *Handler) Delete(c echo.Context) error {
	return nil
}
func (h *Handler) Update(c echo.Context) error {
	return nil
}
