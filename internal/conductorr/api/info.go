package api

import (
	"net/http"

	"github.com/lsnow99/conductorr/pkg/constant"
)

func GetReleaseProfileCfg(w http.ResponseWriter, r *http.Request) {
	response := map[string]interface{}{
		"rip_types":        constant.RipTypes,
		"resolution_types": constant.ResolutionTypes,
	}
	Respond(w, r, nil, response, true)
}

func CheckAuth(w http.ResponseWriter, r *http.Request) {
	tok, err := r.Cookie(UserAuthKey)
	if err != nil {
		Respond(w, r, nil, false, false)
		return
	}
	Respond(w, r, nil, checkToken(tok.Value), false)
}