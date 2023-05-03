package main

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	botToken := "your_bot_token"
	chatID := int64(0000)

	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		panic(err)
	}

	addr := "x.x.x.x:22"
	timeout := time.Second
	//bot.Debug = true

	for {
		_, err := net.DialTimeout("tcp", addr, timeout)
		if err != nil {
			msg := tgbotapi.NewMessage(chatID, fmt.Sprintf("Не могу подключиться, ПАМАГИТЕ!! %s: %v", addr, err))
			_, err := bot.Send(msg)
			if err != nil {
				fmt.Println("Error sending message:", err)
			}
		} else {
			log.Printf("Отправилось chat_id %d\n", chatID)
			fmt.Printf("Всё ОК, СПИ СПОКОЙНО БРАТ %s\n", addr)

			msg := tgbotapi.NewMessage(chatID, "ВСЁ ОК, СПИ СПОКОЙНО БРАТ")
			_, err := bot.Send(msg)
			if err != nil {
				fmt.Println("Error sending message:", err)
			}
		}
		time.Sleep(1 * time.Hour)
	}

}
