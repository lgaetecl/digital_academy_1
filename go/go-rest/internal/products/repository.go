package products

import "errors"

//Repository ...
type Repository interface {
	Create(products Product) error
	FindByID(id int32) (*Product, error)
	FindAll() *[]Product
	Delete(id int32) error
	Update(id int32, products Product) (*Product, error)
}

type repository struct {
	productsMap map[int32]*Product
}

//NewRepository ...
func NewRepository(productsMap map[int32]*Product) Repository {
	return &repository{
		productsMap: productsMap,
	}
}

func (r *repository) Create(product Product) error {
	if len(product.Name) == 0 {
		return errors.New("Producto debe tener nombre")
	}

	if len(product.Description) == 0 {
		return errors.New("Producto debe tener descripcion")
	}

	id := int32(len(r.productsMap) + 1)
	r.productsMap[id] = &product

	return nil
}

func (r *repository) FindByID(id int32) (*Product, error) {
	return nil, nil
}

func (r *repository) FindAll() *[]Product {
	return nil
}

func (r *repository) Delete(id int32) error {
	return nil
}

func (r *repository) Update(id int32, products Product) (*Product, error) {
	return nil, nil
}
