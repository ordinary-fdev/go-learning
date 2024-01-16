package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ordinary-fdev/go-learning/models"
)

func Login(ctx *gin.Context) {
	user := models.NewUserController()
	message, err := user.Login(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"message": message,
		})
	}
}

func Register(ctx *gin.Context) {
	user := models.NewUserController()

	message, err := user.Register(ctx)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"message": message,
		})
	}
}

func RefreshToken(ctx *gin.Context) {
	user := models.NewUserController()
	message := user.RefreshToken(ctx)

	// if err != nil {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{
	// 		"error": err.Error(),
	// 	})
	// } else {
	ctx.JSON(http.StatusOK, gin.H{
		"message": message,
	})
	// }

}

func GetUserBooks(ctx *gin.Context) {
	user := models.NewUserController()
	message, err := user.GetUserBooks(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"message": message,
		})
	}
}
