// @Author Gopher
// @Date 2025/01/31 11:43:00
// @Desc 
package cart

import (
	"errors"
)

var (
	ErrItemAlreadyExistInCart = errors.New("商品已经存在")
	ErrCountInvalid           = errors.New("数量不能是负值")
)
