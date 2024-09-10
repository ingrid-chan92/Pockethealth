package handlers

import (
	"fmt"
	"io"
	"net/http"
)

// Given a file id and a header attribute tag, search and return the associated header attribute
func QueryHeaderAttribute(w http.ResponseWriter, r *http.Request) {
	fileId := r.PathValue("id")
	headerTag := r.URL.Query().Get("tag")

	// error handling: Tag must be passed in as a query param
	if headerTag == "" {
		w.WriteHeader(http.StatusBadRequest)
	}

	// error handling: File Id must exist

	// at this point, all input should be valid
	io.WriteString(w, fmt.Sprintf("Query Header Attribute id: %s, attribute: %s\n", fileId, headerTag))
}

// Given a file id, return the file rendered as a png
func GetImage(w http.ResponseWriter, r *http.Request) {
	fileId := r.PathValue("id")
	io.WriteString(w, fmt.Sprintf("Get Image id: %s\n", fileId))
}
