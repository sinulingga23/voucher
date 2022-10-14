package domain

type (
	Redemption struct {
		VoucherId string
		VoucherName string
		Qtty int
		PointEachVoucher int
		TotalPoint int
	}

	CreateRedemption struct {
		VoucherId string
		Qtty int
	}
)