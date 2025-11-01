package pantrydb

import (
	"context"
	"database/sql"
	"log"

	"github.com/google/uuid"
)

type PantryItem struct {
	UUID     uuid.UUID
	Name     string
	Quantity int32
}

func (PantryItem) Create(db *sql.DB, name string) (*PantryItem, error) {
	p := PantryItem{}
	p.UUID = uuid.New()
	p.Name = name
	p.Quantity = 0

	uuidBin, err := p.UUID.MarshalBinary()
	if err != nil {
		return nil, err
	}

	_, err = db.Exec("INSERT INTO PantryItems (uuid, name, quantity) VALUES (?, ?, ?);", uuidBin, p.Name, p.Quantity)
	return &p, err
}

func init_db(db *sql.DB) error {
	tx, err := db.BeginTx(context.Background(), &sql.TxOptions{})
	if err != nil {
		return err
	}
	tx.Exec("CREATE TABLE IF NOT EXISTS PantryItems (uuid BLOB PRIMARY KEY, name TEXT UNIQUE, quantity INTEGER);")
	err = tx.Commit()
	return err
}

func main() {
	var db *sql.DB

	db, err := sql.Open("sqlite3", "file:pantry.db")
	if err != nil {
		log.Fatal(err)
	}

	if err = init_db(db); err != nil {
		log.Fatal(err)
	}

}
