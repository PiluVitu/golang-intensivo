package web

import (
	"encoding/json"
	"gobooks/internal/service"
	"net/http"
	"strconv"
)

type BookHandlers struct {
	service *service.BookServices
}

func NewBookHandlers(service *service.BookServices) *BookHandlers {
	return &BookHandlers{service: service}
}

func (h *BookHandlers) CreateBook(w http.ResponseWriter, r *http.Request) {
	var book service.Book

	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := h.service.CreateBook(&book); err != nil {
		http.Error(w, "Faild to create book", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}

func (h *BookHandlers) GetBooks(w http.ResponseWriter, r *http.Request) {
	books, err := h.service.GetBooks()
	if err != nil {
		http.Error(w, "Failed to get Books", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func (h *BookHandlers) GetBooksByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	book, err := h.service.GetBookByID(id)
	if err != nil {
		http.Error(w, "Failed to get book", http.StatusInternalServerError)
		return
	}
	if book == nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

// UpdateBook lida com a requisição PUT /books/{id}.
func (h *BookHandlers) UpdateBook(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid book ID", http.StatusBadRequest)
		return
	}

	var book service.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "invalid request payload", http.StatusBadRequest)
		return
	}
	book.ID = id

	if err := h.service.UpdateBook(&book); err != nil {
		http.Error(w, "failed to update book", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)
}

// DeleteBook lida com a requisição DELETE /books/{id}.
func (h *BookHandlers) DeleteBook(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid book ID", http.StatusBadRequest)
		return
	}

	if err := h.service.DeleteBook(id); err != nil {
		http.Error(w, "failed to delete book", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
