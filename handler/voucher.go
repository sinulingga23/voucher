package handler

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sinulingga23/voucher/payload"
	"github.com/sinulingga23/voucher/utility"
)

func (h Handler) CreateVoucher(c *gin.Context) {
	createVoucher := payload.CreateVoucher{}
	
	if err := c.Bind(&createVoucher); err != nil {
		c.JSON(http.StatusBadRequest, struct {
			Message string `json:"message"`
		}{Message: err.Error()})
		return
	}

	createdVoucher, err := h.voucherService.Create(context.TODO(), createVoucher)
	if err != nil {
		if err == utility.ErrVoucherNameIsEmpty || err == utility.ErrVoucherCostInInPointInvalid ||
			err == utility.ErrVoucherStockInvalid || err == utility.ErrVoucherBrandIdNotExists ||
			err == utility.ErrVoucherExpiratonDateInvalid {
			
			c.JSON(http.StatusBadRequest, struct {
				Message string `json:"message"`
			}{Message: err.Error()})
			return
		}

		c.JSON(http.StatusInternalServerError, struct {
			Message string `json:"message"`
		}{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, struct {
		Data payload.Voucher `json:"data"`
	}{Data: *createdVoucher})
	return
}

func (h Handler) FindVoucherById(c *gin.Context) {
	id := c.Request.URL.Query().Get("id")
	
	currentVoucher, err := h.voucherService.FindById(context.TODO(), id)
	if err != nil {
		if err == utility.ErrVoucherNotExists {
			c.JSON(http.StatusNotFound, struct {
				Message string `json:"message"`
			}{Message: err.Error()})
			return
		}

		c.JSON(http.StatusInternalServerError, struct {
			Message string `json:"message"`
		}{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, struct {
		Data payload.Voucher `json:"data"`
	}{Data: *currentVoucher})
	return
}