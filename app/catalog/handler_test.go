package catalog_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mytheresa/go-hiring-challenge/app/catalog"
	"github.com/mytheresa/go-hiring-challenge/models"
	"github.com/shopspring/decimal"
)

func TestGetProductByCode(t *testing.T) {
	repo := &MockProductFetcher{
		Product: &models.Product{
			Code:     "PROD001",
			Price:    decimal.NewFromFloat(10.99),
			Category: models.Category{Code: "CLOTHING", Name: "Clothing"},
			Variants: []models.Variant{
				{Name: "A", SKU: "X", Price: decimal.Decimal{}},
			},
		},
	}

	handler := catalog.NewCatalogHandler(repo)

	req := httptest.NewRequest("GET", "/catalog/PROD001", nil)
	req.SetPathValue("code", "PROD001")
	w := httptest.NewRecorder()

	handler.HandleGetByCode(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}
}
