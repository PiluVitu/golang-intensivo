package main

import (
	"database/sql"
	"gobooks/internal/cli"
	"gobooks/internal/service"
	"gobooks/internal/web"
	"log"
	"net/http"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

// TODO: Implementar os metodos da cli no http
// TODO: Refatorar para aceitar postgres
// TODO: Fazer dockerização
// TODO: Fazer ci com testes, lint e build de multiplas versões com GHActions
// TODO: Adicionar lib de cli e substituir o switch case por ela(COBRA CLI)
// TODO: Refatorar o codigo para não deletar o valor que ja tem em um campo que não foi atualizado(testa update)
// TODO: Fazer criação da tabela de maneira automatica pelo docker se não pela aplicação
// TODO:  Jogar metodos http para a cli
func main() {
	db, err := sql.Open("sqlite3", "./books.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	bookService := service.NewBookService(db)

	bookHandler := web.NewBookHandlers(bookService)

	if len(os.Args) > 1 && (os.Args[1] == "search" || os.Args[1] == "simulate") {
		bookCLI := cli.NewBookCLI(bookService)
		bookCLI.Run()
		return
	}

	router := http.NewServeMux()

	router.HandleFunc("GET /books", bookHandler.GetBooks)
	router.HandleFunc("POST /books", bookHandler.CreateBook)
	router.HandleFunc("GET /books/{id}", bookHandler.GetBooksByID)
	router.HandleFunc("PUT /books/{id}", bookHandler.UpdateBook)
	router.HandleFunc("DELETE /books/{id}", bookHandler.DeleteBook)

	log.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
