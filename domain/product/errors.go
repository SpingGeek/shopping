// @Author Gopher
// @Date 2025/01/30 23:57:35
// @Desc
package product

import "errors"

var (
	ErrProductNotFound         = errors.New("商品没有找到")
	ErrProductStockIsNotEnough = errors.New("商品库存不足")
)
