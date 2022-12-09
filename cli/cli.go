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
	if !mapMethods[method] {
		return nil, errors.New("Method not valid")
	}

	var filename string
	if isCreate(method) {
		if len(os.Args) < 3 {
			return nil, errors.New("Invalid query")
		}
		filename = os.Args[2]
	}

	return &Method{
		Method:   method,
		Filename: filename,
		Dir:      "./",
	}, nil
}

func isCreate(method string) bool {
	return method == "create"
}
