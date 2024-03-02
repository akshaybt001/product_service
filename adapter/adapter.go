package adapter

import (
	"github.com/akshaybt001/product_service/entities"
	"gorm.io/gorm"
)

type ProductAdapter struct {
	DB *gorm.DB
}

// AddProduct implements AdapterInterface.
func (*ProductAdapter) AddProduct(req entities.Products) (entities.Products, error) {
	panic("unimplemented")
}

func NewProductAdapter(db *gorm.DB) *ProductAdapter {
	return &ProductAdapter{
		DB: db,
	}
}

// func (product *ProductAdapter)AddProduct(req entities.Products)(entities.Products,error){
// 	var res entities.Products
// 	quer
// }
