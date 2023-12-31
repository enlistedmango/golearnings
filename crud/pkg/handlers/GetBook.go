package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/jmccallum-rbi/golearnings/crud/pkg/mocks"
	"net/http"
	"strconv"
)

func GetBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	for _, book := range mocks.Books {
		if book.Id == id {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			json.NewEncoder(w).Encode(book)
			break
		}
	}

}
