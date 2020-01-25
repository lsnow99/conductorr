package main

import (
	"bytes"
	"encoding/xml"
	"log"
	"os"
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

// Filebot struct to communicate with Filebot
type Filebot struct {
	config *schema.FilebotConfiguration
}

/*
RunFilebot attempts to exec into the filebot pod and run the filebot
*/
func (f *Filebot) RunFilebot(DownloadDirectory string) string {

	config, err := rest.InClusterConfig()

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	pods, err := clientset.CoreV1().Pods(f.config.FbNamespace).List(metav1.ListOptions{})

	if err != nil {
		panic(err)
	}

	var podName string

	for _, pod := range pods.Items {
		log.Printf("Found pod: %s", pod.GetName())
		if strings.HasPrefix(pod.GetName(), f.config.FbDeploymentName) {
			podName = pod.GetName()
		}
	}

	cmd := []string{
		"/bin/sh",
		"-c",
		strings.Join([]string{
			"FILEBOT_OPTS=-Dapplication.dir=/valinor/plex",
			"/opt/filebot/filebot",
			"-script", "fn:amc",
			"--output", f.config.FbOutputDir,
			"--action", f.config.FbAction,
			"--conflict", "override", "-non-strict",
			"--log-file", f.config.FbAmcLog,
			"--def",
			boolArg("unsorted", f.config.FbUnsorted),
			boolArg("music", false),
			stringArg("subtitles", f.config.FbSubtitlesLocale),
			boolArg("artwork", f.config.FbArtwork),
			boolArg("extras", f.config.FbExtras),
			stringArg("kodi", f.config.FbKodi),
			stringArg("plex", f.config.FbPlex),
			stringArg("emby", f.config.FbEmby),
			stringArg("emby", f.config.FbEmby),
			stringArg("pushover", f.config.FbPushover),
			stringArg("pushbullet", f.config.FbPushbullet),
			stringArg("discord", f.config.FbDiscord),
			stringArg("gmail", f.config.FbGmail),
			stringArg("mail", f.config.FbMail),
			stringArg("mailto", f.config.FbMailto),
			boolArg("reportError", f.config.FbReportError),
			boolArg("storeReport", f.config.FbStoreReport),
			stringArg("extractFolder", f.config.FbExtractFolder),
			boolArg("skipExtract", f.config.FbSkipExtract),
			boolArg("deleteAfterExtract", f.config.FbDeleteAfterExtract),
			boolArg("clean", f.config.FbClean),
			stringArg("exec", f.config.FbExec),
			stringArg("ignore", f.config.FbIgnore),
			// "minFileSize=" + string(f.config.FbMinFileSize),
			// "minLengthMS=" + string(f.config.FbMinLengthMs),
			// "excludeList=" + f.config.FbExcludeList,
			`"` + DownloadDirectory + `"`,
		}, " "),
	}
	log.Printf("Running filebot with cmd: %s", strings.Join(cmd, " "))

	req := clientset.CoreV1().RESTClient().Post().Resource("pods").Name(podName).Namespace(f.config.FbNamespace).SubResource("exec")
	option := &v1.PodExecOptions{
		Command: cmd,
		TTY:     false,
		Stderr:  true,
		Stdout:  true,
		Stdin:   false,
	}
	req.VersionedParams(
		option,
		scheme.ParameterCodec,
	)
	exec, err := remotecommand.NewSPDYExecutor(config, "POST", req.URL())
	if err != nil {
		panic(err)
	}
	stdBuf := &bytes.Buffer{}
	errBuf := &bytes.Buffer{}
	err = exec.Stream(remotecommand.StreamOptions{
		Stdout: stdBuf,
		Stderr: errBuf,
	})
	output := stdBuf.String() + errBuf.String()
	if err != nil {
		output += err.Error()
	}
	return output
}

/*
SaveConfiguration save a filebot configuration to the database
*/
func (f *Filebot) SaveConfiguration(config *schema.FilebotConfiguration) {
	defaultConfig := schema.FilebotConfiguration{}
	defaultConfig.FilebotConfigurationID = true
	config.FilebotConfigurationID = true
	_, err := db.Model(&defaultConfig).SelectOrInsert()
	if err != nil {
		panic(err)
	}

	err = db.Update(config)
	if err != nil {
		panic(err)
	}
	f.config = config
}

/*
LoadConfiguration load a configuration from cache and optionally refresh cache
*/
func (f *Filebot) LoadConfiguration(refreshCache bool) *schema.FilebotConfiguration {
	if refreshCache {
		f.config = &schema.FilebotConfiguration{}
		f.config.FilebotConfigurationID = true
		err := db.Select(f.config)
		if err == pg.ErrNoRows {
			db.Insert(f.config)
		} else if err != nil {
			panic(err)
		}
	}
	return f.config
}

/*
GetNewDirectory searches Filebot's history.xml file to find the new location for
the media it processed by matching the original location to the record
*/
// func (f *Filebot) GetNewDirectory(downloadDir string) (string, schema.Sequence) {
// 	downloadDir = strings.TrimRight(downloadDir, "/")
// 	oRoot := ""
// 	var ourSeq schema.Sequence

// 	historyFile, err := os.Open(EndWithSlash(os.Getenv("FB_DIRECTORY")) + "history.xml")
// 	// if os.Open returns an error then handle it
// 	if err != nil {
// 		panic(err)
// 	}

// 	// defer the closing of our historyFile so that we can parse it later on
// 	defer historyFile.Close()

// 	byteValue, err := ioutil.ReadAll(historyFile)

// 	if err != nil {
// 		panic(err)
// 	}

// 	// we initialize our History array
// 	var history schema.History
// 	// we unmarshal our byteArray which contains our
// 	// historyFile content into 'history' which we defined above
// 	xml.Unmarshal(byteValue, &history)

// 	for i := 0; i < len(history.Sequences); i++ {
// 		seq := history.Sequences[i]
// 		for j := 0; j < len(seq.Renames); j++ {
// 			ren := seq.Renames[j]
// 			switch getPathInfo(downloadDir) {
// 			case FILE:
// 				if ren.Dir+"/"+ren.From == downloadDir {
// 					ourSeq = history.Sequences[i]
// 					oRoot = upDir(upDir(ren.To))
// 				}
// 			case DIRECTORY:
// 				if strings.HasPrefix(ren.Dir, downloadDir) {
// 					ourSeq = history.Sequences[i]
// 					oRoot = upDir(upDir(ren.To))
// 				}
// 			default:
// 				return "no path found", ourSeq
// 			}
// 		}
// 	}

// 	return oRoot, ourSeq
// }

/*
GetNewDirectory searches Filebot's history.xml file to find the new location for
the media it processed by matching the original location to the record
*/
func (f *Filebot) GetNewDirectory(downloadDir string) (string, schema.Sequence) {
	downloadDir = strings.TrimRight(downloadDir, "/")
	oRoot := ""
	var ourSeq schema.Sequence

	config, err := rest.InClusterConfig()

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	pods, err := clientset.CoreV1().Pods(f.config.FbNamespace).List(metav1.ListOptions{})

	if err != nil {
		panic(err)
	}

	var podName string

	for _, pod := range pods.Items {
		log.Printf("Found pod: %s", pod.GetName())
		if strings.HasPrefix(pod.GetName(), f.config.FbDeploymentName) {
			podName = pod.GetName()
		}
	}

	cmd := []string{
		"/bin/cat",
		os.Getenv("FB_DIRECTORY") + "/history.xml",
	}

	req := clientset.CoreV1().RESTClient().Post().Resource("pods").Name(podName).Namespace(f.config.FbNamespace).SubResource("exec")
	option := &v1.PodExecOptions{
		Command: cmd,
		TTY:     false,
		Stderr:  true,
		Stdout:  true,
		Stdin:   false,
	}
	req.VersionedParams(
		option,
		scheme.ParameterCodec,
	)
	exec, err := remotecommand.NewSPDYExecutor(config, "POST", req.URL())
	if err != nil {
		panic(err)
	}
	stdBuf := &bytes.Buffer{}
	errBuf := &bytes.Buffer{}
	err = exec.Stream(remotecommand.StreamOptions{
		Stdout: stdBuf,
		Stderr: errBuf,
	})
	output := stdBuf.String() + errBuf.String()
	if err != nil {
		panic(err)
	}

	// we initialize our History array
	var history schema.History
	// we unmarshal our byteArray which contains our
	// historyFile content into 'history' which we defined above
	xml.Unmarshal([]byte(output), &history)

	for i := 0; i < len(history.Sequences); i++ {
		seq := history.Sequences[i]
		for j := 0; j < len(seq.Renames); j++ {
			ren := seq.Renames[j]
			switch f.execGetPathInfo(downloadDir) {
			case FILE:
				if ren.Dir+"/"+ren.From == downloadDir {
					ourSeq = history.Sequences[i]
					oRoot = upDir(upDir(ren.To))
				}
			case DIRECTORY:
				if strings.HasPrefix(ren.Dir, downloadDir) {
					ourSeq = history.Sequences[i]
					oRoot = upDir(upDir(ren.To))
				}
			default:
				return "no path found", ourSeq
			}
		}
	}

	return oRoot, ourSeq
}

func upDir(path string) string {
	path = strings.TrimRight(path, "/")
	upDir := path[:strings.LastIndex(path, "/")]
	return upDir
}

// PathInfo enum for filsystem object type
type PathInfo int

const (
	// FILE type file
	FILE PathInfo = iota
	// DIRECTORY type directory
	DIRECTORY
	// NEITHER type neither
	NEITHER
)

// getPathInfo utility to get the type of a filesystem object
func getPathInfo(path string) PathInfo {
	fi, err := os.Stat(path)
	switch {
	case err != nil:
		// handle the error and return
		log.Printf("Broken Path: %s", path)
		panic(err)
		// return NEITHER
	case fi.IsDir():
		// it's a directory
		return DIRECTORY
	default:
		// it's not a directory
		return FILE
	}
}

// execGetPathInfo get the file/folder info 
func (f *Filebot) execGetPathInfo(path string) PathInfo {
	config, err := rest.InClusterConfig()

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	pods, err := clientset.CoreV1().Pods(f.config.FbNamespace).List(metav1.ListOptions{})

	if err != nil {
		panic(err)
	}

	var podName string

	for _, pod := range pods.Items {
		log.Printf("Found pod: %s", pod.GetName())
		if strings.HasPrefix(pod.GetName(), f.config.FbDeploymentName) {
			podName = pod.GetName()
		}
	}

	cmd := []string{
		"/bin/sh",
		"-c",
		`if [[ -d ` + path + ` ]]; then
		echo "directory"
	elif [[ -f ` + path + ` ]]; then
		echo "file"
	else
		echo "invalid"
	fi`,
	}

	req := clientset.CoreV1().RESTClient().Post().Resource("pods").Name(podName).Namespace(f.config.FbNamespace).SubResource("exec")
	option := &v1.PodExecOptions{
		Command: cmd,
		TTY:     false,
		Stderr:  true,
		Stdout:  true,
		Stdin:   false,
	}
	req.VersionedParams(
		option,
		scheme.ParameterCodec,
	)
	exec, err := remotecommand.NewSPDYExecutor(config, "POST", req.URL())
	if err != nil {
		panic(err)
	}
	stdBuf := &bytes.Buffer{}
	errBuf := &bytes.Buffer{}
	err = exec.Stream(remotecommand.StreamOptions{
		Stdout: stdBuf,
		Stderr: errBuf,
	})
	output := stdBuf.String() + errBuf.String()
	log.Printf("is a: ;%s;", output)
	if err != nil {
		panic(err)
	}

	switch output {
	case "directory":
		return DIRECTORY
	case "file":
		return FILE
	default:
		return NEITHER
	}
}

// boolArg quick utility to turn a bool into a "y" or "n" string
func boolArg(option string, test bool) string {
	ans := option + "="
	if test {
		ans += "y"
	} else {
		ans += "n"
	}
	return ans
}

func stringArg(option string, value string) string {
	str := ""
	if len(value) > 0 {
		str = option + "=" + value
	}
	return str
}
