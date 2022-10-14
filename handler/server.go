package handler

import "github.com/sinulingga23/voucher/service"

type Handler struct {
	brandService service.BrandService
}

func NewHandler(brandService service.BrandService) Handler {
	return Handler{brandService:  brandService}
}