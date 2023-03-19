package dbstore

import (
	"context"
	"crypto/rand"
	"database/sql"
	"fmt"
	"io"
	"log"
	"math/big"
	"net/url"
	"os"
	"path/filepath"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/database/sqlite"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/huandu/go-sqlbuilder"
	_ "github.com/jackc/pgx/v4"
	"github.com/lsnow99/conductorr"
	"github.com/lsnow99/conductorr/internal/conductorr/integration"
	"github.com/lsnow99/conductorr/internal/conductorr/settings"
	_ "modernc.org/sqlite"
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
		db, err = sql.Open("sqlite", settings.DBPath)
		if err != nil {
			return err
		}

		db.SetMaxOpenConns(1)

		driver, err = sqlite.WithInstance(db, &sqlite.Config{
			DatabaseName: "main",
			NoTxWrap:     true,
		})
		if err != nil {
			return err
		}
	}

	migrationPath, err := filepath.Abs(settings.MigrationsPath)
	if err != nil {
		return err
	}

	if settings.BuildMode == "binary" {
		path, err := integration.MkdirTemp("conductorr_migrations")
		if err != nil {
			return err
		}
		migrationsFS := conductorr.Migrations
		files, err := migrationsFS.ReadDir("migrations")
		if err != nil {
			return err
		}
		for _, migration := range files {
			outFile, err := os.Create(filepath.Join(path, migration.Name()))
			if err != nil {
				return err
			}
			inFile, err := migrationsFS.Open(filepath.Join("migrations", migration.Name()))
			if err != nil {
				return err
			}
			_, err = io.Copy(outFile, inFile)
			if err != nil {
				return err
			}
			if err := inFile.Close(); err != nil {
				return err
			}
			if err := outFile.Close(); err != nil {
				return err
			}
		}
		migrationPath = path
	}

  log.Println("running migrations")

	migrationPath, err = CreateMigrationsFor(sqliteKey, migrationPath)
	if err != nil {
		return err
	}

	sqlbuilder.DefaultFieldMapper = sqlbuilder.SnakeCaseMapper

	pragmas := url.Values{}
	pragmas.Add("_pragma", "foreign_keys = on")

	m, err := migrate.NewWithDatabaseInstance("file://"+migrationPath, "conductorrdb?"+pragmas.Encode(), driver)
	if err != nil {
		return err
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return err
	}

  log.Println("migrations ran without Error", migrationPath)

	// cleanup temporary directory if created
	if settings.BuildMode == "binary" {
		if err := os.RemoveAll(migrationPath); err != nil {
			return err
		}
	}

	if err := initSystem(); err != nil {
		return err
	}

	if err := initUser(); err != nil {
		return err
	}

	return nil
}

func Close() error {
	return db.Close()
}

func NewTx(ctx context.Context, opts sql.TxOptions) (*sql.Tx, error) {
	return db.BeginTx(ctx, &opts)
}

func initUser() error {
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

var charList = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func initSystem() error {
	row := db.QueryRow(`SELECT jwt_secret FROM system_data WHERE id = ?`, 1)
	var nullSecret sql.NullString
	if err := row.Scan(&nullSecret); err != nil {
		return err
	}
	secret := nullSecret.String
	if len(secret) < 32 {
		secret = ""
		for i := 0; i < 128; i++ {
			index, err := rand.Int(rand.Reader, big.NewInt(int64(len(charList))))
			if err != nil {
				return err
			}
			secret += string(charList[index.Int64()])
		}
	}
	settings.JWTSecret = secret

	_, err := db.Exec(`UPDATE system_data SET jwt_secret = ? where id = ?`, secret, 1)
	return err
}

func ptrToNullString(str *string) (nstr sql.NullString) {
	if str == nil {
		return nstr
	}
	nstr.String = *str
	nstr.Valid = true
	return nstr
}

func ptrToNullInt32(i *int) (ni sql.NullInt32) {
	if i == nil {
		return ni
	}
	ni.Int32 = int32(*i)
	ni.Valid = true
	return ni
}

func ptrToNullTime(t *time.Time) (nt sql.NullTime) {
	if t == nil {
		return nt
	}
	nt.Time = *t
	nt.Valid = true
	return nt
}
