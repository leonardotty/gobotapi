package main

import (
	"fmt"
	"github.com/Squirrel-Network/gobotapi"
	"github.com/Squirrel-Network/gobotapi/methods"
	"github.com/Squirrel-Network/gobotapi/types"
)

func main() {
	client := gobotapi.NewClient("YOUR_TOKEN")
	// Add listener to receive only start commands with alias "/", ";", "."
	client.OnCommand("start", []string{"/", ";", "."}, func(update types.Message) {
		_, err := client.Invoke(&methods.SendMessage{
			ChatID: update.Chat.ID,
			Text:   "Hello, I'm a bot!",
		})
		// Check if there is an error
		if err != nil {
			panic(err)
		}
	})
	// Add listener to receive only stop commands with the default alias
	client.OnCommand("stop", nil, func(update types.Message) {
		fmt.Println(fmt.Sprintf("Stop command received from %d", update.Chat.ID))
	})
	// Start and idle the bot
	client.Run()
}
