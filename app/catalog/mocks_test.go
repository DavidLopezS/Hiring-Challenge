package catalog_test

import (
	"errors"

	"github.com/mytheresa/go-hiring-challenge/models"
)

type MockProductFetcher struct {
	Products []models.Product
	Total    int64
	Product  *models.Product
	Err      error
}

func (m *MockProductFetcher) GetAllProducts(offset, limit int, category string, priceLt float64) ([]models.Product, int64, error) {
	if m.Err != nil {
		return nil, 0, m.Err
	}
	return m.Products, m.Total, nil
}

func (m *MockProductFetcher) GetProductByCode(code string) (*models.Product, error) {
	if m.Err != nil {
		return nil, m.Err
	}
	if m.Product == nil {
		return nil, errors.New("not found")
	}
	return m.Product, nil
}

type MockCategoriesRepository struct {
	Categories []models.Category
	Created    *models.Category
	Err        error
}

func (m *MockCategoriesRepository) GetAll() ([]models.Category, error) {
	if m.Err != nil {
		return nil, m.Err
	}
	return m.Categories, nil
}

func (m *MockCategoriesRepository) Create(c *models.Category) error {
	if m.Err != nil {
		return m.Err
	}
	m.Created = c
	return nil
}
