package service

import (
	"context"
	"fmt"

	"github.com/vslitvinov/entebot/internal/models"
)


type AccountService interface {
	Create(ctx context.Context, ac models.Accounty) (string, error)
	GetByID(ctx context.Context, aid string) (models.Accounty, error) 
	GetByEmail(ctx context.Context, email string) (models.Accounty, error)
	Delete(ctx context.Context, aid, sid string) error
}

type SessionService interface {
	Create(ctx context.Context, aid, provider string, d Device) (models.Session, error)
	GetByID(ctx context.Context, sid string) (models.Session, error)
	GetAll(ctx context.Context, aid string) ([]models.Session, error)
	Finish(ctx context.Context, sid, currSid string) error 
	FinishAll(ctx context.Context, aid, sid string) error
}


type Auth struct {
	accountService AccountService
	sessionService SessionService
}

func NewAuth(accountService AccountService, sessionService SessionService) *Auth {
	return &Auth{
		accountService: accountService,
		sessionService: sessionService,
	}
}

func (a *Auth) EmailSingIn(ctx context.Context, email, password string, d Device)  (models.Session, error) {

	var ms models.Session 

	ma, err := a.accountService.GetByEmail(ctx,email)
	if err != nil {
		return ms, fmt.Errorf("service.auth.EmailSingIn.GetByEmail %w", err)
	}

	err = ma.CompareHashAndPassword(password)
	if err != nil {
		return ms, fmt.Errorf("service.auth.EmailSingIn.CompareHashAndPassword %w", err)
	}

	ms, err = a.sessionService.Create(ctx,ma.ID,"email",d)
	if err != nil {
		return ms, fmt.Errorf("service.auth.EmailSingIn.CreateSession %w", err)
	}

	return ms, nil
}

func (a *Auth) CheckSession(ctx context.Context, sid string) (models.Session, error) {

	var s models.Session

	s, err := a.sessionService.GetByID(ctx, sid)
	if err != nil {
		return s, fmt.Errorf("service.auth.CheckSession.GetByID %w", err)
	}

	return s, nil

}

func (a *Auth) SingUp(ctx context.Context, ma models.Accounty) (string,error) {

	var id string

	id, err := a.accountService.Create(ctx, ma)
	if err != nil {
		return id, fmt.Errorf("service.auth.SingUp.Create %w", err)
	}

	return id, nil

}

func (a *Auth) LogOut(ctx context.Context, sid string) error {
	if err := a.sessionService.Finish(ctx, sid, ""); err != nil {
		return fmt.Errorf("authService.LogOut.session.Finish: %w", err)
	}

	return nil
}
