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

package speech_test

import (
	"context"

	speech "cloud.google.com/go/speech/apiv1p1beta1"
	"google.golang.org/api/iterator"
	speechpb "google.golang.org/genproto/googleapis/cloud/speech/v1p1beta1"
)

func ExampleNewAdaptationClient() {
	ctx := context.Background()
	c, err := speech.NewAdaptationClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleAdaptationClient_CreatePhraseSet() {
	ctx := context.Background()
	c, err := speech.NewAdaptationClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &speechpb.CreatePhraseSetRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.CreatePhraseSet(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleAdaptationClient_GetPhraseSet() {
	ctx := context.Background()
	c, err := speech.NewAdaptationClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &speechpb.GetPhraseSetRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GetPhraseSet(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleAdaptationClient_ListPhraseSet() {
	ctx := context.Background()
	c, err := speech.NewAdaptationClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &speechpb.ListPhraseSetRequest{
		// TODO: Fill request struct fields.
	}
	it := c.ListPhraseSet(ctx, req)
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

func ExampleAdaptationClient_UpdatePhraseSet() {
	ctx := context.Background()
	c, err := speech.NewAdaptationClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &speechpb.UpdatePhraseSetRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.UpdatePhraseSet(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleAdaptationClient_DeletePhraseSet() {
	ctx := context.Background()
	c, err := speech.NewAdaptationClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &speechpb.DeletePhraseSetRequest{
		// TODO: Fill request struct fields.
	}
	err = c.DeletePhraseSet(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleAdaptationClient_CreateCustomClass() {
	ctx := context.Background()
	c, err := speech.NewAdaptationClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &speechpb.CreateCustomClassRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.CreateCustomClass(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleAdaptationClient_GetCustomClass() {
	ctx := context.Background()
	c, err := speech.NewAdaptationClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &speechpb.GetCustomClassRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GetCustomClass(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleAdaptationClient_ListCustomClasses() {
	ctx := context.Background()
	c, err := speech.NewAdaptationClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &speechpb.ListCustomClassesRequest{
		// TODO: Fill request struct fields.
	}
	it := c.ListCustomClasses(ctx, req)
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

func ExampleAdaptationClient_UpdateCustomClass() {
	ctx := context.Background()
	c, err := speech.NewAdaptationClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &speechpb.UpdateCustomClassRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.UpdateCustomClass(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleAdaptationClient_DeleteCustomClass() {
	ctx := context.Background()
	c, err := speech.NewAdaptationClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &speechpb.DeleteCustomClassRequest{
		// TODO: Fill request struct fields.
	}
	err = c.DeleteCustomClass(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}
