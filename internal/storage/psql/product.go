package psql

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/vslitvinov/entebot/internal/models"
)

type Product struct {
	pool *pgxpool.Pool
}

func NewProduct(db *pgxpool.Pool) *Product {
	return &Product{
		pool: db,
	}
}

func (p *Product) Create(ctx context.Context, mp models.Product) (string, error) {

	sql := `INSERT INTO product (name, price, category_id, subscription, level, icon, url, archive)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	RETURNING id`

	row := p.pool.QueryRow(ctx, sql,
		mp.Name,
		mp.Price,
		mp.СategoryID,
		mp.Subscription,
		mp.Level,
		mp.Icon,
		mp.Url,
		mp.ISArchive,
	)
	var id string
	err := row.Scan(&id)
	if err != nil {
		return id, fmt.Errorf("storage.Product.Create.Scan %w", err)
	}
	return id, nil

}

func (p *Product) DeleteByID(ctx context.Context, pid string) error {

	sql := `DELETE FROM product WHERE id=$1`

	_, err := p.pool.Query(ctx, sql, pid)
	if err != nil {
		return fmt.Errorf("storage.Product.Delete.Query %w", err)
	}

	return nil

}

func (p *Product) GetByID(ctx context.Context, pid string) (models.Product, error) {
	
	sql := `SELECT * FROM product WHERE id = $1`

	row := p.pool.QueryRow(ctx, sql, pid)

	var mp = models.Product{}

	err := row.Scan(
		&mp.ID,
		&mp.Name,
		&mp.Price,
		&mp.СategoryID,
		&mp.Subscription,
		&mp.Level,
		&mp.Icon,
		&mp.Url,
		&mp.ISArchive,
	)
	if err != nil {
		return mp, fmt.Errorf("storage.Product.FindByID.Scan %w", err)
	}
	return mp, nil
	
}

func (p *Product) ArchiveByID(ctx context.Context, pid string, archive bool) error {

	sql := `UPDATE product
	SET is_archive='$1'
	WHERE id='$2'`

	_, err := p.pool.Query(ctx, sql, pid,archive)

	if err != nil {
		return fmt.Errorf("storage.Product.Archive.Query %w", err)
	}
	return nil 

}

func (p *Product) GetByCategory(ctx context.Context, category string) ([]models.Product, error) {

	sql := `SELECT * FROM product WHERE category_id = $1`

	var mps []models.Product

	rows, err := p.pool.Query(ctx,sql,category)

	if err != nil {
		return mps, fmt.Errorf("storage.Product.GetByCategory.Query %w", err)
	}

	for rows.Next() {

		 mp := models.Product{}

		err := rows.Scan(
			&mp.ID,
			&mp.Name,
			&mp.Price,
			&mp.СategoryID,
			&mp.Subscription,
			&mp.Level,
			&mp.Icon,
			&mp.Url,
			&mp.ISArchive,
		)
		if err != nil {
			return mps, fmt.Errorf("storage.Product.GetByCategory.Query.Scan %w", err)
		}

		mps = append(mps,mp)
	}

	return mps, nil

}
