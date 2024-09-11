package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/ingrid-chan92/Pockethealth/persistence"
	"github.com/suyashkumar/dicom"
	"github.com/suyashkumar/dicom/pkg/tag"
)

type HeaderAttributeElement struct {
	TagName string
	Element string
}

// Given a file id and a header attribute tag, search and return the associated header attribute
func QueryHeaderAttribute(db persistence.Database, w http.ResponseWriter, r *http.Request) {
	fileId := r.PathValue("id")
	tagValue := r.URL.Query().Get("tag")

	// error handling: Tag must follow a specific format
	tagInfo, err := ParseTagInfo(tagValue)
	if err != nil {
		fmt.Printf("error parsing tags %s: %s\n", tagInfo, err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// error handling: File Id must exist in metadata
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

	// Given valid tag and existing file, retrieve the header attribute element
	data, err := dicom.ParseFile(metadata.FilePath, nil)
	if err != nil {
		fmt.Printf("Error parsing dicom file for id %s: %v\n", fileId, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	element, err := data.FindElementByTag(tagInfo.Tag)
	if err != nil {
		fmt.Printf("Could not find tag %s in file %s: %s", tagInfo.Name, fileId, err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	rawResponse := HeaderAttributeElement{
		TagName: tagInfo.Name,
		Element: element.Value.String(),
	}
	response, _ := json.Marshal(rawResponse)

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func ParseTagInfo(tagParam string) (*tag.Info, error) {
	// validate format of the string: (XXXX, YYYY), where X/Y are Hexadecimal numbers
	if tagParam == "" {
		return nil, fmt.Errorf("tag cannot be empty")
	}
	match, _ := regexp.MatchString("\\([0-9a-fA-F]{4},[0-9a-fA-F]{4}\\)", tagParam)
	if !match {
		return nil, fmt.Errorf("tag is not formatted correctly")
	}

	// At this point, the string is in the correct format. Extract tag values
	splitTag := strings.Split(tagParam[1:len(tagParam)-1], ",")

	// Convert Tag values into Hexadecimal
	group, err := strconv.ParseUint(splitTag[0], 16, 16)
	if err != nil {
		return nil, fmt.Errorf("failed to parse Group from tag %s: %s", tagParam, err)
	}
	element, err := strconv.ParseUint(splitTag[1], 16, 16)
	if err != nil {
		return nil, fmt.Errorf("failed to parse Element from tag %s: %s", tagParam, err)
	}

	// Final validation: Check the tag exists and retrieve info
	parsedTag := tag.Tag{
		Group:   uint16(group),
		Element: uint16(element),
	}
	tagInfo, err := tag.Find(parsedTag)
	if err != nil {
		return nil, fmt.Errorf("failed to find tag %s: %s", parsedTag, err)
	}

	return &tagInfo, nil
}
