package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/jmccallum-rbi/golearnings/crud/pkg/mocks"
	"net/http"
	"strconv"
)

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	for index, book := range mocks.Books {
		if book.Id == id {
			mocks.Books = append(mocks.Books[:index], mocks.Books[index+1:]...)

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode("Deleted Book")
			break
		}
	}
}
