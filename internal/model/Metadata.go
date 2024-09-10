package model

type Metadata struct {
	Id           string `sql:"id"`
	UserId       int    `sql:"userId"`
	FileLocation string `sql:"fileLocation"`
}
