package model

type Metadata struct {
	Id           string `sql:"id"`
	FileLocation string `sql:"filepath"`
	UserId       int    `sql:"userId"`
}
