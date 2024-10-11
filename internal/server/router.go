package server

import (
	"github.com/MatheusPMatos/api-aluga-quadras/internal/handlers"
	"github.com/gin-gonic/gin"
)

func router(router *gin.Engine, cmd handlers.Comander) {
	main := router.Group("/")
	{
		unAuth := main.Group("/user")
		{
			unAuth.POST("", cmd.User.Create)
		}

		user := main.Group("/user", cmd.Midlewares.Auth())
		{
			user.GET("/:id", cmd.User.GetById)
			user.DELETE("/:id", cmd.User.Delete)
			user.PUT("", cmd.User.Edit)
		}

	}
}
