package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/ordinary-fdev/go-learning/models"
)

func GellAllBooks(ctx *gin.Context) {
	book := models.NewBookController()

	books, err := book.GetAllBooks()
	if err != nil {
		ctx.JSON(404, gin.H{
			"error": err,
		})
	}
	ctx.JSON(200, gin.H{
		"books": books,
	})
}

func CreateBook(ctx *gin.Context) {
	book := models.NewBookController()
	id, err := book.CreateBook(ctx)

	if err != nil {
		ctx.JSON(404, gin.H{
			"error": err,
		})
	}
	ctx.JSON(200, gin.H{
		"bookID": id,
	})
}

func GetBookById(ctx *gin.Context) {
	book := models.NewBookController()
	id, err := book.GetBookById(ctx)

	if err != nil {
		ctx.JSON(404, gin.H{
			"error": err,
		})
	}
	ctx.JSON(200, gin.H{
		"bookID": id,
	})

}

func UpdateBook(ctx *gin.Context) {
	book := models.NewBookController()

	id, err := book.UpdateBook(ctx)

	if err != nil {
		ctx.JSON(404, gin.H{
			"error": err,
		})
	}
	ctx.JSON(200, gin.H{
		"bookID": id,
	})
}

func DeleteBook(ctx *gin.Context) {
	book := models.NewBookController()

	id, err := book.DeleteBook(ctx)

	if err != nil {
		ctx.JSON(404, gin.H{
			"error": err,
		})
	}
	ctx.JSON(200, gin.H{
		"bookID": id,
	})
}

func AssignBookToUser(ctx *gin.Context) {
	book := models.NewBookController()
	assignedID, err := book.AssignBookToUser(ctx)

	if err != nil {
		ctx.JSON(404, gin.H{
			"error": err,
		})
	}
	ctx.JSON(200, gin.H{
		"bookID": assignedID,
	})
}
