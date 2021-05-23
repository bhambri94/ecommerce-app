package zoro

type ZoroReview struct {
	Name   string `json:"name"`
	Paging struct {
		TotalResults      int `json:"total_results"`
		PagesTotal        int `json:"pages_total"`
		PageSize          int `json:"page_size"`
		CurrentPageNumber int `json:"current_page_number"`
	} `json:"paging"`
	Results []struct {
		PageID  string        `json:"page_id"`
		Reviews []interface{} `json:"reviews"`
		Rollup  struct {
			AverageRating float64 `json:"average_rating"`
		} `json:"rollup"`
	} `json:"results"`
}
