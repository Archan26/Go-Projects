package controllers

import (
	"encoding/json"
	"fmt"
	"golangproject/GO-Bookstore/pkg/models"
	"golangproject/GO-Bookstore/pkg/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var NewBook models.Book

func GetBook(res http.ResponseWriter, req *http.Request) {
	newBooks := models.GetAllBooks()
	response, _ := json.Marshal(newBooks)
	res.Header().Set("Content-Type", "pkglication/json")
	res.WriteHeader(http.StatusOK)
	res.Write(response)
}

func GetBookByID(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	bookDetails, _ := models.GetBookByID(ID)
	response, _ := json.Marshal(bookDetails)
	res.Header().Set("Content-Type", "pkglication/json")
	res.WriteHeader(http.StatusOK)
	res.Write(response)
}

func CreateBook(res http.ResponseWriter, req *http.Request) {
	CreateBook := &models.Book{}
	utils.ParseBody(req, CreateBook)
	b := CreateBook.CreateBook()
	response, _ := json.Marshal(b)
	res.WriteHeader(http.StatusOK)
	res.Write(response)
}

func DeletBook(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	fmt.Printf("%v", ID)
	book := models.DeleteBook(ID)
	response, _ := json.Marshal(book)
	res.Header().Set("Content-Type", "pkglication/json")
	res.WriteHeader(http.StatusOK)
	res.Write(response)
}

func UpdateBook(res http.ResponseWriter, req *http.Request) {
	var updateBook = &models.Book{}
	utils.ParseBody(req, updateBook)
	vars := mux.Vars(req)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	bookDetails, db := models.GetBookByID(ID)
	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}
	db.Save(&bookDetails)
	response, _ := json.Marshal(bookDetails)
	res.Header().Set("Content-Type", "pkglication/json")
	res.WriteHeader(http.StatusOK)
	res.Write(response)
}
