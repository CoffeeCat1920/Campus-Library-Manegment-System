package api

import (
	"CLH/internal/database"
	"net/http"

	"github.com/spf13/afero"
)

type Api interface {
	GetAllBooksHandler(w http.ResponseWriter, h *http.Request)
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
