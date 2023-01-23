package services

import (
	"log"
	"os"

	"github.com/leomirandadev/migration_sql/internal/database"
	"github.com/leomirandadev/migration_sql/internal/entities"
	"github.com/leomirandadev/migration_sql/internal/file_manager"
)

func New(method *entities.Method, fileManager file_manager.IFileReader) IService {
	return &servicesImpl{method, fileManager}
}

type IService interface {
	Exec()
	create()
	up()
	down()
}

type servicesImpl struct {
	method      *entities.Method
	fileManager file_manager.IFileReader
}

func (s *servicesImpl) Exec() {
	switch s.method.Method {
	case "create":
		s.create()
	case "up":
		s.up()
	case "down":
		s.down()
	case "down-group":
		s.downGroup()
	}
}

func (s *servicesImpl) create() {
	err := s.fileManager.Create(s.method.Dir, s.method.Filename)
	if err != nil {
		log.Fatal("It was not possible to create file")
	}

	log.Printf("New migration on file %s/%s.sql", s.method.Dir, s.method.Filename)
}

func (s *servicesImpl) up() {
	db := database.NewSql(s.method.Driver, s.method.Connection)

	migrations, err := db.GetAllMigrations()
	if err != nil {
		log.Fatal(err)
	}

	files, err := os.ReadDir(s.method.Dir)
	if err != nil {
		log.Fatal(err)
	}

	var newMigrations []string = getJustNewMigrations(files, migrations)

	runnerGroup := createRunnerGroupID()

	for _, migration := range newMigrations {
		query, err := s.fileManager.Read(s.method.Dir+"/"+migration+".sql", "up")
		if err != nil {
			log.Fatal(err)
		}

		if err := db.RunMigration(migration, query, runnerGroup); err != nil {
			log.Fatal(err)
		}

		log.Printf("Migrate %s.sql", migration)
	}

	log.Printf("All migrations were executed")
}

func (s *servicesImpl) down() {
	db := database.NewSql(s.method.Driver, s.method.Connection)

	migration, err := db.GetLatestMigration()
	if err != nil {
		log.Fatal(err)
	}

	query, err := s.fileManager.Read(s.method.Dir+"/"+migration+".sql", "down")
	if err != nil {
		log.Fatal(err)
	}

	if err := db.RunRollback(migration, query); err != nil {
		log.Fatal(err)
	}

	log.Printf("Rollback on %s.sql migration", migration)
}

func (s *servicesImpl) downGroup() {
	db := database.NewSql(s.method.Driver, s.method.Connection)

	migrations, err := db.GetLatestRunnerGroupMigrations()
	if err != nil {
		log.Fatal(err)
	}

	for _, migration := range migrations {
		query, err := s.fileManager.Read(s.method.Dir+"/"+migration+".sql", "down")
		if err != nil {
			log.Fatal(err)
		}

		if err := db.RunRollback(migration, query); err != nil {
			log.Fatal(err)
		}

		log.Printf("Rollback on %s.sql migration", migration)
	}

}
