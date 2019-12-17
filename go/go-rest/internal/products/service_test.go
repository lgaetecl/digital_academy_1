package products_test

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"go-tdd/internal/products"
	"go-tdd/internal/products/mocks"
	"testing"
)

func TestServiceTest(t *testing.T) {

	repo := &mocks.Repository{}
	service := products.NewService(repo)

	repo.On("Delete", int32(15)).Return(nil)
	repo.On("Delete", int32(10)).Return(errors.New("Error"))

	t.Run("Debe poder eliminar un producto correctamente", func(t *testing.T) {
		err := service.Delete(15)

		assert.Nil(t, err)
	})

	t.Run("Debe fallar al eliminar en caso de que el repositorio falle", func(t *testing.T) {
		err := service.Delete(10)

		assert.NotNil(t, err)
	})
}
