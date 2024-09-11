package handlers

import (
	"bytes"
	"fmt"
	"image/png"
	"net/http"

	"github.com/ingrid-chan92/Pockethealth/persistence"
	"github.com/suyashkumar/dicom"
	"github.com/suyashkumar/dicom/pkg/tag"
)

// Given a file id, return the file rendered as a png
func GetImage(db persistence.Database, w http.ResponseWriter, r *http.Request) {
	fileId := r.PathValue("id")

	// error handling: File Id must exist
	metadata, err := db.GetMetadata(fileId)
	if err != nil {
		fmt.Printf("Error retrieving metadata for id %s: %s\n", fileId, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if metadata == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Retrieve Dicom file, and get the Pixel Data
	data, err := dicom.ParseFile(metadata.FilePath, nil)
	if err != nil {
		fmt.Printf("Error parsing dicom file for id %s: %v\n", fileId, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	pixelDataElement, err := data.FindElementByTag(tag.PixelData)
	if err != nil {
		fmt.Printf("Could not find Pixel Data Tag in file %s: %s", fileId, err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Convert pixel data into PNG
	pixelDataInfo := dicom.MustGetPixelDataInfo(pixelDataElement.Value)
	image, _ := pixelDataInfo.Frames[0].GetImage()
	var buffer bytes.Buffer
	if png.Encode(&buffer, image) != nil {
		fmt.Printf("Error encoding image as PNG for file %s: %s\n", fileId, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(buffer.Bytes())
}
