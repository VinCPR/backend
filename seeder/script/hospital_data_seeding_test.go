package script

import (
	"testing"
)

func TestSeedHospitalData(t *testing.T) {
	ClearDataDBMigration("file://../../db/migration", config.DBUrl)
	SeedHospitalData(config.BasePath, server)
}
