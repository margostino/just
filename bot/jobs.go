package bot

import (
	"encoding/json"
	"fmt"
	"github.com/margostino/just/common"
	"github.com/margostino/just/domain"
	"net/http"
	"os"
)

var justApiBaseUrl = os.Getenv("JUST_API_BASE_URL")

func GetJobs(timePeriod string) string {
	var reply string

	client := &http.Client{}
	url := fmt.Sprintf("%s?timePeriod=%s", justApiBaseUrl, timePeriod)
	request, err := http.NewRequest(http.MethodGet, url, nil)
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("JUST_API_SECRET")))
	response, err := client.Do(request)
	if !common.IsError(err, "when calling Just API") && response.StatusCode == 200 {
		var jobs []*domain.JobPosition
		err = json.NewDecoder(response.Body).Decode(&jobs)

		if !common.IsError(err, "when decoding Just API response") && len(jobs) > 0 {
			for _, job := range jobs {
				reply += fmt.Sprintf("🔔 New Job! \n"+
					"Title: %s\n"+
					"Datetime: %s\n"+
					"Company: %s\n"+
					"Location: %s\n"+
					"<a href='%s'>Link</a>\n----------\n",
					job.Title,
					job.Datetime,
					job.Company,
					job.Location,
					job.Link)
			}
		} else if len(jobs) == 0 {
			reply = "no jobs found!"
		} else {
			reply = fmt.Sprintf("cannot decoding jobs response: %s", err.Error())
		}

	} else {
		reply = fmt.Sprintf("cannot get jobs (status %d)", response.StatusCode)
	}

	return reply
}
