// Package user @Author Gopher
// @Date 2025/1/30 20:41:00
// @Desc
package user

import "regexp"

// 匹配以字母开头，后面跟着 7-29 个字母、数字或下划线的字符串。
var usernameRegex = regexp.MustCompile("^[A-Za-z][A-Za-z0-9_]{7,29}$")

// 匹配以字母开头，后面跟着 7-29 个字母、数字或下划线的字符串，该字符串必须包含至少一个数字。
var passwordRegex = regexp.MustCompile(`^[A-Za-z][A-Za-z0-9_]{7,29}$`)

func ValidateUserName(name string) bool {
	return usernameRegex.MatchString(name)
}

func ValidatePassword(password string) bool {
	return passwordRegex.MatchString(password)

}
