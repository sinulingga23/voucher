package repository

import (
	"context"
	"strings"
	"testing"

	"github.com/sinulingga23/voucher/domain"
	"github.com/sinulingga23/voucher/utility"
)

func TestBrandRepository_Create_Success(t *testing.T) {
	db, err := utility.ConnectToMySQL()
	if err != nil {
		t.Fatal(err.Error())
	}

	wantName := "Indomaret"
	wantUrlLogo := "www.cloud-service.com/endpoint-brand-misalnya"
	wantDescription := "Indomaret adalah brand yang bergerak di bidang retail."
	wantAddress := "Jl. Sudirman, Jakarta"

	brandRepository := NewBrandRepository(db)

	createdBrand, err := brandRepository.Create(context.TODO(), domain.CreateBrand{
		Name: "Indomaret",
		UrlLogo: "www.cloud-service.com/endpoint-brand-misalnya",
		Description: "Indomaret adalah brand yang bergerak di bidang retail.",
		Address: "Jl. Sudirman, Jakarta",
	})
	if err != nil {
		t.Fatal(err.Error())
	}

	if strings.Compare(wantName, createdBrand.Name) != 0 {
		t.Fatalf("got %q want %q\n", createdBrand.Name, wantName)
	}

	if strings.Compare(wantUrlLogo, createdBrand.UrlLogo) != 0 {
		t.Fatalf("got %q want %q\n", createdBrand.UrlLogo, wantUrlLogo)
	}

	if strings.Compare(wantDescription, createdBrand.Description) != 0 {
		t.Fatalf("got %q want %q\n", createdBrand.Description, wantDescription)
	}

	if strings.Compare(wantAddress, createdBrand.Address) != 0 {
		t.Fatalf("got %q want %q\n", createdBrand.Address, wantAddress)
	}
}