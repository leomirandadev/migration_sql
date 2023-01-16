package main

import (
	"log"

	"github.com/leomirandadev/migration_sql/internal/cli"
	"github.com/leomirandadev/migration_sql/internal/file_manager"
	"github.com/leomirandadev/migration_sql/internal/services"
)

func main() {
	args, err := cli.HandleParams()
	if err != nil {
		log.Fatal(err)
	}

	fileManager := file_manager.New()

	srv := services.New(args, fileManager)

	srv.Exec()

}
