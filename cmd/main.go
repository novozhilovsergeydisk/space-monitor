package main

import (
	"log"
	"os"
	"time"

	"github.com/sergionov/space-monitor/internal/bot"
	"github.com/sergionov/space-monitor/internal/email"
	"github.com/sergionov/space-monitor/internal/monitor"
	"github.com/sergionov/space-monitor/internal/storage"
	"github.com/sergionov/space-monitor/internal/webhook"
)

func main() {
	// Initialize dependencies
	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	emailFrom := os.Getenv("EMAIL_FROM")
	emailPassword := os.Getenv("EMAIL_PASSWORD")
	emailTo := os.Getenv("EMAIL_TO")

	botClient := bot.New(botToken)
	emailClient := email.New(emailFrom, emailPassword, emailTo)
	storageClient := storage.New()
	webhookClient := webhook.New(botClient)

	// Start monitoring
	monitor := monitor.New(botClient, emailClient, storageClient)
	for {
		monitor.CheckDiskSpace()
		time.Sleep(5 * time.Minute)
	}
}
