package database

type IDatabase interface {
	GetLatestMigration() (string, error)
	GetAllMigrations() (map[string]bool, error)
	RunMigration(name, query, runnerGroup string) error
	RunRollback(name, query string) error
}
