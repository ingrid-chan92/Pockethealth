package persistence

import (
	"database/sql"

	"github.com/ingrid-chan92/Pockethealth/internal/model"
	_ "github.com/mattn/go-sqlite3"
)

const MetadataDatabaseLocation = "persistence/dicommetadata.db"

// For simplicity, only one user Id exists in this system
const UserId = 1

type Database interface {
	Connect() error
	Disconnect()
	GetMetadata(id string) (*model.Metadata, error)
	CreateMetadata(id string, filepath string) error
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

	row := d.db.QueryRow("SELECT id, filepath, userId FROM metadata WHERE id = $1", id)
	err := row.Scan(&result.Id, &result.FilePath, &result.UserId)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (d *dbImpl) CreateMetadata(id string, filepath string) error {
	_, err := d.db.Exec("INSERT INTO metadata (id, filepath, userId) VALUES (?, ?, ?);", id, filepath, UserId)
	if err != nil {
		return err
	}
	return nil
}
