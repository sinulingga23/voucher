package domain

type (
	Redemption struct {
		Id string
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