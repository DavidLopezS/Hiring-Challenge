package catalog

import "github.com/mytheresa/go-hiring-challenge/models"

type Response struct {
	Products []Product `json:"products"`
}

type Category struct {
	Code string `json:"code"`
	Name string `json:"name`
}

type Product struct {
	Code     string   `json:"code"`
	Price    float64  `json:"price"`
	Category Category `json:"category"`
}

func NewProductDTO(p models.Product) Product {
	return Product{
		Code:  p.Code,
		Price: p.Price.InexactFloat64(),
		Category: Category{
			Code: p.Category.Code,
			Name: p.Category.Name,
		},
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
