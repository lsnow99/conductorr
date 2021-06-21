package routes

import (
	"net/http"

	"github.com/lsnow99/conductorr/constant"
)

func GetReleaseProfileCfg(w http.ResponseWriter, r *http.Request) {
	response := map[string][]string{
		"rip_types":        constant.RipTypes,
		"quality_types":    constant.QualityTypes,
		"resolution_types": constant.ResolutionTypes,
	}
	Respond(w, r.Host, nil, response, true)
}

func CheckAuth(w http.ResponseWriter, r *http.Request) {
	tok, err := r.Cookie(UserAuthKey)
	if err != nil {
		Respond(w, r.Host, nil, false, false)
		return
	}
	Respond(w, r.Host, nil, checkToken(tok.Value), false)
}