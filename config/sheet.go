package config

import (
	"context"
	"github.com/margostino/earth-station-api/common"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
	"log"
	"os"
	"strconv"
)

func GetUrls() map[string]string {
	urls := make(map[string]string)

	ctx := context.Background()
	api, err := sheets.NewService(ctx, option.WithAPIKey(os.Getenv("GSHEET_API_KEY")))

	if !common.IsError(err, "when creating new Google API Service") {
		spreadsheetId := os.Getenv("SPREADSHEET_ID")
		readRange := os.Getenv("SPREADSHEET_RANGE")
		resp, err := api.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()

		if !common.IsError(err, "unable to retrieve data from sheet") && len(resp.Values) > 0 {
			for _, row := range resp.Values {
				var isEnabled bool
				if len(row) == 3 {
					isEnabled, err = strconv.ParseBool(row[2].(string))
					common.SilentCheck(err, "when fetching feed urls configuration")
				} else {
					log.Printf("Configuration sheet for Feed Urls is not valid. It must have 3 columns. It has %d\n", len(row))
				}

				if isEnabled {
					urls[row[1].(string)] = row[0].(string)
				}

			}
		}
	}

	return urls
}
