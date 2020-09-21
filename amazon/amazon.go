package amazon

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/PuerkitoBio/goquery"
	"github.com/anaskhan96/soup"
)

func GetTopThreeSellerPrices(productUrls []string) [][]interface{} {
	var finalValues [][]interface{}
	loc, _ := time.LoadLocation("America/Bogota")
	currentTime := time.Now().In(loc)
	i := 0
	for i < len(productUrls) {
		time.Sleep(200 * time.Millisecond)
		var row []interface{}
		OutputProductLink := ""
		OutputProductASIN := ""
		var OutputSellerRow []interface{}
		for len(productUrls[i]) < 12 {
			productUrls[i] = "0" + productUrls[i]
		}
		row = append(row, currentTime.Format("2006-01-02 15:04:05"))
		row = append(row, productUrls[i])
		Url := "https://www.amazon.com/s?k=" + productUrls[i] + "&ref=nb_sb_noss"
		row = append(row, Url)
		fmt.Println(Url)
		resp, err := soup.Get(Url)
		if err != nil {
			fmt.Println("Nothing found")
		} else {
			i := 0
			for i < 6 {
				a, b := CheckUnSponseredProduct(resp, strconv.Itoa(i))
				if a != "" {
					OutputProductASIN = a
					OutputProductLink = b
					break
				}
				i++
			}
			if OutputProductASIN == "" {
				doc := soup.HTMLParse(resp)
				ProductLink := doc.Find("div", "data-index", "0")
				if ProductLink.Error == nil {
					OutputProductASIN = ProductLink.Attrs()["data-asin"]
					fmt.Println(OutputProductASIN)
					FinalLink := ProductLink.Find("a")
					if FinalLink.Error == nil {
						OutputProductLink = FinalLink.Attrs()["href"]
					}
				}
			}
		}
		if OutputProductLink != "" {
			row = append(row, "https://www.amazon.com"+OutputProductLink)
		} else {
			row = append(row, OutputProductLink)
		}

		// if OutputProductLink != "" {
		// 	ProductURL := "https://www.amazon.com" + OutputProductLink
		// 	resp, err := soup.Get(ProductURL)
		// 	if err != nil {
		// 		fmt.Println("Nothing found")
		// 	} else {
		// 		doc := soup.HTMLParse(resp)
		// 		TableWithASIN := doc.Find("table", "id", "productDetails_detailBullets_sections1")
		// 		if TableWithASIN.Error == nil {
		// 			TableBody := TableWithASIN.Find("tbody")
		// 			if TableBody.Error == nil {
		// 				TableRows := TableBody.FindAll("tr")
		// 				if len(TableRows) > 0 {
		// 					for _, row := range TableRows {
		// 						SingleRow := row.Find("th")
		// 						if SingleRow.Error == nil {
		// 							THValue := SingleRow.Text()
		// 							if strings.Contains(THValue, "ASIN") {
		// 								OutputProductASIN = stripSpaces(row.Find("td").Text())
		// 							}
		// 						}
		// 					}
		// 				}
		// 			}
		// 		}
		// 	}
		// }
		row = append(row, OutputProductASIN)
		if OutputProductASIN != "" {
			SellerPageURL := "https://www.amazon.com/gp/offer-listing/" + OutputProductASIN + "/ref=dp_olp_ALL_mbc?ie=UTF8&condition=ALL"
			resp, err := soup.Get(SellerPageURL)
			if err != nil {
				fmt.Println("Nothing found")
			} else {
				doc := soup.HTMLParse(resp)
				AllSellerColumns := doc.Find("div", "id", "olpOfferListColumn")
				if AllSellerColumns.Error == nil {
					AllSellerRows := AllSellerColumns.FindAll("div", "class", "olpOffer")
					if len(AllSellerRows) > 0 {
						for _, row := range AllSellerRows {
							OutputProductPrice := ""
							OutputShippingInfo := ""
							OutputProductCondition := ""
							OutputDelivery := ""
							OutputSellerInfo := ""
							Price := row.Find("div", "class", "olpPriceColumn")
							if Price.Error == nil {
								PriceText := Price.Find("span")
								if PriceText.Error == nil {
									OutputProductPrice = stripSpaces(PriceText.Text())
								}
								ShippingInfo := Price.Find("b")
								if ShippingInfo.Error == nil {
									OutputShippingInfo = stripSpaces(ShippingInfo.Text())
								}

							}

							Condition := row.Find("div", "class", "olpConditionColumn")
							if Condition.Error == nil {
								ConditionText := Condition.Find("span")
								if ConditionText.Error == nil {
									OutputProductCondition = stripSpaces(ConditionText.Text())
								}
							}

							Delivery := row.Find("div", "class", "olpDeliveryColumn")
							if Delivery.Error == nil {
								DeliveryText := Delivery.Find("span")
								if DeliveryText.Error == nil {
									OutputDelivery = stripSpaces(DeliveryText.Text())
								}
							}

							Seller := row.Find("div", "class", "olpSellerColumn")
							if Seller.Error == nil {
								SellerTextSpan := Delivery.Find("span")
								if SellerTextSpan.Error == nil {
									OutputSellerInfo = stripSpaces(SellerTextSpan.Text())
								}
								SellerTextA := Delivery.Find("a")
								if SellerTextA.Error == nil {
									OutputSellerInfo = OutputSellerInfo + "\t" + stripSpaces(SellerTextA.Text())
								}
								SellerTextP := Delivery.Find("p")
								if SellerTextP.Error == nil {
									OutputSellerInfo = OutputSellerInfo + "\t" + stripSpaces(SellerTextP.Text())
								}
							}
							OutputSellerRow = append(OutputSellerRow, OutputProductPrice, OutputShippingInfo, OutputProductCondition, OutputDelivery)

						}
					}

				}
			}
		}
		iter := 0
		for iter < len(OutputSellerRow) {
			row = append(row, OutputSellerRow[iter])
			iter++
		}
		i++
		finalValues = append(finalValues, row)
	}
	return finalValues
}

func FindProductLink(URL string) {
	res, err := http.Get(URL)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		fmt.Printf("status code error: %d %s", res.StatusCode, res.Status)
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	doc.Find("div").Find("data-index").Find("0").Each(func(i int, s *goquery.Selection) {
		Value, _ := s.Attr("href")
		fmt.Println(Value)
	})
}

func stripSpaces(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, str)
}

func CheckUnSponseredProduct(resp string, productIndex string) (string, string) {
	OutputProductASIN := ""
	OutputProductLink := ""
	Flag := false
	doc := soup.HTMLParse(resp)
	ProductLink := doc.Find("div", "data-index", productIndex)
	SpanTextCopy := ""
	if ProductLink.Error == nil {
		NonSponserCheck := ProductLink.FindAll("span", "class", "s-label-popover-hover")
		if len(NonSponserCheck) > 0 {
			for _, spanText := range NonSponserCheck {
				SpanS := spanText.Find("span")
				if SpanS.Error == nil {
					SpanTextCopy = SpanS.Text()
					if stripSpaces(SpanTextCopy) == "Sponsored" {
						Flag = false
					}
				}
			}
		} else {
			Flag = true
		}
	}
	if Flag {
		OutputProductASIN = ProductLink.Attrs()["data-asin"]
		fmt.Println(OutputProductASIN)
		FinalLink := ProductLink.Find("a")
		if FinalLink.Error == nil {
			OutputProductLink = FinalLink.Attrs()["href"]
		}
		return OutputProductASIN, OutputProductLink
	} else {
		return OutputProductASIN, OutputProductLink
	}
}
