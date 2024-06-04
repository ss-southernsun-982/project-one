package controllers

import (
	"go-api/internal/constants"
	"go-api/internal/models"
	"go-api/internal/services"
	"go-api/pkgs"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ShopController interface {
	FindAll(ctx *gin.Context)
	Save(ctx *gin.Context)
}

type ShopControllerImpl struct {
	shopService services.ShopService
}

func NewShopController(shopService services.ShopService) ShopController {
	return &ShopControllerImpl{shopService: shopService}
}

func (s *ShopControllerImpl) FindAll(ctx *gin.Context) {
	shops := s.shopService.FindAll()
	ctx.JSON(http.StatusOK, pkgs.BuildResponse(constants.Success, shops))
}

func (s *ShopControllerImpl) Save(ctx *gin.Context) {
	var shop models.ShopModel
	if err := ctx.ShouldBindJSON(&shop); err != nil {
		pkgs.PanicException(constants.BadRequest, err)
	}
	shopModel := s.shopService.Save(shop)
	ctx.JSON(http.StatusCreated, pkgs.BuildResponse(constants.Created, shopModel))
}
