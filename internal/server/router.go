package server

import "github.com/gin-gonic/gin"

func router(router *gin.Engine) {
	main := router.Group("/")
	{
		aluno := main.Group("/user")
		{
			aluno.GET("")
			aluno.GET("/:id")
			aluno.DELETE("/:id")
			aluno.POST("")
			aluno.PUT("")
		}

	}
}
