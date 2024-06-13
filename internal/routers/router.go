package routers

import (
	c "GO-API-SERVER/internal/controllers"
	// "net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine { // *gin.Engine: là kiểu giá trị mà hàm trả về, cụ thể là trả về một con trỏ tới một đối tượng gin.Engine (đại diện cho một router trong framework gin)
	router := gin.Default()

	// Sử dụng giống router.use()
	v1 := router.Group("/v1/go")
	{
		v1.GET("/users", c.NewUserController().GetByUserID)

	}

	return router
}
