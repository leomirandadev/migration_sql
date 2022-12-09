package file

import (
	"fmt"
	"os"
	"time"
)

func Create(dir, filename string) error {
	filepath := createFilepath(dir, filename)

	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	file.Write(contentData())

	return nil
}

func Read(method string) (string, error) {

	return "", nil
}

func createFilepath(dir, filename string) string {
	return fmt.Sprintf("%s/%v_%s.sql", dir, time.Now().Unix(), filename)
}

func contentData() []byte {
	var data string = "--migrate-up--\n\n\n--migrate-down--\n"

	return []byte(data)
}
