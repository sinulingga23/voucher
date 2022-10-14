package service

import (
	"context"

	"github.com/sinulingga23/voucher/domain"
	"github.com/sinulingga23/voucher/payload"
	"github.com/sinulingga23/voucher/repository"
	"github.com/sinulingga23/voucher/utility"
)

type TransactionService struct {
	transactionRepository repository.TransactionRepository
	voucherRepository repository.VoucherRepository
}

func NewTransactionService(
	transactionRepository repository.TransactionRepository,
	voucherRepository repository.VoucherRepository) TransactionService {
	return TransactionService{transactionRepository: transactionRepository, voucherRepository: voucherRepository}
}

func (t TransactionService) CreateRedemption(
	ctx context.Context, 
	createRedemption payload.CreateRedemption) (*payload.Redemption, error) {

	checkVoucher := t.voucherRepository.IsExistById(ctx, createRedemption.VoucherId)
	if !checkVoucher {
		return nil, utility.ErrVoucherNotExists
	}

	currentVoucher, err := t.voucherRepository.FindById(ctx, createRedemption.VoucherId)
	if err != nil {
		return nil, err
	}

	if currentVoucher.Stock < createRedemption.Qtty {
		return nil, utility.ErrVoucherIsNotEnough
	}

	redemption, err := t.transactionRepository.CreateRedemption(ctx, domain.CreateRedemption{
		VoucherId: createRedemption.VoucherId,
		Qtty: createRedemption.Qtty,
	})
	if err != nil {
		return nil, err
	}

	return &payload.Redemption{
		VoucherId: redemption.VoucherId,
		VoucherName: redemption.VoucherName,
		Qtty: redemption.Qtty,
		PointEachVoucher: redemption.PointEachVoucher,
		TotalPoint: redemption.TotalPoint,
	}, nil
}