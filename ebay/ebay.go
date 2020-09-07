package ebay

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

func GetPageDescription(productUrls []string) [][]interface{} {
	var finalValues [][]interface{}
	loc, _ := time.LoadLocation("America/Bogota")
	currentTime := time.Now().In(loc)
	iterator := 0
	for iterator < len(productUrls) {
		var row []interface{}
		row = append(row, currentTime.Format("2006-01-02 15:04:05"), productUrls[iterator])
		OutputProductTitle := ""
		OutputProductPrice := ""
		OutputQuanity := ""
		OutputItemsFeedback := ""
		OutputNotification := ""
		OutputDelDate := ""
		OutputReturnMessage := ""
		OutputDescription := ""
		DescriptionURL := ""
		fmt.Println(OutputDescription)
		OutputImageURLS := ""
		// OutputHeader := ""
		resp, err := soup.Get(productUrls[iterator])
		if err != nil {
			fmt.Println("Nothing found")
		} else {
			doc := soup.HTMLParse(resp)
			ProductTitle := doc.Find("h1", "class", "it-ttl")
			if ProductTitle.Error == nil {
				OutputProductTitle = ProductTitle.Text()
			}
			ProductPrice := doc.Find("span", "id", "prcIsum")
			if ProductPrice.Error == nil {
				OutputProductPrice = stripSpaces(ProductPrice.Text())
			}
			QtySubText := doc.Find("span", "id", "qtySubTxt")
			if QtySubText.Error == nil {
				Qty := QtySubText.Find("span")
				if Qty.Error == nil {
					OutputQuanity = stripSpaces(Qty.Text())
				}
			}
			ItemsFeedback := doc.Find("span", "class", "vi-qtyS-hot-red")
			if ItemsFeedback.Error == nil {
				ItemQty := ItemsFeedback.Find("a")
				if ItemQty.Error == nil {
					OutputItemsFeedback = ItemQty.Text()
				}
			}
			Notification := doc.Find("div", "id", "vi_notification_new")
			if Notification.Error == nil {
				NotificationSpan := Notification.Find("span")
				if NotificationSpan.Error == nil {
					OutputNotification = NotificationSpan.Text()
				}
			}
			ReturnMessage := doc.Find("span", "id", "vi-ret-accrd-txt")
			if ReturnMessage.Error == nil {
				OutputReturnMessage = ReturnMessage.Text()
			}
			ImageURLS := FindAllImages(productUrls[iterator])
			if len(ImageURLS) > 0 {
				OutputImageURLS = ImageURLS
			}

			ItemIndex := strings.Index(productUrls[iterator], "itm")
			if ItemIndex != -1 {
				ItemID := productUrls[iterator][ItemIndex+4:]
				fmt.Println("https://www.ebay.com/itm/getrates?item=" + ItemID + "&_trksid=p2047675.l2682&quantity=1&country=1&zipCode=77478&co=0")
				resp, err = soup.Get("https://www.ebay.com/itm/getrates?item=" + ItemID + "&_trksid=p2047675.l2682&quantity=1&country=1&zipCode=77478&co=0")
				start := strings.Index(resp, "sh-del-frst")
				if start != -1 {
					end := strings.Index(resp[start:len(resp)], "<span")
					if end != -1 {
						OutputDelDate = resp[start+17:start+end] + " "
						start = strings.Index(resp, "vi-acc-del-range")
						if start != -1 {
							end = strings.Index(resp[start:len(resp)], "</span>")
							if end != -1 {
								OutputDelDate = OutputDelDate + resp[start+22:start+end-4]
							}
						}

					}
				}
				currenttime := time.Now()
				epoch := currenttime.Unix()
				DescriptionURL = "https://vi.vipr.ebaydesc.com/ws/eBayISAPI.dll?ViewItemDescV4&item=" + ItemID + "&t=" + strconv.Itoa(int(epoch)) + "000&seller=onebigoutlet&excSoj=1&excTrk=1&lsite=0&ittenable=false&domain=ebay.com&descgauge=1&cspheader=1&oneClk=2&secureDesc=1&oversion=819e9518"
				// _, OutputDescription, _ = ExampleScrape(DescriptionURL)
			}
		}
		//, OutputHeader, "OutputDescription"
		row = append(row, OutputProductTitle, OutputProductPrice, OutputQuanity, OutputItemsFeedback, OutputNotification, OutputDelDate, OutputReturnMessage, OutputImageURLS, DescriptionURL)
		finalValues = append(finalValues, row)
		iterator++
	}
	return finalValues
}

func ExampleScrape(URL string) (string, string, string) {
	res, err := http.Get(URL)
	fmt.Println(URL)
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
	Header := ""
	Description := ""
	FeaturesLI := ""
	doc.Find("body").Find("h2").Each(func(i int, s *goquery.Selection) {
		Header = Header + s.Text() + "\t"
	})
	fmt.Println(Header)
	doc.Find("body").Each(func(i int, s *goquery.Selection) {

		if !strings.Contains(Description, s.Find("strong").Text()) {
			if len(s.Find("strong").Text()) > 5 && len(s.Find("strong").Text()) > 15 && !strings.Contains(s.Find("strong").Text(), "margin") {
				if !strings.Contains(Description, s.Find("strong").Text()[5:15]) {
					Description = Description + s.Find("strong").Text()
				}
			}
		}
		if !strings.Contains(Description, s.Find("p").Text()) {
			if len(s.Find("p").Text()) > 5 && len(s.Find("p").Text()) > 15 && !strings.Contains(s.Find("p").Text(), "margin") {
				if !strings.Contains(Description, s.Find("p").Text()[5:15]) {
					Description = Description + s.Find("p").Text()
				}
			}
		}
		if !strings.Contains(Description, s.Find("b").Text()) {
			if len(s.Find("b").Text()) > 5 && len(s.Find("b").Text()) > 15 && !strings.Contains(s.Find("b").Text(), "margin") {
				if !strings.Contains(Description, s.Find("b").Text()[5:15]) {
					Description = Description + s.Find("b").Text()
				}
			}
		}
		if !strings.Contains(Description, s.Find("i").Text()) {
			if len(s.Find("i").Text()) > 5 && len(s.Find("i").Text()) > 15 && !strings.Contains(s.Find("i").Text(), "margin") {
				if !strings.Contains(Description, s.Find("i").Text()[5:15]) {
					Description = Description + s.Find("i").Text()
				}
			}
		}
		if !strings.Contains(Description, s.Find("div").Text()) {
			if len(s.Find("div").Text()) > 5 && len(s.Find("div").Text()) > 15 && !strings.Contains(s.Find("div").Text(), "margin") {
				if !strings.Contains(Description, s.Find("div").Text()[5:15]) {
					Description = Description + s.Find("div").Text()
				}
			}
		}

		// Description = Description + s.Find("div").Text() + "\t"
		// Description = Description + s.Find("strong").Text() + "\t"
		// Description = Description + s.Find("p").Text() + "\t"
		// Description = Description + s.Find("i").Text() + "\t"
		// Description = Description + s.Find("b").Text() + "\t"
	})
	fmt.Println(Description)
	Iterator := 1
	doc.Find("body").Find("li").Each(func(i int, s *goquery.Selection) {
		FeaturesLI = FeaturesLI + strconv.Itoa(Iterator) + ". " + s.Text() + "\t"
		Iterator++
	})
	fmt.Printf(Description)
	return Header, Description, ""
}

func FindAllImages(URL string) string {
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
	Images := ""
	doc.Find("#vi_main_img_fs").Find("img").Each(func(i int, s *goquery.Selection) {
		Value, ImageAvailable := s.Attr("src")
		Index := strings.Index(Value, "s-l")
		if Index != -1 {
			Value = Value[:Index] + "s-l1000.jpg"
		}
		if ImageAvailable {
			Images = Images + Value + " - "
		}
	})
	return Images
}

func stripSpaces(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, str)
}
