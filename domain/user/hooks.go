// Package user @Author Gopher
// @Date 2025/1/30 20:38:00
// @Desc
package user

import (
	"gorm.io/gorm"
	"shopping/utils/hash"
)

func (u *User) BeforeSave(tx *gorm.DB) (err error) {

	if u.Salt == "" {
		salt := hash.CreateSalt()
		hashPassword, err := hash.HashPassword(u.Password + salt)

		if err != nil {
			return nil
		}

		u.Password = hashPassword
		u.Salt = salt
	}

	return
}
