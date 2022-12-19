package psql

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/vslitvinov/entebot/internal/models"
)

type Bot struct  {
	pool  *pgxpool.Pool
}

func NewBot(db *pgxpool.Pool) *Bot {
	return &Bot{db}
}

func (b *Bot) Create(ctx context.Context, mb models.Bot) (string, error) {
	sql := `INSERT INTO bot (owner, bot_type, exchange, api_key, balance, purchase_amount, number_of_purchases, balance_replenishment, add_profit)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	RETURNING id`

	row := b.pool.QueryRow(ctx, sql,
		mb.Owner,
		mb.BotType,
		mb.Exchange,
		mb.APIKey,
		mb.Balance,
		mb.PurchaseAmount,
		mb.NumberOfPurchases,
		mb.BalanceReplenishment,
		mb.AddProfit,
	)
	var id string
	err := row.Scan(&id)
	if err != nil {
		return id, fmt.Errorf("storage.Bot.Create.Scan %w", err)
	}
	return id, nil
}

func (b *Bot) GetByOwner(ctx context.Context, owid string) ([]models.Bot, error) {

	sql := `SELECT * FROM accounts WHERE owner = $1`

	var mbs []models.Bot

	rows, err := b.pool.Query(ctx, sql, owid)

	if err != nil {
		return mbs, fmt.Errorf("storage.Bot.GetByOwner.Query %w", err)
	}

	for rows.Next() {

		var mb = models.Bot{}

		err := rows.Scan(
			&mb.ID,
			&mb.Owner,
			&mb.BotType,
			&mb.Exchange,
			&mb.APIKey,
			&mb.Balance,
			&mb.PurchaseAmount,
			&mb.NumberOfPurchases,
			&mb.BalanceReplenishment,
			&mb.AddProfit,
		)
		if err != nil {
			return mbs, fmt.Errorf("storage.Bot.GetByOwner.Scan %w", err)
		}

		mbs = append(mbs, mb)

	}

	return mbs, nil

}

func (b *Bot) GetByID(ctx context.Context, bid string) (models.Bot, error) {
	sql := `SELECT * FROM accounts WHERE id = $1`

	row := b.pool.QueryRow(ctx, sql, bid)

	var mb = models.Bot{}

	err := row.Scan(
		&mb.ID,
		&mb.Owner,
		&mb.BotType,
		&mb.Exchange,
		&mb.APIKey,
		&mb.Balance,
		&mb.PurchaseAmount,
		&mb.NumberOfPurchases,
		&mb.BalanceReplenishment,
		&mb.AddProfit,
	)
	if err != nil {
		return mb, fmt.Errorf("storage.pool.FindByID.Scan %w", err)
	}
	return mb, nil
}

func (b *Bot) UpdateByID(ctx context.Context, mb models.Bot)  error {

	sql := `UPDATE bot
	SET 
	owner='$1', 
	bot_type='$2', 
	exchange='$3', 
	api_key='$4', 
	balance='$5', 
	purchase_amount='$6', 
	number_of_purchases='$7', 
	balance_replenishment='$8', 
	add_profit='$9'
	WHERE id='$10'`

	_, err := b.pool.Query(ctx, sql, 
		mb.Owner,
		mb.BotType,
		mb.Exchange,
		mb.APIKey,
		mb.Balance,
		mb.PurchaseAmount,
		mb.NumberOfPurchases,
		mb.BalanceReplenishment,
		mb.AddProfit,
		mb.ID,
	)

	if err != nil {
		return fmt.Errorf("storage.bot.UpdateByID.Query %w", err)
	}

	return nil

}

func (b *Bot) Delete(ctx context.Context, bid string) error {

	sql := `DELETE FROM bot WHERE id=$1`

	_, err := b.pool.Query(ctx, sql, bid)
	if err != nil {
		return fmt.Errorf("storage.Bot.Delete %w", err)
	}

	return nil

}

