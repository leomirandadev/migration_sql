package cli

import (
	"errors"
	"flag"

	"github.com/leomirandadev/migration_sql/internal/entities"
)

const DEFAULT_DRIVER = "mysql"
const DEFAULT_DIR = "./"

func HandleParams() (*entities.Method, error) {

	dirname := flag.String("dir", DEFAULT_DIR, "to inform your path where your files are")
	methodUp := flag.String("up", "", "to inform your driver")
	methodDown := flag.String("down", "", "to inform your driver")
	methodCreate := flag.String("create", "", "to inform the name of migration that will be create")
	connection := flag.String("conn", "", "to inform your database connection when you are use \"up\" or \"down\" method")

	flag.Parse()

	method, err := findTheMethodString(*methodCreate, *methodUp, *methodDown)
	if err != nil {
		return nil, err
	}

	if isMigrateMethod(method) && *connection == "" {
		return nil, errors.New("connection not found")
	}

	return &entities.Method{
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
