package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/margostino/just/common"
)

var commands = common.NewStringSlice("/jobs_1d", "/jobs_1w")

func SanitizeInput(input string) string {
	return common.NewString(input).
		ToLower().
		Trim(" ").
		Value()
}

func IsValidInput(message *tgbotapi.Message) bool {
	input := message.Text
	sanitizedInput := SanitizeInput(input)
	return commands.Contains(sanitizedInput)
}

func extractIds(input string, prefix string) string {
	return common.NewString(input).
		ToLower().
		TrimPrefix(prefix).
		Split(" ").
		Join(",").
		Value()
}
