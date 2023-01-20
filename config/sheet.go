package config

import (
	"context"
	"github.com/margostino/just/common"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
	"log"
	"os"
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
				if len(row) == 2 {
					config["url"] = row[0].(string)
					config["pagination_factor"] = row[1].(string)
				} else {
					log.Printf("Configuration sheet for Feed Urls is not valid. It must have 3 columns. It has %d\n", len(row))
				}
			}
		}
	}

	return config
}
