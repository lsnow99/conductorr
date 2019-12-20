package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/go-pg/pg/v9"
	"github.com/gorilla/mux"
	"github.com/lsnow2017/conductorr/internal/constants"
	"github.com/lsnow2017/conductorr/internal/schema"

	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

// HomeHandler api welcome message
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Conductorr Server API!\n")
}

// isFirstRun return true if the reset_account variable is set or db empty
func isFirstRun() bool {
	sysConfig := &schema.SystemConfiguration{}
	sysConfig.SystemConfigurationID = true

	// Check if there isn't already a user entry in the database
	err := db.Select(sysConfig)
	firstRun := (err == pg.ErrNoRows)
	if err != nil && err != pg.ErrNoRows {
		panic(err)
	}

	resetting := os.Getenv("RESET_ACCOUNT") == "1"

	return resetting || firstRun
}

// FirstRunHandler allow frontend to ask if it is the first run
func FirstRunHandler(w http.ResponseWriter, r *http.Request) {

	responseObj := map[string]bool{"first_run": isFirstRun()}
	responseJSON, err := json.Marshal(responseObj)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(responseJSON))
}

// SignupHandler write a new user to the database if none exists, or RESET_ACCOUNT = 1
func SignupHandler(w http.ResponseWriter, r *http.Request) {

	if !isFirstRun() {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Account already exists. Set RESET_ACCOUNT env to reset")
		return
	}

	sysConfig := &schema.SystemConfiguration{}
	creds := &schema.Credentials{}
	err := json.NewDecoder(r.Body).Decode(creds)
	if err != nil {
		// If there is something wrong with the request body, return a 400 status
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(err.Error())
		return
	}
	// Salt and hash the password using the bcrypt algorithm
	// The second argument is the cost of hashing, which we arbitrarily set as 8 (this value can be more or less, depending on the computing power you wish to utilize)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(creds.Password), 8)
	sysConfig.Password = hashedPassword
	sysConfig.Username = creds.Username

	err = db.Insert(sysConfig)
	if err != nil {
		panic(err)
	}

	ss := GenerateToken()

	fmt.Fprint(w, ss)
}

// GenerateToken generate a new JWT token
func GenerateToken() string {
	mySigningKey := []byte(os.Getenv("JWT_SIGNING_KEY"))

	// Create the Claims
	expirationTime := time.Now().Add(48 * time.Hour)
	claims := &jwt.StandardClaims{
		ExpiresAt: expirationTime.Unix(),
		Issuer:    "Conductorr server",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	if err != nil {
		panic(err)
	}
	return ss
}

// LoginHandler handle a login request
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	sysConfig := &schema.SystemConfiguration{}
	creds := &schema.Credentials{}
	err := json.NewDecoder(r.Body).Decode(creds)
	if err != nil {
		panic(err)
	}

	sysConfig.Username = creds.Username
	err = db.Model(sysConfig).Where("username = ?", creds.Username).Select()
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Username unrecognized")
		return
	}

	err = bcrypt.CompareHashAndPassword(sysConfig.Password, []byte(creds.Password))
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Incorrect password")
		return
	}

	ss := GenerateToken()

	fmt.Fprint(w, ss)
}

// JWTRefreshHandler refresh JWT if authenticated
func JWTRefreshHandler(w http.ResponseWriter, r *http.Request) {
	ss := GenerateToken()

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(ss))
}

// LinkHandler handle OnGrabbed linking events from sonarr/radarr
func LinkHandler(w http.ResponseWriter, r *http.Request) {
	lr := &schema.LinkRequest{}
	err := json.NewDecoder(r.Body).Decode(lr)
	if err != nil {
		// If there is something wrong with the request body, return a 400 status
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(err.Error())
		return
	}
	job := &schema.Jobs{}

	// if lr.DownloadClientIdentifier == "NZBGet" {
	job.NZBLinkerID = lr.NZBLinkerID
	// } else if lr.DownloadClientIdentifier == "rTorrent" {
	job.TorrentLinkerID = lr.TorrentLinkerID
	// }
	job.Status = "DOWNLOADING"
	job.GrabberInternalID = lr.ContentIdentifier
	job.ContentType = lr.ContentCategory
	job.Title = lr.ContentName
	job.ImdbID = lr.ImdbID
	job.ReleaseTitle = lr.ReleaseTitle

	err = db.Insert(job)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
}

// ImportHandler run the filebot command
func ImportHandler(w http.ResponseWriter, r *http.Request) {
	ir := &schema.ImportRequest{}
	err := json.NewDecoder(r.Body).Decode(ir)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(err.Error())
		return
	}
	job := &schema.Jobs{}
	err = db.Model(job).Where("nzb_linker_id = ?", ir.DownloadContentID).WhereOr("torrent_linker_id = ?", ir.DownloadContentID).Limit(1).Select()
	if err == pg.ErrNoRows {
		w.WriteHeader(http.StatusNotFound)
		log.Printf("Could not find a match for this job ID: (nzb: %s), (tor: %s)", job.NZBLinkerID, job.TorrentLinkerID)
		return
	} else if err != nil {
		panic(err)
	}
	job.DownloadDirectory = ir.DownloadDirectory

	fbConfig := schema.FilebotConfiguration{}
	fbConfig.FilebotConfigurationID = true
	err = db.Select(fbConfig)
	if err == pg.ErrNoRows {
		w.WriteHeader(http.StatusNotFound)
		fmt.Println("Please set filebot connection options in Conductorr settings!")
		return
	}
	job.Status = "FILEBOT"
	job.TimeFilebotStarted = time.Now()
	updateJob(job)
	filebot.RunFilebot(job.DownloadDirectory)
	newPath, _ := filebot.GetNewDirectory(job.DownloadDirectory)
	log.Printf("New path identified as: %s", newPath)

	if job.ContentType == sonarr.LoadConfiguration(false).SonarrCategory {
		sonarr.NotifyNewPath(newPath, job.GrabberInternalID)
	} else if job.ContentType == radarr.LoadConfiguration(false).RadarrCategory {
		radarr.NotifyNewPath(newPath, job.GrabberInternalID)
	}

	if job.DownloadClient == "NZBGet" {
		err = os.RemoveAll(job.DownloadDirectory)
		if err != nil {
			// util.LogAllError("Error deleting original file after Filebot copy"+
			// err.Error(), w)
		} else {
			// util.LogAllInfo("Successfully deleted: "+job.DownloadDirectory, w)
		}
	}

	id := plex.GetLibraryID(newPath)
	log.Printf("Library ID: %d", id)
	job.Status = "PLEX"
	job.TimeScanStarted = time.Now()
	updateJob(job)
	plex.ScanPlex(newPath, id)
	log.Println("Done scanning Plex")
	job.Status = "DONE"
	job.TimeScanDone = time.Now()
	updateJob(job)

	w.WriteHeader(http.StatusOK)
}

func updateJob(job *schema.Jobs) {
	err := db.Update(job)
	if err != nil {
		panic(err)
	}
}

// GetJobsHandler get all jobs
func GetJobsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pageNum, err := strconv.Atoi(vars["page"])
	if err != nil {
		panic(err)
	}

	type JobData struct {
		Jobs []schema.Jobs	`json:"data"`
		Total int			`json:"total"`
	}

	jobData := JobData{}

	baseQuery := db.Model(&jobData.Jobs).
		Column("title", "job_id", "imdb_id", "release_title", "content_type", "status").
		Where("title ILIKE ?", "%"+vars["filter"]+"%").
		Order(vars["sort_column"] + " " + vars["sort_order"])

	count, err := baseQuery.Limit(20).Offset((pageNum - 1)*20).SelectAndCount()
	if err != nil {
		panic(err)
	}
	
	jobData.Total = count

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(jobData); err != nil {
		panic(err)
	}
}

// SetConfigHandler update config in database
func SetConfigHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	switch vars["service"] {
	case "filebot":
		newConfig := &schema.FilebotConfiguration{}
		newConfig.FilebotConfigurationID = true
		err := json.NewDecoder(r.Body).Decode(newConfig)
		if err != nil {
			panic(err)
		}
		filebot.SaveConfiguration(newConfig)
		w.WriteHeader(http.StatusOK)
		break

	case "sonarr":
		newConfig := &schema.SonarrConfiguration{}
		err := json.NewDecoder(r.Body).Decode(newConfig)
		if err != nil {
			panic(err)
		}
		newConfig.SonarrURL = EndWithSlash(newConfig.SonarrURL)
		sonarr.SaveConfiguration(newConfig)
		w.WriteHeader(http.StatusOK)
		break

	case "radarr":
		newConfig := &schema.RadarrConfiguration{}
		newConfig.RadarrConfigurationID = true
		err := json.NewDecoder(r.Body).Decode(newConfig)
		if err != nil {
			panic(err)
		}
		newConfig.RadarrURL = EndWithSlash(newConfig.RadarrURL)
		radarr.SaveConfiguration(newConfig)
		w.WriteHeader(http.StatusOK)
		break

	case "plex":
		newConfig := &schema.PlexConfiguration{}
		newConfig.PlexConfigurationID = true
		err := json.NewDecoder(r.Body).Decode(newConfig)
		if err != nil {
			panic(err)
		}
		newConfig.PlexBaseURL = EndWithSlash(newConfig.PlexBaseURL)
		plex.SaveConfiguration(newConfig)
		w.WriteHeader(http.StatusOK)
		break

	default:
		w.Header().Set("Content-Type", "text")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error: %s", "Invalid service")
		return
	}
}

// EndWithSlash utility function to make sure a url/path ends in a '/'
func EndWithSlash(url string) string {
	if strings.HasSuffix(url, "/") {
		return url
	}
	return url + "/"
}

func populateConfigs(defaults []schema.Configurator, saved interface{}) []schema.Configurator {
	reflection := reflect.TypeOf(saved)
	v := reflect.ValueOf(saved)

	for i := 0; i < len(defaults); i++ {
		tagName := defaults[i].Property
		fieldNum := -1
		for j := 0; j < v.Type().NumField(); j++ {
			candidateName, _ := reflection.Field(j).Tag.Lookup("pg")
			if candidateName == tagName {
				fieldNum = j
			}
		}
		if fieldNum == -1 {
			log.Printf("Error matching config: %s", tagName)
			continue
		}
		switch defaults[i].Datatype {
		case "bool":
			defaults[i].BoolValue = reflect.ValueOf(saved).Field(fieldNum).Bool()
			break
		case "int":
			defaults[i].IntValue = reflect.ValueOf(saved).Field(fieldNum).Int()
			break
		case "string":
			defaults[i].StringValue = reflect.ValueOf(saved).Field(fieldNum).String()
			break
		}
	}
	return defaults
}

// TestConfigHandler test the connection to the service
func TestConfigHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	success := false

	switch vars["service"] {
	case "sonarr":
		config := &schema.SonarrConfiguration{}
		config.SonarrConfigurationID = true
		err := json.NewDecoder(r.Body).Decode(config)
		if err != nil {
			panic(err)
		}
		config.SonarrURL = EndWithSlash(config.SonarrURL)

		success = TestSonarrConnection(config)
	case "radarr":
		config := &schema.RadarrConfiguration{}
		config.RadarrConfigurationID = true
		err := json.NewDecoder(r.Body).Decode(config)
		if err != nil {
			panic(err)
		}
		config.RadarrURL = EndWithSlash(config.RadarrURL)

		success = TestRadarrConnection(config)
	case "plex":
		config := &schema.PlexConfiguration{}
		config.PlexConfigurationID = true
		err := json.NewDecoder(r.Body).Decode(config)
		if err != nil {
			panic(err)
		}
		config.PlexBaseURL = EndWithSlash(config.PlexBaseURL)

		success = TestPlexConnection(config)
	}

	responseObj := map[string]bool{"test_success": success}
	responseJSON, err := json.Marshal(responseObj)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	if success {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusFailedDependency)
	}
	fmt.Fprint(w, string(responseJSON))
}

// GetConfigHandler return config form data
func GetConfigHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	var configObj []schema.Configurator

	switch vars["service"] {
	case "filebot":
		configObj = constants.FilebotConfig
		filebotConfig := filebot.LoadConfiguration(false)
		configObj = populateConfigs(configObj, *filebotConfig)
		break
	case "sonarr":
		configObj = constants.SonarrConfig
		sonConfig := sonarr.LoadConfiguration(false)
		configObj = populateConfigs(configObj, *sonConfig)
		break
	case "radarr":
		configObj = constants.RadarrConfig
		radConfig := radarr.LoadConfiguration(false)
		configObj = populateConfigs(configObj, *radConfig)
		break
	case "plex":
		configObj = constants.PlexConfig
		plexConfig := plex.LoadConfiguration(false)
		configObj = populateConfigs(configObj, *plexConfig)
		break
	default:
		w.Header().Set("Content-Type", "text")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error: %s", "Invalid service")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(configObj); err != nil {
		panic(err)
	}
}
