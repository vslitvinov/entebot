package models

import (
	"time"
)

type Session struct {
	ID        string    `json:"id"`
	AccountID string    `json:"accountId"`
	Provider  string    `json:"provider"`
	UserAgent string    `json:"userAgent"`
	IP        string    `json:"ip"`
	TTL       int       `json:"ttl"`
	ExpiresAt int64     `json:"expiresAt"`
	CreatedAt time.Time `json:"createdAt"`
}

func NewSession(aid, provider, userAgent, ip string, ttl time.Duration) (Session, error) {

	now := time.Now()

	return Session{
		AccountID: aid,
		Provider:  provider,
		UserAgent: userAgent,
		IP:        ip,
		TTL:       int(ttl.Seconds()),
		ExpiresAt: now.Add(ttl).Unix(),
	}, nil
}
