package api

import (
	"fmt"
	"github.com/margostino/just/collector"
	"github.com/margostino/just/common"
	"net/http"
)

func getUrl(r *http.Request, index int) string {
	params := getQueryParams(r)
	baseUrl := "dummy"
	pagination := fmt.Sprintf("position=1&pageNum=0&start=%d", index)
	return fmt.Sprintf("%s?keywords=%s&location=%s&f_TPR=%s&%s", baseUrl, params.Keywords, params.Location, params.TimePeriod, pagination)
}

func getQueryParam(r *http.Request, param string) string {
	return common.NewString(r.URL.Query().Get(param)).ReplaceAll(" ", "%20").Value()
}

func getQueryParams(r *http.Request) *collector.QueryParams {
	keywords := getQueryParam(r, "keywords")
	location := getQueryParam(r, "location")
	timePeriod := getQueryParam(r, "f_TPR")
	return &collector.QueryParams{
		Keywords:   keywords,
		Location:   location,
		TimePeriod: timePeriod,
	}
}
