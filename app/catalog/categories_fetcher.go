package catalog

import "github.com/mytheresa/go-hiring-challenge/models"

type CategoriesFetcher interface {
	GetAll() ([]models.Category, error)
	Create(c *models.Category) error
}
