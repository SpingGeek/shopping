// Package user @Author Gopher
// @Date 2025/1/30 20:42:00
// @Desc
package user

import (
	"gorm.io/gorm"
	"log"
)

// Repository 结构体
type Repository struct {
	db *gorm.DB
}

// NewUserRepository 实例化
func NewUserRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

// Migration 创建表
func (r *Repository) Migration() {
	err := r.db.AutoMigrate(&User{})
	if err != nil {
		log.Print(err)
	}
}

// Create 创建用户
func (r *Repository) Create(u *User) error {
	result := r.db.Create(u)

	return result.Error
}

// GetByName 根据用户名查找用户
func (r *Repository) GetByName(name string) (User, error) {
	var user User
	err := r.db.Where("Username = ?", name).Where("IsDeleted = ?", 0.).First(&user, "Username = ?", name).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}

// InsertSampleData 添加测试数据
func (r *Repository) InsertSampleData() {

	user := NewUser("admin", "admin", "admin")
	user.IsAdmin = true
	r.db.Where(User{Username: user.Username}).Attrs(
		User{
			Username: user.Username, Password: user.Password}).FirstOrCreate(&user)
}

// Update 更新用户
func (r *Repository) Update(u *User) error {
	return r.db.Save(&u).Error
}
