package payload

type (
	Voucher struct {
		Id string `json:"id"`
		BrandId string `json:"brandId"`
		Name string `json:"name"`
		CostInPoint int `json:"costInPoint"`
		Stock int `json:"stock"`
		ExpirationDate string `json:"expirationDate"`
	}

	CreateVoucher struct {
		BrandId string `json:"brandId"`
		Name string `json:"name"`
		CostInPoint int `json:"costInPoint"`
		Stock int `json:"stock"`
		ExpirationDate string `json:"expirationDate"`
	}
)