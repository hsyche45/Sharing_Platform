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

package compute_test

import (
	"context"

	compute "cloud.google.com/go/compute/apiv1"
	computepb "google.golang.org/genproto/googleapis/cloud/compute/v1"
)

func ExampleNewNetworkEndpointGroupsRESTClient() {
	ctx := context.Background()
	c, err := compute.NewNetworkEndpointGroupsRESTClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleNetworkEndpointGroupsClient_AggregatedList() {
	ctx := context.Background()
	c, err := compute.NewNetworkEndpointGroupsRESTClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &computepb.AggregatedListNetworkEndpointGroupsRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.AggregatedList(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleNetworkEndpointGroupsClient_AttachNetworkEndpoints() {
	ctx := context.Background()
	c, err := compute.NewNetworkEndpointGroupsRESTClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &computepb.AttachNetworkEndpointsNetworkEndpointGroupRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.AttachNetworkEndpoints(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleNetworkEndpointGroupsClient_Delete() {
	ctx := context.Background()
	c, err := compute.NewNetworkEndpointGroupsRESTClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &computepb.DeleteNetworkEndpointGroupRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.Delete(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleNetworkEndpointGroupsClient_DetachNetworkEndpoints() {
	ctx := context.Background()
	c, err := compute.NewNetworkEndpointGroupsRESTClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &computepb.DetachNetworkEndpointsNetworkEndpointGroupRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.DetachNetworkEndpoints(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleNetworkEndpointGroupsClient_Get() {
	ctx := context.Background()
	c, err := compute.NewNetworkEndpointGroupsRESTClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &computepb.GetNetworkEndpointGroupRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.Get(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleNetworkEndpointGroupsClient_Insert() {
	ctx := context.Background()
	c, err := compute.NewNetworkEndpointGroupsRESTClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &computepb.InsertNetworkEndpointGroupRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.Insert(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleNetworkEndpointGroupsClient_List() {
	ctx := context.Background()
	c, err := compute.NewNetworkEndpointGroupsRESTClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &computepb.ListNetworkEndpointGroupsRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.List(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleNetworkEndpointGroupsClient_ListNetworkEndpoints() {
	ctx := context.Background()
	c, err := compute.NewNetworkEndpointGroupsRESTClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &computepb.ListNetworkEndpointsNetworkEndpointGroupsRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.ListNetworkEndpoints(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleNetworkEndpointGroupsClient_TestIamPermissions() {
	ctx := context.Background()
	c, err := compute.NewNetworkEndpointGroupsRESTClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &computepb.TestIamPermissionsNetworkEndpointGroupRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.TestIamPermissions(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}
