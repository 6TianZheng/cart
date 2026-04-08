package service

import (
	"cart/basic/config"
	"cart/handler/request"
	"cart/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GoodsList(c *gin.Context) {
	var form request.GoodsList
	// This will infer what binder to use depending on the content-type header.
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":  "参数有误",
			"code": 400,
		})
		return
	}

	list, err := model.GoodsList(config.DB, form.CategoryId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":  "参数有误",
			"code": 400,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":  "查询成功",
		"code": 200,
		"list": list,
	})
	return
}

func GoodsDetail(c *gin.Context) {
	var form request.GoodsDetail
	// This will infer what binder to use depending on the content-type header.
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":  "参数有误",
			"code": 400,
		})
		return
	}

	var goods model.Goods

	detail, err := goods.GoodsDetail(config.DB, form.GoodsId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":  "参数有误",
			"code": 400,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":  "查询成功",
		"code": 200,
		"list": detail,
	})
	return
}

func CartAdd(c *gin.Context) {
	var form request.CartAdd
	// This will infer what binder to use depending on the content-type header.
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":  "参数有误",
			"code": 400,
		})
		return
	}

	var goods model.Goods

	err := goods.FindGoodsById(config.DB, form.GoodsId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":  "商品不存在",
			"code": 400,
		})
		return
	}

	cart := model.Cart{
		GoodsId: form.GoodsId,
	}

	err = cart.CartAdd(config.DB)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":  "添加失败",
			"code": 400,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":  "购物车添加成功",
		"code": 200,
	})
	return

}

func CartList(c *gin.Context) {
	list, err := model.CartList(config.DB)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":  "添加失败",
			"code": 400,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":  "添加成功",
		"code": 200,
		"list": list,
	})
	return
}

func AddressAdd(c *gin.Context) {
	var form request.AddressAdd
	// This will infer what binder to use depending on the content-type header.
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":  "参数有误",
			"code": 400,
		})
		return
	}

	address := model.Address{
		Name: form.Name,
	}

	err := address.AddressAdd(config.DB)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":  "地址添加失败",
			"code": 400,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":  "地址添加成功",
		"code": 200,
	})
	return

}

func OrderAdd(c *gin.Context) {
	var form request.OrderAdd
	// This will infer what binder to use depending on the content-type header.
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":  "参数有误",
			"code": 400,
		})
		return
	}

	order := model.Order{
		UserId:  form.UserId,
		OrderSn: form.OrderSn,
		PayType: form.PayType,
		Total:   form.Total,
		Status:  form.Status,
	}

	err := order.OrderAdd(config.DB)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":  "订单添加失败",
			"code": 400,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":  "订单添加成功",
		"code": 200,
	})
	return
}
