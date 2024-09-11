package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ingrid-chan92/Pockethealth/persistence"
)

// Debug Function
// Given a file id, return the metadata as JSON if it exists
func GetMetadata(db persistence.Database, w http.ResponseWriter, r *http.Request) {
	fileId := r.PathValue("id")

	// error handling: File Id must exist
	metadata, err := db.GetMetadata(fileId)
	if err != nil {
		fmt.Printf("Error retrieving metadata for id %s: %s", fileId, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if metadata == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// convert to json
	response, err := json.Marshal(metadata)
	if err != nil {
		fmt.Printf("Error marshalling metadata into json: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(response)
}
