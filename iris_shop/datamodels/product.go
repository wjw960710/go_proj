package datamodels

type Product struct {
	ID           int64  `json:"id" sql:"id" tmp:"id"`
	ProductName  string `json:"productName" sql:"product_name" tmp:"productName"`
	ProductNum   int64  `json:"productNum" sql:"product_num" tmp:"productNum"`
	ProductImage string `json:"productImage" sql:"product_image" tmp:"productImage"`
	ProductUrl   string `json:"productUrl" sql:"product_url" tmp:"productUrl"`
}
