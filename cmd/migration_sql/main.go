package main

import (
	"log"

	"github.com/leomirandadev/migration_sql/internal/cli"
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
		execUp(args.Dir, args.Driver, args.Connection)
	case "down":
		execDown(args.Dir, args.Driver, args.Connection)
	}

}
