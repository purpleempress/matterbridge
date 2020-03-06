// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// WorkbookRangeFormatRequestBuilder is request builder for WorkbookRangeFormat
type WorkbookRangeFormatRequestBuilder struct{ BaseRequestBuilder }

// Request returns WorkbookRangeFormatRequest
func (b *WorkbookRangeFormatRequestBuilder) Request() *WorkbookRangeFormatRequest {
	return &WorkbookRangeFormatRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// WorkbookRangeFormatRequest is request for WorkbookRangeFormat
type WorkbookRangeFormatRequest struct{ BaseRequest }

// Get performs GET request for WorkbookRangeFormat
func (r *WorkbookRangeFormatRequest) Get(ctx context.Context) (resObj *WorkbookRangeFormat, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for WorkbookRangeFormat
func (r *WorkbookRangeFormatRequest) Update(ctx context.Context, reqObj *WorkbookRangeFormat) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for WorkbookRangeFormat
func (r *WorkbookRangeFormatRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// Borders returns request builder for WorkbookRangeBorder collection
func (b *WorkbookRangeFormatRequestBuilder) Borders() *WorkbookRangeFormatBordersCollectionRequestBuilder {
	bb := &WorkbookRangeFormatBordersCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/borders"
	return bb
}

// WorkbookRangeFormatBordersCollectionRequestBuilder is request builder for WorkbookRangeBorder collection
type WorkbookRangeFormatBordersCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for WorkbookRangeBorder collection
func (b *WorkbookRangeFormatBordersCollectionRequestBuilder) Request() *WorkbookRangeFormatBordersCollectionRequest {
	return &WorkbookRangeFormatBordersCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for WorkbookRangeBorder item
func (b *WorkbookRangeFormatBordersCollectionRequestBuilder) ID(id string) *WorkbookRangeBorderRequestBuilder {
	bb := &WorkbookRangeBorderRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// WorkbookRangeFormatBordersCollectionRequest is request for WorkbookRangeBorder collection
type WorkbookRangeFormatBordersCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for WorkbookRangeBorder collection
func (r *WorkbookRangeFormatBordersCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]WorkbookRangeBorder, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []WorkbookRangeBorder
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			errRes := &ErrorResponse{Response: res}
			err := jsonx.Unmarshal(b, errRes)
			if err != nil {
				return nil, fmt.Errorf("%s: %s", res.Status, string(b))
			}
			return nil, errRes
		}
		var (
			paging Paging
			value  []WorkbookRangeBorder
		)
		err := jsonx.NewDecoder(res.Body).Decode(&paging)
		if err != nil {
			return nil, err
		}
		err = jsonx.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if len(paging.NextLink) == 0 {
			return values, nil
		}
		req, err = http.NewRequest("GET", paging.NextLink, nil)
		if ctx != nil {
			req = req.WithContext(ctx)
		}
		res, err = r.client.Do(req)
		if err != nil {
			return nil, err
		}
	}
}

// Get performs GET request for WorkbookRangeBorder collection
func (r *WorkbookRangeFormatBordersCollectionRequest) Get(ctx context.Context) ([]WorkbookRangeBorder, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for WorkbookRangeBorder collection
func (r *WorkbookRangeFormatBordersCollectionRequest) Add(ctx context.Context, reqObj *WorkbookRangeBorder) (resObj *WorkbookRangeBorder, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Fill is navigation property
func (b *WorkbookRangeFormatRequestBuilder) Fill() *WorkbookRangeFillRequestBuilder {
	bb := &WorkbookRangeFillRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/fill"
	return bb
}

// Font is navigation property
func (b *WorkbookRangeFormatRequestBuilder) Font() *WorkbookRangeFontRequestBuilder {
	bb := &WorkbookRangeFontRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/font"
	return bb
}

// Protection is navigation property
func (b *WorkbookRangeFormatRequestBuilder) Protection() *WorkbookFormatProtectionRequestBuilder {
	bb := &WorkbookFormatProtectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/protection"
	return bb
}