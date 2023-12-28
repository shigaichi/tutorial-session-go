package main

import (
	"log"

	"github.com/cockroachdb/errors"
	"github.com/glebarez/sqlite"
	"github.com/shigaichi/tutorial-session-go/internal/migrate"
	"gorm.io/gorm"
)

func main() {
	err := startServer()
	if err != nil {
		log.Fatalf("error %+v\n", err)
	}
}

func startServer() error {
	db, err := gorm.Open(sqlite.Open("db/db.sqlite3"), &gorm.Config{})
	if err != nil {
		return errors.Wrap(err, "failed to open db")
	}
	err = migrate.Migrate(db)
	if err != nil {
		return errors.Wrap(err, "failed to migration")
	}
	return nil
}
