package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

const (
	publicAPIBaseURL = "https://www.vinted.pl"
	apiVersion       = "v2"
	itemsEndpoint    = "catalog/items"
	ItemsApiUrl      = publicAPIBaseURL + "catalog/items?"
)

var (
	baseAPIURL = fmt.Sprintf("%s/api/%s/", publicAPIBaseURL, apiVersion)
)

// VintedClient manages communication with the Vinted API.
type VintedClient struct {
	httpClient *http.Client
	cookies    []*http.Cookie
}

// NewVintedClient initializes a new Vinted API client.
func NewVintedClient() (*VintedClient, error) {
	client := &VintedClient{
		httpClient: &http.Client{},
	}
	if err := client.fetchCookies(); err != nil {
		return nil, fmt.Errorf("failed to fetch cookies: %w", err)
	}
	return client, nil
}

func (c *VintedClient) fetchCookies() error {
	req, err := http.NewRequest("GET", publicAPIBaseURL, nil)
	if err != nil {
		return fmt.Errorf("creating request for cookies: %w", err)
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("fetching cookies: %w", err)
	}
	defer resp.Body.Close()

	c.cookies = resp.Cookies()
	return nil
}

func (c *VintedClient) makeRequestWithCookies(method, url string, body url.Values) (*http.Response, error) {
	req, err := http.NewRequest(method, url, strings.NewReader(body.Encode()))
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}

	for _, cookie := range c.cookies {
		req.AddCookie(cookie)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("performing request: %w", err)
	}
	return resp, nil
}

// FindItems searches for items on Vinted based on the provided query parameters.
func (c *VintedClient) FindItems(params IQueryParams) (*ItemsResponse, error) {
	url := c.buildItemsQuery(params)
	resp, err := c.makeRequestWithCookies("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("fetching items: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("fetching items failed with status: %s", resp.Status)
	}

	var items ItemsResponse
	if err := json.NewDecoder(resp.Body).Decode(&items); err != nil {
		return nil, fmt.Errorf("decoding response: %w", err)
	}
	return &items, nil
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
