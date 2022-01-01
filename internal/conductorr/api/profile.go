package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/lsnow99/conductorr/internal/conductorr/dbstore"
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

func NewProfileFromDB(dbProfile *dbstore.Profile) Profile {
	profile := Profile{}
	profile.ID = dbProfile.ID
	profile.Name = dbProfile.Name.String
	profile.Filter = dbProfile.Filter.String
	profile.Sorter = dbProfile.Sorter.String
	return profile
}

func GetProfiles(w http.ResponseWriter, r *http.Request) {
	dbProfiles, err := dbstore.GetProfiles()
	if err != nil {
		Respond(w, r, err, nil, true)
		return
	}

	profiles := make([]Profile, len(dbProfiles))
	for i, dbProfile := range dbProfiles {
		profiles[i] = NewProfileFromDB(dbProfile)
	}

	Respond(w, r, nil, profiles, true)
}

func GetProfile(w http.ResponseWriter, r *http.Request) {
	profileIDStr := mux.Vars(r)["id"]
	profileID, err := strconv.Atoi(profileIDStr)
	if err != nil {
		Respond(w, r, err, nil, true)
		return
	}
	profile, err := dbstore.GetProfileByID(profileID)
	Respond(w, r, err, NewProfileFromDB(&profile), true)
}

func CreateProfile(w http.ResponseWriter, r *http.Request) {
	profile := NewProfile{}
	if err := json.NewDecoder(r.Body).Decode(&profile); err != nil {
		Respond(w, r, err, nil, true)
		return
	}

	err := dbstore.CreateProfile(profile.Name)
	Respond(w, r, err, nil, true)
}

func UpdateProfile(w http.ResponseWriter, r *http.Request) {
	profile := Profile{}
	if err := json.NewDecoder(r.Body).Decode(&profile); err != nil {
		Respond(w, r, err, nil, true)
		return
	}
	idStr := mux.Vars(r)["id"]
	idInt, err := strconv.Atoi(idStr)
	if err != nil {
		Respond(w, r, err, nil, true)
		return
	}

	err = dbstore.UpdateProfile(idInt, profile.Name, profile.Filter, profile.Sorter)
	Respond(w, r, err, nil, true)
}

func DeleteProfile(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	idInt, err := strconv.Atoi(idStr)
	if err != nil {
		Respond(w, r, err, nil, true)
		return
	}

	err = dbstore.DeleteProfile(idInt)
	Respond(w, r, err, nil, true)
}

func GetRawProfileScript(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]
	scriptType := mux.Vars(r)["script_type"]

	profile, err := dbstore.GetProfileByName(name)
	if err != nil {
		Respond(w, r, err, nil, true)
		return
	}

	fmt.Println(profile)
	
	if scriptType == "filter" {
		fmt.Println("returning", profile.Filter.String)
		w.Write([]byte(profile.Filter.String))
		return
	} else if scriptType == "sorter" {
		fmt.Println("returning", profile.Sorter.String)
		w.Write([]byte(profile.Sorter.String))
		return
	} else {
		Respond(w, r, fmt.Errorf("unrecognized script type %s", scriptType), nil, true)
		return
	}
}
