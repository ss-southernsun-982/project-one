package models

type ShopModel struct {
	ShopID   int    `json:"shop_id" binding:"required" errormgs:"shop_id is required"`
	ShopName string `json:"shop_name" binding:"required,min=2,max=150" errormgs:"shop_name is required"`
	ShopURL  string `json:"shop_url" binding:"required,url" errormgs:"shop_url is required"`
}
