package costway

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/anaskhan96/soup"
)

func GetPageDescription(productUrls []string) [][]interface{} {
	var finalValues [][]interface{}
	loc, _ := time.LoadLocation("America/Bogota")
	iterator := 0
	for iterator < len(productUrls) {
		time.Sleep(1 * time.Second)
		var row []interface{}
		currentTime := time.Now().In(loc)
		row = append(row, currentTime.Format("2006-01-02 15:04:05"), productUrls[iterator])
		OutputItemNumber := ""
		OutputListPrice := ""
		OutputPercOff := ""
		OutputZipcode := ""
		OutputCity := ""
		OutputShippingTime := ""
		OutputRatingValue := ""
		OutputReviewCount := ""
		OutputItemOutOfStock := ""
		OutputZipcode, OutputCity, OutputShippingTime = GetDeliveryDetails(productUrls[iterator])

		resp, err := soup.Get(productUrls[iterator])
		if err != nil {
			fmt.Println("Nothing found")
		} else {
			doc := soup.HTMLParse(resp)
			ItemNumber := doc.Find("div", "class", "orat")
			if ItemNumber.Error == nil {
				ItemNumberValue := ItemNumber.FindAll("span")
				for _, inumber := range ItemNumberValue {
					if len(inumber.Text()) > 4 {
						if inumber.Text()[0:4] == "Item" {
							OutputItemNumber = inumber.Text()
							OutputItemNumber = strings.ReplaceAll(OutputItemNumber, "Item No: ", "")
							break
						}
					}
				}
			}
			PricingBox := doc.Find("div", "class", "price-box")
			if PricingBox.Error == nil {
				OldPrice := PricingBox.Find("p", "class", "old-price")
				if OldPrice.Error == nil {
					ListPriceValue := OldPrice.Find("span", "class", "price")
					if ListPriceValue.Error == nil {
						OutputListPrice = ListPriceValue.Text()
					}
					PercPriceValue := OldPrice.Find("span", "class", "savePrice")
					if PercPriceValue.Error == nil {
						OutputPercOff = PercPriceValue.Text()
					}
				}
			}
			RatingValue := doc.Find("meta", "itemprop", "ratingValue")
			if RatingValue.Error == nil {
				OutputRatingValue = RatingValue.Attrs()["content"]
			}
			ReviewCount := doc.Find("meta", "itemprop", "reviewCount")
			if ReviewCount.Error == nil {
				OutputReviewCount = ReviewCount.Attrs()["content"]
			}
			ItemOutOfStock := doc.Find("input", "class", "buy")
			if ItemOutOfStock.Error == nil {
				if ItemOutOfStock.Attrs()["value"] == "Buy Now" {
					OutputItemOutOfStock = "Item is Available"
				}
			} else {
				OutputItemOutOfStock = "Item is OOS"
			}

		}
		row = append(row, OutputItemNumber, OutputListPrice, OutputPercOff, OutputZipcode, OutputCity, OutputShippingTime, OutputRatingValue, OutputReviewCount, OutputItemOutOfStock)
		finalValues = append(finalValues, row)
		iterator++
	}
	return finalValues
}

type DeliveryResposne struct {
	Status bool `json:"status"`
	Data   struct {
		Zipcode string `json:"zipcode"`
		City    string `json:"city"`
		Tell    string `json:"tell"`
	} `json:"data"`
	Message string `json:"message"`
}

func GetDeliveryDetails(product string) (string, string, string) {
	url := "https://www.costway.com/hycatalog/product/shipping"
	method := "POST"

	payload := strings.NewReader("cpbh=SP36126&zipcode=77478")

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		return "", "", ""
	}
	req.Header.Add("authority", "www.costway.com")
	req.Header.Add("accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Add("x-requested-with", "XMLHttpRequest")
	req.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36")
	req.Header.Add("content-type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Add("origin", "https://www.costway.com")
	req.Header.Add("sec-fetch-site", "same-origin")
	req.Header.Add("sec-fetch-mode", "cors")
	req.Header.Add("sec-fetch-dest", "empty")
	req.Header.Add("referer", product)
	req.Header.Add("accept-language", "en-GB,en-US;q=0.9,en;q=0.8")
	req.Header.Add("cookie", "is_browser=1; _lc2_fpi=ffd29af5c7a5--01emhmbw2szs65944jr74zftr2; _scid=a26dc5d5-f3e1-463b-a7b3-599c6e2dba02; _fbp=fb.1.1602614587968.1598998946; _hjid=f831f0f0-f84e-4b28-a2b9-480e32d2cc19; rskxRunCookie=0; rCookie=mfmk9sxhvwoucqp1yut35pkek55kx9; _pin_unauth=dWlkPVpXRTJPV0ZrTldNdFltWTJaUzAwWmpjekxXSXpORGt0WVRrME0yWmpNRFF3TVdWaQ; __utma=186024123.1743455298.1602614588.1602614588.1602614622.2; __utmz=186024123.1602614622.2.2.utmcsr=google|utmccn=(organic)|utmcmd=organic|utmctr=(not%20provided); is_browser=1; _li_dcdm_c=.costway.com; _hjTLDTest=1; targetipec01af42fb7c8d58af6dbbcaade21611=https://www.costway.com; ip_ec01af42fb7c8d58af6dbbcaade21611=ec01af42fb7c8d58af6dbbcaade21611; subreg=1; feedexport=160498465044395755580576; feedexport_fee=28; __cfduid=dca7efe590d55885da875318ac8508fd81607164187; frontend=3vh4ctji4ajij3sd5pie425gc0; bannerslider_user_code_impress3=5356a3930c5b1fa136a1c2001d24e0c4; PHPSESSID=mm2hoig29dvkq7fslmqegbhm70; _sctr=1|1607106600000; _gid=GA1.2.456919963.1607164205; _hjIncludedInPageviewSample=1; _hjAbsoluteSessionInProgress=0; _hjIncludedInSessionSample=0; _hjShownFeedbackMessage=true; external_no_cache=1; _derived_epik=dj0yJnU9UTBBRUNweTA3LUtvTERnbkNGeGRubG1HOTFESTAwcEgmbj1TeEZzb1ZUQURMeEY2S1lWUTBoMFlBJm09MSZ0PUFBQUFBRl9MWXNVJnJtPTEmcnQ9QUFBQUFGX0xZc1U; _uetsid=d5d4fc1036e411ebbc8eed7cb136ebe1; _uetvid=d5d5511036e411eb916393c6cd6ba1ad; lastRskxRun=1607164616562; _ga=GA1.1.87856650.1604497632; _ga_FB2NZFH559=GS1.1.1607164202.15.1.1607165957.0")

	res, err := client.Do(req)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	var delResponse DeliveryResposne
	err = json.Unmarshal(body, &delResponse)
	return delResponse.Data.Zipcode, delResponse.Data.City, delResponse.Data.Tell
}
