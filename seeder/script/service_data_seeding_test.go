package script

import (
	"testing"
)

func TestSeedServiceData(t *testing.T) {
	ClearDataDBMigration("file://../../db/migration", config.DBUrl)
	SeedHospitalData(config.BasePath, server)
	SeedSpecialtyData(config.BasePath, server)
	SeedServiceData(config.BasePath, server)
}
