// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package impersonate

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
	"time"

	"google.golang.org/api/option"
)

func TestTokenSource_serviceAccount(t *testing.T) {
	ctx := context.Background()
	tests := []struct {
		name            string
		targetPrincipal string
		scopes          []string
		lifetime        time.Duration
		wantErr         bool
	}{
		{
			name:    "missing targetPrincipal",
			wantErr: true,
		},
		{
			name:            "missing scopes",
			targetPrincipal: "foo@project-id.iam.gserviceaccount.com",
			wantErr:         true,
		},
		{
			name:            "lifetime over max",
			targetPrincipal: "foo@project-id.iam.gserviceaccount.com",
			scopes:          []string{"scope"},
			lifetime:        3601 * time.Second,
			wantErr:         true,
		},
		{
			name:            "works",
			targetPrincipal: "foo@project-id.iam.gserviceaccount.com",
			scopes:          []string{"scope"},
			wantErr:         false,
		},
	}

	for _, tt := range tests {
		name := tt.name
		t.Run(name, func(t *testing.T) {
			saTok := "sa-token"
			client := &http.Client{
				Transport: RoundTripFn(func(req *http.Request) *http.Response {
					if strings.Contains(req.URL.Path, "generateAccessToken") {
						resp := generateAccessTokenResp{
							AccessToken: saTok,
							ExpireTime:  time.Now().Format(time.RFC3339),
						}
						b, err := json.Marshal(&resp)
						if err != nil {
							t.Fatalf("unable to marshal response: %v", err)
						}
						return &http.Response{
							StatusCode: 200,
							Body:       ioutil.NopCloser(bytes.NewReader(b)),
							Header:     http.Header{},
						}
					}
					return nil
				}),
			}
			ts, err := CredentialsTokenSource(ctx, CredentialsConfig{
				TargetPrincipal: tt.targetPrincipal,
				Scopes:          tt.scopes,
				Lifetime:        tt.lifetime,
			}, option.WithHTTPClient(client))
			if tt.wantErr && err != nil {
				return
			}
			if err != nil {
				t.Fatal(err)
			}
			tok, err := ts.Token()
			if err != nil {
				t.Fatal(err)
			}
			if tok.AccessToken != saTok {
				t.Fatalf("got %q, want %q", tok.AccessToken, saTok)
			}
		})
	}
}

type RoundTripFn func(req *http.Request) *http.Response

func (f RoundTripFn) RoundTrip(req *http.Request) (*http.Response, error) { return f(req), nil }
