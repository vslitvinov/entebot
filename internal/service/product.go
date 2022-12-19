package service

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/vslitvinov/entebot/internal/models"
	"github.com/vslitvinov/entebot/internal/storage"
)

type ProductStorage interface {
	Create(ctx context.Context, mp models.Product) (string, error) 
	DeleteByID(ctx context.Context, pid string) error 
	GetByID(ctx context.Context, pid string) (models.Product, error) 
	ArchiveByID(ctx context.Context, pid string, archive bool) error 
	GetByCategory(ctx context.Context, category string) ([]models.Product, error) 
}


type Product struct {
	storage ProductStorage
}

func NewProduct(db *pgxpool.Pool) *Product{

	return &Product {
		storage: storage.NewProduct(db,true),
	}

}

func (p *Product) Create(ctx context.Context, mp models.Product) (string, error) {

	var pid string

	pid, err := p.storage.Create(ctx,mp)
	if err != nil {
		return pid, fmt.Errorf("service.Product.Create.storage.Create %w", err)
	}

	return pid, nil

}

func (p *Product) DeleteByID(ctx context.Context, pid string) error {
	err := p.storage.DeleteByID(ctx,pid)
	if err != nil {
		return err 
	}
	return nil
}

func (p *Product) GetByID(ctx context.Context, pid string) (models.Product, error) {

	var mp models.Product 

	mp, err := p.storage.GetByID(ctx, pid)
	if err != nil {
		return mp, err 
	}

	return mp, err

}

func (p *Product) ArchiveByID(ctx context.Context, pid string, archive bool) error {

	err := p.storage.ArchiveByID(ctx,pid,archive)
	if err != nil {
		return err
	}

	return nil

}

func (p *Product) GetByCategory(ctx context.Context, category string) ([]models.Product, error) {

	var mps []models.Product

	mps,err := p.storage.GetByCategory(ctx,category)
	if err != nil {
		return mps, fmt.Errorf("service.Product.GetByCategory.storage.GetByCategory %w", err)
	}

	return mps, nil

}


