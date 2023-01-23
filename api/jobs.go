package api

import (
	"encoding/json"
	"github.com/margostino/just/common"
	"github.com/margostino/just/config"
	"github.com/margostino/just/domain"
	"github.com/margostino/just/executor"
	"log"
	"net/http"
)

type QueryParams struct {
	Keywords         string
	Location         string
	TimePeriod       string
	PaginationFactor string
	Calls            string
	Mode             string
}

func Jobs(w http.ResponseWriter, r *http.Request) {
	var jobs = make([]*domain.JobPosition, 0)

	var configuration = config.GetConfig()

	w.Header().Set("Content-Type", "application/json")

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
	if params.Calls != "" {
		configuration["calls"] = params.Calls
	}
	if params.Mode != "" {
		configuration["mode"] = params.Mode
	}

	//ctx, cancel := context.WithCancel(context.Background())
	//defer cancel()

	jobs = executor.AsyncCall(configuration)

	if len(jobs) == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	} else {
		jsonResp, err := json.Marshal(jobs)
		if err != nil {
			log.Printf("Error happened in JSON marshal. Err: %s\n", err)
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write(jsonResp)
		}
	}

	return
}

func getQueryParam(r *http.Request, param string) string {
	return common.NewString(r.URL.Query().Get(param)).ReplaceAll(" ", "%20").Value()
}

func getQueryParams(r *http.Request) *QueryParams {
	keywords := getQueryParam(r, "keywords")
	location := getQueryParam(r, "location")
	timePeriod := config.GetTimePeriodParam(getQueryParam(r, "timePeriod"))
	mode := config.GetModeParam(getQueryParam(r, "mode"))
	paginationFactor := getQueryParam(r, "paginationFactor")
	calls := getQueryParam(r, "calls")

	return &QueryParams{
		Calls:            calls,
		Keywords:         keywords,
		Location:         location,
		Mode:             mode,
		TimePeriod:       timePeriod,
		PaginationFactor: paginationFactor,
	}
}
