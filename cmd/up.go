package main

import (
	"log"
	"os"
	"strings"

	"github.com/leomirandadev/migration_sql/internal/database"
	"github.com/leomirandadev/migration_sql/internal/file"
)

func execUp(dir, driver, connection string) {
	db := database.New(driver, connection)

	migrations, err := db.GetAllMigrations()
	if err != nil {
		log.Fatal(err)
	}

	files, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	var newMigrations []string = make([]string, 0, len(files)-len(migrations))
	for _, file := range files {
		migration := strings.Split(file.Name(), ".")[0]
		if !migrations[migration] {
			newMigrations = append(newMigrations, migration)
		}
	}

	for _, migration := range newMigrations {
		query, err := file.Read(dir+"/"+migration+".sql", "up")
		if err != nil {
			log.Fatal(err)
		}

		if err := db.RunMigration(migration, query); err != nil {
			log.Fatal(err)
		}

		log.Printf("Migrate %s.sql", migration)
	}

	log.Printf("All migrations were executed")
}
