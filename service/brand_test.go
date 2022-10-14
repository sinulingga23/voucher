package service

import (
	"context"
	"strings"
	"testing"

	"github.com/sinulingga23/voucher/payload"
	"github.com/sinulingga23/voucher/repository"
	"github.com/sinulingga23/voucher/utility"
)

func TestBrandService_Create_Success(t *testing.T) {
	// Prepare Integration Test
	db, err := utility.ConnectToMySQL()
	if err != nil {
		t.Fatal(err.Error())
	}

	brandRepository := repository.NewBrandRepository(db)
	brandService := NewBrandService(brandRepository)

	wantName := "Unipin"
	wantUrlLogo := "wwww.another-cloud-storage.com/thepath/the-file"
	wantDescription := "Unipin menyediakan semua item gamemu ?"
	wantAddress := "Jakarta"

	createdBrand, err := brandService.Create(context.TODO(), payload.CreateBrand{
		Name: "Unipin",
		UrlLogo: "wwww.another-cloud-storage.com/thepath/the-file",
		Description: "Unipin menyediakan semua item gamemu ?",
		Address: "Jakarta",
	})
	if err != nil {
		t.Fatal(err.Error())
	}

	if strings.Compare(wantName, createdBrand.Name) != 0 {
		t.Fatalf("got %q want %q", createdBrand.Name, wantName)
	}

	if strings.Compare(wantDescription, createdBrand.Description) != 0 {
		t.Fatalf("got %q want %q", createdBrand.Description, wantDescription)
	}

	if strings.Compare(wantUrlLogo, createdBrand.UrlLogo) != 0 {
		t.Fatalf("got %q want %q", createdBrand.UrlLogo, wantUrlLogo)
	}

	if strings.Compare(wantAddress, createdBrand.Address) != 0 {
		t.Fatalf("got %q want %q", createdBrand.Address, wantAddress)
	}
}