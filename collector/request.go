package collector

import (
	"fmt"
	"github.com/margostino/just/common"
	"net/http"
)

func GetUrl(r *http.Request, baseUrl string, index int) string {
	params := getQueryParams(r)
	pagination := fmt.Sprintf("position=1&pageNum=0&start=%d", index)
	return fmt.Sprintf("%s?keywords=%s&location=%s&f_TPR=%s&%s", baseUrl, params.Keywords, params.Location, params.TimePeriod, pagination)
}

func getQueryParam(r *http.Request, param string) string {
	return common.NewString(r.URL.Query().Get(param)).ReplaceAll(" ", "%20").Value()
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
