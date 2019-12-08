package main

import (
	"github.com/go-pg/pg/v9/orm"
	migrations "github.com/robinjoseph08/go-pg-migrations/v2"
)

func init() {
	up := func(db orm.DB) error {
		_, err := db.Exec(`
			CREATE TABLE filbot_configuration (
				filebot_configuration_id bool NOT NULL DEFAULT true,
				fb_output_dir text,
				fb_subtitles_locale varchar(16),
				fb_action varchar(16),
				fb_exclude_list text,
				fb_amc_log text,
				fb_kodi text,
				fb_plex text,
				fb_emby text,
				fb_pushover text,
				fb_discord text,
				fb_gmail text,
				fb_mail text,
				fb_mailto text,
				fb_report_error bool DEFAULT true,
				fb_store_report bool DEFAULT true,
				fb_extract_folder text,
				fb_exec text,
				fb_ignore text,
				fb_artwork bool DEFAULT true,
				fb_skip_extract bool DEFAULT true,
				fb_delete_after_extract bool DEFAULT true,
				fb_clean bool DEFAULT true,
				fb_unsorted bool DEFAULT true,
				fb_home_dir text,
				fb_extras bool DEFAULT true,
				fb_deployment_name text,
				fb_namespace text,
				CONSTRAINT filebot_config_pk PRIMARY KEY (filebot_configuration_id),
				CONSTRAINT filebot_config_one_row CHECK (filebot_configuration_id)
			
			);
		`)
		return err
	}

	down := func(db orm.DB) error {
		_, err := db.Exec("DROP TABLE filebot_configuration")
		return err
	}

	opts := migrations.MigrationOptions{}

	migrations.Register("20191208092438_create_filebot_configuration_table", up, down, opts)
}
