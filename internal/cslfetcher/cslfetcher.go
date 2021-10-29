package cslfetcher

import (
	"fmt"

	"github.com/lsnow99/conductorr/csllib"
	"github.com/lsnow99/conductorr/dbstore"
)

func ConductorrFetcher(is csllib.ImportableScript, importPath string, allowInsecureRequests bool) (string, error) {
	if _, ok := is.(csllib.FileScript); ok {
		return "", fmt.Errorf("cannot resolve import from file")
	}

	if ps, ok := is.(csllib.ProfileScript); ok {
		name := ps.GetName()
		scriptType := ps.GetScriptType()
		profile, err := dbstore.GetProfileByName(name)
		if err != nil {
			return "", err
		}

		if scriptType == "filter" {
			return profile.Filter.String, nil
		} else if scriptType == "sorter" {
			return profile.Sorter.String, nil
		} else {
			return "", fmt.Errorf("unexpexcted script type: %s", scriptType)
		}
	}

	return csllib.DefaultFetcher(is, importPath, allowInsecureRequests)
}
