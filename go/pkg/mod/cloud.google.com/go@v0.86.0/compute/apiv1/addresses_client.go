// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package compute

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	httptransport "google.golang.org/api/transport/http"
	computepb "google.golang.org/genproto/googleapis/cloud/compute/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/encoding/protojson"
)

var newAddressesClientHook clientHook

// AddressesCallOptions contains the retry settings for each method of AddressesClient.
type AddressesCallOptions struct {
	AggregatedList []gax.CallOption
	Delete         []gax.CallOption
	Get            []gax.CallOption
	Insert         []gax.CallOption
	List           []gax.CallOption
}

// internalAddressesClient is an interface that defines the methods availaible from Google Compute Engine API.
type internalAddressesClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	AggregatedList(context.Context, *computepb.AggregatedListAddressesRequest, ...gax.CallOption) (*computepb.AddressAggregatedList, error)
	Delete(context.Context, *computepb.DeleteAddressRequest, ...gax.CallOption) (*computepb.Operation, error)
	Get(context.Context, *computepb.GetAddressRequest, ...gax.CallOption) (*computepb.Address, error)
	Insert(context.Context, *computepb.InsertAddressRequest, ...gax.CallOption) (*computepb.Operation, error)
	List(context.Context, *computepb.ListAddressesRequest, ...gax.CallOption) (*computepb.AddressList, error)
}

// AddressesClient is a client for interacting with Google Compute Engine API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// The Addresses API.
type AddressesClient struct {
	// The internal transport-dependent client.
	internalClient internalAddressesClient

	// The call options for this service.
	CallOptions *AddressesCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *AddressesClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *AddressesClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *AddressesClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// AggregatedList retrieves an aggregated list of addresses.
func (c *AddressesClient) AggregatedList(ctx context.Context, req *computepb.AggregatedListAddressesRequest, opts ...gax.CallOption) (*computepb.AddressAggregatedList, error) {
	return c.internalClient.AggregatedList(ctx, req, opts...)
}

// Delete deletes the specified address resource.
func (c *AddressesClient) Delete(ctx context.Context, req *computepb.DeleteAddressRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.Delete(ctx, req, opts...)
}

// Get returns the specified address resource.
func (c *AddressesClient) Get(ctx context.Context, req *computepb.GetAddressRequest, opts ...gax.CallOption) (*computepb.Address, error) {
	return c.internalClient.Get(ctx, req, opts...)
}

// Insert creates an address resource in the specified project by using the data included in the request.
func (c *AddressesClient) Insert(ctx context.Context, req *computepb.InsertAddressRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.Insert(ctx, req, opts...)
}

// List retrieves a list of addresses contained within the specified region.
func (c *AddressesClient) List(ctx context.Context, req *computepb.ListAddressesRequest, opts ...gax.CallOption) (*computepb.AddressList, error) {
	return c.internalClient.List(ctx, req, opts...)
}

// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type addressesRESTClient struct {
	// The http endpoint to connect to.
	endpoint string

	// The http client.
	httpClient *http.Client

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewAddressesRESTClient creates a new addresses rest client.
//
// The Addresses API.
func NewAddressesRESTClient(ctx context.Context, opts ...option.ClientOption) (*AddressesClient, error) {
	clientOpts := append(defaultAddressesRESTClientOptions(), opts...)
	httpClient, endpoint, err := httptransport.NewClient(ctx, clientOpts...)
	if err != nil {
		return nil, err
	}

	c := &addressesRESTClient{
		endpoint:   endpoint,
		httpClient: httpClient,
	}
	c.setGoogleClientInfo()

	return &AddressesClient{internalClient: c, CallOptions: &AddressesCallOptions{}}, nil
}

func defaultAddressesRESTClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("compute.googleapis.com"),
		internaloption.WithDefaultMTLSEndpoint("compute.mtls.googleapis.com"),
		internaloption.WithDefaultAudience("https://compute.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
	}
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *addressesRESTClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "rest", "UNKNOWN")
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *addressesRESTClient) Close() error {
	// Replace httpClient with nil to force cleanup.
	c.httpClient = nil
	return nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *addressesRESTClient) Connection() *grpc.ClientConn {
	return nil
}

// AggregatedList retrieves an aggregated list of addresses.
func (c *addressesRESTClient) AggregatedList(ctx context.Context, req *computepb.AggregatedListAddressesRequest, opts ...gax.CallOption) (*computepb.AddressAggregatedList, error) {
	m := protojson.MarshalOptions{AllowPartial: true, EmitUnpopulated: true}
	jsonReq, err := m.Marshal(req)
	if err != nil {
		return nil, err
	}

	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/aggregated/addresses", req.GetProject())

	params := url.Values{}
	if req != nil && req.Filter != nil {
		params.Add("filter", fmt.Sprintf("%v", req.GetFilter()))
	}
	if req != nil && req.IncludeAllScopes != nil {
		params.Add("includeAllScopes", fmt.Sprintf("%v", req.GetIncludeAllScopes()))
	}
	if req != nil && req.MaxResults != nil {
		params.Add("maxResults", fmt.Sprintf("%v", req.GetMaxResults()))
	}
	if req != nil && req.OrderBy != nil {
		params.Add("orderBy", fmt.Sprintf("%v", req.GetOrderBy()))
	}
	if req != nil && req.PageToken != nil {
		params.Add("pageToken", fmt.Sprintf("%v", req.GetPageToken()))
	}
	if req != nil && req.ReturnPartialSuccess != nil {
		params.Add("returnPartialSuccess", fmt.Sprintf("%v", req.GetReturnPartialSuccess()))
	}

	baseUrl.RawQuery = params.Encode()

	httpReq, err := http.NewRequest("GET", baseUrl.String(), bytes.NewReader(jsonReq))
	if err != nil {
		return nil, err
	}
	httpReq = httpReq.WithContext(ctx)
	// Set the headers
	for k, v := range c.xGoogMetadata {
		httpReq.Header[k] = v
	}
	httpReq.Header["Content-Type"] = []string{"application/json"}

	httpRsp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpRsp.Body.Close()

	if httpRsp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(httpRsp.Status)
	}

	buf, err := ioutil.ReadAll(httpRsp.Body)
	if err != nil {
		return nil, err
	}

	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	rsp := &computepb.AddressAggregatedList{}

	return rsp, unm.Unmarshal(buf, rsp)
}

// Delete deletes the specified address resource.
func (c *addressesRESTClient) Delete(ctx context.Context, req *computepb.DeleteAddressRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	m := protojson.MarshalOptions{AllowPartial: true, EmitUnpopulated: true}
	jsonReq, err := m.Marshal(req)
	if err != nil {
		return nil, err
	}

	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/regions/%v/addresses/%v", req.GetProject(), req.GetRegion(), req.GetAddress())

	params := url.Values{}
	if req != nil && req.RequestId != nil {
		params.Add("requestId", fmt.Sprintf("%v", req.GetRequestId()))
	}

	baseUrl.RawQuery = params.Encode()

	httpReq, err := http.NewRequest("DELETE", baseUrl.String(), bytes.NewReader(jsonReq))
	if err != nil {
		return nil, err
	}
	httpReq = httpReq.WithContext(ctx)
	// Set the headers
	for k, v := range c.xGoogMetadata {
		httpReq.Header[k] = v
	}
	httpReq.Header["Content-Type"] = []string{"application/json"}

	httpRsp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpRsp.Body.Close()

	if httpRsp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(httpRsp.Status)
	}

	buf, err := ioutil.ReadAll(httpRsp.Body)
	if err != nil {
		return nil, err
	}

	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	rsp := &computepb.Operation{}

	return rsp, unm.Unmarshal(buf, rsp)
}

// Get returns the specified address resource.
func (c *addressesRESTClient) Get(ctx context.Context, req *computepb.GetAddressRequest, opts ...gax.CallOption) (*computepb.Address, error) {
	m := protojson.MarshalOptions{AllowPartial: true, EmitUnpopulated: true}
	jsonReq, err := m.Marshal(req)
	if err != nil {
		return nil, err
	}

	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/regions/%v/addresses/%v", req.GetProject(), req.GetRegion(), req.GetAddress())

	httpReq, err := http.NewRequest("GET", baseUrl.String(), bytes.NewReader(jsonReq))
	if err != nil {
		return nil, err
	}
	httpReq = httpReq.WithContext(ctx)
	// Set the headers
	for k, v := range c.xGoogMetadata {
		httpReq.Header[k] = v
	}
	httpReq.Header["Content-Type"] = []string{"application/json"}

	httpRsp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpRsp.Body.Close()

	if httpRsp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(httpRsp.Status)
	}

	buf, err := ioutil.ReadAll(httpRsp.Body)
	if err != nil {
		return nil, err
	}

	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	rsp := &computepb.Address{}

	return rsp, unm.Unmarshal(buf, rsp)
}

// Insert creates an address resource in the specified project by using the data included in the request.
func (c *addressesRESTClient) Insert(ctx context.Context, req *computepb.InsertAddressRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	m := protojson.MarshalOptions{AllowPartial: true, EmitUnpopulated: true}
	body := req.GetAddressResource()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/regions/%v/addresses", req.GetProject(), req.GetRegion())

	params := url.Values{}
	if req != nil && req.RequestId != nil {
		params.Add("requestId", fmt.Sprintf("%v", req.GetRequestId()))
	}

	baseUrl.RawQuery = params.Encode()

	httpReq, err := http.NewRequest("POST", baseUrl.String(), bytes.NewReader(jsonReq))
	if err != nil {
		return nil, err
	}
	httpReq = httpReq.WithContext(ctx)
	// Set the headers
	for k, v := range c.xGoogMetadata {
		httpReq.Header[k] = v
	}
	httpReq.Header["Content-Type"] = []string{"application/json"}

	httpRsp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpRsp.Body.Close()

	if httpRsp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(httpRsp.Status)
	}

	buf, err := ioutil.ReadAll(httpRsp.Body)
	if err != nil {
		return nil, err
	}

	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	rsp := &computepb.Operation{}

	return rsp, unm.Unmarshal(buf, rsp)
}

// List retrieves a list of addresses contained within the specified region.
func (c *addressesRESTClient) List(ctx context.Context, req *computepb.ListAddressesRequest, opts ...gax.CallOption) (*computepb.AddressList, error) {
	m := protojson.MarshalOptions{AllowPartial: true, EmitUnpopulated: true}
	jsonReq, err := m.Marshal(req)
	if err != nil {
		return nil, err
	}

	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/regions/%v/addresses", req.GetProject(), req.GetRegion())

	params := url.Values{}
	if req != nil && req.Filter != nil {
		params.Add("filter", fmt.Sprintf("%v", req.GetFilter()))
	}
	if req != nil && req.MaxResults != nil {
		params.Add("maxResults", fmt.Sprintf("%v", req.GetMaxResults()))
	}
	if req != nil && req.OrderBy != nil {
		params.Add("orderBy", fmt.Sprintf("%v", req.GetOrderBy()))
	}
	if req != nil && req.PageToken != nil {
		params.Add("pageToken", fmt.Sprintf("%v", req.GetPageToken()))
	}
	if req != nil && req.ReturnPartialSuccess != nil {
		params.Add("returnPartialSuccess", fmt.Sprintf("%v", req.GetReturnPartialSuccess()))
	}

	baseUrl.RawQuery = params.Encode()

	httpReq, err := http.NewRequest("GET", baseUrl.String(), bytes.NewReader(jsonReq))
	if err != nil {
		return nil, err
	}
	httpReq = httpReq.WithContext(ctx)
	// Set the headers
	for k, v := range c.xGoogMetadata {
		httpReq.Header[k] = v
	}
	httpReq.Header["Content-Type"] = []string{"application/json"}

	httpRsp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpRsp.Body.Close()

	if httpRsp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(httpRsp.Status)
	}

	buf, err := ioutil.ReadAll(httpRsp.Body)
	if err != nil {
		return nil, err
	}

	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	rsp := &computepb.AddressList{}

	return rsp, unm.Unmarshal(buf, rsp)
}
