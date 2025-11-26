package catalog

import "github.com/mytheresa/go-hiring-challenge/models"

type Response struct {
	Total    int       `json:"total"`
	Products []Product `json:"products"`
}

type Category struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type Variant struct {
	Name  string  `json:"name"`
	SKU   string  `json:"sku"`
	Price float64 `json:"price"`
}

type Product struct {
	Code     string    `json:"code"`
	Price    float64   `json:"price"`
	Category Category  `json:"category"`
	Variants []Variant `json:"variants"`
}

func NewVariantDTO(v models.Variant, productPrice float64) Variant {
	price := productPrice
	if !v.Price.IsZero() {
		price = v.Price.InexactFloat64()
	}

	return Variant{
		Name:  v.Name,
		SKU:   v.SKU,
		Price: price,
	}
}

func NewProductDTO(p models.Product) Product {
	variants := make([]Variant, len(p.Variants))
	for i, v := range p.Variants {
		variants[i] = NewVariantDTO(v, p.Price.InexactFloat64())
	}

	return Product{
		Code:  p.Code,
		Price: p.Price.InexactFloat64(),
		Category: Category{
			Code: p.Category.Code,
			Name: p.Category.Name,
		},
		Variants: variants,
	}
}

func NewResponseDTO(products []models.Product, total int64) Response {
	dtoList := make([]Product, len(products))
	for i, p := range products {
		dtoList[i] = NewProductDTO(p)
	}

	return Response{
		Total:    int(total),
		Products: dtoList,
	}
}
