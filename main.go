package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/anaskhan96/soup"
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
	router.GET("/v1/homedepot/search/query=:queryString", handleHomedepotSearch)
	router.GET("/v1/ebay/search", handleEbaySearch)
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
			loc, _ := time.LoadLocation("America/Bogota")
			currentTime := time.Now().In(loc)
			CSVName := "HomeDepotCSV" + currentTime.Format("2006-01-02 15:04:05") + ".csv"
			f, err := os.Create(CSVName)
			if err != nil {
				log.Fatal(err)
			}
			writer := csv.NewWriter(f)
			defer writer.Flush()

			stringfinalValues := make([][]string, len(finalValues)+5)
			header := []string{"HomeDepot_Refresh_time", "StoreId", "Item Id", "Category", "Store Quantity", "Online Quantity", "Returnable", "Savings Center", "SpecialPrice", "Upc", "ProductLabel", "BrandName", "IsLimitedQuantity", "OriginalPrice", "DollarOff", "AvailabilityType", "BossEstimatedShippingEndDate", "SthEstimatedShippingStartDate", "SthEstimatedShippingEndDate", "FreeShippingThreshold", "ExcludedShipStates", "FreeShippingMessage", "BossEstimatedShippingStartDate", "WebURL", "TotalReviews", "AverageRating", "Description", "BuyOnlineShipToStoreEligible", "IsTopSeller", "BuyOnlinePickupInStoreEligible", "ModelNumber", "VendorNumber", "AttributeValue", "DimensionName", "DimensionValue Name", "DiscountEndDate", "PromoLongDescription", "DiscountStartDate", "Image Links"}
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
			loc, _ = time.LoadLocation("America/Bogota")
			currentTime = time.Now().In(loc)
			ctx.Response.Header.Set("Content-Disposition", "attachment;filename="+"HomeDepotCSV"+currentTime.Format("2006-01-02 15:04:05")+".csv")
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

func handleHomedepotSearch(ctx *fasthttp.RequestCtx) {
	sugar.Infof("calling ecommerce manager api for homedepot search!")
	loc, _ := time.LoadLocation("America/Bogota")
	queryString := ctx.UserValue("queryString")
	if queryString != nil {
		queryString = queryString.(string)[1 : len(queryString.(string))-1]
		sugar.Infof("queryString for search is := " + queryString.(string))
	}
	currentTime := time.Now().In(loc)
	CSVName := "HomeDepotSearchCSV" + currentTime.Format("2006-01-02 15:04:05") + ".csv"
	f, err := os.Create(CSVName)
	if err != nil {
		log.Fatal(err)
	}
	writer := csv.NewWriter(f)
	defer writer.Flush()

	header := []string{"HD_Refresh_time", "Product HDs", "Description", "Brand", "Model Number", "Price", "Ratings", "Reviews Count"}
	writer.Write(header)

	loopCounter := 1
	pageCounter := 0
	firstApiCall := true
	productMap := make(map[string]bool)
	for interator := 0; interator < loopCounter; interator++ {
		var allProdCountInt int
		url := "https://www.homedepot.com/b/N-5yc1v/Ntk-EnrichedProductInfo/Ntt-" + queryString.(string) + "?NCNI-5&experienceName=default&Nao=" + strconv.Itoa(pageCounter*24) + "&Ns=None&storeSelection=6312,284,249,6356,258"
		resp, err := soup.Get(url)
		if err != nil {
			fmt.Println("Unable to call the homedepot apis")
		}
		pageCounter++
		doc := soup.HTMLParse(resp)
		if firstApiCall {
			fmt.Println(url)
			var allProdCount string
			allProductRoot := doc.Find("span", "id", "allProdCount")
			if allProductRoot.Error == nil {
				allProdCount = allProductRoot.Text()
			} else {
				break
			}
			fmt.Println(allProdCount)
			allProdCount = strings.Replace(allProdCount, ",", "", -1)
			allProdCountInt, err = strconv.Atoi(allProdCount)
			if err != nil {
				fmt.Println("Unable to convert all count to Integer")
				break
			}
			if allProdCountInt%24 > 0 {
				fmt.Println(allProdCountInt)
				loopCounter = allProdCountInt / 24
				loopCounter++
				fmt.Println(loopCounter)
			} else {
				loopCounter = allProdCountInt / 24
			}
			firstApiCall = false
		}
		stringfinalValues := make([][]string, (loopCounter*24)+5)
		links := doc.Find("div", "class", "pod-plp__container")
		if links.Error == nil {
			products := links.FindAll("div", "data-component", "productpod")
			i := 0
			for _, link := range products {
				currentTime := time.Now().In(loc)
				if link.Error == nil {
					if !productMap[link.Attrs()["data-productid"]] {
						stringfinalValues[i] = append(stringfinalValues[i], currentTime.Format("2006-01-02 15:04:05"), link.Attrs()["data-productid"])
						productMap[link.Attrs()["data-productid"]] = true
						innerTile := link.Find("div", "class", "pod-inner")
						if innerTile.Error == nil {
							tileInfo := innerTile.Find("div", "class", "plp-pod__info")
							if tileInfo.Error == nil {
								tileDescription := tileInfo.Find("div", "class", "pod-plp__description")
								if tileDescription.Error == nil {
									AnchorTag := tileDescription.Find("a")
									if AnchorTag.Error == nil {
										SpanText := AnchorTag.Find("span", "class", "pod-plp__brand-name")
										ProductTitle := AnchorTag.Text()
										ProductTitle = strings.Replace(ProductTitle, "\n", "", -1)
										ProductTitle = strings.Replace(ProductTitle, "     ", "", -1)
										stringfinalValues[i] = append(stringfinalValues[i], ProductTitle)
										if SpanText.Error == nil {
											BrandCopy := SpanText.Text()
											stringfinalValues[i] = append(stringfinalValues[i], BrandCopy)
										}
									}
								}
								tileModel := tileInfo.Find("div", "class", "pod-plp__model")
								if tileModel.Error == nil {
									ModelNumber := tileModel.Text()
									ModelNumber = strings.Replace(ModelNumber, "\n", "", -1)
									ModelNumber = strings.Replace(ModelNumber, "      ", "", -1)
									ModelNumber = strings.Replace(ModelNumber, " &nbsp;", "", -1)
									stringfinalValues[i] = append(stringfinalValues[i], ModelNumber)
								}
								priceWrapper1 := tileInfo.Find("div", "class", "price__wrapper")
								if priceWrapper1.Error == nil {
									priceWrapper2 := priceWrapper1.Find("div", "class", "if__overflow")
									if priceWrapper2.Error == nil {
										priceWrapper3 := priceWrapper2.Find("div")
										if priceWrapper3.Error == nil {
											priceWrapper4 := priceWrapper3.Find("div")
											if priceWrapper4.Error == nil {
												Price := priceWrapper4.Text()
												Cents := priceWrapper4.FindAll("span", "class", "price__format")
												for iter, c := range Cents {
													if iter == 1 {
														if Cents[iter].Error == nil {
															Price = Price + "." + c.Text()
															stringfinalValues[i] = append(stringfinalValues[i], Price)
														}
													}
												}

											}
										}
									}
								}
								RatingsWrapper := tileInfo.Find("div", "class", "pod-plp__ratings")
								if RatingsWrapper.Error == nil {
									Anchor := RatingsWrapper.FindAll("a")
									for iter, a := range Anchor {
										if iter == 0 {
											if Anchor[iter].Error == nil {
												RatingsSpan := a.Find("span", "class", "stars")
												if RatingsSpan.Error == nil {
													Ratings := RatingsSpan.Attrs()["rel"]
													stringfinalValues[i] = append(stringfinalValues[i], Ratings)
												}
											}
										}
										if iter == 1 {
											if Anchor[iter].Error == nil {
												ReviewsCount := a.Text()
												ReviewsCount = strings.Replace(ReviewsCount, "\n", "", -1)
												ReviewsCount = strings.Replace(ReviewsCount, "(", "", -1)
												ReviewsCount = strings.Replace(ReviewsCount, ")", "", -1)
												stringfinalValues[i] = append(stringfinalValues[i], ReviewsCount)
											}
										}
									}

								}
							}
						}
					}
					writer.Write(stringfinalValues[i])
					i++
				} else {
					continue
				}
			}
		} else {
			continue
		}

	}
	ctx.Response.SetStatusCode(200)
	ctx.Response.Header.Set("Content-Type", "text/csv")
	currentTime = time.Now().In(loc)
	ctx.Response.Header.Set("Content-Disposition", "attachment;filename="+"HomeDepotCSV"+currentTime.Format("2006-01-02 15:04:05")+".csv")
	ctx.SendFile(CSVName)
	err = os.Remove(CSVName)
	if err != nil {
		fmt.Println("Unable to delete file")
	} else {
		fmt.Println("File Deleted")

	}

}

func handleEbaySearch(ctx *fasthttp.RequestCtx) {
	resp, err := soup.Get("https://www.ebay.com/sch/i.html?_from=R40&_trksid=p2499334.m570.l1313&_nkw=iphone&_sacat=20349")

	if err != nil {
		// os.Exit(1)
		fmt.Println("Nothing found")
	} else {
		doc := soup.HTMLParse(resp)
		allProdCount := doc.Find("h1", "class", "srp-controls__count-heading")
		count := allProdCount.FindAll("span")
		for _, c := range count {
			fmt.Println(c.Text())
		}
		fmt.Println(count)
		// span id allProdCount
		// links := doc.Find("div", "class", "pod-plp__container")
		products := doc.FindAll("a", "class", "s-item__link")
		// data - component = "productpod"
		// pod-plp__container--alignment-resetwith__certona
		for _, link := range products {
			fmt.Println(link.Attrs()["href"])
		}
		// fmt.Println(products)
	}

}
