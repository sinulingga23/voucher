package handler

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sinulingga23/voucher/payload"
	"github.com/sinulingga23/voucher/utility"
)

func (h Handler) CreateBrand(c *gin.Context) {
	createBrand := payload.CreateBrand{}
	
	if err := c.Bind(&createBrand); err != nil {
		c.JSON(http.StatusBadRequest, struct {
			Message string `json:"message"`
		}{Message: err.Error()})
		return
	}

	createdBrand, err := h.brandService.Create(context.TODO(), createBrand)
	if err != nil {
		if err == utility.ErrBrandNameIsEmpty || err == utility.ErrBrandLogoIsEmpty {
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
		Data payload.Brand `json:"data"`
	}{Data: *createdBrand})
	return
}