package catalog

import "github.com/mytheresa/go-hiring-challenge/models"

type ProductFetcher interface {
	GetAllProducts() ([]models.Product, error)
}
