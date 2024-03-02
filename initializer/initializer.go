package initializer

import (
	"github.com/akshaybt001/product_service/adapter"
	"github.com/akshaybt001/product_service/service"
	"gorm.io/gorm"
)

func Initialize(db *gorm.DB) *service.ProductService {
	adapter := adapter.NewProductAdapter(db)
	service := service.NewProductService(adapter)

	return service
}
