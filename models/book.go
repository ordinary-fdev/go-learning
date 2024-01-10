package models

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	database "github.com/ordinary-fdev/go-learning/db"
)

type Book struct {
	Id       string
	Title    string
	Author   string
	Quantity int
}

type IBook interface {
	GetAllBooks()
	CreateBook()
	GetBookById()
	DeleteBook()
	UpdateBook()
}

func NewBookController() *Book {
	return &Book{}
}

func (book *Book) GetAllBooks() ([]Book, error) {
	getAllBooks_sql := "select * from BOOKS"
	db := database.GetDb()
	rows, err := db.Query(getAllBooks_sql)
	if err != nil {
		fmt.Println("Error reading records: ", err.Error())
	}
	defer rows.Close()
	var books []Book
	for rows.Next() {
		err := rows.Scan(&book.Id, &book.Title, &book.Author, &book.Quantity)
		if err != nil {
			fmt.Println("Error reading rows: " + err.Error())
			return nil, err
		}
		books = append(books, *book)
	}
	return books, nil
}

func (book *Book) GetBookById(ctx *gin.Context) (Book, error) {

	var book_id = ctx.Param("id")

	query := "select * from BOOKS where book_id =" + book_id

	db := database.GetDb()

	rows, err := db.Query(query)

	if err != nil {
		fmt.Println("Error reading records: ", err.Error())
		return *book, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&book.Id, &book.Title, &book.Author, &book.Quantity)
		if err != nil {
			fmt.Println("Error reading rows: " + err.Error())
		}
	}

	return *book, nil
}

func (book *Book) CreateBook(ctx *gin.Context) (int, error) {
	// Get the JSON data from the request body

	err := ctx.BindJSON(&book)
	if err != nil {
		fmt.Println("Can not bind json")
		return -1, err
	}

	insertProduct_sql := fmt.Sprintf("INSERT INTO BOOKS (Title, Author, Quantity) VALUES ('%s' ,'%s', %d ); select ID = convert(bigint, SCOPE_IDENTITY()) ", book.Title, book.Author, book.Quantity)

	db := database.GetDb()
	rows, err := db.Query(insertProduct_sql)

	if err != nil {
		fmt.Println("Can not bind json")
		return -1, err
	}

	var lastInsertId1 int
	for rows.Next() {
		rows.Scan(&lastInsertId1)
	}
	defer rows.Close()
	return lastInsertId1, err

}

func (book *Book) UpdateBook(ctx *gin.Context) (int, error) {
	bookID, _ := strconv.Atoi(ctx.Param("id"))

	err := ctx.BindJSON(&book)

	if err != nil {
		fmt.Println("Can not bind json")
		return -1, err
	}

	if bookID == -1 {
		return -1, nil
	}

	db := database.GetDb()

	checkIfBookExistQuery := fmt.Sprintf("SELECT COUNT(*) FROM BOOKS WHERE Book_id=%d", bookID)

	var dataExist bool

	rows, err := db.Query(checkIfBookExistQuery)

	if err != nil {
		fmt.Println("Error in finding data")
	}

	for rows.Next() {
		err := rows.Scan(&dataExist)
		if err != nil {
			fmt.Println("Error in finding data")
			return -1, err
		}
		if !dataExist {
			fmt.Println("data doesn't exist finding data")
			return -1, err
		}
	}

	query := fmt.Sprintf("UPDATE BOOKS SET Title='%s', Author='%s', Quantity=%d WHERE book_id = %d", book.Title, book.Author, book.Quantity, bookID)

	rows, err = db.Query(query)

	if err != nil {
		fmt.Println("Can not bind json")
		return -1, err
	}

	defer rows.Close()
	return bookID, err
}

func (book *Book) DeleteBook(ctx *gin.Context) (int, error) {
	bookID, _ := strconv.Atoi(ctx.Param("id"))

	err := ctx.BindJSON(&book)

	if err != nil {
		fmt.Println("Can not bind json")
		return -1, err
	}

	if bookID == -1 {
		return -1, nil
	}

	db := database.GetDb()

	checkIfBookExistQuery := fmt.Sprintf("SELECT COUNT(*) FROM BOOKS WHERE Book_id=%d", bookID)

	var dataExist bool

	rows, err := db.Query(checkIfBookExistQuery)

	if err != nil {
		fmt.Println("Error in finding data")
	}

	for rows.Next() {
		err := rows.Scan(&dataExist)
		if err != nil {
			fmt.Println("Error in finding data")
			return -1, err
		}
		if !dataExist {
			fmt.Println("data doesn't exist finding data")
			return -1, err
		}
	}
	query := fmt.Sprintf("DELETE FROM BOOKS WHERE Book_id=%d", bookID)
	rows, err = db.Query(query)
	if err != nil {
		fmt.Println("Can not bind json")
		return -1, err
	}

	defer rows.Close()
	return bookID, err
}
