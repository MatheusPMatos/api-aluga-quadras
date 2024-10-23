package server

import (
	"github.com/MatheusPMatos/api-aluga-quadras/config"
	"github.com/MatheusPMatos/api-aluga-quadras/internal/handlers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewServer(db *gorm.DB, envs config.Environments) *gin.Engine {
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})
	router(r, handlers.NewComander(db, envs))
	return r
}
