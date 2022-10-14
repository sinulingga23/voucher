package repository

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/sinulingga23/voucher/domain"
)

type BrandRepository struct {
	db *sql.DB
}

func NewBrandRepository(db *sql.DB) BrandRepository {
	return BrandRepository{db: db}
}

func (b BrandRepository) Create(ctx context.Context,  createBrand domain.CreateBrand) (*domain.Brand, error) {
	trx, err := b.db.Begin()
	if err != nil {
		return nil, err
	}
	defer trx.Commit()
	
	result, err := trx.Exec(`INSERT INTO brands (id, name, url_logo, description, address) VALUES (?,?,?,?,?)`,
		uuid.New().String(),
		createBrand.Name,
		createBrand.UrlLogo,
		createBrand.Description,
		createBrand.Address)
	if err != nil {
		trx.Rollback()
		return nil, err
	}

	lastInsertedId, err := result.LastInsertId()
	if err != nil {
		trx.Rollback()
		return nil, err
	}

	createdBrand := &domain.Brand{}
	row := trx.QueryRow("SELECT id, name, url_logo, description, address FROM brands WHERE id = ?", lastInsertedId)
	err = row.Scan(
		&createdBrand.Id,
		&createdBrand.Name,
		&createdBrand.UrlLogo,
		&createdBrand.Description, 
		&createdBrand.Address)
	if err != nil {
		trx.Rollback()
		return nil, err
	}

	return createdBrand, nil
}