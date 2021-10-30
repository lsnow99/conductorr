package api

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/lsnow99/conductorr/csllib"
	"github.com/lsnow99/conductorr/internal/cslfetcher"
	"github.com/lsnow99/conductorr/settings"
)

func LoadCSL(w http.ResponseWriter, r *http.Request) {
	allowedEncodings := strings.Split(r.Header.Get("Accept-Encoding"), ",")
	filename := "dist/csl.wasm"
	for _, encoding := range allowedEncodings {
		if strings.TrimSpace(encoding) == "br" {
			w.Header().Add("Content-Encoding", "br")
			filename = "dist/csl.wasm.br"
		}
	}
	w.Header().Add("Content-Type", "application/wasm")
	if settings.DebugMode {
		file, err := os.OpenFile(filename, os.O_RDONLY, os.ModePerm)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusOK)
		io.Copy(w, file)
	}
}

func FetchScript(w http.ResponseWriter, r *http.Request) {
	importPath := r.URL.Query().Get("importPath")
	if importPath == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	cslpm := csllib.NewCSLPackageManager(cslfetcher.ConductorrFetcher, settings.AllowInsecureRequests)
	script, err := cslpm.Resolve(importPath)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte(script))
}