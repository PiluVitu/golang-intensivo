package main

import (
	"database/sql"
	"fmt"
	"gobooks/internal/service"
	"gobooks/internal/web"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./books.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	bookService := service.NewBookService(db)

	bookHandler := web.NewBookHandlers(bookService)

	router := http.NewServeMux()

	router.HandleFunc("GET /books", bookHandler.GetBooks)
	router.HandleFunc("POST /books", bookHandler.CreateBook)
	router.HandleFunc("GET /books/{id}", bookHandler.GetBooksByID)
	router.HandleFunc("PUT /books/{id}", bookHandler.UpdateBook)
	router.HandleFunc("DELETE /books/{id}", bookHandler.DeleteBook)

	http.ListenAndServe(":8080", router)

	fmt.Print("Server está rodando na porta:", 8080)
}
