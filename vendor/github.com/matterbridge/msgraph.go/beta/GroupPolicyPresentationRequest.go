// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// GroupPolicyPresentationRequestBuilder is request builder for GroupPolicyPresentation
type GroupPolicyPresentationRequestBuilder struct{ BaseRequestBuilder }

// Request returns GroupPolicyPresentationRequest
func (b *GroupPolicyPresentationRequestBuilder) Request() *GroupPolicyPresentationRequest {
	return &GroupPolicyPresentationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// GroupPolicyPresentationRequest is request for GroupPolicyPresentation
type GroupPolicyPresentationRequest struct{ BaseRequest }

// Get performs GET request for GroupPolicyPresentation
func (r *GroupPolicyPresentationRequest) Get(ctx context.Context) (resObj *GroupPolicyPresentation, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for GroupPolicyPresentation
func (r *GroupPolicyPresentationRequest) Update(ctx context.Context, reqObj *GroupPolicyPresentation) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for GroupPolicyPresentation
func (r *GroupPolicyPresentationRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// Definition is navigation property
func (b *GroupPolicyPresentationRequestBuilder) Definition() *GroupPolicyDefinitionRequestBuilder {
	bb := &GroupPolicyDefinitionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/definition"
	return bb
}