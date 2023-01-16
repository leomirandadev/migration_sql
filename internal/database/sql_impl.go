package database

import (
	"database/sql"
	"errors"
	"log"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

type implDatabase struct {
	conn *sql.DB
}

const TABLE_MIGRATIONS = "migration_sql_db"

func NewSql(driver, connection string) IDatabase {
	conn, err := sql.Open(driver, connection)
	if err != nil {
		log.Fatal(err)
	}

	db := &implDatabase{
		conn: conn,
	}

	if err := db.createTableIfNotExists(); err != nil {
		log.Fatal(err)
	}

	return db
}

func (d *implDatabase) GetLatestMigration() (string, error) {
	var data string

	d.conn.QueryRow(`
		SELECT migration FROM ` + TABLE_MIGRATIONS + ` ORDER BY id DESC LIMIT 1
	`).Scan(&data)

	if data == "" {
		return "", errors.New("error to get the latest migration")
	}

	return data, nil
}

func (d *implDatabase) GetAllMigrations() (map[string]bool, error) {
	migrations := make(map[string]bool)

	rows, err := d.conn.Query("SELECT migration FROM " + TABLE_MIGRATIONS)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var migrationName string
		if err := rows.Scan(&migrationName); err != nil {
			return nil, err
		}

		migrations[migrationName] = true
	}

	return migrations, nil
}

func (d *implDatabase) RunMigration(name, query string) error {
	_, err := d.conn.Exec(query)
	if err != nil {
		return err
	}

	_, err = d.conn.Exec(`
		INSERT INTO `+TABLE_MIGRATIONS+` (migration) VALUES (?)
	`, name)
	if err != nil {
		return err
	}

	return nil
}

func (d *implDatabase) RunRollback(name, query string) error {
	_, err := d.conn.Exec(query)
	if err != nil {
		return err
	}

	_, err = d.conn.Exec(`
		DELETE FROM `+TABLE_MIGRATIONS+` WHERE migration = ?
	`, name)
	if err != nil {
		return err
	}
	return nil
}

func (d *implDatabase) createTableIfNotExists() error {
	_, err := d.conn.Exec(`
		CREATE TABLE IF NOT EXISTS ` + TABLE_MIGRATIONS + ` (
			id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
			migration VARCHAR(100),
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)
	`)
	return err
}
