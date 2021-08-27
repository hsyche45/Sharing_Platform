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

var newPacketMirroringsClientHook clientHook

// PacketMirroringsCallOptions contains the retry settings for each method of PacketMirroringsClient.
type PacketMirroringsCallOptions struct {
	AggregatedList     []gax.CallOption
	Delete             []gax.CallOption
	Get                []gax.CallOption
	Insert             []gax.CallOption
	List               []gax.CallOption
	Patch              []gax.CallOption
	TestIamPermissions []gax.CallOption
}

// internalPacketMirroringsClient is an interface that defines the methods availaible from Google Compute Engine API.
type internalPacketMirroringsClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	AggregatedList(context.Context, *computepb.AggregatedListPacketMirroringsRequest, ...gax.CallOption) (*computepb.PacketMirroringAggregatedList, error)
	Delete(context.Context, *computepb.DeletePacketMirroringRequest, ...gax.CallOption) (*computepb.Operation, error)
	Get(context.Context, *computepb.GetPacketMirroringRequest, ...gax.CallOption) (*computepb.PacketMirroring, error)
	Insert(context.Context, *computepb.InsertPacketMirroringRequest, ...gax.CallOption) (*computepb.Operation, error)
	List(context.Context, *computepb.ListPacketMirroringsRequest, ...gax.CallOption) (*computepb.PacketMirroringList, error)
	Patch(context.Context, *computepb.PatchPacketMirroringRequest, ...gax.CallOption) (*computepb.Operation, error)
	TestIamPermissions(context.Context, *computepb.TestIamPermissionsPacketMirroringRequest, ...gax.CallOption) (*computepb.TestPermissionsResponse, error)
}

// PacketMirroringsClient is a client for interacting with Google Compute Engine API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// The PacketMirrorings API.
type PacketMirroringsClient struct {
	// The internal transport-dependent client.
	internalClient internalPacketMirroringsClient

	// The call options for this service.
	CallOptions *PacketMirroringsCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *PacketMirroringsClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *PacketMirroringsClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *PacketMirroringsClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// AggregatedList retrieves an aggregated list of packetMirrorings.
func (c *PacketMirroringsClient) AggregatedList(ctx context.Context, req *computepb.AggregatedListPacketMirroringsRequest, opts ...gax.CallOption) (*computepb.PacketMirroringAggregatedList, error) {
	return c.internalClient.AggregatedList(ctx, req, opts...)
}

// Delete deletes the specified PacketMirroring resource.
func (c *PacketMirroringsClient) Delete(ctx context.Context, req *computepb.DeletePacketMirroringRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.Delete(ctx, req, opts...)
}

// Get returns the specified PacketMirroring resource.
func (c *PacketMirroringsClient) Get(ctx context.Context, req *computepb.GetPacketMirroringRequest, opts ...gax.CallOption) (*computepb.PacketMirroring, error) {
	return c.internalClient.Get(ctx, req, opts...)
}

// Insert creates a PacketMirroring resource in the specified project and region using the data included in the request.
func (c *PacketMirroringsClient) Insert(ctx context.Context, req *computepb.InsertPacketMirroringRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.Insert(ctx, req, opts...)
}

// List retrieves a list of PacketMirroring resources available to the specified project and region.
func (c *PacketMirroringsClient) List(ctx context.Context, req *computepb.ListPacketMirroringsRequest, opts ...gax.CallOption) (*computepb.PacketMirroringList, error) {
	return c.internalClient.List(ctx, req, opts...)
}

// Patch patches the specified PacketMirroring resource with the data included in the request. This method supports PATCH semantics and uses JSON merge patch format and processing rules.
func (c *PacketMirroringsClient) Patch(ctx context.Context, req *computepb.PatchPacketMirroringRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.Patch(ctx, req, opts...)
}

// TestIamPermissions returns permissions that a caller has on the specified resource.
func (c *PacketMirroringsClient) TestIamPermissions(ctx context.Context, req *computepb.TestIamPermissionsPacketMirroringRequest, opts ...gax.CallOption) (*computepb.TestPermissionsResponse, error) {
	return c.internalClient.TestIamPermissions(ctx, req, opts...)
}

// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type packetMirroringsRESTClient struct {
	// The http endpoint to connect to.
	endpoint string

	// The http client.
	httpClient *http.Client

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewPacketMirroringsRESTClient creates a new packet mirrorings rest client.
//
// The PacketMirrorings API.
func NewPacketMirroringsRESTClient(ctx context.Context, opts ...option.ClientOption) (*PacketMirroringsClient, error) {
	clientOpts := append(defaultPacketMirroringsRESTClientOptions(), opts...)
	httpClient, endpoint, err := httptransport.NewClient(ctx, clientOpts...)
	if err != nil {
		return nil, err
	}

	c := &packetMirroringsRESTClient{
		endpoint:   endpoint,
		httpClient: httpClient,
	}
	c.setGoogleClientInfo()

	return &PacketMirroringsClient{internalClient: c, CallOptions: &PacketMirroringsCallOptions{}}, nil
}

func defaultPacketMirroringsRESTClientOptions() []option.ClientOption {
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
func (c *packetMirroringsRESTClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "rest", "UNKNOWN")
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *packetMirroringsRESTClient) Close() error {
	// Replace httpClient with nil to force cleanup.
	c.httpClient = nil
	return nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *packetMirroringsRESTClient) Connection() *grpc.ClientConn {
	return nil
}

// AggregatedList retrieves an aggregated list of packetMirrorings.
func (c *packetMirroringsRESTClient) AggregatedList(ctx context.Context, req *computepb.AggregatedListPacketMirroringsRequest, opts ...gax.CallOption) (*computepb.PacketMirroringAggregatedList, error) {
	m := protojson.MarshalOptions{AllowPartial: true, EmitUnpopulated: true}
	jsonReq, err := m.Marshal(req)
	if err != nil {
		return nil, err
	}

	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/aggregated/packetMirrorings", req.GetProject())

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
	rsp := &computepb.PacketMirroringAggregatedList{}

	return rsp, unm.Unmarshal(buf, rsp)
}

// Delete deletes the specified PacketMirroring resource.
func (c *packetMirroringsRESTClient) Delete(ctx context.Context, req *computepb.DeletePacketMirroringRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	m := protojson.MarshalOptions{AllowPartial: true, EmitUnpopulated: true}
	jsonReq, err := m.Marshal(req)
	if err != nil {
		return nil, err
	}

	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/regions/%v/packetMirrorings/%v", req.GetProject(), req.GetRegion(), req.GetPacketMirroring())

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

// Get returns the specified PacketMirroring resource.
func (c *packetMirroringsRESTClient) Get(ctx context.Context, req *computepb.GetPacketMirroringRequest, opts ...gax.CallOption) (*computepb.PacketMirroring, error) {
	m := protojson.MarshalOptions{AllowPartial: true, EmitUnpopulated: true}
	jsonReq, err := m.Marshal(req)
	if err != nil {
		return nil, err
	}

	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/regions/%v/packetMirrorings/%v", req.GetProject(), req.GetRegion(), req.GetPacketMirroring())

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
	rsp := &computepb.PacketMirroring{}

	return rsp, unm.Unmarshal(buf, rsp)
}

// Insert creates a PacketMirroring resource in the specified project and region using the data included in the request.
func (c *packetMirroringsRESTClient) Insert(ctx context.Context, req *computepb.InsertPacketMirroringRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	m := protojson.MarshalOptions{AllowPartial: true, EmitUnpopulated: true}
	body := req.GetPacketMirroringResource()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/regions/%v/packetMirrorings", req.GetProject(), req.GetRegion())

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

// List retrieves a list of PacketMirroring resources available to the specified project and region.
func (c *packetMirroringsRESTClient) List(ctx context.Context, req *computepb.ListPacketMirroringsRequest, opts ...gax.CallOption) (*computepb.PacketMirroringList, error) {
	m := protojson.MarshalOptions{AllowPartial: true, EmitUnpopulated: true}
	jsonReq, err := m.Marshal(req)
	if err != nil {
		return nil, err
	}

	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/regions/%v/packetMirrorings", req.GetProject(), req.GetRegion())

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
	rsp := &computepb.PacketMirroringList{}

	return rsp, unm.Unmarshal(buf, rsp)
}

// Patch patches the specified PacketMirroring resource with the data included in the request. This method supports PATCH semantics and uses JSON merge patch format and processing rules.
func (c *packetMirroringsRESTClient) Patch(ctx context.Context, req *computepb.PatchPacketMirroringRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	m := protojson.MarshalOptions{AllowPartial: true, EmitUnpopulated: true}
	body := req.GetPacketMirroringResource()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/regions/%v/packetMirrorings/%v", req.GetProject(), req.GetRegion(), req.GetPacketMirroring())

	params := url.Values{}
	if req != nil && req.RequestId != nil {
		params.Add("requestId", fmt.Sprintf("%v", req.GetRequestId()))
	}

	baseUrl.RawQuery = params.Encode()

	httpReq, err := http.NewRequest("PATCH", baseUrl.String(), bytes.NewReader(jsonReq))
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

// TestIamPermissions returns permissions that a caller has on the specified resource.
func (c *packetMirroringsRESTClient) TestIamPermissions(ctx context.Context, req *computepb.TestIamPermissionsPacketMirroringRequest, opts ...gax.CallOption) (*computepb.TestPermissionsResponse, error) {
	m := protojson.MarshalOptions{AllowPartial: true, EmitUnpopulated: true}
	body := req.GetTestPermissionsRequestResource()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/regions/%v/packetMirrorings/%v/testIamPermissions", req.GetProject(), req.GetRegion(), req.GetResource())

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
	rsp := &computepb.TestPermissionsResponse{}

	return rsp, unm.Unmarshal(buf, rsp)
}
