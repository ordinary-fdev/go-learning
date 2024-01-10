package models

import (
	"fmt"

	"github.com/gin-gonic/gin"
	database "github.com/ordinary-fdev/go-learning/db"
	"github.com/ordinary-fdev/go-learning/services"
)

type User struct {
	UserName string
	Password string
	Token    string
}

type IUser interface {
	Register()
	Login()
}

func NewUserController() *User {
	return &User{}
}

func (user *User) Register(ctx *gin.Context) (string, error) {
	err := ctx.BindJSON(&user)
	if err != nil {
		fmt.Println("Error in parsing body")
	}
	db := database.GetDb()
	checkIfBookExistQuery := fmt.Sprintf("SELECT COUNT(*) FROM USERS WHERE username='%s'", user.UserName)
	var dataExist bool
	rows, err := db.Query(checkIfBookExistQuery)
	if err != nil {
		fmt.Println("Error in finding data")
	}
	for rows.Next() {
		err := rows.Scan(&dataExist)
		if err != nil {
			fmt.Println("Error in finding data")
			return "", err
		}
		if dataExist {
			return "", fmt.Errorf("username already exists")
		}
	}
	query := fmt.Sprintf("INSERT INTO USERS (USERNAME,PASSWORD) VALUES('%s','%s')", user.UserName, user.Password)
	rows, err = db.Query(query)
	if err != nil {
		return "", fmt.Errorf("error in inserting data")
	}
	defer rows.Close()
	return "User Created Successfully", nil
}

func (user *User) Login(ctx *gin.Context) (interface{}, error) {
	err := ctx.BindJSON(&user)
	if err != nil {
		return "", fmt.Errorf("error in parsing body")
	}
	db := database.GetDb()
	checkUserInfo := fmt.Sprintf("SELECT COUNT(*) FROM USERS WHERE username='%s' and password = '%s'", user.UserName, user.Password)
	var dataExist bool
	rows, err := db.Query(checkUserInfo)
	if err != nil {
		fmt.Println("Error in finding data")
	}
	for rows.Next() {
		err := rows.Scan(&dataExist)
		if err != nil {
			fmt.Println("Error in finding data")
			return "", err
		}
		if dataExist {
			tokenstring, err := services.GenerateToken(user.UserName)
			if err != nil {
				return "", fmt.Errorf("error in token generation")
			}
			ctx.SetCookie("user", tokenstring, 3600, "/", "localhost", false, true)
			return tokenstring, nil
		}
	}
	return "", nil
}
