package slack

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/slack-go/slack"
)

func NewClient(token string) *slack.Client {
	return slack.New(token)
}

func getTomorrow730AM() int64 {
	now := time.Now()
	tomorrow := now.AddDate(0, 0, 1)
	tomorrow730 := time.Date(tomorrow.Year(), tomorrow.Month(), tomorrow.Day(), 7, 30, 0, 0, now.Location())

	// Calculate difference in minutes
	diffMinutes := int64(tomorrow730.Sub(now).Minutes())
	return diffMinutes
}

func SetSlackStatus(api *slack.Client) error {
	profile := slack.UserProfile{
		StatusText:  "Descansando até amanhã às 7:30",
		StatusEmoji: ":no_entry_sign:",
	}

	err := api.SetUserCustomStatusContext(context.Background(), profile.StatusText, profile.StatusEmoji, 0)
	if err != nil {
		return fmt.Errorf("falha ao definir o status: %v", err)
	}

	err = api.SetUserPresenceContext(context.Background(), "away")
	if err != nil {
		log.Printf("Erro ao definir presença: %v", err)
		return fmt.Errorf("falha ao definir presença: %v", err)
	}

	dndEndTime := getTomorrow730AM()

	_, err = api.SetSnoozeContext(context.Background(), int(dndEndTime))
	if err != nil {
		return fmt.Errorf("falha ao pausar notificações: %v", err)
	}

	fmt.Println("Status definido com sucesso: 'Descansando até amanhã às 7:30' com emoji 🚫, presença 'away' e notificações pausadas.")
	return nil
}

func ClearSlackStatus(api *slack.Client) error {
	err := api.SetUserCustomStatusContext(context.Background(), "", "", 0)
	if err != nil {
		return fmt.Errorf("falha ao limpar o status: %v", err)
	}

	err = api.SetUserPresenceContext(context.Background(), "auto")
	if err != nil {
		log.Printf("Erro ao definir presença: %v", err)
		return fmt.Errorf("falha ao definir presença: %v", err)
	}

	_, err = api.EndSnoozeContext(context.Background())
	if err != nil {
		return fmt.Errorf("falha ao retomar notificações: %v", err)
	}

	fmt.Println("Status limpo com sucesso.")
	return nil
}

func SetLunchTime(api *slack.Client) error {
	profile := slack.UserProfile{
		StatusText:  "Almoço até às 13:45",
		StatusEmoji: ":spaghetti:",
	}

	err := api.SetUserCustomStatusContext(context.Background(), profile.StatusText, profile.StatusEmoji, 0)
	if err != nil {
		return fmt.Errorf("falha ao definir o status: %v", err)
	}

	err = api.SetUserPresenceContext(context.Background(), "away")
	if err != nil {
		log.Printf("Erro ao definir presença: %v", err)
		return fmt.Errorf("falha ao definir presença: %v", err)
	}

	fmt.Printf("Status definido com sucesso: %s com emoji %s, presença 'away'.\n", profile.StatusText, "🍝")
	return nil
}
