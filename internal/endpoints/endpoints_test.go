package endpoints_test

import (
	"os"
	"testing"
	"tinytrail/test/utils"

	migrate "github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()

	os.Exit(code)
}

func setup() {
	config := utils.LoadTestConfig()
	m, err := migrate.New(
		"file://../../migrations",
		config.DatabaseURL,
	)
	if err != nil {
		panic(err)
	}

	m.Up()
}

func teardown() {
	config := utils.LoadTestConfig()
	m, err := migrate.New(
		"file://../../migrations",
		config.DatabaseURL,
	)
	if err != nil {
		panic(err)
	}

	m.Down()
	// if err = m.Down(); err != nil {
	// 	panic(err)
	// }
}
