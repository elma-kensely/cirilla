package commands

import (
	"gopkg.in/telegram-bot-api.v4"
	"log"
	"strings"
)

//Command structure to represent commands
type Command struct {
	Function    func(bot *tgbotapi.BotAPI, args []string, PassString bool, update tgbotapi.Update) (err error)
	Description string
	PassString  bool
	Admin       bool
}

//Init loads all commands
func Init() map[string]Command {
	return map[string]Command{
		"/say": {
			Function:    commandSay,
			Description: "Say as Cirilla",
			PassString:  true,
			Admin:       true,
		},
	}
}

//ExecuteCommand executes command
func ExecuteCommand(update tgbotapi.Update, Commands map[string]Command, bot *tgbotapi.BotAPI) {
	CommandName := strings.Split(update.Message.Text, " ")[0]

	if cmd, ok := Commands[CommandName]; ok {
		var args []string
		if cmd.PassString {
			args = append(args, strings.SplitN(update.Message.Text, " ", 2)[1])
		} else {

			Arguements := strings.Split(update.Message.Text, " ")[1:]
			for _, i := range Arguements {
				args = append(args, i)
			}
		}
		if cmd.Admin {
			chatConfig := update.Message.Chat.ChatConfig()
			admins, _ := bot.GetChatAdministrators(chatConfig)

			var found = false
			for _, admin := range admins {
				if admin.User.ID == update.Message.From.ID {
					found = true
					break
				}
			}

			if !found {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "You are not authorised to use that Command")
				msg.ReplyToMessageID = update.Message.MessageID
				bot.Send(msg)
				return
			}
		}
		err := cmd.Function(bot, args, cmd.PassString, update)
		if err != nil {
			log.Println("Command : ", CommandName, " Failed to execute")
		}
	} else {
		log.Println("Unknown Command : ", CommandName)
	}
}
