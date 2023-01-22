package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/margostino/just/common"
	"strings"
)

const (
	JobsLastWeek = "jobs_1w"
	JobsLastDay  = "jobs_1d"
)

func Reply(message *tgbotapi.Message) string {
	var reply string

	//if message.ReplyToMessage != nil {
	//	input := message.ReplyToMessage.Text
	//	reply = PushReply(input)
	//} else {
	//
	//}

	input := message.Text
	command := common.NewString(input).
		ReplaceAll("/", "").
		ReplaceAll("_", " ").
		Value()
	sanitizedInput := SanitizeInput(command)
	commands := strings.Split(sanitizedInput, " ")

	if len(commands) > 0 {
		switch commands[0] {
		case JobsLastDay:
			reply = GetJobs(JobsLastDay)
		case JobsLastWeek:
			reply = GetJobs(JobsLastWeek)
		default:
			reply = "👌"
		}
	} else {
		reply = "🙈 Invalid command!"
	}

	return reply
}
