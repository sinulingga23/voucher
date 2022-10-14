package domain

type (
	Voucher struct {
		Id string
		BrandId string
		Name string
		CostInPoint int
		Stock int
		ExpirationDate string
	}

	CreateVoucher struct {
		BrandId string
		Name string
		CostInPoint int
		Stock int
		ExpirationDate string
	}
)