package endpoints_test

import (
	"os"
	"testing"

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
	m, err := migrate.New(
		"file://../../migrations",
		"postgresql://postgres:postgres@localhost:5122/tinytrail_test",
	)
	if err != nil {
		panic(err)
	}

	m.Up()
	// if err = m.Up(); err != nil {
	// 	panic(err)
	// }
}

func teardown() {
	m, err := migrate.New(
		"file://../../migrations",
		"postgresql://postgres:postgres@localhost:5122/tinytrail_test",
	)
	if err != nil {
		panic(err)
	}

	m.Down()
	// if err = m.Down(); err != nil {
	// 	panic(err)
	// }
}
