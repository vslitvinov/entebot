package storage

import (
	"context"
	"fmt"
	"sync"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/vslitvinov/entebot/internal/models"
	"github.com/vslitvinov/entebot/internal/storage/psql"
)


type PSQLBotStorage interface {
	Create(ctx context.Context, mb models.Bot) (string, error)
	GetByOwner(ctx context.Context, owid string) ([]models.Bot, error) 
	GetByID(ctx context.Context, bid string) (models.Bot, error)
	UpdateByID(ctx context.Context, mb models.Bot)  error 
	Delete(ctx context.Context, bid string) error
}

type Bot struct {
	db PSQLBotStorage
	cache sync.Map
	isCache bool 
}

func NewBot(db *pgxpool.Pool, isCache bool) *Bot{
	return &Bot{
		db: psql.NewBot(db),
		cache: sync.Map{},
		isCache: isCache,
	}
}

func (b *Bot) Create(ctx context.Context, mb models.Bot) (string, error) {

	var bid string

	bid, err := b.db.Create(ctx,mb)
	if err != nil {
		return bid, fmt.Errorf("storage.Bot.Create.db.Create %w", err)
	}

	if b.isCache {
		mb.ID = bid
		b.cache.Store(bid,mb)
	}

	return bid, nil 

}

func (b *Bot) GetByOwner(ctx context.Context, owid string) ([]models.Bot, error) {

	var mbs []models.Bot 

	if b.isCache {

		b.cache.Range(func(key any,value any) bool {

			if value.(models.Bot).Owner == owid {
				mb := value.(models.Bot)
				mbs = append(mbs, mb)
			}
			return true
		})

	} else {

		mbs, err := b.db.GetByOwner(ctx,owid)
		if err != nil {
			return mbs, fmt.Errorf("storage.Bot.GetByOwner.db.GetByOwner %w", err)
		}

	}

	return mbs, nil 
 
}

func (b *Bot) GetByID(ctx context.Context, bid string) (models.Bot, error) {

	var mb models.Bot

	if b.isCache {
		mbC, ok := b.cache.Load(bid)
		if !ok {
			return mb, fmt.Errorf("storage.Bot.GetByID.cache.Load %w", nil)
		}
		mb = mbC.(models.Bot)
	} else {
		mb, err := b.db.GetByID(ctx,bid) 
		if err != nil {
			return mb, fmt.Errorf("storage.Bot.GetByID.db.GetByID %w", err)
		}
	}

	return mb, nil

}

func (b *Bot) UpdateByID(ctx context.Context, mb models.Bot) error {

	err := b.db.UpdateByID(ctx,mb)
	if err != nil {
		return fmt.Errorf("storage.Bot.UpdateByID.db.UpdateByID %w", err)
	}

	if b.isCache {
		_, ok := b.cache.LoadAndDelete(mb.ID)
		if !ok {
			return fmt.Errorf("storage.Bot.UpdateByID.cache.LoadAndDelete %w", nil)
		}
		b.cache.Store(mb.ID, mb)
	}

	return nil 

}

func (b *Bot) Delete(ctx context.Context, bid string) error {

	err := b.db.Delete(ctx,bid)
	if err != nil {
		return fmt.Errorf("storage.Bot.Delete.db.Delete %w", err)
	}

	_, ok := b.cache.LoadAndDelete(bid)
	if !ok {
		return fmt.Errorf("storage.Bot.Delete.cache.LoadAndDelete %w", err)
	}

	return nil
 
}