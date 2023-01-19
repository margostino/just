package nasa

import (
	"github.com/margostino/earth-station-api/config"
	"github.com/margostino/earth-station-api/nasa"
	"log"
	"net/http"
	"strings"
)

func Nasa(w http.ResponseWriter, r *http.Request) {

	log.Printf("Request with UserAgent %s and Path %s\n", r.UserAgent(), r.URL.Path)

	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

	if r.URL.Path == "/api/nasa/topics" || strings.HasPrefix(r.URL.Path, "/api/nasa/topics/") {
		feedUrls := config.GetUrls()

		if r.URL.Path == "/api/nasa/topics" || r.URL.Path == "/api/nasa/topics/" {
			nasa.GetTopics(feedUrls, w)
		} else {
			topic := strings.Split(r.URL.Path, "/")[4]
			nasa.GetTopic(topic, feedUrls, w)
		}

	} else {
		w.WriteHeader(http.StatusNotFound)
	}

	return
}
