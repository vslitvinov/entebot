package storage

import (
	"context"
	"fmt"
	"sync"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/vslitvinov/entebot/internal/models"
	"github.com/vslitvinov/entebot/internal/storage/psql"
)

type PSQLProductStorage interface {
	Create(ctx context.Context, mp models.Product) (string, error) 
	DeleteByID(ctx context.Context, pid string) error
	GetByID(ctx context.Context, pid string) (models.Product, error)
	ArchiveByID(ctx context.Context, pid string, archive bool) error
	GetByCategory(ctx context.Context, category string) ([]models.Product, error) 
}

type Product struct {
	db PSQLProductStorage
	cache sync.Map
	isCache bool
}

func NewProduct(db *pgxpool.Pool, isCache bool) *Product {
	return &Product{
		db: psql.NewProduct(db),
		cache: sync.Map{},
		isCache: isCache,
	}
}


func (p *Product) Create(ctx context.Context, mp models.Product) (string, error) {

	var pid string

	pid,err := p.db.Create(ctx,mp)

	if err != nil {
		return pid, fmt.Errorf("storage.Product.Create.db.Create %w", err)
	}

	if p.isCache {

		mp.ID = pid 
		p.cache.Store(pid,mp)

	}

	return pid, nil

}

func (p *Product) DeleteByID(ctx context.Context, pid string) error {

	err := p.db.DeleteByID(ctx,pid)
	if err != nil {
		return fmt.Errorf("storage.Product.DeleteByID.db.DeleteByID %w", err)
	}


	if p.isCache {
		p.cache.Delete(pid)
	}

	return nil 

}

func (p *Product) GetByID(ctx context.Context, pid string) (models.Product, error) {

	var mp models.Product

	if p.isCache {
		mpC,ok := p.cache.Load(pid)
		if !ok {
			return mp, fmt.Errorf("storage.Product.GetByID.cache.Load %w", nil)
		}
		return mpC.(models.Product), nil
	} else {
		mp, err := p.db.GetByID(ctx,pid)
		if err != nil {
			return mp, fmt.Errorf("storage.Product.GetByID.db.GetByID %w", err)
		}
	}

	return mp, nil

}

func (p *Product) ArchiveByID(ctx context.Context, pid string, archive bool) error {

	err := p.db.ArchiveByID(ctx,pid,archive)
	if err != nil {
		return err 
	}

	if p.isCache {
		mp, ok := p.cache.Load(pid)
		if !ok {
			return fmt.Errorf("storage.Product.ArchiveByID.cache.Load %w", nil)
		}
		nmp := mp.(models.Product)
		nmp.ISArchive = archive
		p.cache.Store(pid, nmp)
	}

	return nil

}

func (p *Product) GetByCategory(ctx context.Context, category string) ([]models.Product, error) {

	var mps []models.Product 

	if p.isCache {
		p.cache.Range(func(key any, value any) bool {

			if value.(models.Product).СategoryID == category {
				mps = append(mps,value.(models.Product))
			}

			return true
		})
	} else {
		mps, err := p.db.GetByCategory(ctx,category)
		if err != nil {
			return mps, fmt.Errorf("storage.Product.GetByCategory.db.GetByCategory %w", nil)
		}
	}

	return mps, nil

}
