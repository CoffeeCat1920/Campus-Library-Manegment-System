package api

import (
	"CLH/internal/database"
	"net/http"

	"github.com/spf13/afero"
)

type BookInfo struct {
	ISBN   string `json:"isbn"`
	Name   string `json:"name"`
	Author string `json:"author"`
	Genre  string `json:"genre"`
}

type Api interface {
	GetAllBooksHandler(w http.ResponseWriter, r *http.Request)
	AddBookHandler(w http.ResponseWriter, r *http.Request)
}

type api struct {
	db database.Service
	fs afero.Fs
}

func NewApi() Api {
	return &api{
		db: database.New(),
		fs: afero.OsFs{},
	}
}
