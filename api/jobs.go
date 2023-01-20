package api

import (
	"encoding/json"
	"github.com/margostino/just/collector"
	"github.com/margostino/just/domain"
	"github.com/margostino/just/parser"
	"github.com/margostino/just/processor"
	"log"
	"net/http"
)

func Jobs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var jobs = make([]*domain.JobPosition, 0)
	var isEnd bool

	var index = 0
	var factor = 50

	for ok := true; ok; ok = !isEnd {
		index += factor
		url := getUrl(r, index)
		content, err, statusCode := collector.Call(url)

		if err != nil {
			w.WriteHeader(statusCode)
			return
		}

		if statusCode == 400 {
			break
		}

		tokens := parser.Parse(string(content))
		partialJobs := processor.Process(tokens)
		jobs = append(jobs, partialJobs...)
	}

	jsonResp, err := json.Marshal(jobs)
	if err != nil {
		log.Printf("Error happened in JSON marshal. Err: %s\n", err)
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write(jsonResp)
	}

	return
}
