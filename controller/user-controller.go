package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/ordinary-fdev/go-learning/models"
)

func Login(ctx *gin.Context) {
	user := models.NewUserController()
	message, err := user.Login(ctx)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
	} else {
		ctx.JSON(200, gin.H{
			"message": message,
		})
	}
}

func Register(ctx *gin.Context) {
	user := models.NewUserController()

	message, err := user.Register(ctx)

	if err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
	} else {
		ctx.JSON(200, gin.H{
			"message": message,
		})
	}
}
