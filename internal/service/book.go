package service

import "database/sql"

type Book struct {
	ID     int
	Title  string
	Author string
	Genre  string
}

type BookServices struct {
	db *sql.DB
}

func NewBookService(db *sql.DB) *BookServices {
	return &BookServices{db: db}
}

func (s *BookServices) CreateBook(book *Book) error {
	query := "Insert into books (title, author, genre) values(?, ?, ?)"
	result, err := s.db.Exec(query, book.Title, book.Author, book.Genre)
	if err != nil {
		return nil
	}
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return err
	}
	book.ID = int(lastInsertID)

	return nil
}

func (s *BookServices) GetBooks() ([]Book, error) {
	query := "Select id, title, author, genre from books"

	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}

	var books []Book
	for rows.Next() {
		var book Book
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Genre)
		if err != nil {
			return nil, err
		}

		books = append(books, book)
	}

	return books, nil
}

func (s *BookServices) GetBooksByID(id int) (*Book, error) {
	query := "select id, title, author, genre from books where id = ?"

	row := s.db.QueryRow(query, id)

	var book Book
	err := row.Scan(&book.ID, &book.Title, &book.Author, &book.Genre)
	if err != nil {
		return nil, err
	}

	return &book, nil
}

func (s *BookServices) UpdateBook(book *Book) error {
	query := "update books set title=?, author=?, genre=? where id=?"

	_, err := s.db.Exec(query, &book.Title, &book.Author, &book.Genre, &book.ID)

	return err
}

func (s *BookServices) DeleteBook(id int) error {
	query := "delete from books where id=?"

	_, err := s.db.Exec(query, id)

	return err
}
