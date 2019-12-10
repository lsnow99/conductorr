package main

import (
	"github.com/go-pg/pg/v9/orm"
	migrations "github.com/robinjoseph08/go-pg-migrations/v2"
)

func init() {
	up := func(db orm.DB) error {
		_, err := db.Exec(`
			ALTER TABLE filebot_configuration ADD fb_pushbullet text;
		`)
		return err
	}

	down := func(db orm.DB) error {
		_, err := db.Exec(`
			ALTER TABLE filebot_configuration DROP COLUMN fb_pushbullet;
	 	`)
		return err
	}

	opts := migrations.MigrationOptions{}

	migrations.Register("20191210063410_add_fb_pushbullet_column", up, down, opts)
}
