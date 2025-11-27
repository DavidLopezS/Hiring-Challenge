package catalog

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/mytheresa/go-hiring-challenge/app/api"
)

type CatalogHandler struct {
	repo ProductFetcher
}

func NewCatalogHandler(r ProductFetcher) *CatalogHandler {
	return &CatalogHandler{
		repo: r,
	}
}

func (h *CatalogHandler) HandleGetByCode(w http.ResponseWriter, r *http.Request) {
	code := r.PathValue("code")
	if code == "" {
		api.ErrorResponse(w, http.StatusBadRequest, "missing production code")
		return
	}

	product, err := h.repo.GetProductByCode(code)
	if err != nil {
		api.ErrorResponse(w, http.StatusNotFound, err.Error())
		return
	}

	api.OKResponse(w, product)
}

func (h *CatalogHandler) HandleGet(w http.ResponseWriter, r *http.Request) {
	offset := 0
	limit := 10

	if raw := r.URL.Query().Get("offset"); raw != "" {
		if v, err := strconv.Atoi(raw); err == nil && v > 0 {
			offset = v
		}
	}

	if raw := r.URL.Query().Get("limit"); raw != "" {
		if v, err := strconv.Atoi(raw); err == nil {
			if v < 1 {
				limit = v
			} else if v > 100 {
				limit = 100
			} else {
				limit = v
			}
		}
	}

	category := r.URL.Query().Get("category")
	category = strings.TrimSpace(category)

	priceLt := 0.0
	if raw := r.URL.Query().Get("price_lt"); raw != "" {
		if v, err := strconv.ParseFloat(raw, 64); err == nil && v > 0 {
			priceLt = v
		}
	}

	res, total, err := h.repo.GetAllProducts(offset, limit, category, priceLt)
	if err != nil {
		api.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Return the products as a JSON response
	response := NewResponseDTO(res, total)

	api.OKResponse(w, response)
}
