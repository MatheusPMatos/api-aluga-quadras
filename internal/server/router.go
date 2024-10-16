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
			unAuth.POST("/auth", cmd.Auth.Login)
		}

		user := main.Group("/user", cmd.Midlewares.Auth())
		{
			user.GET("/:id", cmd.User.GetById)
			user.DELETE("/:id", cmd.User.Delete)
			user.PUT("", cmd.User.Edit)
			user.GET("/info", cmd.User.UserInfo)
		}

		products := main.Group("products")
		{
			products.GET("")
		}

		reservas := main.Group("/reserva")
		{
			reservas.GET("/byproduct/:id", cmd.Rservation.GetByProduct)
			reservas.GET("/byuser/:id", cmd.Rservation.GetByUser)
			reservas.POST("", cmd.Rservation.Create)
			reservas.PUT("", cmd.Rservation.Edit)
			reservas.DELETE("/:id", cmd.Rservation.Delete)
		}

	}
}
