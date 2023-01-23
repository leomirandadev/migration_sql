package services

import (
	"io/fs"
	"strings"

	"github.com/google/uuid"
)

func createRunnerGroupID() string {
	newID := uuid.New()
	return newID.String()
}

func getJustNewMigrations(files []fs.DirEntry, migrations map[string]bool) []string {
	var newMigrations []string = make([]string, 0, len(files)-len(migrations))

	for _, file := range files {
		filenameSplitted := strings.Split(file.Name(), ".")
		if len(filenameSplitted) != 2 {
			continue
		}

		migration := filenameSplitted[0]
		ext := filenameSplitted[1]
		if ext == "sql" && !migrations[migration] {
			newMigrations = append(newMigrations, migration)
		}
	}

	return newMigrations
}
