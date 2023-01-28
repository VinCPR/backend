package script

import "testing"

func TestSeedAdminData(t *testing.T) {
	ClearDataDBMigration("file://../../db/migration", config.DBUrl)
	SeedAdminData(config.BasePath, server)
}
