package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ordinary-fdev/go-learning/controller"
	"github.com/ordinary-fdev/go-learning/middleware"
)

func InitializeRoutes() {
	r := gin.Default()

	protectedGrp := r.Group("books")

	protectedGrp.Use(middleware.ValidateRequest())
	protectedGrp.GET("/", controller.GellAllBooks)
	protectedGrp.POST("/createBook", controller.CreateBook)
	protectedGrp.GET("/getBookById/:id", controller.GetBookById)
	protectedGrp.POST("/updateBook/:id", controller.UpdateBook)
	protectedGrp.POST("/deleteBook/:id", controller.DeleteBook)
	protectedGrp.POST("/assingBook", controller.AssignBookToUser)
	protectedGrp.POST("/userBooks", controller.GetUserBooks)

	r.POST("/login", controller.Login)
	r.POST("/register", controller.Register)
	r.POST("/refreshToken", controller.RefreshToken)
	http.ListenAndServe(":8080", r)
}
