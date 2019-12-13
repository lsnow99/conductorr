package main

import (
	"bytes"
	"encoding/json"
	"github.com/lsnow2017/conductorr/internal/schema"
	"github.com/go-pg/pg/v9"
	"net/http"
	"strconv"
)

// Sonarr struct for communicating with sonarr
type Sonarr struct {
	config *schema.SonarrConfiguration
}

/*
SaveConfiguration save a sonarr configuration to the database
*/
func (s Sonarr) SaveConfiguration(config *schema.SonarrConfiguration) {
	defaultConfig := schema.SonarrConfiguration{}
	defaultConfig.SonarrConfigurationID = true
	config.SonarrConfigurationID = true
	_, err := db.Model(&defaultConfig).SelectOrInsert()
	if err != nil {
		panic(err)
	}

	err = db.Update(config)
	if err != nil {
		panic(err)
	}
	s.config = config
}

/*
LoadConfiguration load a configuration from cache and optionally refresh cache
*/
func (s Sonarr) LoadConfiguration(refreshCache bool) *schema.SonarrConfiguration {
	if refreshCache {
		s.config.SonarrConfigurationID = true
		err := db.Select(s.config)
		if err == pg.ErrNoRows {
			db.Insert(s.config)
		} else if err != nil {
			panic(err)
		}
	}
	return s.config
}

/*
TestSonarrConnection test the connection given a Sonarr configuration
*/
func TestSonarrConnection(config *schema.SonarrConfiguration) bool {
	resp, err := http.Get(config.SonarrURL + "system/status?apikey=" + config.SonarrAPIKey)
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
NotifyNewPath update Sonarr with the new path for the content
*/
func (s Sonarr) NotifyNewPath(newPath string, contentID int64) {
	/*
	 * API Request #1 - Get existing data
	 */
	//// util.LogAllInfo("Performing API request #1 (get old data):", w)
	req1URL := s.config.SonarrURL + "series" + "/" + strconv.Itoa(int(contentID)) + "?apikey="
	//// util.LogAllInfo(req1URL+"(api key hidden)", w)
	req1URL += s.config.SonarrAPIKey
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
	req2URL := s.config.SonarrURL + "series" + "/" + strconv.Itoa(int(contentID)) + "?apikey="
	//util.LogAllInfo(req2URL, w)
	req2URL += s.config.SonarrAPIKey
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
	req3URL := s.config.SonarrURL + "command?apikey=(api key here)&name=refreshSeries" +
		"&" + "seriesId" + "=" + strconv.Itoa(int(contentID))
	//util.LogAllInfo(req3URL, w)
	req3URL = s.config.SonarrURL + "command?apikey=" + s.config.SonarrURL + "&name=refreshSeries" +
		"&" + "seriesId" + "=" + strconv.Itoa(int(contentID))
	reqJSON := make(map[string]interface{})
	reqJSON["name"] = "refreshSeries"
	reqJSON["seriesId"] = int(contentID)
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
