package main

import (
	"log"

	"github.com/leomirandadev/migration_sql/internal/database"
	"github.com/leomirandadev/migration_sql/internal/file"
)

func execDown(dir, driver, connection string) {
	db := database.New(driver, connection)

	migration, err := db.GetLatestMigration()
	if err != nil {
		log.Fatal(err)
	}

	query, err := file.Read(dir+"/"+migration+".sql", "down")
	if err != nil {
		log.Fatal(err)
	}

	if err := db.RunRollback(migration, query); err != nil {
		log.Fatal(err)
	}
}
