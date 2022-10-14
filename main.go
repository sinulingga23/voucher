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

	brandService := service.NewBrandService(brandRepository)
	voucherService := service.NewVoucherService(voucherRepository, brandRepository)

	handler := handler.NewHandler(brandService, voucherService)

	router.POST("/brand", handler.CreateBrand)
	router.POST("/voucher", handler.CreateVoucher)

	http.ListenAndServe(":8080", router)	
}