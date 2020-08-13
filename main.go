package main

import (
	"fmt"
	"log"

	"github.com/bhambri94/ecommerce-app/configs"
	"github.com/bhambri94/ecommerce-app/homedepot"
	"github.com/bhambri94/ecommerce-app/sheets"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
)

var (
	logger, _ = zap.NewProduction()
	sugar     = logger.Sugar()
)

func main() {
	sugar.Infof("starting ecommerce manager app server...")
	defer logger.Sync() // flushes buffer, if any

	router := fasthttprouter.New()
	router.GET("/v1/homedepot/multiplestore", handleMultipleStore)
	router.GET("/v1/homedepot/multipleproduct", handleMultipleProduct)
	log.Fatal(fasthttp.ListenAndServe(":8081", router.Handler))
}

func handleMultipleStore(ctx *fasthttp.RequestCtx) {
	sugar.Infof("calling ecommerce manager api with multiple stores!")
	configs.SetConfig()
	requestBody := sheets.BatchGet(configs.Configurations.MultipleStoreRequestSheetNameWithRange)
	fmt.Println(requestBody)
	finalValues := homedepot.GetHomeDepotMultipleStoreData(requestBody)
	sheets.ClearSheet(configs.Configurations.ClearMultipleStoreResponseSheetNameWithRange)
	sheets.BatchWrite(configs.Configurations.MultipleStoreResponseSheetNameWithRange, finalValues)
	ctx.Response.Header.Set("Content-Type", "application/json")
	ctx.Response.SetStatusCode(200)
	ctx.SetBody([]byte("Sheet has been updated: https://docs.google.com/spreadsheets/d/1Dr1qTo-o89MtKEBKF4QfLl0Esw8HJNE_cj5_dcZdzEk/edit#gid=2009338596"))
}

func handleMultipleProduct(ctx *fasthttp.RequestCtx) {
	sugar.Infof("calling ecommerce manager api with multiple products!")
	configs.SetConfig()
	requestBody := sheets.BatchGet(configs.Configurations.MultipleProductRequestSheetNameWithRange)
	fmt.Println(requestBody)
	finalValues := homedepot.GetHomeDepotMultipleProductData(requestBody)
	sheets.ClearSheet(configs.Configurations.ClearMultipleProductResponseSheetNameWithRange)
	sheets.BatchWrite(configs.Configurations.MultipleProductResponseSheetNameWithRange, finalValues)
	ctx.Response.Header.Set("Content-Type", "application/json")
	ctx.Response.SetStatusCode(200)

	ctx.SetBody([]byte("Sheet has been updated: https://docs.google.com/spreadsheets/d/1Dr1qTo-o89MtKEBKF4QfLl0Esw8HJNE_cj5_dcZdzEk/edit#gid=2009338596"))
}
