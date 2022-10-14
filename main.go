package main

import (
	"net/http"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/sinulingga23/voucher/handler"
	"github.com/sinulingga23/voucher/repository"
	"github.com/sinulingga23/voucher/service"
	"github.com/sinulingga23/voucher/utility"
)

func main() {
	router := gin.Default()

	db, err := utility.ConnectToMySQL()
	if err != nil {
		log.Fatal(err.Error())
	}

	brandRepository := repository.NewBrandRepository(db)
	voucherRepository := repository.NewVoucherRepository(db)
	transactionRepository := repository.NewTransactionRepository(db)

	brandService := service.NewBrandService(brandRepository)
	voucherService := service.NewVoucherService(voucherRepository, brandRepository)
	transactionService := service.NewTransactionService(transactionRepository, voucherRepository)

	handler := handler.NewHandler(brandService, voucherService, transactionService)

	router.POST("/brand", handler.CreateBrand)
	router.POST("/voucher", handler.CreateVoucher)
	router.POST("/transaction/redemption", handler.CreateRedemption)

	router.GET("/voucher", handler.FindVoucherById)
	router.GET("/voucher/brand", handler.FindAllVoucherByBrandId)
	router.GET("/transaction/redemption", handler.FindRedemptionById)

	http.ListenAndServe(":8080", router)	
}