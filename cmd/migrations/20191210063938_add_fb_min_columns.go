package main

import (
	"github.com/go-pg/pg/v9/orm"
	migrations "github.com/robinjoseph08/go-pg-migrations/v2"
)

func init() {
	up := func(db orm.DB) error {
		_, err := db.Exec(`
			ALTER TABLE filebot_configuration ADD fb_min_file_size integer;
			ALTER TABLE filebot_configuration ADD fb_min_length_ms integer;
		`)
		return err
	}

	down := func(db orm.DB) error {
		_, err := db.Exec(`
		ALTER TABLE filebot_configuration DROP COLUMN fb_min_file_size;
		ALTER TABLE filebot_configuration DROP COLUMN fb_min_length_ms;
	`)
		return err
	}

	opts := migrations.MigrationOptions{}

	migrations.Register("20191210063938_add_fb_min_columns", up, down, opts)
}
