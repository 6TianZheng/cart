package request

type GoodsList struct {
	CategoryId int `form:"categoryId"  binding:"required"`
}

type GoodsDetail struct {
	GoodsId int `form:"goodsId"  binding:"required"`
}

type CartAdd struct {
	GoodsId int `form:"goodsId"  binding:"required"`
}

type AddressAdd struct {
	Name string `form:"name"  binding:"required"`
}

type OrderAdd struct {
	UserId  int     `form:"userId"  binding:"required"`
	OrderSn string  `form:"orderSn"  binding:"required"`
	PayType int     `form:"payType"  binding:"required"`
	Total   float64 `form:"total"  binding:"required"`
	Status  int     `form:"status"  binding:"required"`
}
