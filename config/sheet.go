package config

import (
	"context"
	"github.com/margostino/just/common"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
	"log"
	"os"
)

const (
	LastDay    = "r86400"
	LastWeek   = "r604800"
	OnsiteMode = "1"
	RemoteMode = "2"
	HybridMode = "3"
)

func GetConfig() map[string]string {
	config := make(map[string]string)

	ctx := context.Background()
	api, err := sheets.NewService(ctx, option.WithAPIKey(os.Getenv("GSHEET_API_KEY")))

	if !common.IsError(err, "when creating new Google API Service") {
		spreadsheetId := os.Getenv("SPREADSHEET_ID")
		readRange := os.Getenv("SPREADSHEET_RANGE")
		resp, err := api.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()

		if !common.IsError(err, "unable to retrieve data from sheet") && len(resp.Values) > 0 {
			for _, row := range resp.Values {
				if len(row) == 7 {
					config["url"] = row[0].(string)
					config["paginationFactor"] = row[1].(string)
					config["keywords"] = row[2].(string)
					config["location"] = row[3].(string)
					config["calls"] = row[5].(string)

					timePeriod := row[4].(string)
					mode := row[6].(string)
					config["timePeriod"] = GetTimePeriodParam(timePeriod)
					config["mode"] = GetModeParam(mode)

				} else {
					log.Printf("Configuration sheet for Feed Urls is not valid. It must have 3 columns. It has %d\n", len(row))
				}
			}
		}
	}

	return config
}

func GetTimePeriodParam(timePeriod string) string {
	var timePeriodParam string
	switch timePeriod {
	case "1w":
		timePeriodParam = LastWeek
		break
	case "1d":
		timePeriodParam = LastDay
		break
	}
	return timePeriodParam
}

func GetModeParam(mode string) string {
	var modeParam string
	switch mode {
	case "onsite":
		modeParam = OnsiteMode
		break
	case "remote":
		modeParam = RemoteMode
		break
	case "hybrid":
		modeParam = HybridMode
		break
	}
	return modeParam
}
