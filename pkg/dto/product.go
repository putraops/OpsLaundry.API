package dto

type NewProductDto struct {
	Name              string `json:"name" form:"name" binding:"required"`
	ProductCategoryId string `json:"product_category_id" form:"product_category_id" binding:"required"`
}
