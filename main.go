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
	brandService := service.NewBrandService(brandRepository)
	handler := handler.NewHandler(brandService)

	router.POST("/brand", handler.CreateBrand)

	http.ListenAndServe(":8080", router)	
}