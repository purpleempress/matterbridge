// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// AdministrativeUnitRequestBuilder is request builder for AdministrativeUnit
type AdministrativeUnitRequestBuilder struct{ BaseRequestBuilder }

// Request returns AdministrativeUnitRequest
func (b *AdministrativeUnitRequestBuilder) Request() *AdministrativeUnitRequest {
	return &AdministrativeUnitRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// AdministrativeUnitRequest is request for AdministrativeUnit
type AdministrativeUnitRequest struct{ BaseRequest }

// Get performs GET request for AdministrativeUnit
func (r *AdministrativeUnitRequest) Get(ctx context.Context) (resObj *AdministrativeUnit, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for AdministrativeUnit
func (r *AdministrativeUnitRequest) Update(ctx context.Context, reqObj *AdministrativeUnit) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for AdministrativeUnit
func (r *AdministrativeUnitRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// Extensions returns request builder for Extension collection
func (b *AdministrativeUnitRequestBuilder) Extensions() *AdministrativeUnitExtensionsCollectionRequestBuilder {
	bb := &AdministrativeUnitExtensionsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/extensions"
	return bb
}

// AdministrativeUnitExtensionsCollectionRequestBuilder is request builder for Extension collection
type AdministrativeUnitExtensionsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for Extension collection
func (b *AdministrativeUnitExtensionsCollectionRequestBuilder) Request() *AdministrativeUnitExtensionsCollectionRequest {
	return &AdministrativeUnitExtensionsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for Extension item
func (b *AdministrativeUnitExtensionsCollectionRequestBuilder) ID(id string) *ExtensionRequestBuilder {
	bb := &ExtensionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// AdministrativeUnitExtensionsCollectionRequest is request for Extension collection
type AdministrativeUnitExtensionsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for Extension collection
func (r *AdministrativeUnitExtensionsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]Extension, error) {
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
	var values []Extension
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
			value  []Extension
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

// Get performs GET request for Extension collection
func (r *AdministrativeUnitExtensionsCollectionRequest) Get(ctx context.Context) ([]Extension, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for Extension collection
func (r *AdministrativeUnitExtensionsCollectionRequest) Add(ctx context.Context, reqObj *Extension) (resObj *Extension, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Members returns request builder for DirectoryObject collection
func (b *AdministrativeUnitRequestBuilder) Members() *AdministrativeUnitMembersCollectionRequestBuilder {
	bb := &AdministrativeUnitMembersCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/members"
	return bb
}

// AdministrativeUnitMembersCollectionRequestBuilder is request builder for DirectoryObject collection
type AdministrativeUnitMembersCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for DirectoryObject collection
func (b *AdministrativeUnitMembersCollectionRequestBuilder) Request() *AdministrativeUnitMembersCollectionRequest {
	return &AdministrativeUnitMembersCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for DirectoryObject item
func (b *AdministrativeUnitMembersCollectionRequestBuilder) ID(id string) *DirectoryObjectRequestBuilder {
	bb := &DirectoryObjectRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// AdministrativeUnitMembersCollectionRequest is request for DirectoryObject collection
type AdministrativeUnitMembersCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for DirectoryObject collection
func (r *AdministrativeUnitMembersCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]DirectoryObject, error) {
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
	var values []DirectoryObject
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
			value  []DirectoryObject
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

// Get performs GET request for DirectoryObject collection
func (r *AdministrativeUnitMembersCollectionRequest) Get(ctx context.Context) ([]DirectoryObject, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for DirectoryObject collection
func (r *AdministrativeUnitMembersCollectionRequest) Add(ctx context.Context, reqObj *DirectoryObject) (resObj *DirectoryObject, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// ScopedRoleMembers returns request builder for ScopedRoleMembership collection
func (b *AdministrativeUnitRequestBuilder) ScopedRoleMembers() *AdministrativeUnitScopedRoleMembersCollectionRequestBuilder {
	bb := &AdministrativeUnitScopedRoleMembersCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/scopedRoleMembers"
	return bb
}

// AdministrativeUnitScopedRoleMembersCollectionRequestBuilder is request builder for ScopedRoleMembership collection
type AdministrativeUnitScopedRoleMembersCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for ScopedRoleMembership collection
func (b *AdministrativeUnitScopedRoleMembersCollectionRequestBuilder) Request() *AdministrativeUnitScopedRoleMembersCollectionRequest {
	return &AdministrativeUnitScopedRoleMembersCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for ScopedRoleMembership item
func (b *AdministrativeUnitScopedRoleMembersCollectionRequestBuilder) ID(id string) *ScopedRoleMembershipRequestBuilder {
	bb := &ScopedRoleMembershipRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// AdministrativeUnitScopedRoleMembersCollectionRequest is request for ScopedRoleMembership collection
type AdministrativeUnitScopedRoleMembersCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for ScopedRoleMembership collection
func (r *AdministrativeUnitScopedRoleMembersCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]ScopedRoleMembership, error) {
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
	var values []ScopedRoleMembership
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
			value  []ScopedRoleMembership
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

// Get performs GET request for ScopedRoleMembership collection
func (r *AdministrativeUnitScopedRoleMembersCollectionRequest) Get(ctx context.Context) ([]ScopedRoleMembership, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for ScopedRoleMembership collection
func (r *AdministrativeUnitScopedRoleMembersCollectionRequest) Add(ctx context.Context, reqObj *ScopedRoleMembership) (resObj *ScopedRoleMembership, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}