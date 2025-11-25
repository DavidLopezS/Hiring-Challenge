package catalog

import "github.com/mytheresa/go-hiring-challenge/models"

type Response struct {
	Products []Product `json:"products"`
}

type Product struct {
	Code  string  `json:"code"`
	Price float64 `json:"price"`
}

func NewProductDTO(p models.Product) Product {
	return Product{
		Code:  p.Code,
		Price: p.Price.InexactFloat64(),
	}
}

func NewResponseDTO(products []models.Product) Response {
	dtoList := make([]Product, len(products))
	for i, p := range products {
		dtoList[i] = NewProductDTO(p)
	}

	return Response{
		Products: dtoList,
	}
}
