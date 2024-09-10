package persistence

import (
	"database/sql"

	"github.com/ingrid-chan92/Pockethealth/internal/model"
	_ "github.com/mattn/go-sqlite3"
)

const MetadataDatabaseLocation = "persistence/dicommetadata.db"

type Database interface {
	Connect() error
	Disconnect()
	GetMetadata(id string) (*model.Metadata, error)
}

type dbImpl struct {
	db *sql.DB
}

func New() Database {
	return &dbImpl{}
}

func (d *dbImpl) Connect() error {
	database, err := sql.Open("sqlite3", MetadataDatabaseLocation)
	if err != nil {
		return err
	}
	d.db = database
	return nil
}

func (d *dbImpl) Disconnect() {
	if d.db != nil {
		d.db.Close()
	}
}

func (d *dbImpl) GetMetadata(id string) (*model.Metadata, error) {
	var result model.Metadata

	row := d.db.QueryRow("SELECT id, userId, fileLocation FROM fileMetadata wHERE id = $1", id)
	err := row.Scan(&result.Id, &result.UserId, &result.FileLocation)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
