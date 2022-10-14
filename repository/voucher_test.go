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
		Name:        "Indomaret",
		UrlLogo:     "www.cloud-service.com/endpoint-brand-misalnya",
		Description: "Indomaret adalah brand yang bergerak di bidang retail.",
		Address:     "Jl. Sudirman, Jakarta",
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
		BrandId:        createdBrand.Id,
		Name:           "Diskon 20 %",
		CostInPoint:    500000,
		Stock:          23,
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
		Name:        "Indomaret",
		UrlLogo:     "www.cloud-service.com/endpoint-brand-misalnya",
		Description: "Indomaret adalah brand yang bergerak di bidang retail.",
		Address:     "Jl. Sudirman, Jakarta",
	})
	if err != nil {
		t.Fatal(err.Error())
	}

	createdVoucher, err := voucherRepository.Create(context.TODO(), domain.CreateVoucher{
		BrandId:        createdBrand.Id,
		Name:           "Diskon 20 %",
		CostInPoint:    500000,
		Stock:          23,
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

func TestVoucherRepository_FindById_Success(t *testing.T) {
	db, err := utility.ConnectToMySQL()
	if err != nil {
		t.Fatal(err.Error())
	}

	brandRepository := NewBrandRepository(db)
	voucherRepository := NewVoucherRepository(db)

	// Prepare new data
	createdBrand, err := brandRepository.Create(context.TODO(), domain.CreateBrand{
		Name:        "Indomaret",
		UrlLogo:     "www.cloud-service.com/endpoint-brand-misalnya",
		Description: "Indomaret adalah brand yang bergerak di bidang retail.",
		Address:     "Jl. Sudirman, Jakarta",
	})
	if err != nil {
		t.Fatal(err.Error())
	}

	wantBrandId := createdBrand.Id
	wantName := "Diskon 100 %"
	wantCostInPoint := 700000
	wantStock := 50
	wantExpirationDate := "2022-10-30"

	createdVoucher, err := voucherRepository.Create(context.TODO(), domain.CreateVoucher{
		BrandId:        createdBrand.Id,
		Name:           "Diskon 100 %",
		CostInPoint:    700000,
		Stock:          50,
		ExpirationDate: "2022-10-30",
	})

	currentVoucher, err := voucherRepository.FindById(context.TODO(), createdVoucher.Id)

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

func TestVoucherRepository_FindAllByBrandId(t *testing.T) {
	// Prepare
	db, err := utility.ConnectToMySQL()
	if err != nil {
		t.Fatal(err.Error())
	}

	brandRepository := NewBrandRepository(db)
	voucherRepository := NewVoucherRepository(db)

	// Add a New Brand for references on voucher
	createdBrand, err := brandRepository.Create(context.TODO(), domain.CreateBrand{
		Name:        "Indomaret",
		UrlLogo:     "www.cloud-service.com/endpoint-brand-misalnya",
		Description: "Indomaret adalah brand yang bergerak di bidang retail.",
		Address:     "Jl. Sudirman, Jakarta",
	})
	if err != nil {
		t.Fatal(err.Error())
	}


	_, err = voucherRepository.Create(context.TODO(), domain.CreateVoucher{
		BrandId:        createdBrand.Id,
		Name:           "Diskon 20 %",
		CostInPoint:    500000,
		Stock:          23,
		ExpirationDate: "2022-10-30",
	})
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = voucherRepository.Create(context.TODO(), domain.CreateVoucher{
		BrandId:        createdBrand.Id,
		Name:           "Diskon 20 + 70 %",
		CostInPoint:    700000,
		Stock:          103,
		ExpirationDate: "2022-11-29",
	})
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = voucherRepository.Create(context.TODO(), domain.CreateVoucher{
		BrandId:        createdBrand.Id,
		Name:           "Diskon 35 %",
		CostInPoint:    200000,
		Stock:          503,
		ExpirationDate: "2022-11-29",
	})
	if err != nil {
		t.Fatal(err.Error())
	}

	wantLength := 3

	vouchers, err := voucherRepository.FindAllByBrandId(context.TODO(), createdBrand.Id)
	if err != nil {
		t.Fatal(err.Error())
	}

	if wantLength != len(vouchers) {
		t.Fatalf("got %v want %v", len(vouchers), wantLength)
	}
}
