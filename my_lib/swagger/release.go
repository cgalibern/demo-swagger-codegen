package swagger

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func ReleasePost(w http.ResponseWriter, r *http.Request) {
	var err error
	var release Release
	var body []byte

	body, err = ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("read body error: " + err.Error())
		return
	}
	if err := json.Unmarshal(body, &release); err != nil {
		log.Printf("Unmarshal body error: " + err.Error())
		return
	}

	message := "Process action " + release.Action + " release name " + release.Release.Name
	responseBody, _ := json.Marshal(DefaultResponse{message})

	log.Printf("%v", message)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(responseBody)
}
