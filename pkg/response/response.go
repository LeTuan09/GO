package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int         `json:"code"`    // Mã lỗi
	Message string      `json:"message"` // Thông báo lỗi
	Data    interface{} `json:"data"`    // Dữ liệu được return
}

func SuccessResponse(c *gin.Context, code int, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Message: msg[code],
		Data: data,
	})
}

func ErrorResponse(c *gin.Context, code int, message string) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Message: msg[code],
		Data: nil,
	})
}