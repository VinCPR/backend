package script

import (
	"testing"
)

func TestSeedAttendingData(t *testing.T) {
	ClearDataDBMigration("file://../../db/migration", config.DBUrl)
	SeedAttendingData(config.BasePath, server)
}
