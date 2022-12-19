package psql

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/vslitvinov/entebot/internal/models"
)

type Account struct {
	pool  *pgxpool.Pool
}

func NewAccount(db *pgxpool.Pool) *Account {
	return &Account{db}
}

func (a *Account) Create(ctx context.Context, ac models.Accounty) (string, error) {

	sql := `INSERT INTO accounts (firstname, lastname, email, phone, username, password)
	VALUES ($1, $2, $3, $4, $5, $6)
	RETURNING id`

	row := a.pool.QueryRow(ctx, sql,
		ac.FirstName,
		ac.LastName,
		ac.Email,
		ac.Phone,
		ac.Username,
		ac.Password,
	)
	var id string
	err := row.Scan(&id)
	if err != nil {
		return id, fmt.Errorf("storage.pool.Create.Scan %w", err)
	}
	return id, nil
}

func (a *Account) FindByID(ctx context.Context, aid string) (models.Accounty, error) {
	sql := `SELECT * FROM accounts WHERE id = $1`

	row := a.pool.QueryRow(ctx, sql, aid)

	var ac = models.Accounty{}

	err := row.Scan(
		&ac.ID,
		&ac.FirstName,
		&ac.LastName,
		&ac.Email,
		&ac.Phone,
		&ac.Password,
		&ac.CreatedAt,
		&ac.UpdatedAt,
		&ac.ISArchive,
		&ac.ISVerified,
	)
	if err != nil {
		return ac, fmt.Errorf("storage.pool.FindByID.Scan %w", err)
	}
	return ac, nil

}

func (a *Account) FindByEmail(ctx context.Context, email string) (models.Accounty, error) {
	sql := `SELECT * FROM accounts WHERE email = $1`

	row := a.pool.QueryRow(ctx, sql, email)

	var ac = models.Accounty{}

	err := row.Scan(
		&ac.ID,
		&ac.FirstName,
		&ac.LastName,
		&ac.Email,
		&ac.Phone,
		&ac.Username,
		&ac.Password,
		&ac.CreatedAt,
		&ac.UpdatedAt,
		&ac.ISArchive,
		&ac.ISVerified,
	)
	if err != nil {
		return ac, fmt.Errorf("storage.pool.FindByEmail.Scan %w", err)
	}
	return ac, nil

}

func (a *Account) Verified(ctx context.Context, aid string, verified bool) error {
	sql := `UPDATE accounts
	SET is_verified='$1', updated_at=CURRENT_TIMESTAMP
	WHERE id='$2'`

	_, err := a.pool.Query(ctx, sql, aid,verified)

	if err != nil {
		return fmt.Errorf("storage.pool.Verified.Query %w", err)
	}

	return nil
}

func (a *Account) Archive(ctx context.Context, aid string, archive bool) error {
	sql := `UPDATE accounts
	SET is_archive='$1', updated_at=CURRENT_TIMESTAMP
	WHERE id='$2'`

	_, err := a.pool.Query(ctx, sql, aid,archive)

	if err != nil {
		return fmt.Errorf("storage.pool.Archive.Query %w", err)
	}
	return nil 

}

