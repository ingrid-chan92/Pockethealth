package handlers

import (
	"fmt"
	"io"
	"net/http"

	"github.com/ingrid-chan92/Pockethealth/persistence"
)

// Given a file id and a header attribute tag, search and return the associated header attribute
func QueryHeaderAttribute(db persistence.Database, w http.ResponseWriter, r *http.Request) {
	fileId := r.PathValue("id")
	headerTag := r.URL.Query().Get("tag")

	// error handling: Tag must be passed in as a query param
	if headerTag == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

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

	// at this point, parameters are all valid

	io.WriteString(w, fmt.Sprintf("Query Header Attribute id: %s, attribute: %s\n", fileId, headerTag))
}
