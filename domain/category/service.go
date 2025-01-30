// Package category @Author Gopher
// @Date 2025/1/30 21:52:00
// @Desc
package category

import (
	"mime/multipart"
	"shopping/utils/csv_helper"
	"shopping/utils/pagination"
)

type Service struct {
	r Repository
}

// NewCategoryService 实例化商品分类service
func NewCategoryService(r Repository) *Service {
	// 生成表
	r.Migration()
	// 插入测试数据
	r.InsertSampleData()
	return &Service{
		r: r,
	}
}

// Create 创建分类
func (c *Service) Create(category *Category) error {
	existCity := c.r.GetByName(category.Name)
	if len(existCity) > 0 {
		return ErrCategoryExistWithName
	}

	err := c.r.Create(category)
	if err != nil {
		return err
	}

	return nil
}

// BulkCreate 批量创建分类
func (c *Service) BulkCreate(fileHeader *multipart.FileHeader) (int, error) {
	categories := make([]*Category, 0)
	bulkCategory, err := csv_helper.ReadCsv(fileHeader)
	if err != nil {
		return 0, err
	}
	for _, categoryVariables := range bulkCategory {
		categories = append(categories, NewCategory(categoryVariables[0], categoryVariables[1]))
	}
	count, err := c.r.BulkCreate(categories)
	if err != nil {
		return count, err
	}
	return count, nil
}

// GetAll 获得分页商品分类
func (c *Service) GetAll(page *pagination.Pages) *pagination.Pages {
	categories, count := c.r.GetAll(page.Page, page.PageSize)
	page.Items = categories
	page.TotalCount = count
	return page
}
