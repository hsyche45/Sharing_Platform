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

package privateca_test

import (
	"context"

	privateca "cloud.google.com/go/security/privateca/apiv1beta1"
	"google.golang.org/api/iterator"
	privatecapb "google.golang.org/genproto/googleapis/cloud/security/privateca/v1beta1"
)

func ExampleNewCertificateAuthorityClient() {
	ctx := context.Background()
	c, err := privateca.NewCertificateAuthorityClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleCertificateAuthorityClient_CreateCertificate() {
	ctx := context.Background()
	c, err := privateca.NewCertificateAuthorityClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &privatecapb.CreateCertificateRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.CreateCertificate(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleCertificateAuthorityClient_GetCertificate() {
	ctx := context.Background()
	c, err := privateca.NewCertificateAuthorityClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &privatecapb.GetCertificateRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GetCertificate(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleCertificateAuthorityClient_ListCertificates() {
	ctx := context.Background()
	c, err := privateca.NewCertificateAuthorityClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &privatecapb.ListCertificatesRequest{
		// TODO: Fill request struct fields.
	}
	it := c.ListCertificates(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleCertificateAuthorityClient_RevokeCertificate() {
	ctx := context.Background()
	c, err := privateca.NewCertificateAuthorityClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &privatecapb.RevokeCertificateRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.RevokeCertificate(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleCertificateAuthorityClient_UpdateCertificate() {
	ctx := context.Background()
	c, err := privateca.NewCertificateAuthorityClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &privatecapb.UpdateCertificateRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.UpdateCertificate(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleCertificateAuthorityClient_ActivateCertificateAuthority() {
	ctx := context.Background()
	c, err := privateca.NewCertificateAuthorityClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &privatecapb.ActivateCertificateAuthorityRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.ActivateCertificateAuthority(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleCertificateAuthorityClient_CreateCertificateAuthority() {
	ctx := context.Background()
	c, err := privateca.NewCertificateAuthorityClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &privatecapb.CreateCertificateAuthorityRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.CreateCertificateAuthority(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleCertificateAuthorityClient_DisableCertificateAuthority() {
	ctx := context.Background()
	c, err := privateca.NewCertificateAuthorityClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &privatecapb.DisableCertificateAuthorityRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.DisableCertificateAuthority(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleCertificateAuthorityClient_EnableCertificateAuthority() {
	ctx := context.Background()
	c, err := privateca.NewCertificateAuthorityClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &privatecapb.EnableCertificateAuthorityRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.EnableCertificateAuthority(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleCertificateAuthorityClient_FetchCertificateAuthorityCsr() {
	ctx := context.Background()
	c, err := privateca.NewCertificateAuthorityClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &privatecapb.FetchCertificateAuthorityCsrRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.FetchCertificateAuthorityCsr(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleCertificateAuthorityClient_GetCertificateAuthority() {
	ctx := context.Background()
	c, err := privateca.NewCertificateAuthorityClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &privatecapb.GetCertificateAuthorityRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GetCertificateAuthority(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleCertificateAuthorityClient_ListCertificateAuthorities() {
	ctx := context.Background()
	c, err := privateca.NewCertificateAuthorityClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &privatecapb.ListCertificateAuthoritiesRequest{
		// TODO: Fill request struct fields.
	}
	it := c.ListCertificateAuthorities(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleCertificateAuthorityClient_RestoreCertificateAuthority() {
	ctx := context.Background()
	c, err := privateca.NewCertificateAuthorityClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &privatecapb.RestoreCertificateAuthorityRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.RestoreCertificateAuthority(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleCertificateAuthorityClient_ScheduleDeleteCertificateAuthority() {
	ctx := context.Background()
	c, err := privateca.NewCertificateAuthorityClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &privatecapb.ScheduleDeleteCertificateAuthorityRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.ScheduleDeleteCertificateAuthority(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleCertificateAuthorityClient_UpdateCertificateAuthority() {
	ctx := context.Background()
	c, err := privateca.NewCertificateAuthorityClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &privatecapb.UpdateCertificateAuthorityRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.UpdateCertificateAuthority(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleCertificateAuthorityClient_GetCertificateRevocationList() {
	ctx := context.Background()
	c, err := privateca.NewCertificateAuthorityClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &privatecapb.GetCertificateRevocationListRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GetCertificateRevocationList(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleCertificateAuthorityClient_ListCertificateRevocationLists() {
	ctx := context.Background()
	c, err := privateca.NewCertificateAuthorityClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &privatecapb.ListCertificateRevocationListsRequest{
		// TODO: Fill request struct fields.
	}
	it := c.ListCertificateRevocationLists(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleCertificateAuthorityClient_UpdateCertificateRevocationList() {
	ctx := context.Background()
	c, err := privateca.NewCertificateAuthorityClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &privatecapb.UpdateCertificateRevocationListRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.UpdateCertificateRevocationList(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleCertificateAuthorityClient_GetReusableConfig() {
	ctx := context.Background()
	c, err := privateca.NewCertificateAuthorityClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &privatecapb.GetReusableConfigRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GetReusableConfig(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleCertificateAuthorityClient_ListReusableConfigs() {
	ctx := context.Background()
	c, err := privateca.NewCertificateAuthorityClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &privatecapb.ListReusableConfigsRequest{
		// TODO: Fill request struct fields.
	}
	it := c.ListReusableConfigs(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}
