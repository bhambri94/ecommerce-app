package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

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
	router.GET("/v1/homedepot/multipleproduct", handleMultipleProduct)
	router.GET("/v1/homedepot/multipleproduct/output=:outputType", handleMultipleProduct)
	log.Fatal(fasthttp.ListenAndServe(":3001", router.Handler))
}

func handleMultipleProduct(ctx *fasthttp.RequestCtx) {
	sugar.Infof("calling ecommerce manager api with multiple products!")
	configs.SetConfig()
	outputType := ctx.UserValue("outputType")
	if outputType != nil {
		sugar.Infof("Output type of api is:= " + outputType.(string))
	}
	productList, err := sheets.BatchGet(configs.Configurations.ProductRequestSheetNameWithRange)
	if err != nil {
		ctx.Response.Header.Set("Content-Type", "application/json")
		ctx.Response.SetStatusCode(200)
		ctx.SetBody([]byte("Failed! Unable to Read Product List from Google Sheet Request"))
		sugar.Infof("calling ecommerce manager api failure!")
	}
	storeList, err := sheets.BatchGet(configs.Configurations.StoreRequestSheetNameWithRange)
	if err != nil {
		ctx.Response.Header.Set("Content-Type", "application/json")
		ctx.Response.SetStatusCode(200)
		ctx.SetBody([]byte("Failed! Unable to Read Store List from Google Sheet Request"))
		sugar.Infof("calling ecommerce manager api failure!")
	}
	fmt.Println(productList)
	fmt.Println(storeList)
	finalValues, err := homedepot.GetHomeDepotMultipleProductData(productList, storeList)
	if err == nil {
		if outputType == nil {
			CSVName := "HomeDepotCSV.csv"
			f, err := os.Create(CSVName)
			if err != nil {
				log.Fatal(err)
			}
			writer := csv.NewWriter(f)
			defer writer.Flush()

			stringfinalValues := make([][]string, len(finalValues)+5)
			header := []string{"StoreId", "Item Id", "Category", "Store Quantity", "Online Quantity", "SpecialPrice", "Upc", "ProductLabel", "BrandName", "IsLimitedQuantity", "OriginalPrice", "DollarOff", "AvailabilityType", "BossEstimatedShippingEndDate", "SthEstimatedShippingStartDate", "SthEstimatedShippingEndDate", "FreeShippingThreshold", "ExcludedShipStates", "FreeShippingMessage", "BossEstimatedShippingStartDate", "WebURL", "TotalReviews", "AverageRating", "Description", "BuyOnlineShipToStoreEligible", "IsTopSeller", "BuyOnlinePickupInStoreEligible", "ModelNumber", "VendorNumber", "AttributeValue", "DimensionName", "DimensionValue Name", "DiscountEndDate", "PromoLongDescription", "DiscountStartDate", "Image Links"}
			writer.Write(header)
			i := 0
			for i < len(finalValues) {
				for _, value := range finalValues[i] {
					a := fmt.Sprintf("%v", value)
					stringfinalValues[i] = append(stringfinalValues[i], a)
				}
				writer.Write(stringfinalValues[i])
				i++
			}
			ctx.Response.SetStatusCode(200)
			ctx.Response.Header.Set("Content-Type", "text/csv")
			ctx.Response.Header.Set("Content-Disposition", "attachment;filename="+"HomeDepotCSV"+strconv.Itoa(int(time.Now().Unix()))+".csv")
			ctx.SendFile(CSVName)
		} else {
			err := sheets.ClearSheet(configs.Configurations.ClearMultipleProductResponseSheetNameWithRange)
			if err != nil {
				ctx.Response.Header.Set("Content-Type", "application/json")
				ctx.Response.SetStatusCode(200)
				ctx.SetBody([]byte("Failed! Unable to Read Store List from Google Sheet Request"))
				sugar.Infof("calling ecommerce manager api failure!")
			}
			err = sheets.BatchWrite(configs.Configurations.MultipleProductResponseSheetNameWithRange, finalValues)
			if err != nil {
				ctx.Response.Header.Set("Content-Type", "application/json")
				ctx.Response.SetStatusCode(200)
				ctx.SetBody([]byte("Failed! Unable to Read Store List from Google Sheet Request"))
				sugar.Infof("calling ecommerce manager api failure!")
			}
			ctx.Response.Header.Set("Content-Type", "application/json")
			ctx.Response.SetStatusCode(200)
			ctx.SetBody([]byte("Success! Sheet has been updated: https://docs.google.com/spreadsheets/d/1Dr1qTo-o89MtKEBKF4QfLl0Esw8HJNE_cj5_dcZdzEk/edit#gid=2009338596"))
		}
		sugar.Infof("calling ecommerce manager api success!")
	} else {
		ctx.Response.Header.Set("Content-Type", "application/json")
		ctx.Response.SetStatusCode(200)
		ctx.SetBody([]byte("Failed! There is an issue with Request sent for api call."))
		sugar.Infof("Failed! ecommerce manager api failed!")
	}
}
