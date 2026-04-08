package response

type GoodsList struct {
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	Images     string  `json:"images"`
	Bio        string  `json:"bio"`
	CategoryId int     `json:"category_id"`
	Status     int     `json:"status"`
}

type GoodsDetail struct {
	Name           string  `json:"name"`
	Price          float64 `json:"price"`
	Images         string  `json:"images"`
	Bio            string  `json:"bio"`
	CategoryId     int     `json:"category_id"`
	Status         int     `json:"status"`
	CategoriesName string  `json:"categoriesName"`
	ShopName       string  `json:"shopName"`
}

type CartList struct {
	GoodsId int `json:"goodsId"`
}
