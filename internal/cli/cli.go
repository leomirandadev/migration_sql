package cli

import (
	"errors"
	"flag"
)

const DEFAULT_DRIVER = "mysql"
const DEFAULT_DIR = "./"

func HandleParams() (*Method, error) {

	dirname := flag.String("dir", DEFAULT_DIR, "")
	methodUp := flag.String("up", "", DEFAULT_DRIVER)
	methodDown := flag.String("down", "", DEFAULT_DRIVER)
	methodCreate := flag.String("create", "", "")
	connection := flag.String("conn", "", "")

	flag.Parse()

	method, err := findTheMethodString(*methodCreate, *methodUp, *methodDown)
	if err != nil {
		return nil, err
	}

	if isMigrateMethod(method) && *connection == "" {
		return nil, errors.New("connection not found")
	}

	return &Method{
		Method:     method,
		Driver:     getDriver(*methodUp, *methodDown),
		Connection: *connection,
		Filename:   *methodCreate,
		Dir:        *dirname,
	}, nil
}

func isMigrateMethod(method string) bool {
	switch method {
	case "up":
		return true
	case "down":
		return true
	default:
		return false
	}
}

func findTheMethodString(create, up, down string) (string, error) {
	if create == "" && !driversAllowed[up] && !driversAllowed[down] {
		return "", errors.New("any method has found")
	}

	method := "down"
	if up != "" {
		method = "up"
	}
	if create != "" {
		method = "create"
	}

	return method, nil
}

func getDriver(methodUp, methodDown string) string {
	if methodDown != "" {
		return methodDown
	}
	if methodUp != "" {
		return methodUp
	}
	return DEFAULT_DRIVER
}
