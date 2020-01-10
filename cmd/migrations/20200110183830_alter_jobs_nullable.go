package main

import (
	"github.com/go-pg/pg/v9/orm"
	migrations "github.com/robinjoseph08/go-pg-migrations/v2"
)

func init() {
	up := func(db orm.DB) error {
		_, err := db.Exec(`
			alter table jobs alter column grabber_internal_id drop NOT NULL;
			alter table jobs alter column torrent_linker_id drop NOT NULL;
			alter table jobs alter column nzb_linker_id drop NOT NULL;
			alter table jobs alter column title drop NOT NULL;
			alter table jobs alter column imdb_id drop NOT NULL;
			alter table jobs alter column release_title drop NOT NULL;
			alter table jobs alter column content_type drop NOT NULL;
			alter table jobs alter column status drop NOT NULL;
			alter table jobs alter column grabbed_quality drop NOT NULL;
			alter table jobs alter column grabbed_size drop NOT NULL;
		`)
		return err
	}

	down := func(db orm.DB) error {
		_, err := db.Exec(`
			alter table jobs alter column grabber_internal_id SET NOT NULL;
			alter table jobs alter column torrent_linker_id SET NOT NULL;
			alter table jobs alter column nzb_linker_id SET NOT NULL;
			alter table jobs alter column title SET NOT NULL;
			alter table jobs alter column imdb_id SET NOT NULL;
			alter table jobs alter column release_title SET NOT NULL;
			alter table jobs alter column content_type SET NOT NULL;
			alter table jobs alter column status SET NOT NULL;
			alter table jobs alter column grabbed_quality SET NOT NULL;
			alter table jobs alter column grabbed_size SET NOT NULL;
		`)
		return err
	}

	opts := migrations.MigrationOptions{}

	migrations.Register("20200110183830_alter_jobs_nullable", up, down, opts)
}
