package handlers

import (
	"encoding/json"
	"go-learning-project/db"
	"go-learning-project/web/utils"
	"net/http"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	books, err := db.GetBookRepo().GetAllBooks()

	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, err.Error(), err)
		return
	}

	utils.SendData(w, 200, books)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	id := queryParams.Get("id")

	book, err := db.GetBookRepo().GetBookById(id)

	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, err.Error(), err)
		return
	}

	utils.SendData(w, 200, book)

}

func InsertBook(w http.ResponseWriter, r *http.Request) {

	var book db.Book

	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		utils.SendError(w, http.StatusBadRequest, "Invalid request payload", err)
		return
	}

	newBook, err := db.GetBookRepo().CreateBook(&book)

	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Failed to insert book", err)
		return
	}

	utils.SendData(w, http.StatusCreated, newBook)

}

func UpdateBookById(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	id := queryParams.Get("id")

	var book *db.Book

	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		utils.SendError(w, http.StatusBadRequest, "Invalid request payload", err)
		return
	}

	book, err := db.GetBookRepo().UpdateBookById(id, book)

	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, err.Error(), err)
		return
	}

	utils.SendData(w, 200, &book)
}

func DeleteBookById(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	id := queryParams.Get("id")

	err := db.GetBookRepo().DeleteBookById(id)

	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, err.Error(), err)
		return
	}

	utils.SendData(w, 200, "book deleted")
}
