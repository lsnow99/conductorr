package main

import (
	"github.com/go-pg/pg/v9/orm"
	migrations "github.com/robinjoseph08/go-pg-migrations/v2"
)

func init() {
	up := func(db orm.DB) error {
		_, err := db.Exec(`
			CREATE TABLE radarr_configuration (
				radarr_configuration_id bool NOT NULL DEFAULT true,
				radarr_url text,
				radarr_api_key varchar(128),
				radarr_category text,
				CONSTRAINT radarr_config_pk PRIMARY KEY (radarr_configuration_id),
				CONSTRAINT radarr_config_one_row CHECK (radarr_configuration_id)
			
			);
		`)
		return err
	}

	down := func(db orm.DB) error {
		_, err := db.Exec("DROP TABLE radarr_configuration")
		return err
	}

	opts := migrations.MigrationOptions{}

	migrations.Register("20191208091050_create_radarr_configuration_table", up, down, opts)
}
