package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-pg/pg/v9"
	"github.com/lsnow2017/conductorr/internal/schema"
)

// Radarr struct for communicating with radarr
type Radarr struct {
	config *schema.RadarrConfiguration
}

/*
SaveConfiguration save a radarr configuration to the database
*/
func (r *Radarr) SaveConfiguration(config *schema.RadarrConfiguration) {
	defaultConfig := schema.RadarrConfiguration{}
	defaultConfig.RadarrConfigurationID = true
	config.RadarrConfigurationID = true
	_, err := db.Model(&defaultConfig).SelectOrInsert()
	if err != nil {
		panic(err)
	}

	err = db.Update(config)
	if err != nil {
		panic(err)
	}
	r.config = config
}

/*
LoadConfiguration load a configuration from cache and optionally refresh cache
*/
func (r *Radarr) LoadConfiguration(refreshCache bool) *schema.RadarrConfiguration {
	if refreshCache {
		r.config = &schema.RadarrConfiguration{}
		r.config.RadarrConfigurationID = true
		err := db.Select(r.config)
		if err == pg.ErrNoRows {
			db.Insert(r.config)
		} else if err != nil {
			panic(err)
		}
	}
	return r.config
}

/*
TestRadarrConnection test the connection given a Radarr configuration
*/
func TestRadarrConnection(config *schema.RadarrConfiguration) bool {
	resp, err := http.Get(config.RadarrURL + "system/status?apikey=" + config.RadarrAPIKey)
	if err != nil {
		return false
	}
	decoder := json.NewDecoder(resp.Body)
	var respJSON map[string]interface{}
	err = decoder.Decode(&respJSON)
	if err != nil {
		return false
	}

	return respJSON["version"] != nil
}

/*
NotifyNewPath update Radarr with the new path for the content
*/
func (r *Radarr) NotifyNewPath(newPath string, contentID int64) {
	/*
	 * API Request #1 - Get existing data
	 */
	// util.LogAllInfo("Performing API request #1 (get old data):", w)
	req1URL := r.config.RadarrURL + "movie" + "/" + strconv.Itoa(int(contentID)) + "?apikey="
	// util.LogAllInfo(req1URL+"(api key hidden)", w)
	req1URL += r.config.RadarrAPIKey
	resp1, err := http.Get(req1URL)
	if err != nil {
		//// util.LogAllError("Received error during call: "+err.Error(), w)
		panic(err)
	}
	defer resp1.Body.Close()
	if resp1.StatusCode < 200 || resp1.StatusCode >= 300 {
		//util.LogAllError("Non-2xx http response code returned from API call. "+
		// "Printing body & aborting API calls 2 and 3:", w)
		//body, err := ioutil.ReadAll(resp1.Body)
		if err != nil {
			//util.LogAllError("Received error while reading body: "+
			// err.Error(), w)
			panic(err)
		}
		//util.LogAllError(string(body), w)
		return
	}
	decoder := json.NewDecoder(resp1.Body)
	var respJSON map[string]interface{}
	err = decoder.Decode(&respJSON)
	if err != nil {
		//util.LogAllError("Error parsing response json: "+err.Error(), w)
		panic(err)
	}
	respJSON["path"] = newPath
	resp1.Body.Close()

	/*
	 * API Request #2 - Post update with new data
	 */

	//util.LogAllInfo("Performing API request #2 (post new data):", w)
	req2URL := r.config.RadarrURL + "movie" + "/" + strconv.Itoa(int(contentID)) + "?apikey="
	//util.LogAllInfo(req2URL, w)
	req2URL += r.config.RadarrAPIKey
	respJSONStr, err := json.Marshal(respJSON)
	if err != nil {
		//util.LogAllError("Error marshaling JSON: "+err.Error(), w)
		panic(err)
	}
	client := &http.Client{}
	request2, err := http.NewRequest("PUT", req2URL,
		bytes.NewBuffer(respJSONStr))
	if err != nil {
		//util.LogAllError("Error building request2: "+err.Error(), w)
		panic(err)
	}
	request2.ContentLength = int64(len(respJSONStr))
	resp2, err := client.Do(request2)
	if err != nil {
		//util.LogAllError("Received error during call: "+err.Error(), w)
		panic(err)
	}
	defer resp2.Body.Close()
	if resp2.StatusCode < 200 || resp2.StatusCode >= 300 {
		//util.LogAllError("Non-2xx http response code returned from API call. "+
		// "Printing body & aborting API calls 3:", w)
		//body, err := ioutil.ReadAll(resp2.Body)
		if err != nil {
			//util.LogAllError("Received error while reading body: "+
			// err.Error(), w)
			panic(err)
		}
		//util.LogAllError(string(body), w)
		return
	}
	resp2.Body.Close()

	/*
	 * API Request #3 - Trigger rescan
	 */
	//util.LogAllInfo("Performing API request #3 (trigger rescan):", w)
	req3URL := r.config.RadarrURL + "command?apikey=(api key here)&name=refreshMovie" +
		"&" + "movieId" + "=" + strconv.Itoa(int(contentID))
	//util.LogAllInfo(req3URL, w)
	req3URL = r.config.RadarrURL + "command?apikey=" + r.config.RadarrAPIKey + "&name=refreshMovie" +
		"&" + "movieId" + "=" + strconv.Itoa(int(contentID))
	reqJSON := make(map[string]interface{})
	reqJSON["name"] = "refreshMovie"
	reqJSON["movieId"] = int(contentID)
	reqJSONStr, err := json.Marshal(reqJSON)
	if err != nil {
		//util.LogAllError("Error marshaling json: "+err.Error(), w)
	}
	resp3, err := http.Post(req3URL, "application/json",
		bytes.NewBuffer(reqJSONStr))

	if err != nil {
		//util.LogAllError("Received error during call: "+err.Error(), w)
		panic(err)
	}
	defer resp3.Body.Close()
	if resp3.StatusCode < 200 || resp3.StatusCode >= 300 {
		//util.LogAllError("Non-2xx http response code returned from API call. "+
		// "Printing body & aborting rescan attempt:", w)
		//body, err := ioutil.ReadAll(resp3.Body)
		if err != nil {
			//util.LogAllError("Received error while reading body: "+
			// err.Error(), w)
			panic(err)
		}
		//util.LogAllError(string(body), w)
		return
	}
	resp3.Body.Close()
}
