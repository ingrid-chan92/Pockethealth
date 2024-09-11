package handlers_test

import (
	"testing"

	"github.com/ingrid-chan92/Pockethealth/internal/handlers"
	"github.com/stretchr/testify/assert"
	"github.com/suyashkumar/dicom/pkg/tag"
)

func TestParseTagInfo_happyPath(t *testing.T) {
	tagInfo, err := handlers.ParseTagInfo("(0002,0002)")
	assert.NoError(t, err)
	assert.Equal(t, tag.Tag{
		Group:   uint16(2),
		Element: uint16(2),
	}, tagInfo.Tag)
	assert.Equal(t, "MediaStorageSOPClassUID", tagInfo.Name)
}

func TestParseTagInfo_happyPath_hex(t *testing.T) {
	tagInfo, err := handlers.ParseTagInfo("(300E,0002)")
	assert.NoError(t, err)
	assert.Equal(t, tag.Tag{
		Group:   uint16(12302),
		Element: uint16(2),
	}, tagInfo.Tag)
	assert.Equal(t, "ApprovalStatus", tagInfo.Name)
}

func TestParseTagInfo_emptyTag(t *testing.T) {
	tagInfo, err := handlers.ParseTagInfo("")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "tag cannot be empty")
	assert.Nil(t, tagInfo)
}

func TestParseTagInfo_invalidFormat(t *testing.T) {
	tagInfo, err := handlers.ParseTagInfo("abcdefghi")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "tag is not formatted correctly")
	assert.Nil(t, tagInfo)
}

func TestParseTagInfo_unknownTag(t *testing.T) {
	tagInfo, err := handlers.ParseTagInfo("(FFFF,FFFF)")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to find tag")
	assert.Nil(t, tagInfo)
}
