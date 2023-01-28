package script

import "testing"

func TestSeedAcademicYearData(t *testing.T) {
	ClearDataDBMigration("file://../../db/migration", config.DBUrl)
	SeedAcademicYearData(config.BasePath, server)
}
