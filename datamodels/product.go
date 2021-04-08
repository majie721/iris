package datamodels

type Product struct {
	ID           int64  `json:"id" sql:"id" form:"id"`
	Name         string `json:"name" sql:"name" form:"product_name"`
	Num          int64  `json:"num" sql:"num" form:"product_num"`
	ProductImage string `json:"product_image" sql:"product_image" form:"product_image"`
	ProductUrl   string `json:"product_url" sql:"product_url" form:"product_url"`
}
