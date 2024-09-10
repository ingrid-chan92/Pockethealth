package handlers

import (
	"io"
	"net/http"

	"github.com/ingrid-chan92/Pockethealth/persistence"
)

// Given a Dicom file, save the file locally to the system
func UploadFile(db persistence.Database, w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Upload File\n")
}
