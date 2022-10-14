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
	
	id := uuid.New().String()
	result, err := trx.Exec(`INSERT INTO brands (id, name, url_logo, description, address) VALUES (?,?,?,?,?)`,
		id,
		createBrand.Name,
		createBrand.UrlLogo,
		createBrand.Description,
		createBrand.Address)
	if err != nil {
		trx.Rollback()
		return nil, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil || rowsAffected == 0 {
		trx.Rollback()
		return nil, err
	}

	createdBrand := &domain.Brand{}
	row := trx.QueryRow("SELECT id, name, url_logo, description, address FROM brands WHERE id = ?", id)
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

	if err := trx.Commit(); err != nil  {
		return nil, err
	}

	return createdBrand, nil
}

func (b BrandRepository) IsExistById(ctx context.Context, id string) bool {
	check := 0

	row := b.db.QueryRow("SELECT COUNT(id) FROM brands WHERE id = ?", id)

	if err := row.Scan(&check); err != nil {
		return false
	}

	if check == 0 {
		return false
	}

	return true
}