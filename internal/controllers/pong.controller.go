package controller

import (
	"github.com/gin-gonic/gin"
)

type Pong struct {}

func NewPong () *Pong {
	return &Pong{}
}

func (p *Pong) Pongc(c *gin.Context) {
	name := c.DefaultQuery("name", "TuanZin")
	uid := c.Query("id")

	c.JSON(200, gin.H{
		"message": "pong::::ping" + name,
		"uid":     uid,
		"user":    []string{"TuanZin", "S1", "S2"},
	})
}

func (p *Pong) Ponga(c *gin.Context) {
	name := c.DefaultQuery("name", "TuanZin")
	uid := c.Query("id")

	c.JSON(200, gin.H{
		"message": "pong::::ping" + name,
		"uid":     uid,
		"user":    []string{"TuanZin", "S1", "S2"},
	})
}
