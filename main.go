package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/ingrid-chan92/Pockethealth/internal/handlers"
)

// Helper function for validate method for each handler
func validateAndRunHandler(method string, handler func(http.ResponseWriter, *http.Request), w http.ResponseWriter, r *http.Request) {
	if r.Method != method {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	handler(w, r)
}

func main() {
	http.HandleFunc("/{id}/headerAttribute", func(w http.ResponseWriter, r *http.Request) {
		validateAndRunHandler(http.MethodGet, handlers.QueryHeaderAttribute, w, r)
	})
	http.HandleFunc("/{id}", func(w http.ResponseWriter, r *http.Request) {
		validateAndRunHandler(http.MethodGet, handlers.GetImage, w, r)
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		validateAndRunHandler(http.MethodPost, handlers.UploadFile, w, r)
	})

	err := http.ListenAndServe(":3333", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("closed\n")
	} else if err != nil {
		fmt.Printf("error: %s\n", err)
		os.Exit(1)
	}
}
