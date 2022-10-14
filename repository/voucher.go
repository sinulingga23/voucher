package repository

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/sinulingga23/voucher/domain"
)


type VoucherRepository struct {
	db *sql.DB
}

func NewVoucherRepository(db *sql.DB) VoucherRepository {
	return VoucherRepository{db: db}
}

func (v VoucherRepository) Create(ctx context.Context, createVoucher domain.CreateVoucher) (*domain.Voucher, error) {
	trx, err := v.db.Begin()
	if err != nil {
		return nil, err
	}

	id := uuid.New().String()
	result, err  := trx.Exec("INSERT INTO vouchers (id, brand_id, name, cost_in_point, stock, expiration_date) VALUES (?,?,?,?,?,?)",
		id,
		createVoucher.BrandId,
		createVoucher.Name,
		createVoucher.CostInPoint,
		createVoucher.Stock,
		createVoucher.ExpirationDate)
	if err != nil {
		trx.Rollback()
		return nil, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil || rowsAffected == 0 {
		trx.Rollback()
		return nil, err
	}

	createdVoucher := &domain.Voucher{}
	row := trx.QueryRow("SELECT id, brand_id, name, cost_in_point, stock, expiration_date FROM vouchers WHERE id = ?", id)
	err = row.Scan(
		&createdVoucher.Id,
		&createdVoucher.BrandId,
		&createdVoucher.Name,
		&createdVoucher.CostInPoint,
		&createdVoucher.Stock,
		&createdVoucher.ExpirationDate,
	)
	if err != nil {
		trx.Rollback()
		return nil, err
	}

	if err := trx.Commit(); err != nil {
		return nil, err
	}

	return createdVoucher, nil
}

func (b VoucherRepository) IsExistById(ctx context.Context, id string) bool {
	check := 0

	row := b.db.QueryRow("SELECT COUNT(id) FROM vouchers WHERE id = ?", id)

	if err := row.Scan(&check); err != nil {
		return false
	}

	if check == 0 {
		return false
	}

	return true
}

func (v VoucherRepository) FindById(ctx context.Context, id string) (*domain.Voucher, error) {
	createdVoucher := &domain.Voucher{}
	row := v.db.QueryRow("SELECT id, brand_id, name, cost_in_point, stock, expiration_date FROM vouchers WHERE id = ?", id)
	err := row.Scan(
		&createdVoucher.Id,
		&createdVoucher.BrandId,
		&createdVoucher.Name,
		&createdVoucher.CostInPoint,
		&createdVoucher.Stock,
		&createdVoucher.ExpirationDate,
	)
	if err != nil {
		return nil, err
	}
	return createdVoucher, nil
} 