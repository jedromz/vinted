package client

type ItemsResponse struct {
	Items                []Item         `json:"items"`
	DominantBrand        interface{}    `json:"dominant_brand"`
	SearchTrackingParams TrackingParams `json:"search_tracking_params"`
	Pagination           Pagination     `json:"pagination"`
	Code                 int            `json:"code"`
}
type Item struct {
	Id                    int64          `json:"id"`
	Title                 string         `json:"title"`
	Price                 string         `json:"price"`
	IsVisible             int            `json:"is_visible"`
	Discount              interface{}    `json:"discount"`
	Currency              string         `json:"currency"`
	BrandTitle            string         `json:"brand_title"`
	IsForSwap             bool           `json:"is_for_swap"`
	User                  User           `json:"user"`
	Url                   string         `json:"url"`
	Promoted              bool           `json:"promoted"`
	Photo                 Photo          `json:"photo"`
	FavouriteCount        int            `json:"favourite_count"`
	IsFavourite           bool           `json:"is_favourite"`
	Badge                 interface{}    `json:"badge"`
	Conversion            interface{}    `json:"conversion"`
	ServiceFee            string         `json:"service_fee"`
	TotalItemPrice        string         `json:"total_item_price"`
	TotalItemPriceRounded interface{}    `json:"total_item_price_rounded"`
	ViewCount             int            `json:"view_count"`
	SizeTitle             string         `json:"size_title"`
	ContentSource         string         `json:"content_source"`
	Status                string         `json:"status"`
	IconBadges            []interface{}  `json:"icon_badges"`
	SearchTrackingParams  TrackingParams `json:"search_tracking_params"`
}
type Pagination struct {
	CurrentPage  int `json:"current_page"`
	TotalPages   int `json:"total_pages"`
	TotalEntries int `json:"total_entries"`
	PerPage      int `json:"per_page"`
	Time         int `json:"time"`
}

type TrackingParams struct {
	Score          float64     `json:"score"`
	MatchedQueries interface{} `json:"matched_queries"`
}
type User struct {
	Id         int    `json:"id"`
	Login      string `json:"login"`
	Business   bool   `json:"business"`
	ProfileUrl string `json:"profile_url"`
	Photo      Photo  `json:"photo"`
}
type HighRes struct {
	Id          string      `json:"id"`
	Timestamp   int         `json:"timestamp"`
	Orientation interface{} `json:"orientation"`
}
type Thumbnail struct {
	Type         string      `json:"type"`
	Url          string      `json:"url"`
	Width        int         `json:"width"`
	Height       int         `json:"height"`
	OriginalSize interface{} `json:"original_size"`
}

type Photo struct {
	Id                  int64       `json:"id"`
	Width               int         `json:"width"`
	Height              int         `json:"height"`
	TempUuid            interface{} `json:"temp_uuid"`
	Url                 string      `json:"url"`
	DominantColor       string      `json:"dominant_color"`
	DominantColorOpaque string      `json:"dominant_color_opaque"`
	Thumbnails          []Thumbnail `json:"thumbnails"`
	IsSuspicious        bool        `json:"is_suspicious"`
	Orientation         interface{} `json:"orientation"`
	HighResolution      HighRes     `json:"high_resolution"`
	FullSizeUrl         string      `json:"full_size_url"`
	IsHidden            bool        `json:"is_hidden"`
	Extra               struct{}    `json:"extra"`
}
