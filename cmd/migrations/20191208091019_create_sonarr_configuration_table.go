package main

import (
	"github.com/go-pg/pg/v9/orm"
	migrations "github.com/robinjoseph08/go-pg-migrations/v2"
)

func init() {
	up := func(db orm.DB) error {
		_, err := db.Exec(`
			CREATE TABLE sonarr_configuration (
				sonarr_configuration_id bool NOT NULL DEFAULT true,
				sonarr_url text,
				sonarr_api_key varchar(128),
				sonarr_category text,
				CONSTRAINT sonarr_config_pk PRIMARY KEY (sonarr_configuration_id),
				CONSTRAINT sonarr_config_one_row CHECK (sonarr_configuration_id)
			
			);
		`)
		return err
	}

	down := func(db orm.DB) error {
		_, err := db.Exec("DROP TABLE sonarr_configuration")
		return err
	}

	opts := migrations.MigrationOptions{}

	migrations.Register("20191208091019_create_sonarr_configuration_table", up, down, opts)
}
