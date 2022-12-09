package file

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
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

func Read(filename, method string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}

	var queryBuilder strings.Builder
	var startReading bool

	reader := bufio.NewReader(file)
	for {
		buf, _, err := reader.ReadLine()
		if err != nil {
			if err != io.EOF {
				return "", err
			}
			break
		}

		line := string(buf)
		if line == "--migrate-"+method+"--" && !startReading {
			startReading = true
			continue
		}

		if startReading {
			if strings.Contains(line, "--migrate-") {
				break
			}
			queryBuilder.WriteString(line)
		}

	}

	return queryBuilder.String(), nil
}

func createFilepath(dir, filename string) string {
	return fmt.Sprintf("%s/%v_%s.sql", dir, time.Now().Unix(), filename)
}

func contentData() []byte {
	var data string = "--migrate-up--\n\n\n--migrate-down--\n"

	return []byte(data)
}
