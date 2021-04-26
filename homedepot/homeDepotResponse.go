package homedepot

type HomeDepotNewResponse struct {
	Data struct {
		Product struct {
			Fulfillment struct {
				Backordered        bool `json:"backordered"`
				FulfillmentOptions []struct {
					Type        string `json:"type"`
					Fulfillable bool   `json:"fulfillable"`
					Services    []struct {
						Type      string `json:"type"`
						Locations []struct {
							IsAnchor  bool `json:"isAnchor"`
							Inventory struct {
								IsLimitedQuantity  bool        `json:"isLimitedQuantity"`
								IsOutOfStock       bool        `json:"isOutOfStock"`
								IsInStock          bool        `json:"isInStock"`
								Quantity           int         `json:"quantity"`
								IsUnavailable      bool        `json:"isUnavailable"`
								MaxAllowedBopisQty interface{} `json:"maxAllowedBopisQty"`
								MinAllowedBopisQty interface{} `json:"minAllowedBopisQty"`
								Typename           string      `json:"__typename"`
							} `json:"inventory"`
							Type                    string      `json:"type"`
							StoreName               string      `json:"storeName"`
							LocationID              string      `json:"locationId"`
							CurbsidePickupFlag      bool        `json:"curbsidePickupFlag"`
							IsBuyInStoreCheckNearBy interface{} `json:"isBuyInStoreCheckNearBy"`
							Distance                interface{} `json:"distance"`
							State                   string      `json:"state"`
							StorePhone              string      `json:"storePhone"`
							Typename                string      `json:"__typename"`
						} `json:"locations"`
						DeliveryTimeline      interface{} `json:"deliveryTimeline"`
						DeliveryDates         interface{} `json:"deliveryDates"`
						DeliveryCharge        interface{} `json:"deliveryCharge"`
						DynamicEta            interface{} `json:"dynamicEta"`
						HasFreeShipping       bool        `json:"hasFreeShipping"`
						FreeDeliveryThreshold interface{} `json:"freeDeliveryThreshold"`
						TotalCharge           interface{} `json:"totalCharge"`
						Typename              string      `json:"__typename"`
					} `json:"services"`
					Typename string `json:"__typename"`
				} `json:"fulfillmentOptions"`
				AnchorStoreStatus       bool        `json:"anchorStoreStatus"`
				AnchorStoreStatusType   string      `json:"anchorStoreStatusType"`
				BackorderedShipDate     interface{} `json:"backorderedShipDate"`
				BossExcludedShipStates  string      `json:"bossExcludedShipStates"`
				ExcludedShipStates      string      `json:"excludedShipStates"`
				SeasonStatusEligible    interface{} `json:"seasonStatusEligible"`
				OnlineStoreStatus       bool        `json:"onlineStoreStatus"`
				OnlineStoreStatusType   string      `json:"onlineStoreStatusType"`
				InStoreAssemblyEligible bool        `json:"inStoreAssemblyEligible"`
				Typename                string      `json:"__typename"`
			} `json:"fulfillment"`
			ItemID      string `json:"itemId"`
			DataSources string `json:"dataSources"`
			Identifiers struct {
				CanonicalURL        string      `json:"canonicalUrl"`
				BrandName           string      `json:"brandName"`
				ItemID              string      `json:"itemId"`
				ModelNumber         string      `json:"modelNumber"`
				ProductLabel        string      `json:"productLabel"`
				StoreSkuNumber      string      `json:"storeSkuNumber"`
				UpcGtin13           string      `json:"upcGtin13"`
				SpecialOrderSku     string      `json:"specialOrderSku"`
				ToolRentalSkuNumber interface{} `json:"toolRentalSkuNumber"`
				RentalCategory      interface{} `json:"rentalCategory"`
				RentalSubCategory   interface{} `json:"rentalSubCategory"`
				Upc                 string      `json:"upc"`
				IsSuperSku          bool        `json:"isSuperSku"`
				ParentID            string      `json:"parentId"`
				ProductType         string      `json:"productType"`
				SampleID            interface{} `json:"sampleId"`
				Typename            string      `json:"__typename"`
			} `json:"identifiers"`
			AvailabilityType struct {
				Discontinued bool   `json:"discontinued"`
				Status       bool   `json:"status"`
				Type         string `json:"type"`
				Buyable      bool   `json:"buyable"`
				Typename     string `json:"__typename"`
			} `json:"availabilityType"`
			Details struct {
				Description string      `json:"description"`
				Collection  interface{} `json:"collection"`
				Highlights  []string    `json:"highlights"`
				Typename    string      `json:"__typename"`
			} `json:"details"`
			Media struct {
				Images []struct {
					URL      string   `json:"url"`
					Sizes    []string `json:"sizes"`
					Type     string   `json:"type"`
					SubType  string   `json:"subType"`
					Typename string   `json:"__typename"`
				} `json:"images"`
				Video []struct {
					ShortDescription string      `json:"shortDescription"`
					Thumbnail        string      `json:"thumbnail"`
					URL              string      `json:"url"`
					VideoStill       string      `json:"videoStill"`
					Link             interface{} `json:"link"`
					Title            string      `json:"title"`
					Type             string      `json:"type"`
					VideoID          string      `json:"videoId"`
					LongDescription  interface{} `json:"longDescription"`
					Typename         string      `json:"__typename"`
				} `json:"video"`
				ThreeSixty []struct {
					ID       string `json:"id"`
					URL      string `json:"url"`
					Typename string `json:"__typename"`
				} `json:"threeSixty"`
				AugmentedRealityLink struct {
					Usdz     interface{} `json:"usdz"`
					Image    string      `json:"image"`
					Typename string      `json:"__typename"`
				} `json:"augmentedRealityLink"`
				Typename string `json:"__typename"`
			} `json:"media"`
			Pricing struct {
				Promotion struct {
					Dates struct {
						End      string `json:"end"`
						Start    string `json:"start"`
						Typename string `json:"__typename"`
					} `json:"dates"`
					Type                    string      `json:"type"`
					Description             interface{} `json:"description"`
					DollarOff               float64     `json:"dollarOff"`
					PercentageOff           float64     `json:"percentageOff"`
					SavingsCenter           string      `json:"savingsCenter"`
					SavingsCenterPromos     interface{} `json:"savingsCenterPromos"`
					SpecialBuySavings       interface{} `json:"specialBuySavings"`
					SpecialBuyDollarOff     interface{} `json:"specialBuyDollarOff"`
					SpecialBuyPercentageOff interface{} `json:"specialBuyPercentageOff"`
					ExperienceTag           interface{} `json:"experienceTag"`
					SubExperienceTag        interface{} `json:"subExperienceTag"`
					AnchorItemList          interface{} `json:"anchorItemList"`
					ItemList                interface{} `json:"itemList"`
					Reward                  interface{} `json:"reward"`
					Typename                string      `json:"__typename"`
				} `json:"promotion"`
				Value                 float64 `json:"value"`
				AlternatePriceDisplay bool    `json:"alternatePriceDisplay"`
				Alternate             struct {
					Bulk interface{} `json:"bulk"`
					Unit struct {
						CaseUnitOfMeasure  interface{} `json:"caseUnitOfMeasure"`
						UnitsOriginalPrice interface{} `json:"unitsOriginalPrice"`
						UnitsPerCase       interface{} `json:"unitsPerCase"`
						Value              interface{} `json:"value"`
						Typename           string      `json:"__typename"`
					} `json:"unit"`
					Typename string `json:"__typename"`
				} `json:"alternate"`
				Original              float64     `json:"original"`
				MapAboveOriginalPrice interface{} `json:"mapAboveOriginalPrice"`
				Message               interface{} `json:"message"`
				PreferredPriceFlag    bool        `json:"preferredPriceFlag"`
				SpecialBuy            interface{} `json:"specialBuy"`
				UnitOfMeasure         string      `json:"unitOfMeasure"`
				Typename              string      `json:"__typename"`
			} `json:"pricing"`
			Reviews struct {
				RatingsReviews struct {
					AverageRating string `json:"averageRating"`
					TotalReviews  string `json:"totalReviews"`
					Typename      string `json:"__typename"`
				} `json:"ratingsReviews"`
				Typename string `json:"__typename"`
			} `json:"reviews"`
			Seo struct {
				SeoKeywords    interface{} `json:"seoKeywords"`
				SeoDescription interface{} `json:"seoDescription"`
				Typename       string      `json:"__typename"`
			} `json:"seo"`
			SpecificationGroup []struct {
				Specifications []struct {
					SpecName  string `json:"specName"`
					SpecValue string `json:"specValue"`
					Typename  string `json:"__typename"`
				} `json:"specifications"`
				SpecTitle string `json:"specTitle"`
				Typename  string `json:"__typename"`
			} `json:"specificationGroup"`
			Taxonomy struct {
				BreadCrumbs []struct {
					Label           string      `json:"label"`
					URL             string      `json:"url"`
					BrowseURL       interface{} `json:"browseUrl"`
					CreativeIconURL interface{} `json:"creativeIconUrl"`
					DeselectURL     interface{} `json:"deselectUrl"`
					DimensionName   string      `json:"dimensionName"`
					RefinementKey   interface{} `json:"refinementKey"`
					Typename        string      `json:"__typename"`
				} `json:"breadCrumbs"`
				BrandLinkURL string `json:"brandLinkUrl"`
				Typename     string `json:"__typename"`
			} `json:"taxonomy"`
			FavoriteDetail struct {
				Count    int    `json:"count"`
				Typename string `json:"__typename"`
			} `json:"favoriteDetail"`
			Info struct {
				HidePrice                bool        `json:"hidePrice"`
				EcoRebate                bool        `json:"ecoRebate"`
				QuantityLimit            int         `json:"quantityLimit"`
				SskMin                   interface{} `json:"sskMin"`
				SskMax                   interface{} `json:"sskMax"`
				UnitOfMeasureCoverage    interface{} `json:"unitOfMeasureCoverage"`
				WasMaxPriceRange         interface{} `json:"wasMaxPriceRange"`
				WasMinPriceRange         interface{} `json:"wasMinPriceRange"`
				FiscalYear               string      `json:"fiscalYear"`
				ProductDepartment        string      `json:"productDepartment"`
				ClassNumber              string      `json:"classNumber"`
				ForProfessionalUseOnly   bool        `json:"forProfessionalUseOnly"`
				GlobalCustomConfigurator interface{} `json:"globalCustomConfigurator"`
				MovingCalculatorEligible bool        `json:"movingCalculatorEligible"`
				Label                    interface{} `json:"label"`
				RecommendationFlags      struct {
					VisualNavigation bool   `json:"visualNavigation"`
					ReqItems         bool   `json:"reqItems"`
					Typename         string `json:"__typename"`
				} `json:"recommendationFlags"`
				ReplacementOMSID          interface{} `json:"replacementOMSID"`
				HasSubscription           bool        `json:"hasSubscription"`
				MinimumOrderQuantity      int         `json:"minimumOrderQuantity"`
				ProjectCalculatorEligible bool        `json:"projectCalculatorEligible"`
				SubClassNumber            string      `json:"subClassNumber"`
				CalculatorType            interface{} `json:"calculatorType"`
				IsLiveGoodsProduct        bool        `json:"isLiveGoodsProduct"`
				ProtectionPlanSku         interface{} `json:"protectionPlanSku"`
				HasServiceAddOns          bool        `json:"hasServiceAddOns"`
				ConsultationType          interface{} `json:"consultationType"`
				Typename                  string      `json:"__typename"`
			} `json:"info"`
			SizeAndFitDetail struct {
				AttributeGroups []struct {
					Attributes []struct {
						AttributeName interface{} `json:"attributeName"`
						Dimensions    string      `json:"dimensions"`
						Typename      string      `json:"__typename"`
					} `json:"attributes"`
					DimensionLabel string      `json:"dimensionLabel"`
					ProductType    interface{} `json:"productType"`
					Typename       string      `json:"__typename"`
				} `json:"attributeGroups"`
				Typename string `json:"__typename"`
			} `json:"sizeAndFitDetail"`
			KeyProductFeatures struct {
				KeyProductFeaturesItems []struct {
					Features []struct {
						Name          string `json:"name"`
						RefinementID  string `json:"refinementId"`
						RefinementURL string `json:"refinementUrl"`
						Value         string `json:"value"`
						Typename      string `json:"__typename"`
					} `json:"features"`
					Typename string `json:"__typename"`
				} `json:"keyProductFeaturesItems"`
				Typename string `json:"__typename"`
			} `json:"keyProductFeatures"`
			SeoDescription  interface{}   `json:"seoDescription"`
			Badges          []interface{} `json:"badges"`
			InstallServices struct {
				ScheduleAMeasure bool   `json:"scheduleAMeasure"`
				Typename         string `json:"__typename"`
			} `json:"installServices"`
			Subscription interface{} `json:"subscription"`
			DataSource   string      `json:"dataSource"`
			Typename     string      `json:"__typename"`
		} `json:"product"`
	} `json:"data"`
}
