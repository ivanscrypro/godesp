package main

import (
	"fmt"
	"log"
	"log/slog"
	"os"
)

type LogMessage struct {
	MessageType string `json:"message_type"`
	Message     string `json:"message"`
	webTitle    string `json:"web_title"`
	htmlTitle   string `json:"html_title"`
	URL         string `json:"url"`
}

func getHelper(message string) {
	usage := "Flags:\n"
	usage += " -t Target Domain.\n"
	usage += " -l - Log filename path. The default value is ./assets/{target}.log\n"
	usage += " -e - The authorization token for the white-box testing. Default value is 'aws'\n"
	usage += "Example:\n"
	usage += " ./godork -t example.com -e aws -l example1\n"
	fmt.Println(usage)
	log.Fatal("The error is " + message)
}

func (m *LogMessage) getLogger() {
	slogger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	switch m.MessageType {
	case "helper":
		getHelper(m.Message)
	case "fatal":
		slogger.Error(m.Message)
	case "regular":
		slogger.Info(m.Message)
	case "error":
		slogger.Warn(m.Message, "webTitle", m.webTitle, "htmlTitle", m.htmlTitle, "URL", m.URL)
	default:
		getHelper(m.Message)
	}

}
