// @Author Gopher
// @Date 2025/01/31 16:52:40
// @Desc 
package category

// 创建分类请求参数类型
type CreateCategoryRequest struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
}

// 创建分类响应参数类型
type CreateCategoryResponse struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
}
