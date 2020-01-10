package main

import (
	"github.com/go-pg/pg/v9/orm"
	migrations "github.com/robinjoseph08/go-pg-migrations/v2"
)

func init() {
	up := func(db orm.DB) error {
		_, err := db.Exec(`
			alter table jobs modify grabber_internal_id null;
			alter table jobs modify torrent_linker_id null;
			alter table jobs modify nzb_linker_id null;
			alter table jobs modify title null;
			alter table jobs modify imdb_id null;
			alter table jobs modify release_title null;
			alter table jobs modify content_type null;
			alter table jobs modify status null;
			alter table jobs modify grabbed_quality null;
			alter table jobs modify grabbed_size null;
		`)
		return err
	}

	down := func(db orm.DB) error {
		_, err := db.Exec(`
			alter table jobs modify grabber_internal_id NOT NULL;
			alter table jobs modify torrent_linker_id NOT NULL;
			alter table jobs modify nzb_linker_id NOT NULL;
			alter table jobs modify title NOT NULL;
			alter table jobs modify imdb_id NOT NULL;
			alter table jobs modify release_title NOT NULL;
			alter table jobs modify content_type NOT NULL;
			alter table jobs modify status NOT NULL;
			alter table jobs modify grabbed_quality NOT NULL;
			alter table jobs modify grabbed_size NOT NULL;
		`)
		return err
	}

	opts := migrations.MigrationOptions{}

	migrations.Register("20200110183830_alter_jobs_nullable", up, down, opts)
}
