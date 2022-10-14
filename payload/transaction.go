package payload

type (
	Redemption struct {
		VoucherId string `json:"voucherId"`
		VoucherName string `json:"voucherName"`
		Qtty int `json:"qtty"`
		PointEachVoucher int `json:"pointEachVoucher"`
		TotalPoint int `json:"totalPoint"`
	}

	CreateRedemption struct {
		VoucherId string `json:"voucherId"`
		Qtty int `json:"qtty"`
	}
)