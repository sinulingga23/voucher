package handler

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sinulingga23/voucher/payload"
	"github.com/sinulingga23/voucher/utility"
)

func (h Handler) CreateRedemption(c *gin.Context) {
	createRedempion := payload.CreateRedemption{}

	if err := c.Bind(&createRedempion); err != nil {
		c.JSON(http.StatusBadRequest, struct {
			Message string `json:"message"`
		}{Message: err.Error()})
		return
	}

	redemption, err := h.transactionService.CreateRedemption(context.TODO(), createRedempion)
	if err != nil {
		if err == utility.ErrVoucherNotExists {
			c.JSON(http.StatusNotFound, struct {
				Message string `json:"message"`
			}{Message: err.Error()})
			return
		}

		if err == utility.ErrVoucherIsNotEnough {
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
		Data payload.Redemption`json:"data"`
	}{Data: *redemption})
	return
}

func (h Handler) FindRedemptionById(c *gin.Context) {
	transactionId := c.Request.URL.Query().Get("transactionId")

	redemption, err := h.transactionService.FindRedemptionById(context.TODO(), transactionId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, struct {
			Message string `json:"message"`
		}{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, struct {
		Data payload.Redemption`json:"data"`
	}{Data: *redemption})
	return
}