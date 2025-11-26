package catalog

import "github.com/mytheresa/go-hiring-challenge/models"

type ProductFetcher interface {
	GetAllProducts(offset, limit int, category string, priceLt float64) ([]models.Product, int64, error)
	GetProductByCode(code string) (*models.Product, error)
}
