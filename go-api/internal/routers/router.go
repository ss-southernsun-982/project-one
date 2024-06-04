package routers

import (
	"fmt"
	"go-api/internal/controllers"
	"go-api/internal/services"
	"go-api/pkgs"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	shopService    services.ShopService       = services.NewShopService()
	shopController controllers.ShopController = controllers.NewShopController(shopService)
)

func setupFileLog() {
	f, err := os.Create("../../logs/log.txt")
	if err != nil {
		fmt.Println(err)
	}
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func NewRouter() *gin.Engine {
	setupFileLog()
	r := gin.Default()
	r.Use(Recover())
	v1 := r.Group("/v1")
	{
		v1.GET("/ping", pong)
		v1.GET("/shops", shopController.FindAll)
		v1.POST("/shops", shopController.Save)
	}
	return r
}

func pong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func Recover() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer pkgs.PanicHandler(ctx)
		ctx.Next()
	}
}
