// @Author Gopher
// @Date 2025/01/31 11:12:18
// @Desc
package product

import (
	"shopping/utils/pagination"
)

// 商品服务结构体
type Service struct {
	productRepository Repository
}

// Service 实例化
func NewService(productRepository Repository) *Service {
	productRepository.Migration()
	return &Service{
		productRepository: productRepository,
	}
}

// 获得所有商品分页
func (c *Service) GetAll(page *pagination.Pages) *pagination.Pages {
	products, count := c.productRepository.GetAll(page.Page, page.PageSize)
	page.Items = products
	page.TotalCount = count
	return page
}

// 创建商品
func (c *Service) CreateProduct(name string, decs string, count int, price float32, cid uint) error {
	newProduct := NewProduct(name, decs, count, price, cid)
	err := c.productRepository.Create(newProduct)
	return err
}

// 删除商品
func (c *Service) DeleteProduct(sku string) error {
	err := c.productRepository.Delete(sku)
	return err
}

// 更新商品
func (c *Service) UpdateProduct(product *Product) error {
	err := c.productRepository.Update(*product)
	return err
}

// 查找商品
func (c *Service) SearchProduct(text string, page *pagination.Pages) *pagination.Pages {
	products, count := c.productRepository.SearchByString(text, page.Page, page.PageSize)
	page.Items = products
	page.TotalCount = count
	return page
}
