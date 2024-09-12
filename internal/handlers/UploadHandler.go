package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/google/uuid"
	"github.com/ingrid-chan92/Pockethealth/persistence"
	"github.com/suyashkumar/dicom"
)

// Store all Dicom files here
const StoragePath = "persistence/dicom"

// Given a Dicom file, save the file locally to the system
func UploadFile(db persistence.Database, w http.ResponseWriter, r *http.Request) {
	// id will serve as both the file id, and file name
	id := uuid.New().String()

	// Extract source file from the request
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		fmt.Printf("Error parsing dicom file: %s\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	source, header, err := r.FormFile("dicom")
	if err != nil {
		fmt.Printf("Error retrieving file: %s\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer source.Close()

	// Validate the file is DICOM format by attempting to parse it
	_, err = dicom.Parse(source, header.Size, nil)
	if err != nil {
		fmt.Printf("Invalid DICOM file received: %s\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// Move cursor back to start of file
	_, err = source.Seek(0, io.SeekStart)
	if err != nil {
		fmt.Printf("Error reading dicom file: %s\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Create destination file
	filepath := fmt.Sprintf("%s/%s", StoragePath, id)
	destination, err := os.Create(filepath)
	if err != nil {
		fmt.Printf("Error creating file: %s\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer destination.Close()

	// Copy Source into Destination
	_, err = io.Copy(destination, source)
	if err != nil {
		fmt.Printf("Error copying source to destination: %s\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// At this point, all files have been processed
	// Save metadata for the file
	err = db.CreateMetadata(id, filepath)
	if err != nil {
		fmt.Printf("Error creating metadata: %s\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = w.Write([]byte(id))
	if err != nil {
		fmt.Printf("Error writing response for file %s: %s\n", id, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
