package homedepot

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type HomeDepotResponse struct {
	Channel      string `json:"channel"`
	FeatureState struct {
		Trackif struct {
			IsEnabled bool   `json:"isEnabled"`
			HostName  string `json:"hostName"`
			Enabled   bool   `json:"enabled"`
		} `json:"trackif"`
		BazaarVoiceFrontEnd struct {
			BazaarvoiceContent string `json:"bazaarvoiceContent"`
			BvapiURL           string `json:"bvapiUrl"`
		} `json:"bazaarVoiceFrontEnd"`
		Threesixty struct {
			IsEnabled bool   `json:"isEnabled"`
			HostName  string `json:"hostName"`
			Enabled   bool   `json:"enabled"`
		} `json:"threesixty"`
		CertonaOnDemandFeature struct {
			Enabled   bool `json:"enabled"`
			IsEnabled bool `json:"isEnabled"`
		} `json:"certonaOnDemandFeature"`
		TextAisleBayFeature struct {
			IsEnabled bool `json:"isEnabled"`
			Enabled   bool `json:"enabled"`
		} `json:"textAisleBayFeature"`
		ApplianceDeliveryChargeFeature struct {
			IsEnabled bool `json:"isEnabled"`
			Enabled   bool `json:"enabled"`
		} `json:"applianceDeliveryChargeFeature"`
		LivePersonFeature struct {
			Enabled bool `json:"enabled"`
		} `json:"livePersonFeature"`
		EnsightenFeature struct {
			IsEnabled           bool   `json:"isEnabled"`
			JavascriptURI       string `json:"javascriptUri"`
			MobileJavascriptURI string `json:"mobileJavascriptUri"`
			Enabled             bool   `json:"enabled"`
		} `json:"ensightenFeature"`
		BazaarVoiceFeature struct {
			JsAPIURL                 string `json:"jsApiUrl"`
			ReviewsBaseContentURI    string `json:"reviewsBaseContentUri"`
			QuestionAnswerSeoEnabled bool   `json:"questionAnswerSeoEnabled"`
			RatingsReviewsSeoEnabled bool   `json:"ratingsReviewsSeoEnabled"`
		} `json:"bazaarVoiceFeature"`
		PipBadgingFeature struct {
			IsEnabled bool `json:"isEnabled"`
			Enabled   bool `json:"enabled"`
		} `json:"pipBadgingFeature"`
		ExperienceSelectorFeature struct {
			IsInspirationalPIPExperienceEnabled bool `json:"isInspirationalPIPExperienceEnabled"`
			InspirationalPIPExperienceEnabled   bool `json:"inspirationalPIPExperienceEnabled"`
		} `json:"experienceSelectorFeature"`
		MccCartApplianceAvailabilityFeature struct {
			IsEnabledForDesktop bool `json:"isEnabledForDesktop"`
			IsEnabledForMobile  bool `json:"isEnabledForMobile"`
			EnabledForDesktop   bool `json:"enabledForDesktop"`
			EnabledForMobile    bool `json:"enabledForMobile"`
		} `json:"mccCartApplianceAvailabilityFeature"`
		NewAisleBayFeature         bool `json:"newAisleBayFeature"`
		BodfsAsServiceLevelFeature bool `json:"bodfsAsServiceLevelFeature"`
		ProjectCalculatorFeature   struct {
			IsFeatureEnabled bool `json:"isFeatureEnabled"`
			FeatureEnabled   bool `json:"featureEnabled"`
		} `json:"projectCalculatorFeature"`
		CertonaReferrerRecommenderFeature struct {
			IsFeatureEnabled bool `json:"isFeatureEnabled"`
			FeatureEnabled   bool `json:"featureEnabled"`
		} `json:"certonaReferrerRecommenderFeature"`
		ServiceAddOnServiceFeature                                                bool   `json:"serviceAddOnServiceFeature"`
		InspirationalReviewSnapShotFeatureEnabled                                 bool   `json:"inspirationalReviewSnapShotFeatureEnabled"`
		AttachLaborFeatureEnabled                                                 bool   `json:"attachLaborFeatureEnabled"`
		SponsoredRecommenderFeatureEnabled                                        bool   `json:"sponsoredRecommenderFeatureEnabled"`
		SponsoredRecommenderMobileFeatureEnabled                                  bool   `json:"sponsoredRecommenderMobileFeatureEnabled"`
		ApplianceBundlingFeatureEnabled                                           bool   `json:"applianceBundlingFeatureEnabled"`
		DynamicRecsAPIsEnabled                                                    bool   `json:"dynamicRecsAPIsEnabled"`
		DynamicRecsRequiredItemsEnabled                                           bool   `json:"dynamicRecsRequiredItemsEnabled"`
		FbrMobileVersionFeatureEnabled                                            bool   `json:"fbrMobileVersionFeatureEnabled"`
		IrgAccessoriesDynamicAPIEnabled                                           bool   `json:"irgAccessoriesDynamicAPIEnabled"`
		ApplianceBundlingAccessoriesFeatureEnabled                                bool   `json:"applianceBundlingAccessoriesFeatureEnabled"`
		ChatbotFeatureEnabled                                                     bool   `json:"chatbotFeatureEnabled"`
		ClientPriceLoadingEnabled                                                 bool   `json:"clientPriceLoadingEnabled"`
		VisuallySimilarFeatureEnabled                                             bool   `json:"visuallySimilarFeatureEnabled"`
		MultistoreFeatureEnabled                                                  bool   `json:"multistoreFeatureEnabled"`
		RefactorHomeServicesCoordinatorEnabled                                    bool   `json:"refactorHomeServicesCoordinatorEnabled"`
		BossDetaFeatureEnabled                                                    bool   `json:"bossDetaFeatureEnabled"`
		AisleAndBayReqInStockCheckFeatureEnabled                                  bool   `json:"aisleAndBayReqInStockCheckFeatureEnabled"`
		RenderingServiceIntegrationEnabled                                        bool   `json:"renderingServiceIntegrationEnabled"`
		NonBuyableScheduleFurnishInstallExperienceRoutedToRenderingServiceEnabled bool   `json:"nonBuyableScheduleFurnishInstallExperienceRoutedToRenderingServiceEnabled"`
		AttachLaborZipCodeEnabled                                                 bool   `json:"attachLaborZipCodeEnabled"`
		FeatureBasedRecommendationMobileEnabled                                   bool   `json:"featureBasedRecommendationMobileEnabled"`
		FeatureBasedRecommendationsMobileTestAndTarget                            bool   `json:"featureBasedRecommendationsMobileTestAndTarget"`
		UseNewGcpHostnameForImagesEnabled                                         bool   `json:"useNewGcpHostnameForImagesEnabled"`
		AutoLocalizationBackendCleanUpFeatureEnabled                              bool   `json:"autoLocalizationBackendCleanUpFeatureEnabled"`
		CrmInputAisleAndBayFeatureEnabled                                         bool   `json:"crmInputAisleAndBayFeatureEnabled"`
		BossSharedSkuEnhancement                                                  bool   `json:"bossSharedSkuEnhancement"`
		DataCollectionVersion                                                     string `json:"dataCollectionVersion"`
		HeaderFooterB2BStaticVersion                                              string `json:"headerFooterB2BStaticVersion"`
		HeaderFooterB2CStaticVersion                                              string `json:"headerFooterB2CStaticVersion"`
		RemoveDuplicatePCPEnabled                                                 bool   `json:"removeDuplicatePCPEnabled"`
		BossSharedSkuEnhancementB2B                                               bool   `json:"bossSharedSkuEnhancementB2B"`
		EmailErrorMessageForOOSItemsEnabled                                       bool   `json:"emailErrorMessageForOOSItemsEnabled"`
		FbrAPIModelVersionsFeatureEnabled                                         bool   `json:"fbrApiModelVersionsFeatureEnabled"`
		ConsumerCreditMessageEnabled                                              bool   `json:"consumerCreditMessageEnabled"`
		PaymentEstimatorEnabled                                                   bool   `json:"paymentEstimatorEnabled"`
		PaymentEstimatorForApplianceEnabled                                       bool   `json:"paymentEstimatorForApplianceEnabled"`
		PaymentEstimatorVersion                                                   string `json:"paymentEstimatorVersion"`
		GoogleLocalInventoryAdFeatureEnabled                                      bool   `json:"googleLocalInventoryAdFeatureEnabled"`
		DesktopSponsoredComponentsEnabled                                         bool   `json:"desktopSponsoredComponentsEnabled"`
		MobileSponsoredComponentsEnabled                                          bool   `json:"mobileSponsoredComponentsEnabled"`
		BopisStoreOffEnabled                                                      bool   `json:"bopisStoreOffEnabled"`
		ShowBannerForPaintColorWallEnabled                                        bool   `json:"showBannerForPaintColorWallEnabled"`
		MobileStickyHeaderEnabled                                                 bool   `json:"mobileStickyHeaderEnabled"`
		BulkPriceAlignEnabled                                                     bool   `json:"bulkPriceAlignEnabled"`
		VehicleFitmentFeatureEnabled                                              bool   `json:"vehicleFitmentFeatureEnabled"`
		FitmentAPIKey                                                             string `json:"fitmentApiKey"`
		SearchFeedbackEnabled                                                     bool   `json:"searchFeedbackEnabled"`
		SeoInternalLinks                                                          struct {
			SeoServiceCrawlablelinksFeatureSwitch bool   `json:"SeoServiceCrawlablelinksFeatureSwitch"`
			SeoServiceCrawlablelinksURL           string `json:"SeoServiceCrawlablelinksUrl"`
			SeoserviceCrawlablelinksExclusions    string `json:"SeoserviceCrawlablelinksExclusions"`
			SeoServiceBucketURL                   string `json:"SeoServiceBucketUrl"`
		} `json:"seoInternalLinks"`
		NewRelicFeature         bool   `json:"newRelicFeature"`
		NewRelicThrottleValue   string `json:"newRelicThrottleValue"`
		ThdNewRelicVersion      string `json:"thdNewRelicVersion"`
		NewRelicAgentUpdate     bool   `json:"newRelicAgentUpdate"`
		PipCustomNewRelic       bool   `json:"pipCustomNewRelic"`
		LinkOverageToAPIEnabled bool   `json:"linkOverageToApiEnabled"`
		SubscribeAndSaveFeature struct {
			IsEnabled bool   `json:"isEnabled"`
			HostName  string `json:"hostName"`
			Enabled   bool   `json:"enabled"`
		} `json:"subscribeAndSaveFeature"`
		StrikeThruPricingForTntEnabled                       bool   `json:"strikeThruPricingForTntEnabled"`
		AllowVisuallySimilarSkusForFbrEnabled                bool   `json:"allowVisuallySimilarSkusForFbrEnabled"`
		InstantCheckoutFeatureEnabled                        bool   `json:"instantCheckoutFeatureEnabled"`
		ShowDefaultIRGInApplianceEnabled                     bool   `json:"showDefaultIRGInApplianceEnabled"`
		IrgShowMultiAtcButtonEnabled                         bool   `json:"irgShowMultiAtcButtonEnabled"`
		FbtDataDynamicallyFromTnTEnabled                     bool   `json:"fbtDataDynamicallyFromTnTEnabled"`
		FullSizeUserSubmittedImagesEnabled                   bool   `json:"fullSizeUserSubmittedImagesEnabled"`
		UgcProsAndConsEnabled                                bool   `json:"ugcProsAndConsEnabled"`
		ReviewsOnMobilePageLoadEnabled                       bool   `json:"reviewsOnMobilePageLoadEnabled"`
		UgcZoneAReact                                        bool   `json:"ugcZoneAReact"`
		IrgTabbedCollectionViewEnabledTestAndTarget          bool   `json:"irgTabbedCollectionViewEnabledTestAndTarget"`
		ThdPackageConfiguratorVersion                        string `json:"thdPackageConfiguratorVersion"`
		PackageConfiguratorEnabled                           bool   `json:"packageConfiguratorEnabled"`
		BopisForApplianceEnabled                             bool   `json:"bopisForApplianceEnabled"`
		PackageBundlingExperienceEnabled                     bool   `json:"packageBundlingExperienceEnabled"`
		MigrateApplianceBundleToDynamicRecsFeatureEnabled    bool   `json:"migrateApplianceBundleToDynamicRecsFeatureEnabled"`
		KitchenPackagesPackageNameTitleEnabled               bool   `json:"kitchenPackagesPackageNameTitleEnabled"`
		DetaZipDomain                                        bool   `json:"detaZipDomain"`
		LocalizationDrawerEnabled                            bool   `json:"localizationDrawerEnabled"`
		BypassBossBopisOverlayEnabled                        bool   `json:"bypassBossBopisOverlayEnabled"`
		MobileCompareEnabled                                 bool   `json:"mobileCompareEnabled"`
		InventoryUrgencyMessagingEnabled                     bool   `json:"inventoryUrgencyMessagingEnabled"`
		SskuHistoryCardsEnabled                              bool   `json:"sskuHistoryCardsEnabled"`
		SskuAvailabilityDisplayEnabled                       bool   `json:"sskuAvailabilityDisplayEnabled"`
		PaintSskuMigrationEnabled                            bool   `json:"paintSskuMigrationEnabled"`
		PaintSskuOverwriteFixEnabled                         bool   `json:"paintSskuOverwriteFixEnabled"`
		AddToCartQuantityEnabled                             bool   `json:"addToCartQuantityEnabled"`
		SskuScaledHoverTooltipsEnabled                       bool   `json:"sskuScaledHoverTooltipsEnabled"`
		ThreeSixtyImageLoadFixEnabled                        bool   `json:"threeSixtyImageLoadFixEnabled"`
		SskuDeadlockFixEnabled                               bool   `json:"sskuDeadlockFixEnabled"`
		DoItYourselfStandaloneEnabled                        bool   `json:"doItYourselfStandaloneEnabled"`
		SskuSwatchCardsEnabled                               bool   `json:"sskuSwatchCardsEnabled"`
		MinimumOrderQuantityEnabled                          bool   `json:"minimumOrderQuantityEnabled"`
		FloorCalcDimensionsEnabled                           bool   `json:"floorCalcDimensionsEnabled"`
		MeasurementGuideEnabled                              bool   `json:"measurementGuideEnabled"`
		ConfigurableCalcEnabled                              bool   `json:"configurableCalcEnabled"`
		SskuSwatchdownEnabled                                bool   `json:"sskuSwatchdownEnabled"`
		DisclaimerMessageTypeEnabled                         bool   `json:"disclaimerMessageTypeEnabled"`
		StickyMediaGalleryEnabled                            bool   `json:"stickyMediaGalleryEnabled"`
		ApplianceMobileShuffleEnabled                        bool   `json:"applianceMobileShuffleEnabled"`
		ClassificationServiceEnabled                         bool   `json:"classificationServiceEnabled"`
		PayPalPaymentOptionEnabled                           bool   `json:"payPalPaymentOptionEnabled"`
		MobileMessageCTAEnabled                              bool   `json:"mobileMessageCTAEnabled"`
		DynamicRecsBadgeAPIEnabled                           bool   `json:"dynamicRecsBadgeApiEnabled"`
		UpdatedB2BAddToListIconEnabled                       bool   `json:"updatedB2BAddToListIconEnabled"`
		CustomProductOfferingEnabledInExperience             string `json:"customProductOfferingEnabledInExperience"`
		CollectionsCustomProductOfferingEnabled              bool   `json:"collectionsCustomProductOfferingEnabled"`
		LlcMockDataServiceEnabled                            bool   `json:"llcMockDataServiceEnabled"`
		UseNewLeadGenEndpoint                                bool   `json:"useNewLeadGenEndpoint"`
		UseWWWLeadGenHost                                    bool   `json:"useWWWLeadGenHost"`
		ShowInstallationOptions                              bool   `json:"showInstallationOptions"`
		BossOffshoreEnabled                                  bool   `json:"bossOffshoreEnabled"`
		StoreMessagingEnabled                                bool   `json:"storeMessagingEnabled"`
		ReactFulfillmentTilesEnabled                         bool   `json:"reactFulfillmentTilesEnabled"`
		ReactFulfillmentTilesVersion                         string `json:"reactFulfillmentTilesVersion"`
		StoreAssemblyEnabled                                 bool   `json:"storeAssemblyEnabled"`
		DualPathForCarpetInstallationEnabled                 bool   `json:"dualPathForCarpetInstallationEnabled"`
		DualPathForHardSurfacesInstallationEnabled           bool   `json:"dualPathForHardSurfacesInstallationEnabled"`
		ScheduleMeasureIdmAttrEnabled                        bool   `json:"scheduleMeasureIdmAttrEnabled"`
		InternalRelatedSearchAndProductsServiceEnabled       bool   `json:"internalRelatedSearchAndProductsServiceEnabled"`
		InternalRelatedSearchAndProductsServiceForAllEnabled bool   `json:"internalRelatedSearchAndProductsServiceForAllEnabled"`
		BodfsDirectAddToCartOverlayAndPIPSourcingEnabled     bool   `json:"bodfsDirectAddToCartOverlayAndPIPSourcingEnabled"`
		BodfsGrillsAssemblyEnabled                           bool   `json:"bodfsGrillsAssemblyEnabled"`
		RepeatPurchaseBannerEnabled                          bool   `json:"repeatPurchaseBannerEnabled"`
		AugmentedRealityEnabled                              bool   `json:"augmentedRealityEnabled"`
		RecommendedGuidesFeatureEnabled                      bool   `json:"recommendedGuidesFeatureEnabled"`
		AddToCartBackOrderAnalytics                          bool   `json:"addToCartBackOrderAnalytics"`
		PageLoadBackOrderableAnalyticsEnabled                bool   `json:"pageLoadBackOrderableAnalyticsEnabled"`
		NearbyLimitedStockBOSSEnabled                        bool   `json:"nearbyLimitedStockBOSSEnabled"`
		SharedSkuOnlineQuantityEnabled                       bool   `json:"sharedSkuOnlineQuantityEnabled"`
		EmailOosAuthUserEnabled                              bool   `json:"emailOosAuthUserEnabled"`
		BuyOneGetOneFreeEnabled                              bool   `json:"buyOneGetOneFreeEnabled"`
		BuyOneGetOneEligibleAnchorSkus                       string `json:"buyOneGetOneEligibleAnchorSkus"`
		SpecialBuyTeaserEnabled                              bool   `json:"specialBuyTeaserEnabled"`
		EventsBadgingEnabled                                 bool   `json:"eventsBadgingEnabled"`
		ThdCustomerCookieEnabled                             bool   `json:"thdCustomerCookieEnabled"`
		ServicesLifeCycleEventFixEnabled                     bool   `json:"servicesLifeCycleEventFixEnabled"`
		LiveGoodsMessagingEnabled                            bool   `json:"liveGoodsMessagingEnabled"`
		PinchToZoomEnabled                                   bool   `json:"pinchToZoomEnabled"`
		RemoveBodfsTodayTomorrow                             bool   `json:"removeBodfsTodayTomorrow"`
		ReportLocalizerToNewRelicEnabled                     bool   `json:"reportLocalizerToNewRelicEnabled"`
		GoToTopButtonEnabled                                 bool   `json:"goToTopButtonEnabled"`
		PinchToZoomMediaGalleryEnabled                       bool   `json:"pinchToZoomMediaGalleryEnabled"`
		CaptureDeliveryZipEnabled                            bool   `json:"captureDeliveryZipEnabled"`
		MultiSkuBulkPromotionEnabled                         bool   `json:"multiSkuBulkPromotionEnabled"`
		MultiSkuBulkPromotionUIEnabled                       bool   `json:"multiSkuBulkPromotionUIEnabled"`
		DisableBuildCookieEnabled                            bool   `json:"disableBuildCookieEnabled"`
		ClickableBrandLinkFeatureEnabled                     bool   `json:"clickableBrandLinkFeatureEnabled"`
		KitchenPackagesAnalyticsEnabled                      bool   `json:"kitchenPackagesAnalyticsEnabled"`
		HomeServicesCoordinatorEnabled                       bool   `json:"homeServicesCoordinatorEnabled"`
		BodfsOverlayTextFeatureEnabled                       bool   `json:"bodfsOverlayTextFeatureEnabled"`
		RichContentMobileEnabled                             bool   `json:"richContentMobileEnabled"`
		BazzarVoiceReactComponentEnabled                     bool   `json:"bazzarVoiceReactComponentEnabled"`
		CustomProductOfferingEnabled                         bool   `json:"customProductOfferingEnabled"`
		IsBopisCutoffEnabled                                 bool   `json:"isBopisCutoffEnabled"`
		BodfsThresholdTotal                                  string `json:"bodfsThresholdTotal"`
		BopisThresholdTotal                                  string `json:"bopisThresholdTotal"`
		BopisThresholdStores                                 string `json:"bopisThresholdStores"`
		KitchenPackagesExperienceFeatureEnabled              bool   `json:"kitchenPackagesExperienceFeatureEnabled"`
		ServerSideIrgEnabled                                 bool   `json:"serverSideIrgEnabled"`
	} `json:"featureState"`
	Fulfillment struct {
		Fulfillable bool `json:"fulfillable"`
		Store       struct {
			BuyOnlinePickUpInStore struct {
				CheckLocalStores            bool `json:"checkLocalStores"`
				IsRestricted                bool `json:"isRestricted"`
				IsStoreAssortmentRestricted bool `json:"isStoreAssortmentRestricted"`
				IsInventoryRestricted       bool `json:"isInventoryRestricted"`
				IsAlphaPromptRestricted     bool `json:"isAlphaPromptRestricted"`
			} `json:"buyOnlinePickUpInStore"`
			PurchasedInStoreOnly bool `json:"purchasedInStoreOnly"`
		} `json:"store"`
		Shipping struct {
			ShipToHome struct {
				HasFreeShipping                          bool   `json:"hasFreeShipping"`
				FreeShippingThreshhold                   string `json:"freeShippingThreshhold"`
				ConversationalEstimatedShippingStartDate string `json:"conversationalEstimatedShippingStartDate"`
				ConversationalEstimatedShippingEndDate   string `json:"conversationalEstimatedShippingEndDate"`
				EstimatedShippingStartDate               string `json:"estimatedShippingStartDate"`
				EstimatedShippingEndDate                 string `json:"estimatedShippingEndDate"`
				RawEstimatedShippingStartDate            string `json:"rawEstimatedShippingStartDate"`
				RawEstimatedShippingEndDate              string `json:"rawEstimatedShippingEndDate"`
			} `json:"shipToHome"`
		} `json:"shipping"`
		BossEstimatedShippingStartDate   string `json:"BossEstimatedShippingStartDate"`
		BossEstimatedShippingEndDate     string `json:"BossEstimatedShippingEndDate"`
		IsZeroStockInLocalStoreBoss      bool   `json:"isZeroStockInLocalStoreBoss"`
		IsPurchasableInLocalStoreBoss    bool   `json:"isPurchasableInLocalStoreBoss"`
		IsPurchasableInLocalStore        bool   `json:"isPurchasableInLocalStore"`
		IsAssortedToLocalStore           bool   `json:"isAssortedToLocalStore"`
		IsPurchasableInLocalStoreInStock bool   `json:"isPurchasableInLocalStoreInStock"`
		IsZeroStockInLocalStoreActive    bool   `json:"isZeroStockInLocalStoreActive"`
	} `json:"fulfillment"`
	LocalStore struct {
		ZipCode              string `json:"zipCode"`
		BopisEligibilityFlag string `json:"bopisEligibilityFlag"`
		BossElgFlg           string `json:"bossElgFlg"`
		BossMsgFlg           string `json:"bossMsgFlg"`
		BodfsEligibilityFlag string `json:"bodfsEligibilityFlag"`
		Phone                string `json:"phone"`
		StoreName            string `json:"storeName"`
		StoreTimeZone        string `json:"storeTimeZone"`
		AssemblyFlag         string `json:"assemblyFlag"`
		CurbsidePickupFlag   string `json:"curbsidePickupFlag"`
		Storenumber          string `json:"storenumber"`
		Marketid             string `json:"marketid"`
		Storeopendate        string `json:"storeopendate"`
		Streetname           string `json:"streetname"`
		Cityname             string `json:"cityname"`
		Statecode            string `json:"statecode"`
		Countrycode          string `json:"countrycode"`
	} `json:"localStore"`
	RemoteHosts struct {
		SecureBrowse    string `json:"secureBrowse"`
		AisleBayHost    string `json:"aisleBayHost"`
		TechShedHost    string `json:"techShedHost"`
		NewTechShedHost string `json:"newTechShedHost"`
	} `json:"remoteHosts"`
	PipCheckoutOptions struct {
		IsPayPalAvailable bool `json:"isPayPalAvailable"`
	} `json:"pipCheckoutOptions"`
	ProtectionPlan struct {
		ItemID     string `json:"itemId"`
		Price      string `json:"price"`
		TermLength string `json:"termLength"`
		Category   string `json:"category"`
		DetailsURL string `json:"detailsUrl"`
	} `json:"protectionPlan"`
	ShippingExclusion struct {
		ExclusionLocations                   []string `json:"exclusionLocations"`
		BossExclusionLocations               []string `json:"bossExclusionLocations"`
		HasExclusionLocations                bool     `json:"hasExclusionLocations"`
		HasSeparateBossSthExclusionLocations bool     `json:"hasSeparateBossSthExclusionLocations"`
	} `json:"shippingExclusion"`
	CorporateInformation struct {
		FiscalYear string `json:"fiscalYear"`
	} `json:"corporateInformation"`
	InstallServices struct {
		ScheduleAMeasure         bool `json:"scheduleAMeasure"`
		ScheduleAMeasureDualPath bool `json:"scheduleAMeasureDualPath"`
		ScheduleAnAppointment    bool `json:"scheduleAnAppointment"`
	} `json:"installServices"`
	Experience                           string  `json:"experience"`
	InStoreAssemblyAvailable             bool    `json:"inStoreAssemblyAvailable"`
	HasApplianceBundle                   bool    `json:"hasApplianceBundle"`
	OveragePercentage                    float64 `json:"overagePercentage"`
	ProjectCalculatorEligible            bool    `json:"projectCalculatorEligible"`
	OverageMultiplierEligible            bool    `json:"overageMultiplierEligible"`
	MovingCalculatorEligible             bool    `json:"movingCalculatorEligible"`
	Env                                  string  `json:"env"`
	BazaarVoiceFrontEndKey               string  `json:"bazaarVoiceFrontEndKey"`
	ShowBopisExperienceForMajorAppliance bool    `json:"showBopisExperienceForMajorAppliance"`
	B2BCustomer                          bool    `json:"b2BCustomer"`
	AnchorStore                          bool    `json:"anchorStore"`
	ExperienceFlavor                     string  `json:"experienceFlavor"`
	PrimaryItemData                      struct {
		ItemID           string `json:"itemId"`
		ItemType         string `json:"itemType"`
		AvailabilityType string `json:"availabilityType"`
		PartNumber       string `json:"partNumber"`
		WebURL           string `json:"webUrl"`
		CanonicalURL     string `json:"canonicalURL"`
		ItemAvailability struct {
			Buyable              bool `json:"buyable"`
			AvailableOnlineStore bool `json:"availableOnlineStore"`
			AvailableInStore     bool `json:"availableInStore"`
			InventoryStatus      bool `json:"inventoryStatus"`
			Backorderable        bool `json:"backorderable"`
			Published            bool `json:"published"`
			DiscontinuedItem     bool `json:"discontinuedItem"`
		} `json:"itemAvailability"`
		AttributeGroups []struct {
			GroupType string `json:"groupType"`
			Entries   []struct {
				Name         string `json:"name"`
				Value        string `json:"value"`
				GUID         string `json:"guid"`
				BulletedAttr bool   `json:"bulletedAttr"`
			} `json:"entries"`
		} `json:"attributeGroups"`
		StoreSkus []struct {
			StoreID             string `json:"storeId"`
			StoreStatus         bool   `json:"storeStatus"`
			FulfillmentOptions2 struct {
				Fulfillable            bool `json:"fulfillable"`
				BuyOnlinePickupInStore struct {
					Status            bool `json:"status"`
					CheckNearbyStores bool `json:"checkNearbyStores"`
				} `json:"buyOnlinePickupInStore"`
				ShipToHome struct {
					Status bool `json:"status"`
				} `json:"shipToHome"`
			} `json:"fulfillmentOptions,omitempty"`
			StoreAvailability2 struct {
				AvailableInLocalStore   bool `json:"availableInLocalStore"`
				ItemAvailable           bool `json:"itemAvailable"`
				BopisElgStore           bool `json:"bopisElgStore"`
				ItemAvilabilityMessages []struct {
					MessageKey   string `json:"messageKey"`
					MessageValue string `json:"messageValue"`
				} `json:"itemAvilabilityMessages"`
			} `json:"storeAvailability,omitempty"`
			ItemID          string `json:"itemId"`
			LowerPriceFlag  bool   `json:"lowerPriceFlag"`
			StoreStatusType string `json:"storeStatusType"`
			Pricing         struct {
				Uom                     string  `json:"uom"`
				UnitsPerCase            string  `json:"unitsPerCase"`
				OriginalPrice           float64 `json:"originalPrice"`
				SpecialPrice            float64 `json:"specialPrice"`
				LowerPrice              bool    `json:"lowerPrice"`
				DollarOff               float64 `json:"dollarOff"`
				PercentageOff           float64 `json:"percentageOff"`
				PriceType               string  `json:"priceType"`
				ItemOnSale              bool    `json:"itemOnSale"`
				AlternatePricingDisplay bool    `json:"alternatePricingDisplay"`
			} `json:"pricing"`
			Inventory struct {
				OnHandQuantity            int `json:"onHandQuantity"`
				ExpectedQuantityAvailable int `json:"expectedQuantityAvailable"`
			} `json:"inventory"`
			FulfillmentOptions struct {
				Fulfillable bool `json:"fulfillable"`
				ShipToHome  struct {
					Status bool `json:"status"`
				} `json:"shipToHome"`
			} `json:"fulfillmentOptions,omitempty"`
			StoreAvailability struct {
				AvailableInLocalStore   bool `json:"availableInLocalStore"`
				ItemAvailable           bool `json:"itemAvailable"`
				ItemAvilabilityMessages []struct {
					MessageKey   string `json:"messageKey"`
					MessageValue string `json:"messageValue"`
				} `json:"itemAvilabilityMessages"`
			} `json:"storeAvailability,omitempty"`
		} `json:"storeSkus"`
		LowerPriceFlag bool `json:"lowerPriceFlag"`
		Rebates        struct {
			HasEcoRebates bool `json:"hasEcoRebates"`
		} `json:"rebates"`
		ProtectionPlanSku string `json:"protectionPlanSku"`
		ItemExtension     struct {
			Pricing struct {
				Uom                     string  `json:"uom"`
				UnitsPerCase            string  `json:"unitsPerCase"`
				OriginalPrice           float64 `json:"originalPrice"`
				SpecialPrice            float64 `json:"specialPrice"`
				LowerPrice              bool    `json:"lowerPrice"`
				DollarOff               float64 `json:"dollarOff"`
				PercentageOff           float64 `json:"percentageOff"`
				PriceType               string  `json:"priceType"`
				ItemOnSale              bool    `json:"itemOnSale"`
				AlternatePricingDisplay bool    `json:"alternatePricingDisplay"`
			} `json:"pricing"`
			RichContentDisplayMode      string `json:"richContentDisplayMode"`
			FunctionalDetailsAttributes []struct {
				Name                string `json:"name"`
				Value               string `json:"value"`
				GUID                string `json:"guid"`
				BulletedAttr        bool   `json:"bulletedAttr"`
				ItemPropValueMapKey string `json:"itemPropValueMapKey,omitempty"`
			} `json:"functionalDetailsAttributes"`
			SupplementalDimensionsAttributes []struct {
				Name         string `json:"name"`
				Value        string `json:"value"`
				GUID         string `json:"guid"`
				BulletedAttr bool   `json:"bulletedAttr"`
			} `json:"supplementalDimensionsAttributes"`
			CategoryID                string `json:"categoryId"`
			ItemCategoryURL           string `json:"itemCategoryUrl"`
			CategoryName              string `json:"categoryName"`
			LocalStoreID              string `json:"localStoreId"`
			B2BOnlyItem               bool   `json:"b2bOnlyItem"`
			QuoteCenterSku            bool   `json:"quoteCenterSku"`
			ScheduleFurnishInstallSku bool   `json:"scheduleFurnishInstallSku"`
			SuperSkuItem              bool   `json:"superSkuItem"`
			BrowseOnlyItem            bool   `json:"browseOnlyItem"`
			OosOnline                 bool   `json:"oosOnline"`
			Merchandise               bool   `json:"merchandise"`
			MajorAppliance            bool   `json:"majorAppliance"`
			Buyable                   bool   `json:"buyable"`
			BackorderedOnline         bool   `json:"backorderedOnline"`
			OnlineStoreSku            struct {
				StoreID            string `json:"storeId"`
				StoreStatus        bool   `json:"storeStatus"`
				FulfillmentOptions struct {
					Fulfillable bool `json:"fulfillable"`
					ShipToHome  struct {
						Status bool `json:"status"`
					} `json:"shipToHome"`
				} `json:"fulfillmentOptions"`
				StoreAvailability struct {
					AvailableInLocalStore   bool `json:"availableInLocalStore"`
					ItemAvailable           bool `json:"itemAvailable"`
					ItemAvilabilityMessages []struct {
						MessageKey   string `json:"messageKey"`
						MessageValue string `json:"messageValue"`
					} `json:"itemAvilabilityMessages"`
				} `json:"storeAvailability"`
				ItemID          string `json:"itemId"`
				LowerPriceFlag  bool   `json:"lowerPriceFlag"`
				StoreStatusType string `json:"storeStatusType"`
				Pricing         struct {
					Uom                     string  `json:"uom"`
					UnitsPerCase            string  `json:"unitsPerCase"`
					OriginalPrice           float64 `json:"originalPrice"`
					SpecialPrice            float64 `json:"specialPrice"`
					LowerPrice              bool    `json:"lowerPrice"`
					DollarOff               float64 `json:"dollarOff"`
					PercentageOff           float64 `json:"percentageOff"`
					PriceType               string  `json:"priceType"`
					ItemOnSale              bool    `json:"itemOnSale"`
					AlternatePricingDisplay bool    `json:"alternatePricingDisplay"`
				} `json:"pricing"`
				Inventory struct {
					OnHandQuantity            int `json:"onHandQuantity"`
					ExpectedQuantityAvailable int `json:"expectedQuantityAvailable"`
				} `json:"inventory"`
			} `json:"onlineStoreSku"`
			OnlineOnlyItem      bool `json:"onlineOnlyItem"`
			LocalStoreInventory struct {
				OnHandQuantity            int `json:"onHandQuantity"`
				ExpectedQuantityAvailable int `json:"expectedQuantityAvailable"`
			} `json:"localStoreInventory"`
			AssortedToLocalStore                              bool `json:"assortedToLocalStore"`
			BuyOnlinePickUpInStoreProductAPIFulfillmentOption struct {
				Status            bool `json:"status"`
				CheckNearbyStores bool `json:"checkNearbyStores"`
			} `json:"buyOnlinePickUpInStoreProductApiFulfillmentOption"`
			ShipToHomeProductAPIFulfillmentOption struct {
				Status bool `json:"status"`
			} `json:"shipToHomeProductApiFulfillmentOption"`
			LocalStoreSku struct {
				StoreID            string `json:"storeId"`
				StoreStatus        bool   `json:"storeStatus"`
				FulfillmentOptions struct {
					Fulfillable            bool `json:"fulfillable"`
					BuyOnlinePickupInStore struct {
						Status            bool `json:"status"`
						CheckNearbyStores bool `json:"checkNearbyStores"`
					} `json:"buyOnlinePickupInStore"`
					ShipToHome struct {
						Status bool `json:"status"`
					} `json:"shipToHome"`
				} `json:"fulfillmentOptions"`
				StoreAvailability struct {
					AvailableInLocalStore   bool `json:"availableInLocalStore"`
					ItemAvailable           bool `json:"itemAvailable"`
					BopisElgStore           bool `json:"bopisElgStore"`
					ItemAvilabilityMessages []struct {
						MessageKey   string `json:"messageKey"`
						MessageValue string `json:"messageValue"`
					} `json:"itemAvilabilityMessages"`
				} `json:"storeAvailability"`
				ItemID          string `json:"itemId"`
				LowerPriceFlag  bool   `json:"lowerPriceFlag"`
				StoreStatusType string `json:"storeStatusType"`
				Pricing         struct {
					Uom                     string  `json:"uom"`
					UnitsPerCase            string  `json:"unitsPerCase"`
					OriginalPrice           float64 `json:"originalPrice"`
					SpecialPrice            float64 `json:"specialPrice"`
					LowerPrice              bool    `json:"lowerPrice"`
					DollarOff               float64 `json:"dollarOff"`
					PercentageOff           float64 `json:"percentageOff"`
					PriceType               string  `json:"priceType"`
					ItemOnSale              bool    `json:"itemOnSale"`
					AlternatePricingDisplay bool    `json:"alternatePricingDisplay"`
				} `json:"pricing"`
				Inventory struct {
					OnHandQuantity            int `json:"onHandQuantity"`
					ExpectedQuantityAvailable int `json:"expectedQuantityAvailable"`
				} `json:"inventory"`
			} `json:"localStoreSku"`
			DisplayPrice           float64 `json:"displayPrice"`
			ConfigurableBlind      bool    `json:"configurableBlind"`
			QuanityType            string  `json:"quanityType"`
			BtnBlock               string  `json:"btnBlock"`
			IsMapPriceScenario     bool    `json:"isMapPriceScenario"`
			IsMstPriceScenario     bool    `json:"isMstPriceScenario"`
			IsValuePricingScenario bool    `json:"isValuePricingScenario"`
			Apiversion             string  `json:"apiversion"`
			SharedItem             bool    `json:"sharedItem"`
			OnlineStoreAvailMsgs   []struct {
				MessageKey   string `json:"messageKey"`
				MessageValue string `json:"messageValue"`
			} `json:"onlineStoreAvailMsgs"`
			BareCushion                   bool `json:"bareCushion"`
			BossDominant                  bool `json:"bossDominant"`
			OnlineStoreFulfillmentOptions struct {
				Fulfillable bool `json:"fulfillable"`
				ShipToHome  struct {
					Status bool `json:"status"`
				} `json:"shipToHome"`
			} `json:"onlineStoreFulfillmentOptions"`
			RegularBlindsOrShadesItem    bool `json:"regularBlindsOrShadesItem"`
			KitchenOrBathFaucetsItem     bool `json:"kitchenOrBathFaucetsItem"`
			MajorAppliancePackageBundle  bool `json:"majorAppliancePackageBundle"`
			AreaRugMeasurementAssistItem bool `json:"areaRugMeasurementAssistItem"`
		} `json:"itemExtension"`
		PartialData bool   `json:"partialData"`
		ProductID   string `json:"productId"`
		Info        struct {
			BrandName                      string  `json:"brandName"`
			ModelNumber                    string  `json:"modelNumber"`
			VendorNumber                   string  `json:"vendorNumber"`
			ProductDepartment              string  `json:"productDepartment"`
			ForProfessionalUseOnly         bool    `json:"forProfessionalUseOnly"`
			ShipType                       int     `json:"shipType"`
			StoreSkuNumber                 string  `json:"storeSkuNumber"`
			SpecialOrderSKU                string  `json:"specialOrderSKU"`
			ShowProduct                    bool    `json:"showProduct"`
			OnlineStatus                   bool    `json:"onlineStatus"`
			BackorderFlag                  bool    `json:"backorderFlag"`
			GenericBrandFlag               bool    `json:"genericBrandFlag"`
			ShowLocalPrice                 bool    `json:"showLocalPrice"`
			HasIrgItems                    bool    `json:"hasIrgItems"`
			HasFbtItems                    bool    `json:"hasFbtItems"`
			HasFbrItems                    bool    `json:"hasFbrItems"`
			RecommendationFeatures         string  `json:"recommendationFeatures"`
			DiyTreatment                   bool    `json:"diyTreatment"`
			Description                    string  `json:"description"`
			ProductLabel                   string  `json:"productLabel"`
			Upc                            string  `json:"upc"`
			ClassNumber                    string  `json:"classNumber"`
			SubClassNumber                 string  `json:"subClassNumber"`
			IsAppliance                    bool    `json:"isAppliance"`
			IsTopSeller                    bool    `json:"isTopSeller"`
			BuyOnlinePickupInStoreEligible bool    `json:"buyOnlinePickupInStoreEligible"`
			BuyOnlineShipToStoreEligible   bool    `json:"buyOnlineShipToStoreEligible"`
			OmsThdSku                      string  `json:"omsThdSku"`
			OveragePercentage              float64 `json:"overagePercentage"`
			IsAutomotiveCategory           bool    `json:"isAutomotiveCategory"`
			IsBogoAnchorItem               bool    `json:"isBogoAnchorItem"`
		} `json:"info"`
		Media struct {
			ImageEnhancementFlag bool `json:"imageEnhancementFlag"`
			MediaList            []struct {
				Location         string `json:"location,omitempty"`
				Height           string `json:"height,omitempty"`
				Width            string `json:"width,omitempty"`
				MediaType        string `json:"mediaType"`
				Title            string `json:"title,omitempty"`
				ShortDescription string `json:"shortDescription,omitempty"`
				VideoID          string `json:"videoId,omitempty"`
				Thumbnail        string `json:"thumbnail,omitempty"`
				Video            string `json:"video,omitempty"`
				VideoStill       string `json:"videoStill,omitempty"`
			} `json:"mediaList"`
		} `json:"media"`
		RatingsReviews struct {
			TotalReviews  string `json:"totalReviews"`
			HasReviews    bool   `json:"hasReviews"`
			AverageRating string `json:"averageRating"`
		} `json:"ratingsReviews"`
		Promotions struct {
		} `json:"promotions"`
		Shipping struct {
			EligibleForFreeShipping        bool   `json:"eligibleForFreeShipping"`
			ItemLevelFreeShip              bool   `json:"itemLevelFreeShip"`
			ExcludedShipStates             string `json:"excludedShipStates"`
			FreeShippingMessageNumber      string `json:"freeShippingMessageNumber"`
			FreeShippingThreshold          string `json:"freeShippingThreshold"`
			FreeShippingMessage            string `json:"freeShippingMessage"`
			SthEstimatedShippingStartDate  string `json:"sthEstimatedShippingStartDate"`
			SthEstimatedShippingEndDate    string `json:"sthEstimatedShippingEndDate"`
			BossEstimatedShippingStartDate string `json:"bossEstimatedShippingStartDate"`
			BossEstimatedShippingEndDate   string `json:"bossEstimatedShippingEndDate"`
		} `json:"shipping"`
		Dimensions []struct {
			DimensionName string `json:"dimensionName"`
			DimensionID   int    `json:"dimensionId"`
			Ancestors     []struct {
				Name string `json:"name"`
				ID   int    `json:"id"`
			} `json:"ancestors"`
			DimensionValue2 struct {
				Name string `json:"name"`
				ID   int64  `json:"id"`
			} `json:"dimensionValue,omitempty"`
			IsDefaultBreadCrumb bool `json:"isDefaultBreadCrumb,omitempty"`
			DimensionValue      struct {
				Name string `json:"name"`
				ID   int64  `json:"id"`
				URL  string `json:"url"`
			} `json:"dimensionValue,omitempty"`
		} `json:"dimensions"`
		DefiningAttributes []struct {
			AttributeName        string `json:"attributeName"`
			AttributeValue       string `json:"attributeValue"`
			AttributeValueSeqNum int    `json:"attributeValueSeqNum"`
		} `json:"definingAttributes"`
		OveragePercentage float64 `json:"overagePercentage"`
		MsbPromotions     struct {
			PromotionEntry []struct {
				Reward struct {
					Tier []struct {
						Tier                 int     `json:"tier"`
						MinPurchaseQuantity  float64 `json:"minPurchaseQuantity"`
						MaxPurchaseQuantity  float64 `json:"maxPurchaseQuantity"`
						RewardPercent        float64 `json:"rewardPercent"`
						MinPurchaseAmount    float64 `json:"minPurchaseAmount"`
						RewardAmountPerItem  float64 `json:"rewardAmountPerItem"`
						RewardAmountPerOrder float64 `json:"rewardAmountPerOrder"`
					} `json:"tier"`
				} `json:"reward"`
				DiscountEndDate      string `json:"discountEndDate"`
				PromoLongDescription string `json:"promoLongDescription,omitempty"`
				Name                 string `json:"name"`
				ShortDescription     string `json:"shortDescription,omitempty"`
				DiscountStartDate    string `json:"discountStartDate"`
				PromotionID          int    `json:"promotionId"`
			} `json:"promotionEntry"`
		} `json:"msbPromotions"`
	} `json:"primaryItemData"`
	Inventory struct {
		Online struct {
			Quantity      int  `json:"quantity"`
			IsOutOfStock  bool `json:"isOutOfStock"`
			IsUnavailable bool `json:"isUnavailable"`
			IsInStock     bool `json:"isInStock"`
			IsBackordered bool `json:"isBackordered"`
		} `json:"online"`
		Store struct {
			Quantity          int  `json:"quantity"`
			IsOutOfStock      bool `json:"isOutOfStock"`
			IsLimitedQuantity bool `json:"isLimitedQuantity"`
			IsUnavailable     bool `json:"isUnavailable"`
			IsInStock         bool `json:"isInStock"`
		} `json:"store"`
	} `json:"inventory"`
}

func GetHomeDepotMultipleStoreData(requestBody [][]string) [][]interface{} {
	currenttime := time.Now()
	epoch := currenttime.Unix()
	var finalValues [][]interface{}
	for i := range requestBody {
		fmt.Println(i)
		storeID := requestBody[i][0]
		Cookie := "HD_DC=origin; bm_sz=D933413D6B61E866E16EEEBBBD5D555F~YAAQyN44fXtwtdhzAQAAztf82AjnLwoltXV6mpwDLtu0mxeGsnZULbuXY7+rrBAN8SGLGzmJagTwwbMomHLd9LH5bT9AYEr/2KMQegiKivy+GIZ8BVN6aw5eNEV8x+xeNmK44C2Pq3F+nwGLmrPtcIKjfsqnApYGrqKlqdjER6o/wy3XGHw8YuPZwOBpWXDwxRt8; THD_NR=1; THD_SESSION=; THD_CACHE_NAV_SESSION=; THD_CACHE_NAV_PERSIST=; check=true; AMCVS_F6421253512D2C100A490D45%40AdobeOrg=1; WORKFLOW=GEO_LOCATION; THD_FORCE_LOC=1; THD_INTERNAL=0; DELIVERY_ZIP=96913; THD_PERSIST=C4%3D" + requestBody[i][0] + "%2BGuam%20-%20Tamuning%20-%20Tamuning%2C%20GU%2B%3A%3BC4_EXP%3D1628609317%3A%3BC24%3D96913%3A%3BC24_EXP%3D1628609317%3A%3BC39%3D1%3B7%3A00-20%3A00%3B2%3B6%3A00-21%3A00%3B3%3B6%3A00-21%3A00%3B4%3B6%3A00-21%3A00%3B5%3B6%3A00-21%3A00%3B6%3B6%3A00-21%3A00%3B7%3B6%3A00-21%3A00%3A%3BC39_EXP%3D1597076917; THD_LOCALIZER=%7B%22WORKFLOW%22%3A%22GEO_LOCATION%22%2C%22THD_FORCE_LOC%22%3A%221%22%2C%22THD_INTERNAL%22%3A%220%22%2C%22THD_STRFINDERZIP%22%3A%2296913%22%2C%22THD_LOCSTORE%22%3A%221710%2BGuam%20-%20Tamuning%20-%20Tamuning%2C%20GU%2B%22%2C%22THD_STORE_HOURS%22%3A%221%3B7%3A00-20%3A00%3B2%3B6%3A00-21%3A00%3B3%3B6%3A00-21%3A00%3B4%3B6%3A00-21%3A00%3B5%3B6%3A00-21%3A00%3B6%3B6%3A00-21%3A00%3B7%3B6%3A00-21%3A00%22%2C%22THD_STORE_HOURS_EXPIRY%22%3A1597076917%7D; _pxvid=291e881c-db1e-11ea-b4bf-0242ac120005; thda.s=72cb0c5c-3dda-a400-f2f5-93f7116d1e63; thda.u=a37756ef-14a8-ed1c-3a67-3a2b717565fb; ak_bmsc=89609659092818FF7E985C3664D7194F7D38DEC86C5F0000A367315F2F2C413B~plb62Mk1slFeK9p8TVLPPQJFbQcEb3lo3NFcbSSoFz7r049tRe+FqubdAxjzlcukyN5/rL9jk0ntf4oXbdtsnJz9tE/siUbOHx8VdtLc95QalM+bE42AqFSU+dTm131uISgALHoziV3+8lJgC7UgJxvw6u4PaaKeeWQJh3pvhF8rOilozHsG79q+k+wGm9j+hNdTP4BRrali+5ppT8OwDG6TleNyvvp3dANYHO4hXtZb4SEcQHDp76AlglOb/VDl3E; _abck=CE6278AB62773D597362F7D2B3DEBC0E~0~YAAQyN44fYlwtdhzAQAABOD82AT1nVwvlPbMn09bgoV4EfbRuZeX1y4NyhKoYy9vz0keGx2Gp+FIo4Bn21ylgAocFRDFwjGma51lpTWbWsNQ8CkW+F+4MdJsFZem+DoO+aS74NkwXcvtZtuLbIerSpTKx6oI+9/8i4hnonCvf/j40RRQwf/8ORVqGzx5byjNc5GoyTqMfdwvDIPg8sgfKvUkv9D7BHRUKF3botyZBRTf7zihiK7CCs15v3lLlgsw+31u0vIP52F4NUb8PjKG2OfDvNb3rNA0ufI0qbdNJdKS/Qdaf2Zz8EEiN5+RSc5J9jEinPavhUJs+g==~-1~-1~-1; _px_f394gi7Fvmc43dfg_user_id=Mjk5NGEyNTEtZGIxZS0xMWVhLTk0ODktNmZlNzhlODI0ZTE1; RES_TRACKINGID=36632330381337079; ResonanceSegment=; RES_SESSIONID=72154150381337079; ajs_user_id=null; ajs_group_id=null; ajs_anonymous_id=%22d9ab20d3-708e-4e9d-884f-2b980df9a686%22; _meta_bing_beaconFired=true; _meta_facebookPixel_beaconFired=true; _gcl_au=1.1.1437797179.1597073320; ftr_ncd=6; QSI_HistorySession=https%3A%2F%2Fwww.homedepot.com%2Fp%2FMilwaukee-M18-FUEL-120-MPH-450-CFM-18-Volt-Lithium-Ion-Brushless-Cordless-Handheld-Blower-Tool-Only-2724-20%2F302752040~1597073320612; AMCV_F6421253512D2C100A490D45%40AdobeOrg=1585540135%7CMCIDTS%7C18485%7CMCMID%7C35670817154508997331213643943760871573%7CMCAAMLH-1597678120%7C12%7CMCAAMB-1597678120%7CRKhpRz8krg2tLO6pguXWp5olkAcUniQYPHaMWWgdJ3xzPWQmdj0y%7CMCOPTOUT-1597080520s%7CNONE%7CMCCIDH%7C172712820%7CvVersion%7C4.4.0; _meta_pinterest_epik=dj0yJnU9d05XaVI3TWlmLVQ0QnFTeHVkRkNnS2toNk1TbHkyQTcmbj1JTlRtQUpFWkZCbGMxTWR3R3BxNVBBJm09MSZ0PUFBQUFBRjh4WjZjJnJtPTEmcnQ9QUFBQUFGOHhaNmM; _meta_revjet_revjet_vid=4961946459097064000; aam_mobile=seg%3D1131078; aam_uuid=35959978601500025281202714362026421984; _meta_criteo_userId=Rj-FFdPI4xtnbz8-uKO266Q5vKUgmWks; _meta_mediaMath_mm_id=d0eb5d22-f864-4d00-b015-f81e4c80ae45; _meta_mediaMath_cid=d0eb5d22-f864-4d00-b015-f81e4c80ae45; _meta_amnet_uid=5573507722726040061; _ga=GA1.2.633436542.1597073321; _gid=GA1.2.568460109.1597073321; QuantumMetricUserID=2fcb28195be36292240261fa3a3fb47d; QuantumMetricSessionID=c22a75f61f04d3f1219648584f95b948; LPVID=QyZTAyZjMxNGI3ZWY3MDBk; LPSID-31564604=u8lPe1sOQfmF6g-G1Gmf9A; _meta_adobe_aam_uuid=35959978601500025281202714362026421984; mbox=session#6803e51076de49bb974712cb45d1cf94#1597075276|PC#6803e51076de49bb974712cb45d1cf94.31_0#1660318117; _px=H5TtyrokynGm0xZ8VnbRvIRtYIJ6MvBSOuTOQR9xIhodFc82766Yn2o7SFbGf1GfZoq1xP0wktwIg0O+G/BNQA==:1000:kAHErd/1qWUpNn5bQLbtw9r1mpBD95/lhqx/z25rAwnTvlmUaRq1sUf4LYiF7fyuGlsPfGC3Zmln6DcDCSrkcER7DTj9448xXSYNl6lsbfkjYY179JPwlpRccdSKClTzmWOa+Fvcmp476p0z09DyP4m0VdyGvxJgNjQ76bWwfDOZ5h36aweu3OMrJhb1vlrqbvbx80Wg3F1nHQAKJW/joVoiXuG9bd0+SK3At3IwfNazpdFINnBsjFI1BhHdEV5ey67DFM5q2o6VpDx/Ed7QYA==; s_pers=%20productnum%3D1%7C1599665318868%3B%20s_nr365%3D1597073417914-New%7C1628609417914%3B%20s_dslv%3D1597073417921%7C1691681417921%3B; s_sess=%20stsh%3D%3B%20s_pv_pName%3Dproductdetails%253E302752040%3B%20s_pv_pType%3Dpip%3B%20s_pv_cmpgn%3D%3B%20s_cc%3Dtrue%3B; _meta_mediaMath_iframe_counter=3; forterToken=640fce6dab2d475c91c80f273376b9ed_1597073417951__UDF43_9ck; _px_4946459675_cs=eyJpZCI6IjI5OTQ1NDMwLWRiMWUtMTFlYS05NDg5LTZmZTc4ZTgyNGUxNSIsInN0b3JhZ2UiOnt9LCJleHBpcmF0aW9uIjoxNTk3MDc1NjkxODE4fQ==; _pxde=6e624cc55fc1a20bd3fb6d477a0232a81102f7d4d84b291714a5942fca59d955:eyJ0aW1lc3RhbXAiOjE1OTcwNzM4OTMzMTF9; akaau=1597074193~id=7585bff3cb571847323335006e1ee492; bm_sv=67B4BD82E3C07565D9C966B534F53950~SRQIMfQu0qB+XuC4kxuDGg1mtgowuHmZFA1ijc4z2RT6fwwaxeXFuKICPhy63+0318AJAAcxo/7PsPaoQ7Sg/jGh8l+icBOajdURDqPQDKAczQY3bV5dG02O+ojfzojJ17OYs876L6XWWPOAxC0c8ypv1+5xp60X7iRJfuPVfrk=; IN_STORE_API_SESSION=TRUE; akaau=1597074756~id=570e6607324d97f43aff334a687ef720; bm_sv=67B4BD82E3C07565D9C966B534F53950~SRQIMfQu0qB+XuC4kxuDGg1mtgowuHmZFA1ijc4z2RT6fwwaxeXFuKICPhy63+0318AJAAcxo/7PsPaoQ7Sg/jGh8l+icBOajdURDqPQDKB01Zf137gHE7JiOZKKo/sLZ5v33xatSoFD0VeF/1OvvX3lI2J8uwwHSHHmyTujUYw=; _abck=CE6278AB62773D597362F7D2B3DEBC0E~-1~YAAQyN44fZ+JtdhzAQAA0T0O2QSpbesPatxFjVR5gmDxtyb1Pmp5BnIeh/+ggFLpWb/whyotqoyArNJSunSQQayxXQVeXIjSi89P4zO8QcMSXRLvADNin46chyAG0jGiQNpG/YvUuLK+DAopT5NO9NvQ2mj7WOqI/FoAwoEjo6Ap2e0KOjllQyhShzo6rImJEnhvaG7xtHxte0yul/OTIahfnwfXGm3DNkacQ5xS4H97HNpnv54PrWSgi+Jmybq5bPrKmKuffIhjS3DDCzbVZ4DZ1uaV52d61LvNuckShPu0+qRjwM0L8dr39N35ZpUEPatLOLfADrK3dg==~0~-1~-1; akaau=1597160858~id=5d84ab426f4dc213c48a00d290401967; _abck=CE6278AB62773D597362F7D2B3DEBC0E~-1~YAAQT9cLF5Ss/45zAQAA3Q8w3gSR0CK+O35uM2K5kZcckE1rq9bg7JHcnFRMZIhKwxuIuTMIx+N6S3MMwU7pbVEQHRxk2+hVw9liPqz/rB2OBPicKXxRwT/qK78fQZqPQQklPInnLwqsPLkuOEGbnfXy/RfKQrypOZWkTyBt1iDm33beCtjInYzQs+skq+2egELaX1D2H2Mp+qC10LSp2mjb2tfxPHKu4Fj4i+Orj7tAEfQ0I19VVbHEaqJW2RRU+C9EbpZJadLtiZm4Sxb00lYt+evrnjtfF0F0npGMz25eEjT16J3320BJKiWAN0fCcoYTarCmLVPmvg==~0~-1~-1"
		url := "https://www.homedepot.com/p/svcs/frontEndModel/" + requestBody[0][1] + "?_=" + strconv.Itoa(int(epoch)) + "000"
		method := "GET"

		client := &http.Client{}
		req, err := http.NewRequest(method, url, nil)

		if err != nil {
			fmt.Println(err)
		}
		req.Header.Add("authority", "www.homedepot.com")
		req.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.89 Safari/537.36")
		req.Header.Add("accept", "*/*")
		req.Header.Add("sec-fetch-site", "same-origin")
		req.Header.Add("sec-fetch-mode", "cors")
		req.Header.Add("sec-fetch-dest", "empty")
		req.Header.Add("referer", "https://www.homedepot.com/p/Milwaukee-M18-FUEL-120-MPH-450-CFM-18-Volt-Lithium-Ion-Brushless-Cordless-Handheld-Blower-Tool-Only-2724-20/302752040")
		req.Header.Add("accept-language", "en-GB,en-US;q=0.9,en;q=0.8")
		req.Header.Add("cookie", Cookie)
		res, err := client.Do(req)
		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(string(body))

		var homeDepotResponse HomeDepotResponse
		err = json.Unmarshal(body, &homeDepotResponse)
		if err != nil {
			fmt.Println("whoops:", err)
		}
		var row []interface{}
		row = append(row, storeID, homeDepotResponse.PrimaryItemData.StoreSkus[0].ItemID, homeDepotResponse.Inventory.Store.Quantity, homeDepotResponse.Inventory.Online.Quantity, homeDepotResponse.PrimaryItemData.StoreSkus[0].Pricing.SpecialPrice, homeDepotResponse.PrimaryItemData.Info.Upc, homeDepotResponse.PrimaryItemData.Info.ProductLabel, homeDepotResponse.PrimaryItemData.Info.BrandName, homeDepotResponse.Inventory.Store.IsLimitedQuantity, homeDepotResponse.PrimaryItemData.StoreSkus[0].Pricing.OriginalPrice, homeDepotResponse.PrimaryItemData.StoreSkus[0].Pricing.DollarOff, homeDepotResponse.PrimaryItemData.AvailabilityType, homeDepotResponse.PrimaryItemData.Shipping.BossEstimatedShippingEndDate, homeDepotResponse.PrimaryItemData.Shipping.SthEstimatedShippingStartDate, homeDepotResponse.PrimaryItemData.Shipping.SthEstimatedShippingEndDate, homeDepotResponse.PrimaryItemData.Shipping.FreeShippingThreshold, homeDepotResponse.PrimaryItemData.Shipping.ExcludedShipStates, homeDepotResponse.PrimaryItemData.Shipping.FreeShippingMessage, homeDepotResponse.PrimaryItemData.Shipping.BossEstimatedShippingStartDate, homeDepotResponse.PrimaryItemData.WebURL, homeDepotResponse.PrimaryItemData.RatingsReviews.TotalReviews, homeDepotResponse.PrimaryItemData.RatingsReviews.AverageRating, homeDepotResponse.PrimaryItemData.Info.Description, homeDepotResponse.PrimaryItemData.Info.BuyOnlineShipToStoreEligible, homeDepotResponse.PrimaryItemData.Info.IsTopSeller, homeDepotResponse.PrimaryItemData.Info.BuyOnlinePickupInStoreEligible, homeDepotResponse.PrimaryItemData.Info.ModelNumber, homeDepotResponse.PrimaryItemData.Info.VendorNumber, homeDepotResponse.PrimaryItemData.AttributeGroups[0].Entries[0].Value, homeDepotResponse.PrimaryItemData.Dimensions[0].DimensionName, homeDepotResponse.PrimaryItemData.Dimensions[1].DimensionValue.Name, homeDepotResponse.PrimaryItemData.MsbPromotions.PromotionEntry[0].DiscountEndDate, homeDepotResponse.PrimaryItemData.MsbPromotions.PromotionEntry[0].PromoLongDescription, homeDepotResponse.PrimaryItemData.MsbPromotions.PromotionEntry[0].DiscountStartDate)
		finalValues = append(finalValues, row)
		// time.Sleep(2 * time.Second)
	}
	return finalValues
}

func GetHomeDepotMultipleProductData(requestBody [][]string) [][]interface{} {
	var finalValues [][]interface{}
	currenttime := time.Now()
	epoch := currenttime.Unix()
	for i := range requestBody {
		storeID := requestBody[0][1]
		Cookie := "HD_DC=origin; bm_sz=D933413D6B61E866E16EEEBBBD5D555F~YAAQyN44fXtwtdhzAQAAztf82AjnLwoltXV6mpwDLtu0mxeGsnZULbuXY7+rrBAN8SGLGzmJagTwwbMomHLd9LH5bT9AYEr/2KMQegiKivy+GIZ8BVN6aw5eNEV8x+xeNmK44C2Pq3F+nwGLmrPtcIKjfsqnApYGrqKlqdjER6o/wy3XGHw8YuPZwOBpWXDwxRt8; THD_NR=1; THD_SESSION=; THD_CACHE_NAV_SESSION=; THD_CACHE_NAV_PERSIST=; check=true; AMCVS_F6421253512D2C100A490D45%40AdobeOrg=1; WORKFLOW=GEO_LOCATION; THD_FORCE_LOC=1; THD_INTERNAL=0; DELIVERY_ZIP=96913; THD_PERSIST=C4%3D" + storeID + "%2BGuam%20-%20Tamuning%20-%20Tamuning%2C%20GU%2B%3A%3BC4_EXP%3D1628609317%3A%3BC24%3D96913%3A%3BC24_EXP%3D1628609317%3A%3BC39%3D1%3B7%3A00-20%3A00%3B2%3B6%3A00-21%3A00%3B3%3B6%3A00-21%3A00%3B4%3B6%3A00-21%3A00%3B5%3B6%3A00-21%3A00%3B6%3B6%3A00-21%3A00%3B7%3B6%3A00-21%3A00%3A%3BC39_EXP%3D1597076917; THD_LOCALIZER=%7B%22WORKFLOW%22%3A%22GEO_LOCATION%22%2C%22THD_FORCE_LOC%22%3A%221%22%2C%22THD_INTERNAL%22%3A%220%22%2C%22THD_STRFINDERZIP%22%3A%2296913%22%2C%22THD_LOCSTORE%22%3A%221710%2BGuam%20-%20Tamuning%20-%20Tamuning%2C%20GU%2B%22%2C%22THD_STORE_HOURS%22%3A%221%3B7%3A00-20%3A00%3B2%3B6%3A00-21%3A00%3B3%3B6%3A00-21%3A00%3B4%3B6%3A00-21%3A00%3B5%3B6%3A00-21%3A00%3B6%3B6%3A00-21%3A00%3B7%3B6%3A00-21%3A00%22%2C%22THD_STORE_HOURS_EXPIRY%22%3A1597076917%7D; _pxvid=291e881c-db1e-11ea-b4bf-0242ac120005; thda.s=72cb0c5c-3dda-a400-f2f5-93f7116d1e63; thda.u=a37756ef-14a8-ed1c-3a67-3a2b717565fb; ak_bmsc=89609659092818FF7E985C3664D7194F7D38DEC86C5F0000A367315F2F2C413B~plb62Mk1slFeK9p8TVLPPQJFbQcEb3lo3NFcbSSoFz7r049tRe+FqubdAxjzlcukyN5/rL9jk0ntf4oXbdtsnJz9tE/siUbOHx8VdtLc95QalM+bE42AqFSU+dTm131uISgALHoziV3+8lJgC7UgJxvw6u4PaaKeeWQJh3pvhF8rOilozHsG79q+k+wGm9j+hNdTP4BRrali+5ppT8OwDG6TleNyvvp3dANYHO4hXtZb4SEcQHDp76AlglOb/VDl3E; _abck=CE6278AB62773D597362F7D2B3DEBC0E~0~YAAQyN44fYlwtdhzAQAABOD82AT1nVwvlPbMn09bgoV4EfbRuZeX1y4NyhKoYy9vz0keGx2Gp+FIo4Bn21ylgAocFRDFwjGma51lpTWbWsNQ8CkW+F+4MdJsFZem+DoO+aS74NkwXcvtZtuLbIerSpTKx6oI+9/8i4hnonCvf/j40RRQwf/8ORVqGzx5byjNc5GoyTqMfdwvDIPg8sgfKvUkv9D7BHRUKF3botyZBRTf7zihiK7CCs15v3lLlgsw+31u0vIP52F4NUb8PjKG2OfDvNb3rNA0ufI0qbdNJdKS/Qdaf2Zz8EEiN5+RSc5J9jEinPavhUJs+g==~-1~-1~-1; _px_f394gi7Fvmc43dfg_user_id=Mjk5NGEyNTEtZGIxZS0xMWVhLTk0ODktNmZlNzhlODI0ZTE1; RES_TRACKINGID=36632330381337079; ResonanceSegment=; RES_SESSIONID=72154150381337079; ajs_user_id=null; ajs_group_id=null; ajs_anonymous_id=%22d9ab20d3-708e-4e9d-884f-2b980df9a686%22; _meta_bing_beaconFired=true; _meta_facebookPixel_beaconFired=true; _gcl_au=1.1.1437797179.1597073320; ftr_ncd=6; QSI_HistorySession=https%3A%2F%2Fwww.homedepot.com%2Fp%2FMilwaukee-M18-FUEL-120-MPH-450-CFM-18-Volt-Lithium-Ion-Brushless-Cordless-Handheld-Blower-Tool-Only-2724-20%2F302752040~1597073320612; AMCV_F6421253512D2C100A490D45%40AdobeOrg=1585540135%7CMCIDTS%7C18485%7CMCMID%7C35670817154508997331213643943760871573%7CMCAAMLH-1597678120%7C12%7CMCAAMB-1597678120%7CRKhpRz8krg2tLO6pguXWp5olkAcUniQYPHaMWWgdJ3xzPWQmdj0y%7CMCOPTOUT-1597080520s%7CNONE%7CMCCIDH%7C172712820%7CvVersion%7C4.4.0; _meta_pinterest_epik=dj0yJnU9d05XaVI3TWlmLVQ0QnFTeHVkRkNnS2toNk1TbHkyQTcmbj1JTlRtQUpFWkZCbGMxTWR3R3BxNVBBJm09MSZ0PUFBQUFBRjh4WjZjJnJtPTEmcnQ9QUFBQUFGOHhaNmM; _meta_revjet_revjet_vid=4961946459097064000; aam_mobile=seg%3D1131078; aam_uuid=35959978601500025281202714362026421984; _meta_criteo_userId=Rj-FFdPI4xtnbz8-uKO266Q5vKUgmWks; _meta_mediaMath_mm_id=d0eb5d22-f864-4d00-b015-f81e4c80ae45; _meta_mediaMath_cid=d0eb5d22-f864-4d00-b015-f81e4c80ae45; _meta_amnet_uid=5573507722726040061; _ga=GA1.2.633436542.1597073321; _gid=GA1.2.568460109.1597073321; QuantumMetricUserID=2fcb28195be36292240261fa3a3fb47d; QuantumMetricSessionID=c22a75f61f04d3f1219648584f95b948; LPVID=QyZTAyZjMxNGI3ZWY3MDBk; LPSID-31564604=u8lPe1sOQfmF6g-G1Gmf9A; _meta_adobe_aam_uuid=35959978601500025281202714362026421984; mbox=session#6803e51076de49bb974712cb45d1cf94#1597075276|PC#6803e51076de49bb974712cb45d1cf94.31_0#1660318117; _px=H5TtyrokynGm0xZ8VnbRvIRtYIJ6MvBSOuTOQR9xIhodFc82766Yn2o7SFbGf1GfZoq1xP0wktwIg0O+G/BNQA==:1000:kAHErd/1qWUpNn5bQLbtw9r1mpBD95/lhqx/z25rAwnTvlmUaRq1sUf4LYiF7fyuGlsPfGC3Zmln6DcDCSrkcER7DTj9448xXSYNl6lsbfkjYY179JPwlpRccdSKClTzmWOa+Fvcmp476p0z09DyP4m0VdyGvxJgNjQ76bWwfDOZ5h36aweu3OMrJhb1vlrqbvbx80Wg3F1nHQAKJW/joVoiXuG9bd0+SK3At3IwfNazpdFINnBsjFI1BhHdEV5ey67DFM5q2o6VpDx/Ed7QYA==; s_pers=%20productnum%3D1%7C1599665318868%3B%20s_nr365%3D1597073417914-New%7C1628609417914%3B%20s_dslv%3D1597073417921%7C1691681417921%3B; s_sess=%20stsh%3D%3B%20s_pv_pName%3Dproductdetails%253E302752040%3B%20s_pv_pType%3Dpip%3B%20s_pv_cmpgn%3D%3B%20s_cc%3Dtrue%3B; _meta_mediaMath_iframe_counter=3; forterToken=640fce6dab2d475c91c80f273376b9ed_1597073417951__UDF43_9ck; _px_4946459675_cs=eyJpZCI6IjI5OTQ1NDMwLWRiMWUtMTFlYS05NDg5LTZmZTc4ZTgyNGUxNSIsInN0b3JhZ2UiOnt9LCJleHBpcmF0aW9uIjoxNTk3MDc1NjkxODE4fQ==; _pxde=6e624cc55fc1a20bd3fb6d477a0232a81102f7d4d84b291714a5942fca59d955:eyJ0aW1lc3RhbXAiOjE1OTcwNzM4OTMzMTF9; akaau=1597074193~id=7585bff3cb571847323335006e1ee492; bm_sv=67B4BD82E3C07565D9C966B534F53950~SRQIMfQu0qB+XuC4kxuDGg1mtgowuHmZFA1ijc4z2RT6fwwaxeXFuKICPhy63+0318AJAAcxo/7PsPaoQ7Sg/jGh8l+icBOajdURDqPQDKAczQY3bV5dG02O+ojfzojJ17OYs876L6XWWPOAxC0c8ypv1+5xp60X7iRJfuPVfrk=; IN_STORE_API_SESSION=TRUE; akaau=1597074756~id=570e6607324d97f43aff334a687ef720; bm_sv=67B4BD82E3C07565D9C966B534F53950~SRQIMfQu0qB+XuC4kxuDGg1mtgowuHmZFA1ijc4z2RT6fwwaxeXFuKICPhy63+0318AJAAcxo/7PsPaoQ7Sg/jGh8l+icBOajdURDqPQDKB01Zf137gHE7JiOZKKo/sLZ5v33xatSoFD0VeF/1OvvX3lI2J8uwwHSHHmyTujUYw=; _abck=CE6278AB62773D597362F7D2B3DEBC0E~-1~YAAQyN44fZ+JtdhzAQAA0T0O2QSpbesPatxFjVR5gmDxtyb1Pmp5BnIeh/+ggFLpWb/whyotqoyArNJSunSQQayxXQVeXIjSi89P4zO8QcMSXRLvADNin46chyAG0jGiQNpG/YvUuLK+DAopT5NO9NvQ2mj7WOqI/FoAwoEjo6Ap2e0KOjllQyhShzo6rImJEnhvaG7xtHxte0yul/OTIahfnwfXGm3DNkacQ5xS4H97HNpnv54PrWSgi+Jmybq5bPrKmKuffIhjS3DDCzbVZ4DZ1uaV52d61LvNuckShPu0+qRjwM0L8dr39N35ZpUEPatLOLfADrK3dg==~0~-1~-1; akaau=1597160858~id=5d84ab426f4dc213c48a00d290401967; _abck=CE6278AB62773D597362F7D2B3DEBC0E~-1~YAAQT9cLF5Ss/45zAQAA3Q8w3gSR0CK+O35uM2K5kZcckE1rq9bg7JHcnFRMZIhKwxuIuTMIx+N6S3MMwU7pbVEQHRxk2+hVw9liPqz/rB2OBPicKXxRwT/qK78fQZqPQQklPInnLwqsPLkuOEGbnfXy/RfKQrypOZWkTyBt1iDm33beCtjInYzQs+skq+2egELaX1D2H2Mp+qC10LSp2mjb2tfxPHKu4Fj4i+Orj7tAEfQ0I19VVbHEaqJW2RRU+C9EbpZJadLtiZm4Sxb00lYt+evrnjtfF0F0npGMz25eEjT16J3320BJKiWAN0fCcoYTarCmLVPmvg==~0~-1~-1"
		url := "https://www.homedepot.com/p/svcs/frontEndModel/" + requestBody[i][0] + "?_=" + strconv.Itoa(int(epoch)) + "000"
		method := "GET"

		client := &http.Client{}
		req, err := http.NewRequest(method, url, nil)

		if err != nil {
			fmt.Println(err)
		}
		req.Header.Add("authority", "www.homedepot.com")
		req.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.89 Safari/537.36")
		req.Header.Add("accept", "*/*")
		req.Header.Add("sec-fetch-site", "same-origin")
		req.Header.Add("sec-fetch-mode", "cors")
		req.Header.Add("sec-fetch-dest", "empty")
		req.Header.Add("referer", "https://www.homedepot.com/p/Milwaukee-M18-FUEL-120-MPH-450-CFM-18-Volt-Lithium-Ion-Brushless-Cordless-Handheld-Blower-Tool-Only-2724-20/302752040")
		req.Header.Add("accept-language", "en-GB,en-US;q=0.9,en;q=0.8")
		req.Header.Add("cookie", Cookie)
		res, err := client.Do(req)
		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(string(body))

		var homeDepotResponse HomeDepotResponse
		err = json.Unmarshal(body, &homeDepotResponse)
		if err != nil {
			fmt.Println("whoops:", err)
		}
		var row []interface{}
		row = append(row, storeID, homeDepotResponse.PrimaryItemData.StoreSkus[0].ItemID, homeDepotResponse.Inventory.Store.Quantity, homeDepotResponse.Inventory.Online.Quantity, homeDepotResponse.PrimaryItemData.StoreSkus[0].Pricing.SpecialPrice, homeDepotResponse.PrimaryItemData.Info.Upc, homeDepotResponse.PrimaryItemData.Info.ProductLabel, homeDepotResponse.PrimaryItemData.Info.BrandName, homeDepotResponse.Inventory.Store.IsLimitedQuantity, homeDepotResponse.PrimaryItemData.StoreSkus[0].Pricing.OriginalPrice, homeDepotResponse.PrimaryItemData.StoreSkus[0].Pricing.DollarOff, homeDepotResponse.PrimaryItemData.AvailabilityType, homeDepotResponse.PrimaryItemData.Shipping.BossEstimatedShippingEndDate, homeDepotResponse.PrimaryItemData.Shipping.SthEstimatedShippingStartDate, homeDepotResponse.PrimaryItemData.Shipping.SthEstimatedShippingEndDate, homeDepotResponse.PrimaryItemData.Shipping.FreeShippingThreshold, homeDepotResponse.PrimaryItemData.Shipping.ExcludedShipStates, homeDepotResponse.PrimaryItemData.Shipping.FreeShippingMessage, homeDepotResponse.PrimaryItemData.Shipping.BossEstimatedShippingStartDate, homeDepotResponse.PrimaryItemData.WebURL, homeDepotResponse.PrimaryItemData.RatingsReviews.TotalReviews, homeDepotResponse.PrimaryItemData.RatingsReviews.AverageRating, homeDepotResponse.PrimaryItemData.Info.Description, homeDepotResponse.PrimaryItemData.Info.BuyOnlineShipToStoreEligible, homeDepotResponse.PrimaryItemData.Info.IsTopSeller, homeDepotResponse.PrimaryItemData.Info.BuyOnlinePickupInStoreEligible, homeDepotResponse.PrimaryItemData.Info.ModelNumber, homeDepotResponse.PrimaryItemData.Info.VendorNumber, homeDepotResponse.PrimaryItemData.AttributeGroups[0].Entries[0].Value, homeDepotResponse.PrimaryItemData.Dimensions[0].DimensionName, homeDepotResponse.PrimaryItemData.Dimensions[1].DimensionValue.Name, homeDepotResponse.PrimaryItemData.MsbPromotions.PromotionEntry[0].DiscountEndDate, homeDepotResponse.PrimaryItemData.MsbPromotions.PromotionEntry[0].PromoLongDescription, homeDepotResponse.PrimaryItemData.MsbPromotions.PromotionEntry[0].DiscountStartDate)
		finalValues = append(finalValues, row)
		// time.Sleep(2 * time.Second)
	}
	return finalValues
}
