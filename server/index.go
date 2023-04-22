package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os/exec"
	"path"
	"strings"
	"time"
)

func Main() {
	http.HandleFunc("/request-download", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("")
		fmt.Println("")

		bodyBytes, err := io.ReadAll(r.Body)
		if err != nil {
			responseNotfound(w)
			return
		}

		var body YoutubeDownloadRequest
		json.Unmarshal(bodyBytes, &body)

		if !validateIsFilled(body) {
			responseBadRequest(w)
			return
		}

		filename := fmt.Sprintf("%d", time.Now().Unix())

		downloadPath := path.Join(body.SaveDir, filename)

		cmd := exec.Command("yt-dlp", "-x", "-o", downloadPath, body.Url)
		fmt.Println(cmd.String())
		err = cmd.Run()
		if err != nil {
			fmt.Println(err)
			responseBadRequest(w)
			return
		}

		cmd = exec.Command("find", body.SaveDir, "-type", "f", "-name", filename+"*", "-not", "-name", "*.mp3")
		fmt.Println(cmd.String())
		downloadedFilenamesBytes, err := cmd.Output()
		if err != nil {
			fmt.Println(err)
			responseBadRequest(w)
			return
		}

		downloadedFilenames := strings.Split(string(downloadedFilenamesBytes), "\n")

		if len(downloadedFilenames) == 0 || downloadedFilenames[0] == "" {
			writeCorsHeaders(w)
			return
		}

		downloadedFilename := downloadedFilenames[0]

		cmd = exec.Command("ffmpeg", "-y", "-i", downloadedFilename, downloadPath+".mp3")
		fmt.Println(cmd.String())
		err = cmd.Run()
		if err != nil {
			fmt.Println(err)
			responseBadRequest(w)
			return
		}

		cmd = exec.Command("rm", downloadedFilename)
		fmt.Println(cmd.String())
		err = cmd.Run()
		if err != nil {
			fmt.Println(err)
			responseBadRequest(w)
			return
		}

		escapedFilename := strings.ReplaceAll(body.Filename, "(", "[")
		escapedFilename = strings.ReplaceAll(escapedFilename, ")", "]")
		escapedFilename = strings.ReplaceAll(escapedFilename, "/", "_")
		renamedPath := path.Join(body.SaveDir, escapedFilename+".mp3")
		cmd = exec.Command("mv", downloadPath+".mp3", renamedPath)
		fmt.Println(cmd.String())
		err = cmd.Run()
		if err != nil {
			fmt.Println(err)
			responseBadRequest(w)
			return
		}

		writeCorsHeaders(w)
	})

	http.ListenAndServe(":5906", nil)
}

func responseNotfound(w http.ResponseWriter) {
	writeCorsHeaders(w)
	w.WriteHeader(404)
}

func responseBadRequest(w http.ResponseWriter) {
	writeCorsHeaders(w)
	w.WriteHeader(400)
}

func writeCorsHeaders(w http.ResponseWriter) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
}

type YoutubeDownloadRequest struct {
	Url      string `json:"url"`
	Filename string `json:"filename"`
	SaveDir  string `json:"save_dir"`
}
