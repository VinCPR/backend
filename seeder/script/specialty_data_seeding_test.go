package script

import (
	"testing"
)

func TestSeedSpecialtyData(t *testing.T) {
	ClearDataDBMigration("file://../../db/migration", config.DBUrl)
	SeedSpecialtyData(config.BasePath, server)
}
