package services

import (
	"go-api/internal/models"
)

var shops = []models.ShopModel{}

type ShopService interface {
	Save(shop models.ShopModel) models.ShopModel
	FindAll() []models.ShopModel
}

type ShopServiceImpl struct {
}

func NewShopService() ShopService {
	return &ShopServiceImpl{}
}

func (s *ShopServiceImpl) Save(shop models.ShopModel) models.ShopModel {
	shops = append(shops, shop)
	return shop
}

func (s *ShopServiceImpl) FindAll() []models.ShopModel {
	return shops
}
