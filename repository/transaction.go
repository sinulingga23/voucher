package repository

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/sinulingga23/voucher/domain"
)

type TransactionRepository struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) TransactionRepository {
	return TransactionRepository{db: db}
}

func (t TransactionRepository) CreateRedemption(
	ctx context.Context,
	createRedemption domain.CreateRedemption) (*domain.Redemption, error) {
	trx, err := t.db.Begin()
	if err != nil {
		return nil, err
	}

	currentVoucher := &domain.Voucher{}
	row := trx.QueryRow("SELECT id, brand_id, name, cost_in_point, stock, expiration_date FROM vouchers WHERE id = ?", createRedemption.VoucherId)
	err = row.Scan(
		&currentVoucher.Id,
		&currentVoucher.BrandId,
		&currentVoucher.Name,
		&currentVoucher.CostInPoint,
		&currentVoucher.Stock,
		&currentVoucher.ExpirationDate,
	)
	if err != nil {
		trx.Rollback()
		return nil, err
	}

	currentVoucher.Stock = currentVoucher.Stock - createRedemption.Qtty
	result, err := trx.Exec("UPDATE vouchers SET stock = ? WHERE id = ?", currentVoucher.Stock, currentVoucher.Id)
	if err != nil {
		trx.Rollback()
		return nil, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil || rowsAffected != 1 {
		trx.Rollback()
		return nil, err
	}

	totalPoint := createRedemption.Qtty * currentVoucher.CostInPoint

	id := uuid.New().String()
	result, err = trx.Exec("INSERT INTO transactions (id, voucher_id, qtty, total_point) VALUES (?, ?, ?, ?)",
		id,
		currentVoucher.Id,
		createRedemption.Qtty,
		totalPoint)
	if err != nil {
		trx.Rollback()
		return nil, err
	}

	if err := trx.Commit(); err != nil {
		return nil, err
	}

	return &domain.Redemption{
		Id: id,
		VoucherId: currentVoucher.Id,
		VoucherName: currentVoucher.Name,
		Qtty: createRedemption.Qtty,
		PointEachVoucher: currentVoucher.CostInPoint,
		TotalPoint: totalPoint,
	}, nil
}

func (t TransactionRepository) FindById(ctx context.Context, id string) (*domain.Redemption, error) {
	trx, err := t.db.Begin()
	if err != nil {
		return nil, err
	}

	currentRedemption := &domain.Redemption{}
	row := trx.QueryRow("SELECT id, voucher_id, qtty, total_point FROM transactions WHERE id = ?", id)
	err = row.Scan(
		&currentRedemption.Id,
		&currentRedemption.VoucherId,
		&currentRedemption.Qtty,
		&currentRedemption.TotalPoint,
	)
	if err != nil {
		trx.Rollback()
		return nil, err
	}

	if err := row.Err(); err != nil {
		trx.Rollback()
		return nil, err
	}

	row = trx.QueryRow("SELECT name, cost_in_point FROM vouchers WHERE id = ?", currentRedemption.VoucherId)
	err = row.Scan(&currentRedemption.VoucherName, &currentRedemption.PointEachVoucher)
	if err != nil {
		trx.Rollback()
		return nil, err
	}

	if err := row.Err(); err != nil {
		trx.Rollback()
		return nil, err
	}

	if err := trx.Commit(); err != nil {
		return nil, err
	}

	return currentRedemption, nil
}