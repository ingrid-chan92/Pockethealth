package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/google/uuid"
	"github.com/ingrid-chan92/Pockethealth/persistence"
)

// Store all Dicom files here
const StoragePath = "persistence/dicom"

// Given a Dicom file, save the file locally to the system
func UploadFile(db persistence.Database, w http.ResponseWriter, r *http.Request) {
	// id will serve as both the file id, and file name
	id := uuid.New().String()

	// Extract source file from the request
	r.ParseMultipartForm(32 << 20)
	source, _, err := r.FormFile("dicom")
	if err != nil {
		fmt.Printf("Error retrieving file: %s\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer source.Close()

	// Validate Dicom File here

	// Create destination file
	filepath := fmt.Sprintf("%s/%s", StoragePath, id)
	destination, err := os.Create(filepath)
	if err != nil {
		fmt.Printf("Error creating file: %s\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer destination.Close()

	// At this point, source and destination exist. Copy Source into Destination
	_, err = io.Copy(destination, source)
	if err != nil {
		fmt.Printf("Error copying source to destination: %s\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Save metadata for the file
	err = db.CreateMetadata(id, filepath)
	if err != nil {
		fmt.Printf("Error creating metadata: %s\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write([]byte(id))
}
