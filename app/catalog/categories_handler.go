package catalog

import (
	"encoding/json"
	"net/http"

	"github.com/mytheresa/go-hiring-challenge/app/api"
	"github.com/mytheresa/go-hiring-challenge/models"
)

type CategoriesHandler struct {
	repo CategoriesFetcher
}

func NewCategoriesHandler(r CategoriesFetcher) *CategoriesHandler {
	return &CategoriesHandler{repo: r}
}

func (h *CategoriesHandler) HandleGetAll(w http.ResponseWriter, r *http.Request) {
	categories, err := h.repo.GetAll()
	if err != nil {
		api.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	api.OKResponse(w, categories)
}

func (h *CategoriesHandler) HandleCreate(w http.ResponseWriter, r *http.Request) {
	var c models.Category
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		api.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if c.Code == "" || c.Name == "" {
		api.ErrorResponse(w, http.StatusBadRequest, "code and name are required")
		return
	}

	if err := h.repo.Create(&c); err != nil {
		api.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	api.OKResponse(w, c)
}
