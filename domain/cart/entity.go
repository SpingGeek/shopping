// @Author Gopher
// @Date 2025/01/31 11:34:20
// @Desc
package cart

import (
	"shopping/domain/product"
	"shopping/domain/user"

	"gorm.io/gorm"
)

// 购物车结构体
type Cart struct {
	gorm.Model
	UserID uint
	User   user.User `gorm:"foreignKey:ID;references:UserID"`
}

// 实例化
func NewCart(uid uint) *Cart {
	return &Cart{
		UserID: uid,
	}
}

// Item
type Item struct {
	gorm.Model
	Product   product.Product `gorm:"foreignKey:ProductID"`
	ProductID uint
	Count     int
	CartID    uint
	Cart      Cart `gorm:"foreignKey:CartID" json:"-"`
}

// 创建Item
func NewCartItem(productId uint, cartId uint, count int) *Item {
	return &Item{
		ProductID: productId,
		Count:     count,
		CartID:    cartId,
	}
}
