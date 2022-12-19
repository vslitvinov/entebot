package storage

import (
	"context"
	"fmt"
	"sync"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/vslitvinov/entebot/internal/models"
	"github.com/vslitvinov/entebot/internal/storage/psql"
)


type PSQLSessionStorage interface {
	Create(ctx context.Context, ms models.Session) (models.Session, error)
	FindByID(ctx context.Context, sid string) (models.Session, error)
	FindAll(ctx context.Context, aid string) ([]models.Session, error)
	Delete(ctx context.Context, sid string) error 
	DeleteAll(ctx context.Context, aid,sid string) error
}

type Session struct {
	db PSQLSessionStorage
	cache sync.Map
	isCache bool
}

func NewSession(db *pgxpool.Pool, isCache bool) *Session{

	return &Session{
		db: psql.NewSession(db),
		cache: sync.Map{},
		isCache: isCache,
	}

}


func (s *Session) Create(ctx context.Context, ms models.Session) (models.Session, error) {

	rms, err := s.db.Create(ctx,ms)
	if err != nil {
		return ms, fmt.Errorf("storage.Create.db.Create %w", err)
	}

	s.cache.Store(rms.ID,rms)

	return rms, nil 

}

func (s *Session) FindByID(ctx context.Context, sid string) (models.Session, error) {

	var rms models.Session

	if s.isCache {

		rms, ok := s.cache.Load(sid)
		if !ok {
			return rms.(models.Session), fmt.Errorf("storage.FindByID.cache.Load %w", nil)
		}

	} else {

		rms, err := s.db.FindByID(ctx,sid)
		if err != nil {
			return rms, fmt.Errorf("storage.FindByID.db.FindByID %w", nil)
		}

	}

	return rms, nil

}

func (s *Session) FindAll(ctx context.Context, aid string) ([]models.Session, error) {
	return []models.Session{},nil
}

func (s *Session) Delete(ctx context.Context, sid string) error {


	err := s.db.Delete(ctx,sid)
	if err != nil {
		return fmt.Errorf("storage.Delete.db.Delete %w", nil)
	}

	if s.isCache {
		s.cache.Delete(sid)
	}

	return nil

}

func (s *Session) DeleteAll(ctx context.Context, aid,sid string) error {
	return nil
}
