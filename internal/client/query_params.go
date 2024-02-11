package client

type IQueryParams interface {
	GetSizeIDs() []int
	GetCatalog() []int
	GetMaterialIDs() []int
	GetColorIDs() []int
	GetBrandIDs() []int
	GetPriceFrom() int
	GetPriceTo() int
	GetStatusIDs() []int
	GetOrder() string
	GetCurrency() string
	GetSearchText() string
	GetPage() int
	GetPerPage() int
}
