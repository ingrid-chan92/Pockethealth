package router

import (
	"net/http"

	"github.com/ingrid-chan92/Pockethealth/internal/handlers"
	"github.com/ingrid-chan92/Pockethealth/persistence"
)

type Router struct {
	db persistence.Database
}

func New(db persistence.Database) Router {
	err := db.Connect()
	if err != nil {
		panic(err)
	}
	return Router{db: db}
}

func (router *Router) QueryHeaderAttribute(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	handlers.QueryHeaderAttribute(router.db, w, r)
}

func (router *Router) GetImage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	handlers.GetImage(router.db, w, r)
}

func (router *Router) UploadFile(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	handlers.UploadFile(router.db, w, r)
}
