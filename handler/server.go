package handler

import "github.com/sinulingga23/voucher/service"

type Handler struct {
	brandService service.BrandService
	voucherService service.VoucherService
	transactionService service.TransactionService
}

func NewHandler(
		brandService service.BrandService,
		voucherService service.VoucherService,
		transactionService service.TransactionService) Handler {
	return Handler{brandService:  brandService, voucherService: voucherService, transactionService: transactionService}
}