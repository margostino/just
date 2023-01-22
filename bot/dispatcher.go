package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"strings"
)

const (
	JobsLastWeek = "jobs_1w"
	JobsLastDay  = "jobs_1d"
)

func Reply(message *tgbotapi.Message) string {
	timePeriod := strings.Split(message.Text, "_")[1]
	reply := GetJobs(timePeriod)
	return reply
}
