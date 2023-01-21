package client

import (
	"fmt"
	"github.com/margostino/just/common"
)

func GetUrl(configuration map[string]string, index int) string {
	//params := getQueryParams(r)
	var baseUrl = sanitizeParameter(configuration, "url")
	var keywords = sanitizeParameter(configuration, "keywords")
	var location = sanitizeParameter(configuration, "location")
	var timePeriod = sanitizeParameter(configuration, "timePeriod")
	pagination := fmt.Sprintf("start=%d", index)
	return fmt.Sprintf("%s?keywords=%s&location=%s&f_TPR=%s&%s", baseUrl, keywords, location, timePeriod, pagination)
}

func sanitizeParameter(configuration map[string]string, param string) string {
	return common.NewString(configuration[param]).ReplaceAll(" ", "%20").Value()
}
