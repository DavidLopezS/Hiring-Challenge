package models

type Category struct {
	ID   uint   `grom:"primariKey"`
	Code string `gorm:"uniqueIndex;not null"`
	Name string `gorm:"not null"`

	Products []Product
}

func (c *Category) TableName() string {
	return "product_categories"
}
