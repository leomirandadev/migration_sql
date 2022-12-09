package cli

import (
	"errors"
	"os"
)

func HandleParams() (*Method, error) {
	if len(os.Args) < 2 {
		return nil, errors.New("Invalid query")
	}

	var method string = os.Args[1]

	switch method {
	case "create":
		return getParamsFromCreate()
	case "up":
		return getParamsToMigrate()
	case "down":
		return getParamsToMigrate()
	}

	return nil, errors.New("No method identify")
}

// ./migration_sql create $(name)
func getParamsFromCreate() (*Method, error) {
	if len(os.Args) < 3 {
		return nil, errors.New("Invalid \"create\" query")
	}

	var dir string = "./"
	if len(os.Args) >= 4 {
		dir = os.Args[3]
	}

	return &Method{
		Method:   os.Args[1],
		Filename: os.Args[2],
		Dir:      dir,
	}, nil
}

// ./migration_sql $(method) $(driver) $(connection)
func getParamsToMigrate() (*Method, error) {
	if len(os.Args) < 4 {
		return nil, errors.New("Invalid \"up\" query")
	}

	var dir string = "./"
	if len(os.Args) >= 5 {
		dir = os.Args[4]
	}

	return &Method{
		Method:     os.Args[1],
		Driver:     os.Args[2],
		Connection: os.Args[3],
		Dir:        dir,
	}, nil
}
