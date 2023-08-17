package main

import (
	"flag"
	"log"

	"github.com/benavio/Go/clients/telegram"
)

const (
	tgBotHost = "api.telegram.org"
)

func main() {
	//tocken = flags.Get(token)
	tgClient = telegram.New(tgBotHost, mustToken())
	//tgClient = telegram.New(token)

	//fetcher = fetcher.New()

	//processor = processor.New()

	// consumer.Start(fetcher, processor)

}

func mustToken() string {
	token := flag.String("token-bot-token", "", "Your token access to tg_bot")
	flag.Parse()
	if *token == "" {
		log.Fatal("invalid token")
	}
	return *token
}
