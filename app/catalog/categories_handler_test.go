package catalog_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mytheresa/go-hiring-challenge/app/catalog"
	"github.com/mytheresa/go-hiring-challenge/models"
)

func TestGetAllCategories(t *testing.T) {
	repo := &MockCategoriesRepository{
		Categories: []models.Category{
			{Code: "CLOTHING", Name: "Clothing"},
		},
	}

	handler := catalog.NewCategoriesHandler(repo)

	req := httptest.NewRequest("GET", "/categories", nil)
	w := httptest.NewRecorder()

	handler.HandleGetAll(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}
}

func TestCreateCategory(t *testing.T) {
	repo := &MockCategoriesRepository{}
	handler := catalog.NewCategoriesHandler(repo)

	body := bytes.NewBufferString(`{"code":"NEW","name":"New Category"}`)
	req := httptest.NewRequest("POST", "/categories", body)
	w := httptest.NewRecorder()

	handler.HandleCreate(w, req)

	if w.Code != http.StatusCreated {
		t.Fatalf("expected 201, got %d", w.Code)
	}
}
