package main

import (
	"fmt"
	"log"
	"os"
	client2 "tasks-distribution/cmd/chatgpt/client"
	"tasks-distribution/cmd/notion/client"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
)

func main() {
	// Загружаем файл .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	notionKey := os.Getenv("NOTION_CLIENT_API_KEY")
	notionDBId := os.Getenv("NOTION_CLIENT_DB_ID")
	chatGPTKey := os.Getenv("CHAT_GPT_CLIENT_API_KEY")
	tgBotKey := os.Getenv("TG_BOT_API_KEY")

	notionClient := client.NewNotionTasksClient(notionKey, notionDBId)
	chatGPTClient := client2.NewChatGPTClient(chatGPTKey)
	bot, err := tgbotapi.NewBotAPI(tgBotKey)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		task, err := chatGPTClient.GetTitle(update.Message.Text)
		if err != nil {
			fmt.Printf("Failed to get title from chat gpt: %w\n", err)
			return
		}
		task.TaskContent = update.Message.Text

		_, err = notionClient.AddNewTask(task)
		if err != nil {
			fmt.Printf("Failed to add new task from notion: %w\n", err)
			return
		}

		// Здесь вы можете добавить функцию для вставки задачи в Notion,
		// например, addTaskToNotion(update.Message.Text)
	}
}
