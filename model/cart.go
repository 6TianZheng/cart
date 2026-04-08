package model

import (
	"cart/handler/response"

	"gorm.io/gorm"
)

type Goods struct {
	gorm.Model
	Name       string  `gorm:"type:varchar(50);comment:'商品名称'"`
	Price      float64 `gorm:"type:decimal(10,2);comment:'商品价格'"`
	Images     string  `gorm:"type:varchar(50);comment:'商品图片'"`
	Bio        string  `gorm:"type:varchar(50);comment:'商品描述'"`
	CategoryId int     `gorm:"type:tinyint(1);comment:'分类id'"`
	ShopId     int     `gorm:"type:tinyint(1);comment:'店铺id'"`
	Stock      int     `gorm:"type:int(11);comment:'商品库存'"`
	Status     int     `gorm:"tinyint(1);comment:'商品状态'"`
}

func (g *Goods) GoodsDetail(db *gorm.DB, id int) ([]*response.GoodsDetail, error) {
	var list []*response.GoodsDetail
	err := db.Debug().Model(&Goods{}).
		Select("goods.*,categories.name as categoriesName,shops.name as shopName").
		Joins("left join categories on categories.id = goods.categoryId").
		Joins("left join shops on shops.id = goods.shopId").Where("id = ?", id).
		Find(&list).Error
	return list, err
}

func (g *Goods) FindGoodsById(db *gorm.DB, id int) error {
	return db.Debug().Where("id = ?", id).First(&g).Error
}

type Category struct {
	gorm.Model
	Name string `gorm:"type:varchar(20);comment:'分类名称'"`
}

type Cart struct {
	gorm.Model
	GoodsId int `gorm:"type:tinyint(1);comment:'商品id'"`
}

func (c *Cart) CartAdd(db *gorm.DB) error {
	return db.Debug().Create(&c).Error
}

type Shop struct {
	gorm.Model
	Name string `gorm:"type:varchar(20);comment:'店铺名称'"`
}

type Order struct {
	gorm.Model
	UserId  int     `gorm:"type:tinyint(1);comment:'用户id'"`
	OrderSn string  `gorm:"type:varchar(20);comment:'订单号'"`
	PayType int     `gorm:"type:tinyint(1);comment:'支付方式'"`
	Total   float64 `gorm:"type:decimal(10,2);comment:'总价钱'"`
	Status  int     `gorm:"type:tinyint(1);comment:'订单状态'"`
}

func (o *Order) OrderAdd(db *gorm.DB) error {
	return db.Debug().Create(&o).Error
}

type OrderItem struct {
	gorm.Model
	OrderId     int     `gorm:"type:tinyint(1);comment:'订单id'"`
	GoodsName   string  `gorm:"type:varchar(30);comment:'商品名称'"`
	GoodsPrice  float64 `gorm:"type:decimal(10,2);comment:'商品价格'"`
	GoodsImages string  `gorm:"type:varchar(50);comment:'商品图片'"`
	GoodsBio    string  `gorm:"type:varchar(50);comment:'商品描述'"`
}

type Address struct {
	gorm.Model
	Name string `gorm:"type:varchar(50);comment:'地址名称'"`
}

func (a *Address) AddressAdd(db *gorm.DB) error {
	return db.Debug().Create(&a).Error
}

func GoodsList(db *gorm.DB, id int) ([]*response.GoodsList, error) {
	var list []*response.GoodsList
	err := db.Debug().
		Model(&Goods{}).
		Select("goods.*,categories.name").
		Joins("left join categories on categories.id = goods.category_id").Where("categories.id = ?", id).
		Find(&list).Error
	return list, err
}

func CartList(db *gorm.DB) ([]*response.CartList, error) {
	var list []*response.CartList

	err := db.Debug().Model(&Cart{}).Select("carts.*").Find(&list).Error
	return list, err
}
