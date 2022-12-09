package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type implDatabase struct {
	conn *sql.DB
}

func New(driver, connection string) IDatabase {
	conn, err := sql.Open(driver, connection)
	if err != nil {
		log.Fatal(err)
	}

	return &implDatabase{
		conn: conn,
	}
}

func (d *implDatabase) GetLatestMigration() (string, error) {
	return "", nil
}

func (d *implDatabase) GetAllMigrations() (map[string]bool, error) {
	migrations := make(map[string]bool)
	return migrations, nil
}

func (d *implDatabase) RunMigration(name, query string) error {
	//run query and if its succeed, store on table migrations
	return nil
}

func (d *implDatabase) RunRollback(name, query string) error {
	//run query and if its succeed, delete on table migrations
	return nil
}
