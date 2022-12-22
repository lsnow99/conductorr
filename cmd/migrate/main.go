package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
)

func main() {

	databasePtr := flag.String("database", "sqlite://conductorr.db", "Database connection string to use for the force command (default sqlite://conductorrdb)")
	sourcePtr := flag.String("source", "file://migrations", "Migrations source (default file://migrations)")
	pathPtr := flag.String("path", "migrations/", "Migrations folder to use (default migrations/)")

	flag.Parse()

	cmd := flag.Arg(0)

	switch cmd {
	case "help":
		flag.Usage()
		return
	case "create":
		name := flag.Arg(1)

		if flag.NArg() == 1 {
			log.Fatal("error: please specify migration name")
		}

		if err := createCmd(*pathPtr, time.Now(), "unix", name, "sql", true, 6, true); err != nil {
			log.Fatal(err)
		}
		return
  case "up":
    migrater, migraterErr := migrate.New(*sourcePtr, *databasePtr)
    defer func() {
      if migraterErr == nil {
        if _, err := migrater.Close(); err != nil {
          log.Println(err)
        }
      }
    }()
    if migraterErr == nil {
      if err := upCmd(migrater, 1); err != nil {
        log.Println(err)
      }
    } else {
      log.Fatal(migraterErr)
    }
  case "down":
    migrater, migraterErr := migrate.New(*sourcePtr, *databasePtr)
    defer func() {
      if migraterErr == nil {
        if _, err := migrater.Close(); err != nil {
          log.Println(err)
        }
      }
    }()
    if migraterErr == nil {
      if err := downCmd(migrater, 1); err != nil {
        log.Println(err)
      }
    } else {
      log.Fatal(migraterErr)
    }
	case "force":
		if flag.NArg() == 1 {
			log.Fatal("error: please specify version argument V")
		}

		v, err := strconv.ParseInt(flag.Arg(1), 10, 64)
		if err != nil {
			log.Fatal("error: can't read version argument V")
		}

		if v < -1 {
			log.Fatal("error: argument V must be >= -1")
		}

		log.Println(*databasePtr)
		migrater, migraterErr := migrate.New(*sourcePtr, *databasePtr)
		defer func() {
			if migraterErr == nil {
				if _, err := migrater.Close(); err != nil {
					log.Println(err)
				}
			}
		}()
		if migraterErr == nil {
			migrater.PrefetchMigrations = 10
			migrater.LockTimeout = time.Duration(int64(15)) * time.Second

			// handle Ctrl+c
			signals := make(chan os.Signal, 1)
			signal.Notify(signals, syscall.SIGINT)
			go func() {
				for range signals {
					log.Println("Stopping after this running migration ...")
					migrater.GracefulStop <- true
					return
				}
			}()
		}

		if migraterErr != nil {
			log.Fatal(migraterErr)
		}

		if err := forceCmd(migrater, int(v)); err != nil {
			log.Fatal(err)
		}
	default:
		log.Fatalln(fmt.Errorf("unrecognized command: %s", cmd))
	}
}
