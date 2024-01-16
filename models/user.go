package models

import (
	"fmt"

	"github.com/gin-gonic/gin"
	database "github.com/ordinary-fdev/go-learning/database"
	"github.com/ordinary-fdev/go-learning/services"
)

type User struct {
	UserName string
	Password string
	Token    string
	Books    []int
}

type IUser interface {
	Register()
	Login()
	RefreshToken()
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
			// ctx.SetCookie("user", tokenstring, 3600, "/", "localhost", false, true)
			return tokenstring, nil
		}
	}
	return "", nil
}

func (user *User) RefreshToken(ctx *gin.Context) interface{} {
	err := ctx.BindJSON(&user)
	if err != nil {
		return fmt.Errorf("error in parsing body")
	}
	tokenstring, _ := services.ValidateRefreshToken(ctx, user.Token)
	return tokenstring
}

func (user *User) GetUserBooks(ctx *gin.Context) (interface{}, error) {
	db := database.GetDb()

	// queryForBooksID := fmt.Sprintf("SELECT Assignments.bookid from Assignments where userid = '%s'", user.UserName)

	username, err := services.ExtractUserNameFromToken(ctx)

	if err != nil {
		return nil, err
	}

	query := fmt.Sprintf("SELECT Books.Book_id, Books.title FROM Assignments INNER JOIN users ON users.username = Assignments.userid INNER JOIN Books on Assignments.bookid = Books.Book_id  WHERE Assignments.userid = '%s'", username)

	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error in parsing body")
	}

	defer rows.Close()
	var assignedBooks []AssignBookBody
	aBook := AssignBookBody{}
	for rows.Next() {
		err := rows.Scan(&aBook.BookID, &aBook.UserID)
		if err != nil {
			fmt.Println("Error reading rows: " + err.Error())
			return nil, err
		}
		assignedBooks = append(assignedBooks, aBook)
	}
	return assignedBooks, nil

}
