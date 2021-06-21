package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/lsnow99/conductorr/dbstore"
)

type NewProfile struct {
	Name string `json:"name"`
}

type Profile struct {
	ID     int    `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	Filter string `json:"filter,omitempty"`
	Sorter string `json:"sorter,omitempty"`
}

func GetProfiles(w http.ResponseWriter, r *http.Request) {
	dbProfiles, err := dbstore.GetProfiles()
	if err != nil {
		Respond(w, r.Host, err, nil, true)
		return
	}

	profiles := make([]Profile, len(dbProfiles))
	for i, dbProfile := range dbProfiles {
		profile := Profile{}
		profile.ID = dbProfile.ID
		profile.Name = dbProfile.Name.String
		profile.Filter = dbProfile.Filter.String
		profile.Sorter = dbProfile.Sorter.String
		profiles[i] = profile
	}

	Respond(w, r.Host, nil, profiles, true)
}

func CreateProfile(w http.ResponseWriter, r *http.Request) {
	profile := NewProfile{}
	if err := json.NewDecoder(r.Body).Decode(&profile); err != nil {
		Respond(w, r.Host, err, nil, true)
		return
	}

	err := dbstore.CreateProfile(profile.Name)
	Respond(w, r.Host, err, nil, true)
}

func UpdateProfile(w http.ResponseWriter, r *http.Request) {
	profile := Profile{}
	if err := json.NewDecoder(r.Body).Decode(&profile); err != nil {
		Respond(w, r.Host, err, nil, true)
		return
	}
	idStr := mux.Vars(r)["id"]
	idInt, err := strconv.Atoi(idStr)
	if err != nil {
		Respond(w, r.Host, err, nil, true)
		return
	}

	err = dbstore.UpdateProfile(idInt, profile.Name, profile.Filter, profile.Sorter)
	Respond(w, r.Host, err, nil, true)
}

func DeleteProfile(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	idInt, err := strconv.Atoi(idStr)
	if err != nil {
		Respond(w, r.Host, err, nil, true)
		return
	}

	err = dbstore.DeleteProfile(idInt)
	Respond(w, r.Host, err, nil, true)
}