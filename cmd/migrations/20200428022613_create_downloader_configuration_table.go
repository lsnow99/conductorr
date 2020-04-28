package main

import (
	"github.com/go-pg/pg/v9/orm"
	migrations "github.com/robinjoseph08/go-pg-migrations/v2"
)

func init() {
	up := func(db orm.DB) error {
		_, err := db.Exec(`
		CREATE TABLE public.downloader_configuration (
			downloader_configuration_id serial NOT NULL,
			name text,
			download_dir text,
			CONSTRAINT downloader_configuration_pk PRIMARY KEY (downloader_configuration_id)
		
		);
		`)
		return err
	}

	down := func(db orm.DB) error {
		_, err := db.Exec(`
		DROP TABLE IF EXISTS public.downloader_configuration;
		`)
		return err
	}

	opts := migrations.MigrationOptions{}

	migrations.Register("20200428022613_create_downloader_configuration_table", up, down, opts)
}
