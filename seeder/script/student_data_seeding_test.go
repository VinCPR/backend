package script

import (
	"testing"
)

func TestSeedStudentData(t *testing.T) {
	ClearDataDBMigration("file://../db/migration", config.DBUrl)
	SeedStudentData(config.BasePath, server)
}
