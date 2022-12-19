package psql

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/vslitvinov/entebot/internal/models"
)

type Session struct {
	pool *pgxpool.Pool
}

func NewSession(db *pgxpool.Pool) *Session {
	return &Session{
		pool: db,
	}
}

// Create new session.
func (s *Session) Create(ctx context.Context, ms models.Session) (models.Session, error) {

	sql := `INSERT INTO session (account_id, provider, user_agent, ip, ttl, expires_at)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id,created_at`

	row := s.pool.QueryRow(ctx, sql,
		ms.AccountID,
		ms.Provider,
		ms.UserAgent,
		ms.IP,
		ms.TTL,
		ms.ExpiresAt,
	)
	err := row.Scan(&ms.ID,&ms.CreatedAt)
	if err != nil {
		return ms, fmt.Errorf("storage.pool.Create.Scan %w", err)
	}
	return ms, nil

}

// FindByID session.
func (s *Session) FindByID(ctx context.Context, sid string) (models.Session, error) {

	sql := `SELECT * FROM session WHERE id = $1`

	row := s.pool.QueryRow(ctx, sql, sid)

	var ms = models.Session{}

	err := row.Scan(
		&ms.ID,
		&ms.AccountID,
		&ms.Provider,
		&ms.UserAgent,
		&ms.IP,
		&ms.TTL,
		&ms.ExpiresAt,
		&ms.CreatedAt,
	)
	if err != nil {
		return ms, fmt.Errorf("storage.pool.FindByID.Scan %w", err)
	}
	return ms, nil

}

// FindAll accounts sessions by provided account id.
func (s *Session) FindAll(ctx context.Context, aid string) ([]models.Session, error) {

	sql := `SELECT * FROM session WHERE account_id = $1`

	var mss []models.Session

	rows, err := s.pool.Query(ctx, sql, aid)
	if err != nil {
		return mss, fmt.Errorf("storage.pool.FindAll %w", err)
	}

	for rows.Next() {

		ms := models.Session{}

		err := rows.Scan(
			&ms.ID,
			&ms.AccountID,
			&ms.Provider,
			&ms.UserAgent,
			&ms.IP,
			&ms.TTL,
			&ms.ExpiresAt,
			&ms.CreatedAt,
		)
		if err != nil {
			return mss, fmt.Errorf("storage.pool.FildAll.rows.Scan %w", err)
		}

		mss = append(mss, ms)

	}

	return mss, nil

}

// Delete session by id.
func (s *Session) Delete(ctx context.Context, sid string) error {

	sql := `DELETE FROM session WHERE id=$1`

	_, err := s.pool.Query(ctx, sql, sid)
	if err != nil {
		return fmt.Errorf("storage.pool.Delete %w", err)
	}

	return nil
}

// DeleteAll account sessions by provided account id excluding current session.
func (s *Session) DeleteAll(ctx context.Context, aid,sid string) error {

	sql := `DELETE FROM session WHERE account_id = $1 AND id != $2`

	_, err := s.pool.Query(ctx, sql, aid,sid)
	if err != nil {
		return fmt.Errorf("storage.pool.DeleteALl %w", err)
	}

	return nil
}
