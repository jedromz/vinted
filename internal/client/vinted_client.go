package client

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

const (
	PublicApiUrl = "https://www.vinted.pl"
	BaseApiUrl   = "https://www.vinted.pl/api/v2/"
	ItemsApiUrl  = BaseApiUrl + "catalog/items?"
)

type VintedClient struct {
	Client  *http.Client
	Cookies []*http.Cookie
}

func New() (*VintedClient, error) {
	v := &VintedClient{
		Client: &http.Client{},
	}
	err := v.fetchCookies()
	if err != nil {
		return nil, errors.New("failed to fetch cookies")
	}
	return v, nil
}

func (c *VintedClient) fetchCookies() error {
	req, err := http.NewRequest("GET", PublicApiUrl, nil)
	if err != nil {
		return errors.New("failed to create request while fetching cookies")
	}
	resp, err := c.Client.Do(req)
	if err != nil {
		return errors.New("failed to get response while fetching cookies")
	}
	defer resp.Body.Close()
	c.Cookies = resp.Cookies()
	return nil
}

func (c *VintedClient) String() string {
	s := ""
	for _, cookie := range c.Cookies {
		s += cookie.Name + "\n"
	}
	return s
}

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

func (c *VintedClient) FindItems(params IQueryParams) (ItemsResponse, error) {
	log.Println(c.buildItemsQuery(params))
	req, err := http.NewRequest("GET", c.buildItemsQuery(params), nil)
	if err != nil {
		return ItemsResponse{}, errors.New("failed to create request while fetching items")
	}

	for _, cookie := range c.Cookies {
		req.AddCookie(cookie)
	}

	resp, err := c.Client.Do(req)

	if err != nil {
		return ItemsResponse{}, errors.New("failed to get response while fetching items")
	}
	defer resp.Body.Close()
	var items ItemsResponse
	if resp.StatusCode != http.StatusOK {
		return ItemsResponse{}, errors.New("failed to get response while fetching items " + resp.Status)
	}
	if resp.Body == nil {
		return ItemsResponse{}, nil
	}

	err = json.NewDecoder(resp.Body).Decode(&items)
	if err != nil {
		return ItemsResponse{}, errors.New("failed to decode response while fetching items")
	}
	return items, nil
}
func (c *VintedClient) buildItemsQuery(params IQueryParams) string {
	values := url.Values{}

	// Helper function to add slice parameters
	addSliceParams := func(key string, ids []int) {
		for _, id := range ids {
			values.Add(key, strconv.Itoa(id))
		}
	}

	addSliceParams("size_ids[]", params.GetSizeIDs())
	addSliceParams("catalog[]", params.GetCatalog())
	addSliceParams("material_ids[]", params.GetMaterialIDs())
	addSliceParams("color_ids[]", params.GetColorIDs())
	addSliceParams("brand_ids[]", params.GetBrandIDs())
	addSliceParams("status_ids[]", params.GetStatusIDs())

	if params.GetPage() > 0 {
		values.Set("page", strconv.Itoa(params.GetPage()))
	} else {
		values.Set("page", "1")
	}
	if params.GetPerPage() > 0 {
		values.Set("per_page", strconv.Itoa(params.GetPerPage()))
	} else {
		values.Set("per_page", "10")
	}
	if params.GetPriceFrom() > 0 {
		values.Set("price_from", strconv.Itoa(params.GetPriceFrom()))
	}
	if params.GetPriceTo() > 0 {
		values.Set("price_to", strconv.Itoa(params.GetPriceTo()))
	}
	if params.GetOrder() != "" {
		values.Set("order", params.GetOrder())
	} else {
		values.Set("order", "newest_first")
	}
	if params.GetCurrency() != "" {
		values.Set("currency", params.GetCurrency())
	}
	if params.GetSearchText() != "" {
		values.Set("search_text", params.GetSearchText())
	}

	queryString := strings.TrimLeft(values.Encode(), "&")
	return ItemsApiUrl + "&" + queryString
}
