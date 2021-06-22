package routes

import (
	"io"
	"net/http"
	"os"
	"strings"

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
