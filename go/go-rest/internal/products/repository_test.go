package products_test

import (
	"go-tdd/internal/products"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRepository(t *testing.T) {

	//Tests de Creacion
	t.Run("Debe guardar un Producto en el map al crearlo", func(t *testing.T) {
		//Creamos una instancia de repo
		productsMap := make(map[int32]*products.Product)
		repo := products.NewRepository(productsMap)

		//Creamos un nuevo producto
		product := products.Product{"Cerveza", "Kunstmann Torobayo", 2700}
		error := repo.Create(product)

		//Verificamos que no hay errores y que se haya guardado correctamente
		assert.Nil(t, error)
		assert.NotEmpty(t, productsMap)
	})

	t.Run("Debe retornar un error al crear producto sin nombre", func(t *testing.T) {
		//Creamos una instancia de repo
		productsMap := make(map[int32]*products.Product)
		repo := products.NewRepository(productsMap)

		//Creamos un nuevo producto
		product := products.Product{
			Description: "Un producto sin nombre",
			Price:       1.00,
		}

		error := repo.Create(product)

		//Verificamos que error no sea nulo
		assert.NotNil(t, error)
	})

	t.Run("Debe retornar un error al crear producto sin descripcion", func(t *testing.T) {
		//Creamos una instancia de repo
		productsMap := make(map[int32]*products.Product)
		repo := products.NewRepository(productsMap)

		//Creamos un nuevo producto
		product := products.Product{
			Name:  "Vino",
			Price: 1.00,
		}

		error := repo.Create(product)

		//Verificamos que error no sea nulo
		assert.NotNil(t, error)
	})

	//TODO: Implementar tests para otros casos
	// Refactorizar para utilizar tabla
}
