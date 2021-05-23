package zoro

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/anaskhan96/soup"
)

func GetPageDescriptionForZoro(productUrls []string) [][]interface{} {
	var finalValues [][]interface{}
	loc, _ := time.LoadLocation("America/Bogota")
	currentTime := time.Now().In(loc)
	iterator := 0
	for iterator < len(productUrls) {
		var row []interface{}
		row = append(row, currentTime.Format("2006-01-02 15:04:05"), productUrls[iterator])
		OutputProductTitle := ""
		OutputProductBrandName := ""
		OutputProductPrice := ""
		OutputZoroNumber := ""
		OutputMFRNumber := ""
		OutputQuanity := ""
		OutputNumberOfReviews := ""
		OutputShippingTime := ""
		OutputDescription := ""
		OutputImageURLS := ""
		OutputItemsFeedback := ""
		OutputNotification := ""
		OutputRatingValue := ""
		resp, err := soup.Get(productUrls[iterator])
		if err != nil {
			fmt.Println("Nothing found")
		} else {
			doc := soup.HTMLParse(resp)
			ProductBrandName := doc.Find("a", "data-za", "product-brand-name")
			if ProductBrandName.Error == nil {
				OutputProductBrandName = ProductBrandName.Text()
			}
			ProductTitle := doc.Find("h1", "data-za", "product-name")
			if ProductTitle.Error == nil {
				OutputProductTitle = ProductTitle.Text()
			}
			fmt.Println(OutputProductTitle)
			ProductDetails := doc.Find("div", "class", "product-details-page__overview")
			if ProductDetails.Error == nil {
				ProductPrice := ProductDetails.Find("span", "class", "product-price__price")
				if ProductPrice.Error == nil {
					OutputProductPrice = stripSpaces(ProductPrice.Text())
				}
			}
			ZoroNumber := doc.Find("span", "data-za", "PDPZoroNo")
			if ZoroNumber.Error == nil {
				OutputZoroNumber = ZoroNumber.Text()
			}

			MFRNumber := doc.Find("span", "data-za", "PDPMfrNo")
			if MFRNumber.Error == nil {
				OutputMFRNumber = MFRNumber.Text()
			}
			Qty := doc.Find("input", "data-za", "quantity-input")
			if Qty.Error == nil {
				if Qty.Error == nil {
					OutputQuanity = stripSpaces(Qty.Attrs()["value"])
				}
			}

			//section["data-testid"="review-snippet"]/[@class="pr-snippet-rating-decimal"]
			Out := doc.Find("section", "class", "pr-review-snippet-container")
			if Out.Error == nil {
				fmt.Println("reached here 1")
				Out1 := doc.Find("div", "class", "pr-snippet-rating-decimal")
				if Out1.Error == nil {
					fmt.Println("reached here 2")
					OutputItemsFeedback = Out1.Text()
				}
			}

			NumberOfReviews := doc.Find("a", "class", "pr-snippet-review-count")
			if NumberOfReviews.Error == nil {
				fmt.Println("reached here 3")
				OutputNumberOfReviews = NumberOfReviews.Text()
			}

			Notification := doc.Find("div", "class", "ships-from-lead-time")
			OutputNotification = ""
			if Notification.Error == nil {
				NotificationSpan := Notification.Find("span")
				if NotificationSpan.Error == nil {
					NotificationStrong := NotificationSpan.Find("strong")
					OutputNotification = NotificationSpan.Text()
					if NotificationStrong.Error == nil {
						OutputNotification = NotificationStrong.Text()
					}
				}
			}
			ShipmentCharge := doc.Find("div", "class", "product-free-shipping")
			OutputShippingTime = ""
			if ShipmentCharge.Error == nil {
				ShipmentChargeSpan := ShipmentCharge.Find("span")
				if ShipmentChargeSpan.Error == nil {
					ShipmentChargeStrong := ShipmentChargeSpan.Find("strong")
					OutputShippingTime = ShipmentChargeSpan.Text()
					if ShipmentChargeStrong.Error == nil {
						OutputShippingTime = OutputShippingTime + ShipmentChargeStrong.Text()
					}
				}
			}

			Description := doc.Find("div", "class", "product-description__text")
			if Description.Error == nil {
				OutputDescription = Description.Text()
			}

			MFRNo := doc.Find("span", "data-za", "PDPMfrNo")
			if MFRNo.Error == nil {
				OutputMFRNumber = MFRNo.Text()
			}

			// OutputImageURLS = FindAllImages(productUrls[iterator])
			ImageParentDiv := doc.FindAll("div", "class", "product-images__heros")
			fmt.Println(len(ImageParentDiv))

			ImageURLParentChildDiv := doc.FindAll("div", "class", "product-images__hero")
			fmt.Println(len(ImageParentDiv))

			OutputImageURLS = ""
			fmt.Println(len(ImageURLParentChildDiv))
			for _, inumber := range ImageURLParentChildDiv {
				if inumber.Error == nil {
					fmt.Println("thn aaya")
					// ImageURLChildDiv := inumber.Find("div")
					// if ImageURLChildDiv.Error == nil {
					ImageLinks := inumber.Find("img")
					// fmt.Println(len(ImageLinks))
					// for _, inumber := range ImageLinks {
					fmt.Println("yhn bhi aaya")
					OutputImageURLS = ImageLinks.Attrs()["src"] + " ,"
					// }
				}
			}

			OutputNumberOfReviews, OutputRatingValue = FindRatingsCount(OutputZoroNumber)

			fmt.Println(OutputImageURLS)
		}

		row = append(row, OutputProductBrandName, OutputProductTitle, OutputProductPrice, OutputQuanity, OutputZoroNumber, OutputMFRNumber, OutputNumberOfReviews, OutputRatingValue, OutputImageURLS, OutputDescription, OutputShippingTime, OutputItemsFeedback, OutputNotification)
		finalValues = append(finalValues, row)
		iterator++
	}
	return finalValues
}

func stripSpaces(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, str)
}

func FindRatingsCount(itemId string) (string, string) {
	req, err := http.NewRequest("GET", "https://display.powerreviews.com/m/297763/l/en_US/product/"+itemId+"/reviews?_noconfig=true&apikey=90d71773-2b19-45c6-a4ec-053f308ea2cd", nil)
	if err != nil {
		// handle err
	}
	req.Header.Set("Sec-Ch-Ua", "\" Not A;Brand\";v=\"99\", \"Chromium\";v=\"90\", \"Google Chrome\";v=\"90\"")
	req.Header.Set("Referer", "https://www.zoro.com/")
	req.Header.Set("Sec-Ch-Ua-Mobile", "?0")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.93 Safari/537.36")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", ""
	}
	var zoroReview ZoroReview
	err = json.Unmarshal(body, &zoroReview)
	if err != nil {
		fmt.Println("whoops:", err)
		return "", ""
	}
	s := fmt.Sprint(zoroReview.Results[0].Rollup.AverageRating)
	return strconv.Itoa(zoroReview.Paging.TotalResults), s
}

//http://localhost:7001/v1/zoro/search
