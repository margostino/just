package collector

import (
	"fmt"
	"github.com/margostino/just/common"
	"net/http"
)

func GetUrl(configuration map[string]string, index int) string {
	//params := getQueryParams(r)
	var baseUrl = sanitizeParameter(configuration, "url")
	var keywords = sanitizeParameter(configuration, "keywords")
	var location = sanitizeParameter(configuration, "location")
	var timePeriod = sanitizeParameter(configuration, "timePeriod")
	pagination := fmt.Sprintf("position=1&pageNum=0&start=%d", index)
	return fmt.Sprintf("%s?keywords=%s&location=%s&f_TPR=%s&%s", baseUrl, keywords, location, timePeriod, pagination)
}

func getQueryParam(r *http.Request, param string) string {
	return common.NewString(r.URL.Query().Get(param)).ReplaceAll(" ", "%20").Value()
}

func sanitizeParameter(configuration map[string]string, param string) string {
	return common.NewString(configuration[param]).ReplaceAll(" ", "%20").Value()
}

func getQueryParams(r *http.Request) *QueryParams {
	keywords := getQueryParam(r, "keywords")
	location := getQueryParam(r, "location")
	timePeriod := getQueryParam(r, "f_TPR")
	return &QueryParams{
		Keywords:   keywords,
		Location:   location,
		TimePeriod: timePeriod,
	}
}
