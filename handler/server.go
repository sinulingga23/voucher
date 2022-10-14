package handler

import "github.com/sinulingga23/voucher/service"

type Handler struct {
	brandService service.BrandService
	voucherService service.VoucherService
}

func NewHandler(
		brandService service.BrandService,
		voucherService service.VoucherService) Handler {
	return Handler{brandService:  brandService, voucherService: voucherService}
}