package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewRouter() *gin.Engine { // *gin.Engine: là kiểu giá trị mà hàm trả về, cụ thể là trả về một con trỏ tới một đối tượng gin.Engine (đại diện cho một router trong framework gin)
	router := gin.Default()

	// Sử dụng giống router.use()
	v1 := router.Group("/v1/go")
	{
		v1.GET("/ping/:name", Pong) // /v1/app/ping
		v1.PUT("/ping", Pong)
		v1.PATCH("/ping", Pong)
		v1.OPTIONS("/ping", Pong)
		v1.HEAD("/ping", Pong)
		v1.DELETE("/ping", Pong)
	}

	v2 := router.Group("/v2/go")
	{
		v2.GET("/ping", Pong) // /v2/app/ping
		v2.PUT("/ping", Pong)
		v2.PATCH("/ping", Pong)
		v2.OPTIONS("/ping", Pong)
		v2.HEAD("/ping", Pong)
		v2.DELETE("/ping", Pong)
	}

	return router
}

// middleware hoặc controller
func Pong(c *gin.Context) {
	// Lấy param và query
	name := c.DefaultQuery("name", "TuanZin")
	uid := c.Query("id")
	c.JSON(http.StatusOK, gin.H{ // map string
		"message": "pong::::ping" + name,
		"uid":     uid,
		"user":    []string{"TuanZin", "S1", "S2"},
	})
}
