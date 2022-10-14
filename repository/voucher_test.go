package repository

import (
	"context"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/sinulingga23/voucher/domain"
	"github.com/sinulingga23/voucher/utility"
)

func TestVoucherRepository_Create_Success(t *testing.T) {
	// Prepare
	db, err := utility.ConnectToMySQL()
	if err != nil {
		t.Fatal(err.Error())
	}


	brandRepository := NewBrandRepository(db)
	voucherRepository := NewVoucherRepository(db)

	// Add a New Brand for references on voucher
	createdBrand, err := brandRepository.Create(context.TODO(), domain.CreateBrand{
		Name: "Indomaret",
		UrlLogo: "www.cloud-service.com/endpoint-brand-misalnya",
		Description: "Indomaret adalah brand yang bergerak di bidang retail.",
		Address: "Jl. Sudirman, Jakarta",
	})
	if err != nil {
		t.Fatal(err.Error())
	}

	wantBrandId := createdBrand.Id
	wantName := "Diskon 20 %"
	wantCostInPoint := 500000
	wantStock := 23
	wantExpirationDate := "2022-10-30"

	createdVoucher, err := voucherRepository.Create(context.TODO(), domain.CreateVoucher{
		BrandId: createdBrand.Id,
		Name: "Diskon 20 %",
		CostInPoint: 500000,
		Stock: 23,
		ExpirationDate: "2022-10-30",
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

func TestVoucherRepository_IsExistsById_Exists(t *testing.T) {
	db, err := utility.ConnectToMySQL()
	if err != nil {
		t.Fatal(err.Error())
	}

	brandRepository := NewBrandRepository(db)
	voucherRepository := NewVoucherRepository(db)
	
	// Prepare new data
	createdBrand, err := brandRepository.Create(context.TODO(), domain.CreateBrand{
		Name: "Indomaret",
		UrlLogo: "www.cloud-service.com/endpoint-brand-misalnya",
		Description: "Indomaret adalah brand yang bergerak di bidang retail.",
		Address: "Jl. Sudirman, Jakarta",
	})
	if err != nil {
		t.Fatal(err.Error())
	}

	createdVoucher, err := voucherRepository.Create(context.TODO(), domain.CreateVoucher{
		BrandId: createdBrand.Id,
		Name: "Diskon 20 %",
		CostInPoint: 500000,
		Stock: 23,
		ExpirationDate: "2022-10-30",
	})

	wantResult := true

	result := voucherRepository.IsExistById(context.TODO(), createdVoucher.Id)

	if wantResult != result {
		t.Fatalf("got %v want %v\n", result, wantResult)
	}
}

func TestVoucherRepository_IsExistsById_NotExists(t *testing.T) {
	db, err := utility.ConnectToMySQL()
	if err != nil {
		t.Fatal(err.Error())
	}

	voucherRepository := NewVoucherRepository(db)

	wantResult := false

	result := voucherRepository.IsExistById(context.TODO(), uuid.New().String())

	if wantResult != result {
		t.Fatalf("got %v want %v\n", result, wantResult)
	}
}