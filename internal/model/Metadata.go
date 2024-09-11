package model

type Metadata struct {
	Id       string `sql:"id"`
	FilePath string `sql:"filepath"`
	UserId   int    `sql:"userId"`
}
