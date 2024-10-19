package main

import (
	"flag"
	"fmt"
	"log"

	"addi/clients/telegram"
)

func main() {
	token := mustToken()
	fmt.Println("Token: ", token)

	tgClient := telegram.New("host", token)
	fmt.Println(tgClient)

	// fetcher = fetcher.New()

	// processor = processor.New()

	// consumer.Start(fetcher, processor)

	fmt.Println("Hello, World!")
}

func mustToken() string {
	token := flag.String("tg-bot-token", "", "token for access to telegram bot")

	flag.Parse()
	if *token == "" {
		log.Fatal()
	}

	return *token
}
