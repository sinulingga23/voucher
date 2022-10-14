package service

import (
	"context"
	"strings"

	"github.com/sinulingga23/voucher/domain"
	"github.com/sinulingga23/voucher/payload"
	"github.com/sinulingga23/voucher/repository"
	"github.com/sinulingga23/voucher/utility"
)

type VoucherService struct {
	voucherRepository repository.VoucherRepository
	brandRepository repository.BrandRepository
}

func NewVoucherService(
	voucherRepository repository.VoucherRepository,
	brandRepository repository.BrandRepository) VoucherService {
	return VoucherService{voucherRepository: voucherRepository, brandRepository: brandRepository}
}

func (v VoucherService) Create(ctx context.Context, createVoucher payload.CreateVoucher) (*payload.Voucher, error) {
	if len(strings.Trim(createVoucher.Name, " ")) == 0 {
		return nil, utility.ErrVoucherNameIsEmpty
	}
	createVoucher.Name = strings.Trim(createVoucher.Name, " ")

	if createVoucher.CostInPoint <= 0 {
		return nil, utility.ErrVoucherCostInInPointInvalid
	}

	if createVoucher.Stock <= 0 {
		return nil, utility.ErrVoucherStockInvalid
	}

	if len(strings.Trim(createVoucher.ExpirationDate, " ")) == 0 {
		return nil, utility.ErrVoucherExpiratonDateInvalid
	}

	checkBrand := v.brandRepository.IsExistById(ctx, createVoucher.BrandId)
	if !checkBrand {
		return nil, utility.ErrVoucherBrandIdNotExists
	}

	createdVoucher, err := v.voucherRepository.Create(ctx, domain.CreateVoucher{
		BrandId: createVoucher.BrandId,
		Name: createVoucher.Name,
		CostInPoint: createVoucher.CostInPoint,
		Stock: createVoucher.Stock,
		ExpirationDate: createVoucher.ExpirationDate,
	})
	if err != nil {
		return nil, utility.ErrInternalError
	}

	return &payload.Voucher{
		Id: createdVoucher.Id,
		BrandId: createdVoucher.BrandId,
		Name: createdVoucher.Name,
		CostInPoint: createdVoucher.CostInPoint,
		Stock: createdVoucher.Stock,
		ExpirationDate: createdVoucher.ExpirationDate,
	}, nil
}

func (v VoucherService) FindById(ctx context.Context, id string) (*payload.Voucher, error) {
	checkVoucher := v.voucherRepository.IsExistById(ctx, id)
	if !checkVoucher {
		return nil, utility.ErrVoucherNotExists
	}

	currentVoucher, err := v.voucherRepository.FindById(ctx, id)
	if err != nil {
		return nil, utility.ErrInternalError
	}

	return &payload.Voucher{
		Id: currentVoucher.Id,
		BrandId: currentVoucher.BrandId,
		Name: currentVoucher.Name,
		CostInPoint: currentVoucher.CostInPoint,
		Stock: currentVoucher.Stock,
		ExpirationDate: currentVoucher.ExpirationDate,
	}, nil
}