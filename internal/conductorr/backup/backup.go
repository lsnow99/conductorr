package backup

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/lsnow99/conductorr/internal/conductorr/dbstore"
	"github.com/lsnow99/conductorr/internal/conductorr/integration"
	"github.com/lsnow99/conductorr/internal/conductorr/settings"
	"github.com/mholt/archiver/v3"
)

type BackupData struct {
	Medias       []*dbstore.Media
	Paths        []*dbstore.Path
	Profiles     []*dbstore.Profile
	MediaServers []*dbstore.MediaServer
	Downloads    []*dbstore.Download
	Downloaders  []*dbstore.Downloader
	Indexers     []*dbstore.Indexer
	Version      string
}

func CreateBackup(ctx context.Context) (id int, err error) {
	var dir, tmpDir string
	tmpDir, err = integration.MkdirTemp("conductorr_backup")
	if err != nil {
		return
	}
	dir, err = ioutil.TempDir(tmpDir, "backup")
	if err != nil {
		return
	}

	defer func() {
		if err != nil {
			os.RemoveAll(dir)
		}
	}()

	// Grab all our persistent data
	var tx *sql.Tx
	tx, err = dbstore.NewTx(ctx, sql.TxOptions{
		ReadOnly: true,
	})

	if err != nil {
		return
	}

	committed := false

	defer func() {
		if err != nil && !committed {
			tx.Rollback()
		}
	}()

	var medias []*dbstore.Media
	medias, err = dbstore.DumpMedias(tx)
	if err != nil {
		return
	}

	var downloads []*dbstore.Download
	downloads, err = dbstore.DumpDownloads(tx)
	if err != nil {
		return
	}

	var mediaServers []*dbstore.MediaServer
	mediaServers, err = dbstore.DumpMediaServers(tx)
	if err != nil {
		return
	}

	var downloaders []*dbstore.Downloader
	downloaders, err = dbstore.DumpDownloaders(tx)
	if err != nil {
		return
	}

	var paths []*dbstore.Path
	paths, err = dbstore.DumpPaths(tx)
	if err != nil {
		return
	}

	var profiles []*dbstore.Profile
	profiles, err = dbstore.DumpProfiles(tx)
	if err != nil {
		return
	}

	var indexers []*dbstore.Indexer
	indexers, err = dbstore.DumpIndexers(tx)
	if err != nil {
		return
	}

	err = tx.Commit()
	if err != nil {
		return
	}
	committed = true

	// Remove the posters; they will be stored as files in the zip to make the JSON marshalling a little less unwieldy
	posterDir := filepath.Join(dir, "posters")
	err = os.Mkdir(posterDir, os.ModePerm)
	if err != nil {
		return
	}

	err = addMediaPosters(medias, posterDir)
	if err != nil {
		return
	}

	// Build the backup data object
	bd := BackupData{
		Medias:       medias,
		Paths:        paths,
		Profiles:     profiles,
		MediaServers: mediaServers,
		Downloads:    downloads,
		Downloaders:  downloaders,
		Indexers:     indexers,
		Version:      settings.Version,
	}

	// Write the backup data object to a file
	var raw []byte
	raw, err = json.Marshal(bd)
	if err != nil {
		return
	}

	dataFile := filepath.Join(dir, "data.json")
	err = ioutil.WriteFile(dataFile, raw, fs.ModePerm)
	if err != nil {
		return
	}

	var outDir string
	outDir, err = ioutil.TempDir(tmpDir, "backup_out")
	if err != nil {
		return
	}

	defer func() {
		if err != nil {
			os.RemoveAll(outDir)
		}
	}()

	outFile := filepath.Join(outDir, "backup.zip")
	err = archiver.Archive([]string{posterDir, dataFile}, outFile)

	id = addBackup(outFile)

	return
}

func addMediaPosters(medias []*dbstore.Media, dir string) error {
	for _, media := range medias {
		name := fmt.Sprintf("%d.jpeg", media.ID)
		f := filepath.Join(dir, name)
		if media.Poster == nil {
			continue
		}
		if err := ioutil.WriteFile(f, *media.Poster, os.ModePerm); err != nil {
			return err
		}
		// Clear the poster field so it isn't written to the JSON
		media.Poster = nil
	}
	return nil
}
