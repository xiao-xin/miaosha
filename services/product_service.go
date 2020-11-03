package services

import (
	"miao/datamodels"
	"miao/repositories"
)

type IProductService interface {
	GetProductByID (int64) *datamodels.Product
	GetAllProduct () []*datamodels.Product
	DeleteProductByID (int64) bool
	InsertProduct (*datamodels.Product) (int64,error)
	UpdateProduct (*datamodels.Product) error
}

type ProductService struct {
	ProductRepository repositories.IProduct
}

// 实例化ProductService
func NewProductService (repository repositories.IProduct) IProductService {
	return &ProductService{
		ProductRepository :repository,
	}
}

// 通过产品ID获取产品
func (p *ProductService) GetProductByID(productID  int64) *datamodels.Product {
	return p.ProductRepository.SelectByKey(productID)
}

// 获取所有产品
func (p *ProductService) GetAllProduct()  []*datamodels.Product {
	return p.ProductRepository.SelectAll()
}

// 根据产品ID删除产品
func (p *ProductService) DeleteProductByID(productID int64) bool {
	return p.ProductRepository.Delete(productID)
}

// 新增产品
func (p *ProductService) InsertProduct(product *datamodels.Product) (int64,error) {
	return p.ProductRepository.Insert(product)
}

// 更新产品
func (p *ProductService) UpdateProduct(product *datamodels.Product) error {
	return p.ProductRepository.Update(product)
}
