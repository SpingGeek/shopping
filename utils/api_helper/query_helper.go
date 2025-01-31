// @Author Gopher
// @Date 2025/01/31 16:39:19
// @Desc
package api_helper

import (
	"shopping/utils/pagination"

	"github.com/gin-gonic/gin"
)

var userIdText = "userId"

// 从context获得用户id
func GetUserId(g *gin.Context) uint {
	return uint(pagination.ParseInt(g.GetString(userIdText), -1))
}
