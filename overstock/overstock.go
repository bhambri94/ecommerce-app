package overstock

import (
	"fmt"
	"time"

	"github.com/anaskhan96/soup"
)

// http://localhost:7001/v1/overstock/search

func GetPageDescription(productUrls []string) [][]interface{} {
	var finalValues [][]interface{}
	loc, _ := time.LoadLocation("America/Bogota")
	iterator := 0
	for iterator < len(productUrls) {
		time.Sleep(1 * time.Second)
		var row []interface{}
		currentTime := time.Now().In(loc)
		row = append(row, currentTime.Format("2006-01-02 15:04:05"), productUrls[iterator])
		OutputItemTitle := ""
		// OutputListPrice := ""
		// OutputPercOff := ""
		// OutputZipcode := ""
		// OutputCity := ""
		// OutputShippingTime := ""
		// OutputRatingValue := ""
		// OutputReviewCount := ""
		// OutputItemOutOfStock := ""
		// OutputZipcode, OutputCity, OutputShippingTime = GetDeliveryDetails(productUrls[iterator])
		fmt.Println(productUrls[iterator])
		resp, err := soup.Get(productUrls[iterator])
		if err != nil {
			fmt.Println("Nothing found")
		} else {
			doc := soup.HTMLParse(resp)
			fmt.Println(resp)
			ItemNumber := doc.Find("h1", "class", "_3Bj68d3")
			if ItemNumber.Error == nil {
				fmt.Println(ItemNumber)
				OutputItemTitle = ItemNumber.Text()
				fmt.Println(OutputItemTitle)
			}
			// PricingBox := doc.Find("div", "class", "price-box")
			// if PricingBox.Error == nil {
			// 	OldPrice := PricingBox.Find("p", "class", "old-price")
			// 	if OldPrice.Error == nil {
			// 		ListPriceValue := OldPrice.Find("span", "class", "price")
			// 		if ListPriceValue.Error == nil {
			// 			OutputListPrice = ListPriceValue.Text()
			// 		}
			// 		PercPriceValue := OldPrice.Find("span", "class", "savePrice")
			// 		if PercPriceValue.Error == nil {
			// 			OutputPercOff = PercPriceValue.Text()
			// 		}
			// 	}
			// }
			// RatingValue := doc.Find("meta", "itemprop", "ratingValue")
			// if RatingValue.Error == nil {
			// 	OutputRatingValue = RatingValue.Attrs()["content"]
			// }
			// ReviewCount := doc.Find("meta", "itemprop", "reviewCount")
			// if ReviewCount.Error == nil {
			// 	OutputReviewCount = ReviewCount.Attrs()["content"]
			// }
			// ItemOutOfStock := doc.Find("input", "class", "buy")
			// if ItemOutOfStock.Error == nil {
			// 	if ItemOutOfStock.Attrs()["value"] == "Buy Now" {
			// 		OutputItemOutOfStock = "Item is Available"
			// 	}
			// } else {
			// 	OutputItemOutOfStock = "Item is OOS"
			// }

		}
		// row = append(row, OutputItemNumber, OutputListPrice, OutputPercOff, OutputZipcode, OutputCity, OutputShippingTime, OutputRatingValue, OutputReviewCount, OutputItemOutOfStock)
		// finalValues = append(finalValues, row)
		iterator++
	}
	return finalValues
}
