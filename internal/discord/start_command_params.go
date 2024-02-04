package discord

type SearchStartParams struct {
	SizeIDs     []int  `json:"SizeIDs"`
	Catalog     []int  `json:"Catalog"`
	MaterialIDs []int  `json:"MaterialIDs"`
	ColorIDs    []int  `json:"ColorIDs"`
	BrandIDs    []int  `json:"BrandIDs"`
	PriceFrom   int    `json:"PriceFrom"`
	PriceTo     int    `json:"PriceTo"`
	StatusIDs   []int  `json:"StatusIDs"`
	Order       string `json:"Order"`
	Currency    string `json:"Currency"`
	SearchText  string `json:"SearchText"`
	Page        int    `json:"Page"`
	PerPage     int    `json:"PerPage"`
}

func (s SearchStartParams) GetSizeIDs() []int {
	return s.SizeIDs
}

func (s SearchStartParams) GetCatalog() []int {
	return s.Catalog
}

func (s SearchStartParams) GetMaterialIDs() []int {
	return s.Catalog
}

func (s SearchStartParams) GetColorIDs() []int {
	return s.ColorIDs
}

func (s SearchStartParams) GetBrandIDs() []int {
	return s.BrandIDs
}

func (s SearchStartParams) GetPriceFrom() int {
	return s.PriceFrom
}

func (s SearchStartParams) GetPriceTo() int {
	return s.PriceTo
}

func (s SearchStartParams) GetStatusIDs() []int {
	return s.StatusIDs
}

func (s SearchStartParams) GetOrder() string {
	return s.Order
}

func (s SearchStartParams) GetCurrency() string {
	return s.Currency
}

func (s SearchStartParams) GetSearchText() string {
	return s.SearchText
}

func (s SearchStartParams) GetPage() int {
	return s.Page
}

func (s SearchStartParams) GetPerPage() int {
	return s.PerPage
}
