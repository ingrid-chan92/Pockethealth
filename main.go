package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/ingrid-chan92/Pockethealth/internal/router"
	"github.com/ingrid-chan92/Pockethealth/persistence"
)

func main() {
	db := persistence.New()
	defer db.Disconnect()
	router := router.New(db)

	http.HandleFunc("/{id}/headerAttribute", router.QueryHeaderAttribute)
	http.HandleFunc("/{id}/image", router.GetImage)
	http.HandleFunc("/{id}", router.GetMetadata)
	http.HandleFunc("/", router.UploadFile)

	err := http.ListenAndServe(":3333", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("Server closed successfully\n")
	} else if err != nil {
		fmt.Printf("HTTP Server Error: %s\n", err)
		os.Exit(1)
	}
}
