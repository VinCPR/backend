package script

import "testing"

func TestSeedGroupData(t *testing.T) {
	ClearDataDBMigration("file://../../db/migration", config.DBUrl)
	SeedAcademicYearData(config.BasePath, server)
	SeedGroupData(config.BasePath, server)
}
