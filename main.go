package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

var token string
var me string
var recipient string

func init() {
	token = os.Getenv("TEXT_BOT_TOKEN")
	me = os.Getenv("TEXT_ME")

	if token == "" {
		fmt.Println("No bot token set")
		os.Exit(3)
	}

	if me == "" {
		fmt.Println("No default Chat ID set")
		os.Exit(3)
	}
}

func main() {
	args := os.Args[1:]

	if args[0] == "me" {
		recipient = me
	} else {
		recipient = os.Getenv("TEXT_" + strings.ToUpper(args[0]))

		if recipient == "" {
			fmt.Println("Unknown recipient")
			os.Exit(3)
		}
	}

	url := "https://api.telegram.org/bot" + token + "/sendMessage?chat_id=" + recipient + "&parse_mode=Markdown&text=" + args[1]

	res, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
	}

	defer res.Body.Close()
}
