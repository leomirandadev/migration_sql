package main

import (
	"log"

	"github.com/leomirandadev/migration_sql/internal/file"
)

func execCreate(dir, filename string) {
	err := file.Create(dir, filename)
	if err != nil {
		log.Fatal("It was not possible to create file")
	}

	log.Printf("New migration on file %s/%s.sql", dir, filename)
}
