package catalog

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/dzt-corp/selling-partner-api-sdk/pkg/runtime"
)

type V2022GetCatalogItemResp struct {
	Body         []byte
	HTTPResponse *http.Response
	Model        *V2022GetCatalogItemResponse
}

// GetCatalogItemWithResponse request returning *GetCatalogItemResponse
func (c *ClientWithResponses) V2022GetCatalogItemWithResponse(ctx context.Context, asin string, params *V2022GetCatalogItemParams) (*V2022GetCatalogItemResp, error) {
	rsp, err := c.V2022GetCatalogItem(ctx, asin, params)
	if err != nil {
		return nil, err
	}
	return ParserV2022GetCatalogItemWithResponse(rsp)
}

func (c *Client) V2022GetCatalogItem(ctx context.Context, asin string, params *V2022GetCatalogItemParams) (*http.Response, error) {
	req, err := V2022NewGetCatalogItemRequest(c.Endpoint, asin, params)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	req.Header.Set("User-Agent", c.UserAgent)
	if c.RequestBefore != nil {
		err = c.RequestBefore(ctx, req)
		if err != nil {
			return nil, err
		}
	}

	rsp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	if c.ResponseAfter != nil {
		err = c.ResponseAfter(ctx, rsp)
		if err != nil {
			return nil, err
		}
	}
	return rsp, nil
}

// NewGetCatalogItemRequest generates requests for GetCatalogItem
func V2022NewGetCatalogItemRequest(endpoint string, asin string, params *V2022GetCatalogItemParams) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParam("simple", false, "asin", asin)
	if err != nil {
		return nil, err
	}

	queryUrl, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	basePath := fmt.Sprintf("/catalog/2022-04-01/items/%s", pathParam0)
	if basePath[0] == '/' {
		basePath = basePath[1:]
	}

	queryUrl, err = queryUrl.Parse(basePath)
	if err != nil {
		return nil, err
	}

	queryValues := queryUrl.Query()
	if queryFrag, err := runtime.StyleParam("form", true, "marketplaceIds", strings.Join(params.MarketplaceIds, ",")); err != nil {
		return nil, err
	} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
		return nil, err
	} else {
		for k, v := range parsed {
			for _, v2 := range v {
				queryValues.Add(k, v2)
			}
		}
	}

	if (len(params.IncludedData)) > 0 {
		queryValues.Add("includedData", strings.Join(params.IncludedData, ","))
	}

	queryUrl.RawQuery = queryValues.Encode()
	req, err := http.NewRequest("GET", queryUrl.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// ParseGetCatalogItemResp parses an HTTP response from a GetCatalogItemWithResponse call
func ParserV2022GetCatalogItemWithResponse(rsp *http.Response) (*V2022GetCatalogItemResp, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	fmt.Println("ParserV2022GetCatalogItemWithResponse ", string(bodyBytes))
	response := &V2022GetCatalogItemResp{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	var dest V2022GetCatalogItemResponse
	if err := json.Unmarshal(bodyBytes, &dest); err != nil {
		return nil, err
	}

	response.Model = &dest

	if rsp.StatusCode >= 300 {
		err = fmt.Errorf(rsp.Status)
	}

	return response, err
}
