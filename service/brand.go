package service

import (
	"context"
	"strings"

	"github.com/sinulingga23/voucher/domain"
	"github.com/sinulingga23/voucher/payload"
	"github.com/sinulingga23/voucher/repository"
	"github.com/sinulingga23/voucher/utility"
)

type BrandService struct {
	brandRepository repository.BrandRepository
}

func NewBrandService(brandRepository repository.BrandRepository) BrandService {
	return BrandService{brandRepository: brandRepository}
}


func (b BrandService) Create(ctx context.Context, createBrand payload.CreateBrand) (*payload.Brand, error) {
	if len(strings.Trim(createBrand.Name, " ")) == 0 {
		return nil, utility.ErrBrandNameIsEmpty
	}
	createBrand.Name = strings.Trim(createBrand.Name, " ")

	if len(strings.Trim(createBrand.UrlLogo, " ")) == 0 {
		return nil, utility.ErrBrandLogoIsEmpty
	}
	createBrand.UrlLogo = strings.Trim(createBrand.UrlLogo, " ")

	createdBrand, err := b.brandRepository.Create(ctx, domain.CreateBrand{
		Name: createBrand.Name,
		UrlLogo: createBrand.UrlLogo,
		Description: createBrand.Description,
		Address: createBrand.Address,
	})
	if err != nil {
		return nil, utility.ErrInternalError
	}

	return &payload.Brand{
		Id: createdBrand.Id,
		Name: createdBrand.Name,
		UrlLogo: createdBrand.UrlLogo,
		Description: createdBrand.Description,
		Address: createdBrand.Address,
	}, nil
}