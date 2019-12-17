package main

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/go-pg/pg/v9"
	"github.com/lsnow2017/conductorr/internal/schema"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/remotecommand"
)

// Plex struct for communicating with sonarr
type Plex struct {
	config *schema.PlexConfiguration
}

/*
SaveConfiguration save a sonarr configuration to the database
*/
func (p Plex) SaveConfiguration(config *schema.PlexConfiguration) {
	defaultConfig := schema.PlexConfiguration{}
	defaultConfig.PlexConfigurationID = true
	config.PlexConfigurationID = true
	_, err := db.Model(&defaultConfig).SelectOrInsert()
	if err != nil {
		panic(err)
	}

	err = db.Update(config)
	if err != nil {
		panic(err)
	}
	p.config = config
}

/*
LoadConfiguration load a configuration from cache and optionally refresh cache
*/
func (p Plex) LoadConfiguration(refreshCache bool) *schema.PlexConfiguration {
	if refreshCache {
		p.config = &schema.PlexConfiguration{}
		p.config.PlexConfigurationID = true
		err := db.Select(p.config)
		if err == pg.ErrNoRows {
			db.Insert(p.config)
		} else if err != nil {
			panic(err)
		}
	}
	return p.config
}

/*
ScanPlex run the Plex scanning command on the newly downloaded media
*/
func (p Plex) ScanPlex(scanDir string, libID int) {
	config, err := rest.InClusterConfig()

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	pods, err := clientset.CoreV1().Pods(p.config.PlexNamespace).List(metav1.ListOptions{})

	var podName string

	for _, pod := range pods.Items {
		if strings.HasPrefix(pod.GetName(), p.config.PlexDeploymentName) {
			podName = pod.GetName()
		}
	}

	cmd := []string{
		"/bin/bash",
		"-c",
		"\"" +
			"LD_LIBRARY_PATH=/usr/lib/plexmediaserver" +
			"/usr/lib/plexmediaserver/Plex Media Scanner " +
			"--scan " +
			"--refresh " +
			"--force " +
			"--section " +
			strconv.Itoa(libID) +
			"--directory " +
			scanDir +
			"\"",
	}

	req := clientset.CoreV1().RESTClient().Post().Resource("pods").Name(podName).Namespace(p.config.PlexNamespace).SubResource("exec")
	option := &v1.PodExecOptions{
		Command: cmd,
		TTY:     true,
		Stderr:  true,
		Stdout:  true,
		Stdin:   true,
	}
	req.VersionedParams(
		option,
		scheme.ParameterCodec,
	)
	exec, err := remotecommand.NewSPDYExecutor(config, "POST", req.URL())
	if err != nil {
		panic(err)
	}
	err = exec.Stream(remotecommand.StreamOptions{
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	})
	if err != nil {
		panic(err)
	}

}

/*
GetLibraryID given a path to some media, get the corresponding library
*/
func (p Plex) GetLibraryID(path string) int {
	// util.LogAllInfo("Performing Plex API request to list libraries:", w)
	reqURL := p.config.PlexBaseURL + "/library/sections?X-Plex-token="
	// util.LogAllInfo(reqURL+"(auth token here)", w)
	reqURL += p.config.PlexAuthToken
	resp, err := http.Get(reqURL)
	if err != nil {
		// util.LogAllError("Received error during call: "+err.Error(), w)
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		panic(err)
	}

	var mediaContainer schema.MediaContainer
	xml.Unmarshal(body, &mediaContainer)

	for i := 0; i < len(mediaContainer.Directory); i++ {
		libPath := mediaContainer.Directory[i].Location.Path
		if strings.HasPrefix(path, libPath) {
			libID := mediaContainer.Directory[i].Location.ID
			idNum, err := strconv.Atoi(libID)
			if err != nil {
				// util.LogAllError("Error converting library ID to integer", w)
				panic(err)
			}
			return idNum
		}
	}

	// util.LogAllError("Could not find a matching library!", w)
	return -1
}
