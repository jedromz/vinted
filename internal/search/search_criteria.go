package search

type QueryParams struct {
	SizeIDs     []int
	Catalog     []int
	MaterialIDs []int
	ColorIDs    []int
	BrandIDs    []int
	PriceFrom   int
	PriceTo     int
	StatusIDs   []int
	Order       string
	Currency    string
	SearchText  string
	Page        int
	PerPage     int
}

func (qp *QueryParams) GetSizeIDs() []int {
	return qp.SizeIDs
}

func (qp *QueryParams) GetCatalog() []int {
	return qp.Catalog
}

func (qp *QueryParams) GetMaterialIDs() []int {
	return qp.MaterialIDs
}

func (qp *QueryParams) GetColorIDs() []int {
	return qp.ColorIDs
}

func (qp *QueryParams) GetBrandIDs() []int {
	return qp.BrandIDs
}

func (qp *QueryParams) GetPriceFrom() int {
	return qp.PriceFrom
}

func (qp *QueryParams) GetPriceTo() int {
	return qp.PriceTo
}

func (qp *QueryParams) GetStatusIDs() []int {
	return qp.StatusIDs
}

func (qp *QueryParams) GetOrder() string {
	return qp.Order
}

func (qp *QueryParams) GetCurrency() string {
	return qp.Currency
}

func (qp *QueryParams) GetSearchText() string {
	return qp.SearchText
}

func (qp *QueryParams) GetPage() int {
	return qp.Page
}

func (qp *QueryParams) GetPerPage() int {
	return qp.PerPage
}
