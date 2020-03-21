// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// MacManagedAppProtectionRequestBuilder is request builder for MacManagedAppProtection
type MacManagedAppProtectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns MacManagedAppProtectionRequest
func (b *MacManagedAppProtectionRequestBuilder) Request() *MacManagedAppProtectionRequest {
	return &MacManagedAppProtectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// MacManagedAppProtectionRequest is request for MacManagedAppProtection
type MacManagedAppProtectionRequest struct{ BaseRequest }

// Get performs GET request for MacManagedAppProtection
func (r *MacManagedAppProtectionRequest) Get(ctx context.Context) (resObj *MacManagedAppProtection, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for MacManagedAppProtection
func (r *MacManagedAppProtectionRequest) Update(ctx context.Context, reqObj *MacManagedAppProtection) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for MacManagedAppProtection
func (r *MacManagedAppProtectionRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}