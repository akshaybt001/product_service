package adapter

import "github.com/akshaybt001/product_service/entities"

type AdapterInterface interface {
	AddProduct(req entities.Products)(entities.Products,error)
}