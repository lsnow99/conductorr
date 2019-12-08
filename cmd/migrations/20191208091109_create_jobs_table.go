package main

import (
	"github.com/go-pg/pg/v9/orm"
	migrations "github.com/robinjoseph08/go-pg-migrations/v2"
)

func init() {
	up := func(db orm.DB) error {
		_, err := db.Exec(`
			CREATE TABLE jobs (
				grabber_internal_id integer NOT NULL,
				job_id serial NOT NULL,
				torrent_linker_id text NOT NULL,
				nzb_linker_id text NOT NULL,
				time_grabbed timestamptz DEFAULT now(),
				title text NOT NULL,
				imdb_id text NOT NULL,
				release_title text NOT NULL,
				content_type text NOT NULL,
				download_client text,
				download_directory text,
				status text NOT NULL,
				filebot_logs text,
				scanner_logs text,
				grabbed_quality text NOT NULL,
				grabbed_size text NOT NULL,
				time_filebot_started timestamptz,
				time_filebot_done timestamptz,
				time_scan_started timestamptz,
				time_scan_done timestamptz,
				CONSTRAINT jobs_pk PRIMARY KEY (job_id)

			);
		`)
		return err
	}

	down := func(db orm.DB) error {
		_, err := db.Exec("DROP TABLE jobs")
		return err
	}

	opts := migrations.MigrationOptions{}

	migrations.Register("20191208091109_create_jobs_table", up, down, opts)
}
