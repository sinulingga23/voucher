package service

import (	
	"context"
	"testing"
	"strings"

	"github.com/sinulingga23/voucher/domain"
	"github.com/sinulingga23/voucher/payload"
	"github.com/sinulingga23/voucher/repository"
	"github.com/sinulingga23/voucher/utility"
)

func TestVoucherService_Create_Success(t *testing.T) {
	db, err := utility.ConnectToMySQL()
	if err != nil {
		t.Fatal(err.Error())
	}

	brandRepository := repository.NewBrandRepository(db)
	voucherRepository := repository.NewVoucherRepository(db)
	voucherService := NewVoucherService(voucherRepository, brandRepository)

	// Prepare New Brand for test purpose
	createdBrand, err := brandRepository.Create(context.TODO(), domain.CreateBrand{
		Name: "Indomaret",
		UrlLogo: "www.cloud-service.com/endpoint-brand-misalnya",
		Description: "Indomaret adalah brand yang bergerak di bidang retail.",
		Address: "Jl. Sudirman, Jakarta",
	})

	wantBrandId := createdBrand.Id
	wantName := "Diskon 50 % + 20 %"
	wantCostInPoint := 650000
	wantStock := 50
	wantExpirationDate := "2022-11-29"

	createdVoucher, err := voucherService.Create(context.TODO(), payload.CreateVoucher{
		BrandId: createdBrand.Id,
		Name: "Diskon 50 % + 20 %",
		CostInPoint: 650000,
		Stock: 50,
		ExpirationDate: "2022-11-29",
	})
	if err != nil {
		t.Fatal(err.Error())
	}

	if strings.Compare(wantBrandId, createdBrand.Id) != 0 {
		t.Fatalf("got %q want %q", createdVoucher.BrandId, wantBrandId)
	}

	if strings.Compare(wantName, createdVoucher.Name) != 0 {
		t.Fatalf("got %q want %q", createdVoucher.Name, wantName)
	}

	if wantCostInPoint != createdVoucher.CostInPoint {
		t.Fatalf("got %q want %q", createdVoucher.CostInPoint, wantCostInPoint)
	}

	if wantStock != createdVoucher.Stock {
		t.Fatalf("got %q want %q", createdVoucher.Stock, wantStock)
	}

	if strings.Compare(wantExpirationDate, createdVoucher.ExpirationDate) != 0 {
		t.Fatalf("got %q want %q", createdVoucher.ExpirationDate, wantExpirationDate)
	}
}

func TestVoucherService_FindById_Found(t *testing.T) {
	db, err := utility.ConnectToMySQL()
	if err != nil {
		t.Fatal(err.Error())
	}

	brandRepository := repository.NewBrandRepository(db)
	voucherRepository := repository.NewVoucherRepository(db)
	voucherService := NewVoucherService(voucherRepository, brandRepository)

	// Prepare New Brand for test purpose
	createdBrand, err := brandRepository.Create(context.TODO(), domain.CreateBrand{
		Name: "Indomaret",
		UrlLogo: "www.cloud-service.com/endpoint-brand-misalnya",
		Description: "Indomaret adalah brand yang bergerak di bidang retail.",
		Address: "Jl. Sudirman, Jakarta",
	})

	wantBrandId := createdBrand.Id
	wantName := "Diskon 50 % + 20 %"
	wantCostInPoint := 650000
	wantStock := 50
	wantExpirationDate := "2022-11-29"

	createdVoucher, err := voucherService.Create(context.TODO(), payload.CreateVoucher{
		BrandId: createdBrand.Id,
		Name: "Diskon 50 % + 20 %",
		CostInPoint: 650000,
		Stock: 50,
		ExpirationDate: "2022-11-29",
	})
	if err != nil {
		t.Fatal(err.Error())
	}

	currentVoucher, err := voucherService.FindById(context.TODO(), createdVoucher.Id)
	if err != nil {
		t.Fatalf(err.Error())
	}

	if strings.Compare(wantBrandId, currentVoucher.BrandId) != 0 {
		t.Fatalf("got %q want %q", currentVoucher.BrandId, wantBrandId)
	}

	if strings.Compare(wantName, currentVoucher.Name) != 0 {
		t.Fatalf("got %q want %q", currentVoucher.Name, wantName)
	}

	if wantCostInPoint != currentVoucher.CostInPoint {
		t.Fatalf("got %q want %q", currentVoucher.CostInPoint, wantCostInPoint)
	}

	if wantStock != currentVoucher.Stock {
		t.Fatalf("got %q want %q", currentVoucher.Stock, wantStock)
	}

	if strings.Compare(wantExpirationDate, currentVoucher.ExpirationDate) != 0 {
		t.Fatalf("got %q want %q", currentVoucher.ExpirationDate, wantExpirationDate)
	}
}