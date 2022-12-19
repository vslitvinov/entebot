package service

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/vslitvinov/entebot/internal/models"
	"github.com/vslitvinov/entebot/internal/storage"
)


type SessionStorage interface {
	Create(ctx context.Context, ms models.Session) (models.Session, error)
	FindByID(ctx context.Context, sid string) (models.Session, error)
	FindAll(ctx context.Context, aid string) ([]models.Session, error)
	Delete(ctx context.Context, sid string) error 
	DeleteAll(ctx context.Context, aid,sid string) error
}


type Session struct {
	storage SessionStorage
}

func NewSession(db *pgxpool.Pool, isCache bool) *Session{
	return &Session{
		storage: storage.NewSession(db,isCache),
	}
}


// Device represents data transfer object with user device data
type Device struct {
	UserAgent string
	IP        string
}

// Create new session
func (s *Session) Create(ctx context.Context, aid, provider string, d Device) (models.Session, error) {

	var se models.Session

	se, err := models.NewSession(aid,provider,d.UserAgent,d.IP,224)
	if err != nil {
		return se, fmt.Errorf("service.session.Create.NewSession %w", err)
	}

	se, err = s.storage.Create(ctx,se)
	if err != nil {
		return se, fmt.Errorf("service.session.Create %w", err)
	}

	return se, nil
}

// GetByID session.
func (s *Session) GetByID(ctx context.Context, sid string) (models.Session, error) {

	var se models.Session

	se,err := s.storage.FindByID(ctx,sid)
	if err != nil {
		return se, fmt.Errorf("service.session.GetByID %w", err)
	}

	return se, nil
}

// GetAll account sessions using provided account id.
func (s *Session) GetAll(ctx context.Context, aid string) ([]models.Session, error) {
	var ses []models.Session

	ses,err := s.storage.FindAll(ctx,aid)
	if err != nil {
		return ses, fmt.Errorf("service.session.GetAll %w", err)
	}

	return ses, nil
}

// Finish session by id excluding current session with id.
func (s *Session) Finish(ctx context.Context, sid, currSid string) error {
	if sid == currSid {
		return fmt.Errorf("service.session.Finish:")
	}

	err := s.storage.Delete(ctx,sid)
	if err != nil {
		return fmt.Errorf("service.Finish %w", err)
	}

	return nil
}

// FinishAll account sessions excluding current session with id.
func (s *Session) FinishAll(ctx context.Context, aid, sid string) error {
	err := s.storage.DeleteAll(ctx,aid,sid)
	if err != nil {
		return fmt.Errorf("service.session.FinishAll %w", err)
	}

	return nil
}