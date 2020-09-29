package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/slack-go/slack"
)

func main() {

	messageSlice := os.Args[1:]
	message := strings.Join(messageSlice, " ")

	if len(messageSlice) == 0 {
		log.Fatal("no message entered!")
	}
	defaultRecipient := "@jon"

	slackToken := os.Getenv("SLACK_TOKEN")
	if len(slackToken) == 0 {
		log.Fatal("ERROR: environment variable SLACK_TOKEN not set")
	}

	slackRecipient := os.Getenv("SLACK_RECIPIENT")
	if len(slackRecipient) == 0 {
		slackRecipient = defaultRecipient
	}

	api := slack.New(slackToken)

	channelID, timestamp, err := api.PostMessage(
		slackRecipient,
		slack.MsgOptionText(message, false),
	)

	if err != nil {
		fmt.Printf("%s\n", err)
	}

	fmt.Printf("Sent message to %s at %s\n", channelID, timestamp)
}
