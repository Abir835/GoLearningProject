package handlers

import (
	"encoding/json"
	"go-learning-project/db"
	"go-learning-project/web/utils"
	"math"
	"net/http"
	"strconv"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()

	page := 1
	limit := 25

	if p := queryParams.Get("page"); p != "" {
		if parsedPage, err := strconv.Atoi(p); err == nil {
			page = parsedPage
		}
	}

	if l := queryParams.Get("limit"); l != "" {
		if parsedLimit, err := strconv.Atoi(l); err == nil {
			limit = parsedLimit
		}
	}

	books, count, err := db.GetBookRepo().GetAllBooks(limit, page)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, err.Error(), err)
		return
	}

	totalPages := int(math.Ceil(float64(count) / float64(limit)))

	response := map[string]interface{}{
		"data":        books,
		"currentPage": page,
		"totalCount":  count,
		"totalPages":  totalPages,
	}

	utils.SendData(w, http.StatusOK, response)
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
