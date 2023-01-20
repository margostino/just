package api

import (
	"encoding/json"
	"github.com/margostino/just/collector"
	"github.com/margostino/just/common"
	"github.com/margostino/just/config"
	"github.com/margostino/just/domain"
	"github.com/margostino/just/parser"
	"github.com/margostino/just/processor"
	"log"
	"net/http"
	"strconv"
)

var configuration = config.GetConfig()
var baseUrl = configuration["url"]

func Jobs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var jobs = make([]*domain.JobPosition, 0)
	var isEnd bool

	var index int

	for ok := true; ok; ok = !isEnd {
		factor, err := strconv.Atoi(configuration["pagination_factor"])

		if common.IsError(err, "invalid configuration when parsing") {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		index += factor
		url := collector.GetUrl(r, baseUrl, index)

		if common.IsError(err, "error calling upstream") {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		content, err, statusCode := collector.Call(url)

		if common.IsError(err, "error status from upstream") {
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
