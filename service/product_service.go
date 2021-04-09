package service

import (
	"iris_stu/datamodels"
	"iris_stu/repositories"
)

type IProductService interface {
	GetProductByID(int64)(*datamodels.Product,error)
	GetAllProduct()([]*datamodels.Product,error)
	Insert(*datamodels.Product)(int64,error)
	Update(*datamodels.Product)(error)
	Delete(int64)(bool)
}

type ProductService struct {
	ProductRepo repositories.IProduct
}

func NewProductService(product repositories.IProduct) IProductService {
	return &ProductService{ProductRepo: product}
}

func (p *ProductService)GetProductByID(id int64)(*datamodels.Product,error){
	return p.ProductRepo.SelectByKey(id)
}

func (p *ProductService)GetAllProduct()([]*datamodels.Product,error){
	return p.ProductRepo.SelectAll()
}
func (p *ProductService)Insert(model *datamodels.Product)(int64,error){
	return p.ProductRepo.Insert(model)
}
func (p *ProductService) Update(model *datamodels.Product)(error){
	return p.ProductRepo.Update(model)
}
func (p *ProductService) Delete(id int64)(bool){
	return p.ProductRepo.Delete(id)
}
