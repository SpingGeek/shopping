// Package user æ@Author Gopher
// @Date 2025/1/30 20:54:00
// @Desc
package user

import "shopping/utils/hash"

// Service 用户 service 结构体
type Service struct {
	r Repository
}

// NewUserService 实例化 Service
func NewUserService(r Repository) *Service {
	r.Migration()
	r.InsertSampleData()
	return &Service{
		r: r,
	}
}

// Create 创建用户
func (c *Service) Create(user *User) error {

	if user.Password != user.Password2 {
		return ErrMismatchedPasswords
	}

	_, err := c.r.GetByName(user.Username)
	if err == nil {
		return ErrUserExistWithName
	}

	if ValidateUserName(user.Username) {
		return ErrInvalidUsername
	}

	if ValidatePassword(user.Password) {
		return ErrInvalidPassword
	}

	err = c.r.Create(user)
	return err
}

// GetUser 查询用户
func (c *Service) GetUser(username string, password string) (User, error) {
	user, err := c.r.GetByName(username)
	if err != nil {
		return User{}, ErrUserNotFound
	}

	match := hash.CheckPasswordHash(password+user.Salt, user.Password)
	if !match {
		return User{}, ErrUserNotFound
	}
	return user, nil
}

// UpdateUser 更新用户
func (c *Service) UpdateUser(user *User) error {
	return c.r.Update(user)
}
