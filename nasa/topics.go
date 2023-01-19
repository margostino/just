package nasa

import (
	"github.com/margostino/earth-station-api/server"
	"github.com/mmcdole/gofeed"
	"log"
	"net/http"
	"reflect"
	"strings"
)

func GetTopics(feedUrls map[string]string, writer http.ResponseWriter) {
	topics := make([]string, 0)
	for topic, _ := range feedUrls {
		topics = append(topics, topic)
	}
	server.WriteResponse(topics, writer)
}

func GetTopic(topic string, feedUrls map[string]string, writer http.ResponseWriter) {

	items := make([]map[string]interface{}, 0)
	feedUrl := feedUrls[topic]
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL(feedUrl)
	if feed != nil {
		for _, entry := range feed.Items {

			element := reflect.ValueOf(entry).Elem()
			item := make(map[string]interface{})

			for i := 0; i < element.NumField(); i++ {
				fieldName := element.Type().Field(i).Name
				fieldValue := element.Field(i).Interface()
				item[strings.ToLower(fieldName)] = fieldValue
			}

			items = append(items, item)
		}
	} else {
		log.Printf("There are no feeds")
	}

	server.WriteResponse(items, writer)

}
