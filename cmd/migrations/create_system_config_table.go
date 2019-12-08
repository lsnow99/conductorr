package migrations

import (
	"github.com/go-pg/pg/v9/orm"
	migrations "github.com/robinjoseph08/go-pg-migrations/v2"
)

func init() {
	up := func(db orm.DB) error {
		_, err := db.Exec(`
			CREATE TABLE public."SystemConfiguration" (
				system_configuration_id bool NOT NULL DEFAULT true,
				username varchar(32),
				password bytea,
				filebot_logs_enabled bool DEFAULT true,
				plex_scanner_logs_enabled bool DEFAULT true,
				CONSTRAINT system_config_pk PRIMARY KEY (system_configuration_id),
				CONSTRAINT configuration_one_row CHECK (system_configuration_id)
			
			);
		`)
		return err
	}

	down := func(db orm.DB) error {
		_, err := db.Exec("DROP TABLE users")
		return err
	}

	opts := migrations.MigrationOptions{}

	migrations.Register("create_system_config_table", up, down, opts)
}