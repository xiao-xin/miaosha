package datamodels

// 商品模型
type Product struct {
	ID int64 `json:"id" sql:"id" form:"id"`
	Name string `json:"name" sql:"name" form:"name"`
	Num int64 `json:"num" sql:"num" form:"num"`
	Image string `json:"image" sql:"image" form:"image"`
	Url  string `json:"url" sql:"url" form:"url"`
	Bo  bool `sql:"bo"`
}
