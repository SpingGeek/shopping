// @Author Gopher
// @Date 2025/01/31 16:39:41
// @Desc 
package api_helper

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

// 错误处理
func HandleError(g *gin.Context, err error) {

	g.JSON(
		http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
	g.Abort()
}
