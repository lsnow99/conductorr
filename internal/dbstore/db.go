package dbstore

import (
	"database/sql"
	"fmt"
	"net/url"
	"path/filepath"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v4"
	"github.com/lsnow99/conductorr/internal/settings"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB
var driver database.Driver

func Init() error {
	var err error
	if settings.UsePG {
		connStr := url.QueryEscape(fmt.Sprintf("postgres://%s:%s@%s:%s", settings.PGUser, settings.PGPass, settings.PGHost, settings.PGPort))
		if settings.PGSSL {
			connStr += "?sslmode=enable"
		} else {
			connStr += "?sslmode=disable"
		}

		db, err = sql.Open("postgres", connStr)
		if err != nil {
			return err
		}

		driver, err = postgres.WithInstance(db, &postgres.Config{})
		if err != nil {
			return err
		}
	} else {
		db, err = sql.Open("sqlite3", settings.DBPath)
		if err != nil {
			return err
		}

		driver, err = sqlite3.WithInstance(db, &sqlite3.Config{
			DatabaseName: "main",
			NoTxWrap: true,
		})
		if err != nil {
			return err
		}
	}

	migrationPath, err := filepath.Abs(settings.MigrationsPath)
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance("file://"+migrationPath, "conductorrdb", driver)
	if err != nil {
		return err
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return err
	}

	rows, err := db.Query(`SELECT count(*) FROM user`)
	if err != nil {
		return err
	}
	defer rows.Close()

	var count int

	for rows.Next() {
		if err := rows.Scan(&count); err != nil {
			return err
		}
	}

	if count == 0 {
		settings.ResetUser = true
	}

	return nil
}

func Close() error {
	return db.Close()
}
