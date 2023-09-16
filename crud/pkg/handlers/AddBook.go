package handlers

import (
	"encoding/json"
	"github.com/jmccallum-rbi/golearnings/crud/pkg/mocks"
	"github.com/jmccallum-rbi/golearnings/crud/pkg/models"
	"io"
	"log"
	"math/rand"
	"net/http"
)

func AddBook(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)

	if err != nil {
		log.Fatalln()
	}

	var book models.Book
	json.Unmarshal(body, &book)

	book.Id = rand.Intn(100)
	mocks.Books = append(mocks.Books, book)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Created")
}
