package rest

import (
	"context"
	"net/http"

	"github.com/vslitvinov/entebot/internal/models"
)

type BotService interface {
	Create(ctx context.Context, mb models.Bot) (string, error)
	GetByOwner(ctx context.Context) ([]models.Bot, error)
	GetByID(ctx context.Context, bid string) (models.Bot, error)
	UpdateByID(ctx context.Context, mb models.Bot) error 
	Delete(ctx context.Context, bid string) error
}

type Bot struct {
	service BotService
}

func NewBot(service BotService) *Bot {

	return &Bot{
		service: service,
	}

}

func (b *Bot) Create(w http.ResponseWriter, r *http.Request) {}

func (b *Bot) GetByOwner(w http.ResponseWriter, r *http.Request) {}

func (b *Bot) GetByID(w http.ResponseWriter, r *http.Request) {}

func (b *Bot) UpdateByID(w http.ResponseWriter, r *http.Request) {}
 
func (b *Bot) Delete(w http.ResponseWriter, r *http.Request) {}