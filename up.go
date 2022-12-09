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
		filename := strings.Split(file.Name(), ".")[0]
		if !migrations[filename] {
			newMigrations = append(newMigrations, filename)
		}
	}

	for _, filename := range newMigrations {
		query, err := file.Read(filename)
		if err != nil {
			log.Fatal(err)
		}

		if err := db.RunMigration(filename, query); err != nil {
			log.Fatal(err)
		}
	}
}
