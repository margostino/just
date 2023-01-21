package api

import (
	"encoding/json"
	"github.com/margostino/just/client"
	"github.com/margostino/just/common"
	"github.com/margostino/just/config"
	"github.com/margostino/just/domain"
	"github.com/margostino/just/parser"
	"github.com/margostino/just/processor"
	"log"
	"net/http"
	"strconv"
)

type QueryParams struct {
	Keywords         string
	Location         string
	TimePeriod       string
	PaginationFactor string
}

var configuration = config.GetConfig()

func Jobs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var jobs = make([]*domain.JobPosition, 0)
	var isEnd bool
	var index int

	params := getQueryParams(r)
	if params.Keywords != "" {
		configuration["keywords"] = params.Keywords
	}
	if params.Location != "" {
		configuration["location"] = params.Location
	}
	if params.TimePeriod != "" {
		configuration["timePeriod"] = params.TimePeriod
	}
	if params.PaginationFactor != "" {
		configuration["paginationFactor"] = params.PaginationFactor
	}

	for ok := true; ok; ok = !isEnd {
		factor, err := strconv.Atoi(configuration["paginationFactor"])

		if common.IsError(err, "invalid configuration when parsing") {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		index += factor
		url := client.GetUrl(configuration, index)

		if common.IsError(err, "error calling upstream") {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		content, err, statusCode := client.Call(url)

		if statusCode == 400 || statusCode == 429 {
			break
		}

		if common.IsError(err, "error status from upstream") {
			w.WriteHeader(statusCode)
			return
		}

		tokens := parser.Parse(string(content))
		partialJobs := processor.Process(tokens)
		jobs = append(jobs, partialJobs...)

		if len(jobs) == 0 {
			w.WriteHeader(http.StatusNotFound)
			return
		}
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

func getQueryParam(r *http.Request, param string) string {
	return common.NewString(r.URL.Query().Get(param)).ReplaceAll(" ", "%20").Value()
}

func getQueryParams(r *http.Request) *QueryParams {
	keywords := getQueryParam(r, "keywords")
	location := getQueryParam(r, "location")
	timePeriod := getQueryParam(r, "timePeriod")
	paginationFactor := getQueryParam(r, "paginationFactor")
	return &QueryParams{
		Keywords:         keywords,
		Location:         location,
		TimePeriod:       timePeriod,
		PaginationFactor: paginationFactor,
	}
}
