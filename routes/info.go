package routes

import (
	"net/http"

	"github.com/lsnow99/conductorr/csl"
)

func GetReleaseProfileCfg(w http.ResponseWriter, r *http.Request) {
	response := map[string][]string{
		"rip_types":        csl.RipTypes,
		"quality_types":    csl.QualityTypes,
		"resolution_types": csl.ResolutionTypes,
	}
	Respond(w, nil, response, true)
}
