package http_handler_test

import (
	"go-tdd/cmd/products/http_handler"
	"go-tdd/internal/products"
	"go-tdd/internal/products/mocks"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {

	e := echo.New()
	service := &mocks.Service{}
	handler := http_handler.NewServerHandler(e, service)

	service.On("FindById", int32(1)).Return(&products.Product{}, nil)

	t.Run("Debe encontrar un producto por ID existente", func(t *testing.T) {

		req := httptest.NewRequest(echo.GET, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/products/:id")
		c.SetParamNames("id")
		c.SetParamValues("1")

		handler.FindById(c)

		assert.Equal(t, 200, rec.Code)
	})
}
