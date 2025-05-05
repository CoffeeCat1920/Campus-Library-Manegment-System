package api

import (
	"CLH/internal/modals"
	"encoding/json"
	"fmt"
	"net/http"
)

func (api *api) GetAllBooksHandler(w http.ResponseWriter, r *http.Request) {
	books, err := api.db.GetAllBooks()
	if err != nil {
		http.Error(w, "Can't get books", http.StatusInternalServerError)
		fmt.Printf("Can't get books cause, %v", err)
		return
	}

	jsonData, err := json.Marshal(books)
	if err != nil {
		http.Error(w, "Can't Marshall json", http.StatusInternalServerError)
		fmt.Printf("Can't Marshall books cause, %v", err)
		return
	}

	fmt.Print(books)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func (api *api) AddBookHandler(w http.ResponseWriter, r *http.Request) {
	var Book modals.Book
	err := json.NewDecoder(r.Body).Decode(&Book)

	if err != nil {
		http.Error(w, "Form Error", http.StatusBadRequest)
		fmt.Print(err.Error())
		return
	}

	db := api.db
	err = db.AddBook(&Book)

	if err != nil {
		http.Error(w, "Can't get book from db", http.StatusInternalServerError)
		fmt.Print(err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
}
