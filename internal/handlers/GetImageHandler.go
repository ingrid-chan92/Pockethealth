package handlers

import (
	"fmt"
	"io"
	"net/http"

	"github.com/ingrid-chan92/Pockethealth/persistence"
)

// Given a file id, return the file rendered as a png
func GetImage(db persistence.Database, w http.ResponseWriter, r *http.Request) {
	fileId := r.PathValue("id")

	// error handling: File Id must exist
	metadata, err := db.GetMetadata(fileId)
	if metadata == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	io.WriteString(w, fmt.Sprintf("Get Image id: %s\n", fileId))
}
