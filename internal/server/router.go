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

		products := main.Group("/products", cmd.Midlewares.Auth())
		{
			products.GET("", cmd.Products.GetAll)
			products.POST("", cmd.Products.Create)
			products.DELETE("", cmd.Products.Delete)
			products.GET("/:id", cmd.Products.GetById)
			products.PUT("", cmd.Products.Update)
		}

		reservas := main.Group("/reserva", cmd.Midlewares.Auth())
		{
			reservas.GET("/byproduct/:id", cmd.Rservation.GetByProduct)
			reservas.GET("/byuser/:id", cmd.Rservation.GetByUser)
			reservas.POST("", cmd.Rservation.Create)
			reservas.PUT("", cmd.Rservation.Edit)
			reservas.DELETE("/:id", cmd.Rservation.Delete)
		}
		schedule := main.Group("/schedule")
		{
			schedule.POST("/:id", cmd.Schedule.GetByProductAndDate)
		}

	}
}
