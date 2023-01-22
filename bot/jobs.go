package bot

import (
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/margostino/just/common"
	"github.com/margostino/just/domain"
	"log"
	"net/http"
	"os"
	"strconv"
)

var justApiBaseUrl = os.Getenv("JUST_API_BASE_URL")
var botApi, _ = newBot()

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
				message := fmt.Sprintf("🔔 New Job! \n"+
					"Title: <a href='%s'>%s</a>\n"+
					"Datetime: %s\n"+
					"Company: %s\n"+
					"Location: %s\n",
					job.Link,
					job.Title,
					job.Datetime,
					job.Company,
					job.Location)
				send(message)
			}
			reply = fmt.Sprintf("✅ Total Jobs: %d", len(jobs))
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

func send(message string) {
	if botApi != nil {
		userId, _ := strconv.ParseInt(os.Getenv("TELEGRAM_ADMIN_USER"), 10, 64)
		msg := tgbotapi.NewMessage(userId, message)
		msg.ReplyMarkup = nil
		msg.ParseMode = "HTML"
		botApi.Send(msg)
	} else {
		log.Printf("Bot initialization failed")
	}
}

func newBot() (*tgbotapi.BotAPI, error) {
	client, error := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_BOT_TOKEN"))
	//bot.Debug = true
	common.SilentCheck(error, "when creating a new BotAPI instance")
	//log.Printf("Authorized on account %s\n", client.Self.UserName)
	return client, error
}
