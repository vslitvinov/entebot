package storage

import (
	"context"
	"fmt"
	"sync"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/vslitvinov/entebot/internal/models"
	"github.com/vslitvinov/entebot/internal/storage/psql"
)


type PSQLAccountStorage interface {
	Create(ctx context.Context, ac models.Accounty) (string, error)
	FindByID(ctx context.Context, aid string) (models.Accounty, error)
	FindByEmail(ctx context.Context, email string) (models.Accounty, error)
	Verified(ctx context.Context, aid string, verified bool) error
	Archive(ctx context.Context, aid string, archive bool) error
}


type Account struct {
	db PSQLAccountStorage
	cache sync.Map
	isCache bool
}

func NewAccount(db *pgxpool.Pool, isCache bool) *Account{

	return &Account{
		db: psql.NewAccount(db),
		cache: sync.Map{},
		isCache: isCache,
	}

}


func (a *Account) Create(ctx context.Context, ac models.Accounty) (string, error) {

	var aid string

	aid, err := a.db.Create(ctx,ac)
	if err != nil {
		return aid, fmt.Errorf("storage.Account.Create.db.Create %w", err)
	}

	if a.isCache {
		ac.ID = aid
		a.cache.Store(aid,ac)
	}

	return aid, nil

}

func (a *Account) FindByID(ctx context.Context, aid string) (models.Accounty, error) {

	var ac models.Accounty

	if a.isCache {
		ac, ok := a.cache.Load(aid)
		if !ok {
			return ac.(models.Accounty), fmt.Errorf("storage.Account.FindByID.cache.Load %w", nil)
		}
	} else {
		ac, err := a.db.FindByID(ctx,aid)
		if err != nil {
			return ac, fmt.Errorf("storage.Account.FindByID.db.FindByID %w", err)
		}
	}

	return ac,nil

}

func (a *Account) FindByEmail(ctx context.Context, email string) (models.Accounty, error) {
	return models.Accounty{},nil
}

func (a *Account) Verified(ctx context.Context, aid string, verified bool) error {
	return nil
}

func (a *Account) Archive(ctx context.Context, aid string, archive bool) error {
	return nil
}





