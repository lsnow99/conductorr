package app

import (
	"fmt"

	"github.com/lsnow99/conductorr/dbstore"
)

func ScanMediaFiles(id int) error {
	media, err := dbstore.GetMediaByID(id)
	if err != nil {
		return err
	}
	if !media.ContentType.Valid {
		return fmt.Errorf("media has null content type")
	}
	_, err = dbstore.GetPath(int(media.PathID.Int32))
	if err != nil {
		return err
	}
	switch media.ContentType.String {
	case "movie":
		
	case "series":

	default:
		return fmt.Errorf("media has unrecognized content type")
	}
	return nil
}
