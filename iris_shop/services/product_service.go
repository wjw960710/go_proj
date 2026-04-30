package services

import (
	"iris_shop/datamodels"
	"iris_shop/repositories"
)

type ProductService interface {
	GetProductByID(productID int64) (*datamodels.Product, error)
	GetAllProduct() ([]*datamodels.Product, error)
	DeleteProduct(productID int64) bool
	InsertProduct(*datamodels.Product) (int64, error)
	UpdateProduct(*datamodels.Product) error
}

type ProductServiceImpl struct {
	productRepository repositories.ProductRepository
}

func NewProductService(repository repositories.ProductRepository) ProductService {
	return &ProductServiceImpl{
		productRepository: repository,
	}
}

func (p *ProductServiceImpl) GetProductByID(productID int64) (*datamodels.Product, error) {
	return p.productRepository.SelectById(productID)
}

func (p *ProductServiceImpl) GetAllProduct() ([]*datamodels.Product, error) {
	return p.productRepository.SelectAll()
}

func (p *ProductServiceImpl) DeleteProduct(productID int64) bool {
	return p.productRepository.Delete(productID)
}

func (p *ProductServiceImpl) InsertProduct(product *datamodels.Product) (int64, error) {
	return p.productRepository.Insert(product)
}

func (p *ProductServiceImpl) UpdateProduct(product *datamodels.Product) error {
	return p.productRepository.Update(product)
}
