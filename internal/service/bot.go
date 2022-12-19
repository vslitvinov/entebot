package service

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/vslitvinov/entebot/internal/models"
	"github.com/vslitvinov/entebot/internal/storage"
)


type BotStorage interface {
	Create(ctx context.Context, mb models.Bot) (string, error)
	GetByOwner(ctx context.Context, owid string) ([]models.Bot, error)
	GetByID(ctx context.Context, bid string) (models.Bot, error)
	UpdateByID(ctx context.Context, mb models.Bot) error
	Delete(ctx context.Context, bid string) error 
}

type BotService interface {
	CheckSession(ctx context.Context, sid string) (models.Session, error)
}

type Bot struct {
	serviceAuth BotService
	storage BotStorage
}

func NewBot(serviceAuth BotService, db *pgxpool.Pool, isCache bool) *Bot {
	return &Bot{
		serviceAuth: serviceAuth,
		storage: storage.NewBot(db,isCache),
	}
}

func (b *Bot) getAccountID(ctx context.Context) (string, error) {

	var aid string

	sid, ok := ctx.Value("session_token").(string)
	if !ok {
		return aid, fmt.Errorf("Service.Bot.getAccountID.ctx.Value %w", nil)
	}

	ms, err := b.serviceAuth.CheckSession(ctx,sid)
	if err != nil {
		return aid, fmt.Errorf("Service.Bot.getAccountID.ServiceAuth.CheckSession %w", err)
	}

	return ms.AccountID, nil
}

func (b *Bot) Create(ctx context.Context, mb models.Bot) (string, error) {

	var bid string

	aid, err := b.getAccountID(ctx)
	if err != nil {
		return bid,  fmt.Errorf("Service.Bot.Create %w", err)
	}

	mb.Owner = aid

	bid, err = b.storage.Create(ctx,mb)
	if err != nil {
		return bid, fmt.Errorf("Service.Bot.Create %w", err)
	}

	return bid, nil

}

func (b *Bot) GetByOwner(ctx context.Context) ([]models.Bot, error) {

	var mbs []models.Bot

	aid, err := b.getAccountID(ctx)
	if err != nil {
		return mbs,  fmt.Errorf("Service.Bot.GetByOwner %w", err)
	}

	mbs, err = b.storage.GetByOwner(ctx,aid)
	if err != nil {
		return mbs, fmt.Errorf("Service.Bot.storage.GetByOwner %w", err)
	}

	return mbs, nil

}

func (b *Bot) GetByID(ctx context.Context, bid string) (models.Bot, error) {

	var mb models.Bot

	_, err := b.getAccountID(ctx)
	if err != nil {
		return mb, fmt.Errorf("Service.Bot.GetByID %w", err)
	}

	mb, err = b.storage.GetByID(ctx,bid)
	if err != nil {
		return mb, fmt.Errorf("Service.Bot.storage.GetByID %w", err)
	}

	return mb, nil

}

func (b *Bot) UpdateByID(ctx context.Context, mb models.Bot) error {

	aid, err := b.getAccountID(ctx)
	if err != nil {
		return fmt.Errorf("Service.Bot.UpdateByID %w", err)
	}

	mb.Owner = aid

	err = b.storage.UpdateByID(ctx,mb)
	if err != nil {
		return fmt.Errorf("Service.Bot.UpdateByID.db.UpdateByID %w", err)
	}

	return nil

}

func (b *Bot) Delete(ctx context.Context, bid string) error {

	_, err := b.getAccountID(ctx)
	if err != nil {
		return fmt.Errorf("Service.Bot.Delete %w", err)
	}

	err = b.storage.Delete(ctx,bid)
	if err != nil {
		return fmt.Errorf("Service.Bot.Delete.db.Delete %w", err)
	}

	return nil

}
