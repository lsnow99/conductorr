package routes

import (
	"encoding/json"
	"log"
	"net/http"
)

type response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Msg     string      `json:"msg"`
}

func Respond(w http.ResponseWriter, err error, data interface{}) {
	r := response{}
	if err != nil {
		r.Success = false
		r.Msg = err.Error()
	} else {
		r.Success = true
	}
	r.Data = data

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(r); err != nil {
		log.Println(err)
	}
}
