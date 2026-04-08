package router

import (
	"cart/handler/service"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()

	r.GET("/goods/list", service.GoodsList)
	r.GET("/goods/detail", service.GoodsDetail)

	r.POST("/cart/add", service.CartAdd)
	r.GET("/cart/list", service.CartList)

	r.POST("/address/add", service.AddressAdd)

	r.POST("/order/add", service.OrderAdd)
	return r
}
