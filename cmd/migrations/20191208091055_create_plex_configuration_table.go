package main

import (
	"github.com/go-pg/pg/v9/orm"
	migrations "github.com/robinjoseph08/go-pg-migrations/v2"
)

func init() {
	up := func(db orm.DB) error {
		_, err := db.Exec(`
			CREATE TABLE plex_configuration (
				plex_configuration_id bool NOT NULL DEFAULT true,
				plex_namespace text,
				plex_deployment_name text,
				plex_auth_token text,
				plex_base_url text,
				CONSTRAINT plex_config_pk PRIMARY KEY (plex_configuration_id),
				CONSTRAINT plex_config_one_row CHECK (plex_configuration_id)
			
			);
`		)
		return err
	}

	down := func(db orm.DB) error {
		_, err := db.Exec("DROP TABLE plex_configuration")
		return err
	}

	opts := migrations.MigrationOptions{}

	migrations.Register("20191208091055_create_plex_configuration_table", up, down, opts)
}
