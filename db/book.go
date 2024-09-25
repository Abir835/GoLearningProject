package db

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/Masterminds/squirrel"
	"strconv"
)

type Book struct {
	Id   *int64 `json:"id,omitempty"`
	Name string `json:"name"   db:"name"`
}

type BookRepository struct {
	table string
}

var bookRepo *BookRepository

func initBookRepo() {
	bookRepo = &BookRepository{
		table: "book",
	}
}

func GetBookRepo() *BookRepository {
	return bookRepo
}

func (r *BookRepository) CreateBook(book *Book) (*Book, error) {
	var newBook Book

	columns := map[string]interface{}{
		"name": book.Name,
	}
	var colNames []string
	var colValues []any

	for colName, colVal := range columns {
		colNames = append(colNames, colName)
		colValues = append(colValues, colVal)
	}

	query, args, err := GetQueryBuilder().
		Insert(r.table).
		Columns(colNames...).
		Values(colValues...).
		Suffix("RETURNING \"id\", name").
		ToSql()
	if err != nil {
		return nil, err
	}

	err = GetWriteDB().QueryRow(query, args...).Scan(
		&newBook.Id,
		&newBook.Name,
	)
	if err != nil {

		return nil, err
	}

	return &newBook, nil
}

func (r *BookRepository) GetAllBooks(limit, page int) ([]*Book, int, error) {
	var books []*Book
	var count int

	offset := (page - 1) * limit

	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM %s", r.table)
	err := GetReadDB().Get(&count, countQuery)
	if err != nil {
		fmt.Println("Failed to count books")
		return nil, 0, err
	}

	query, args, err := GetQueryBuilder().
		Select("*").
		From(r.table).
		OrderBy("name ASC").
		Limit(uint64(limit)).
		Offset(uint64(offset)).
		ToSql()
	if err != nil {
		fmt.Println("Failed to create books query")
		return nil, 0, err
	}

	err = GetReadDB().Select(&books, query, args...)
	if err != nil {
		fmt.Println("Failed to get books")
		return nil, 0, err
	}
	return books, count, nil
}

func (r *BookRepository) GetBookById(id string) (*Book, error) {
	var book Book
	bookId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Invalid ID:", err)
		return nil, err
	}

	query, args, err := GetQueryBuilder().
		Select("*").
		From(r.table).
		Where(squirrel.Eq{"id": bookId}).
		ToSql()

	if err != nil {
		fmt.Println("Failed to create query:", err)
		return nil, err
	}

	err = GetReadDB().QueryRow(query, args...).Scan(
		&book.Id,
		&book.Name,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("no book found with id %d", bookId)
		}
		fmt.Println("Failed to get book:", err)
		return nil, err
	}

	return &book, nil
}

func (r *BookRepository) DeleteBookById(id string) error {
	bookId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Invalid ID:", err)
		return err
	}

	query, args, err := GetQueryBuilder().
		Delete(r.table).
		Where(squirrel.Eq{"id": bookId}).
		ToSql()

	if err != nil {
		fmt.Println("Failed to create delete query:", err)
		return err
	}

	result, err := GetWriteDB().Exec(query, args...)
	if err != nil {
		fmt.Println("Failed to delete book:", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Println("Failed to check rows affected:", err)
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no book found with id %d", bookId)
	}

	return nil
}

func (r *BookRepository) UpdateBookById(id string, updatedBook *Book) (*Book, error) {

	bookId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Invalid ID:", err)
		return nil, err
	}

	columns := map[string]interface{}{
		"name": updatedBook.Name,
	}

	query, args, err := GetQueryBuilder().
		Update(r.table).
		SetMap(columns).
		Where(squirrel.Eq{"id": bookId}).
		Suffix("RETURNING id, name").
		ToSql()

	if err != nil {
		fmt.Println("Failed to create update query:", err)
		return nil, err
	}

	err = GetWriteDB().QueryRow(query, args...).Scan(
		&updatedBook.Id,
		&updatedBook.Name,
	)

	if err != nil {
		fmt.Println("Failed to update book:", err)
		return nil, err
	}

	return updatedBook, nil
}
