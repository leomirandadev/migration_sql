package main

import (
	"log"

	"github.com/leomirandadev/migration_sql/cli"
	"github.com/leomirandadev/migration_sql/file"
)

func main() {
	args, err := cli.HandleParams()
	if err != nil {
		log.Fatal(err)
	}

	switch args.Method {
	case "create":
		execCreate(args.Dir, args.Filename)
	case "up":
		execUp()
	case "down":
		execDown()
	}

}

func execCreate(dir, filename string) {
	err := file.Create(dir, filename)
	if err != nil {
		log.Fatal("It was not possible to create file")
	}
}

func execUp() {
	// create connection with database
	// read migrations done
	// find migration that does not run yet
	// read new migrations file
	// execute migrations
	// store on database new migrations
}

func execDown() {
	// create connection with database
	// read the last migration
	// read file
	// execute rollback
	// remove from database
}
