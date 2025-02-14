package monitor

import (
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"strings"

	"github.com/sergionov/space-monitor/internal/bot"
	"github.com/sergionov/space-monitor/internal/email"
	"github.com/sergionov/space-monitor/internal/storage"
)

type Monitor struct {
	botClient   *bot.Bot
	emailClient *email.Email
	storage     *storage.Storage
}

func New(botClient *bot.Bot, emailClient *email.Email, storage *storage.Storage) *Monitor {
	return &Monitor{
		botClient:   botClient,
		emailClient: emailClient,
		storage:     storage,
	}
}

func (m *Monitor) CheckDiskSpace() {
	out, err := exec.Command("df", "-h").Output()
	if err != nil {
		log.Printf("Error checking disk space: %v", err)
		return
	}

	lines := strings.Split(string(out), "\n")
	for _, line := range lines[1:] {
		fields := strings.Fields(line)
		if len(fields) < 5 {
			continue
		}

		usedPercent, err := strconv.Atoi(strings.TrimSuffix(fields[4], "%"))
		if err != nil {
			log.Printf("Error parsing disk usage: %v", err)
			continue
		}

		if usedPercent > 85 {
			message := fmt.Sprintf("Warning: Disk usage on %s is %d%%", fields[0], usedPercent)
			m.botClient.SendMessage(message)
			m.emailClient.Send("High Disk Usage Alert", message)
			m.storage.Log(message)
		}
	}
}
