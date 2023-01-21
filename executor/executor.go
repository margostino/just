package executor

import (
	"github.com/margostino/just/client"
	"github.com/margostino/just/common"
	"github.com/margostino/just/domain"
	"github.com/margostino/just/parser"
	"github.com/margostino/just/processor"
	"strconv"
	"sync"
)

type QueryParams struct {
	Keywords         string
	Location         string
	TimePeriod       string
	PaginationFactor string
}

func AsyncCall(config map[string]string) []*domain.JobPosition {
	var offset int
	var wg sync.WaitGroup
	var jobs = make([]*domain.JobPosition, 0)

	calls, err := strconv.Atoi(config["calls"])
	if common.IsError(err, "invalid configuration for calls") {
		return jobs
	}

	jobsChannel := make(chan []*domain.JobPosition, calls)
	defer close(jobsChannel)

	//go func() {
	//	for partial := range jobsChannel {
	//		log.Printf("Got partial results: %d", len(partial))
	//		jobs = append(jobs, partial...)
	//	}
	//}()

	wg.Add(calls)
	for i := 1; i <= calls; i++ {

		factor, err := strconv.Atoi(config["paginationFactor"])

		if common.IsError(err, "invalid configuration when parsing") {
			wg.Done()
			return jobs
		}

		offset += factor
		go func(i int) {
			defer wg.Done()
			url := client.GetUrl(config, i)
			content, err, _ := client.Call(url)
			common.SilentCheck(err, "error calling upstream")

			tokens := parser.Parse(string(content))
			if len(tokens) > 0 {
				partialJobs := processor.Process(tokens)
				//log.Printf("jobs from offset %d: %d", i, len(partialJobs))
				jobsChannel <- partialJobs
			} else {
				jobsChannel <- make([]*domain.JobPosition, 0)
				//log.Printf("NO jobs from offset %d", i)
			}

		}(offset)

	}

	for i := 1; i <= calls; i++ {
		partial := <-jobsChannel
		//log.Printf("Got partial results: %d", len(partial))
		jobs = append(jobs, partial...)
	}

	wg.Wait()

	return jobs
}
