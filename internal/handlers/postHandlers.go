package handlers

import (
	"io"
	"net/http"
)

// Given a Dicom file, save the file locally to the system
func uploadFile(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Upload File\n")
}
