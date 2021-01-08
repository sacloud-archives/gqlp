// Copyright 2021 The gqlp Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package graph

//go:generate gqlgen generate

import (
	"fmt"
	"net/http"
	"os"

	"github.com/sacloud/gqlp/graph/model"
	"github.com/sacloud/gqlp/version"
	"github.com/sacloud/libsacloud/v2"
	"github.com/sacloud/libsacloud/v2/helper/api"
	"github.com/sacloud/libsacloud/v2/sacloud"
)

var httpClient = &http.Client{Transport: http.DefaultTransport}

type Resolver struct{}

func (r *Resolver) APICaller() sacloud.APICaller {
	return api.NewCaller(&api.CallerOptions{
		AccessToken:          os.Getenv("SAKURACLOUD_ACCESS_TOKEN"),
		AccessTokenSecret:    os.Getenv("SAKURACLOUD_ACCESS_TOKEN_SECRET"),
		HTTPClient:           httpClient,
		HTTPRequestRateLimit: 5,
		OpenTelemetry:        os.Getenv("TRACE") != "",
		UserAgent:            fmt.Sprintf("sacloud/gqlp/v%s (+https://github.com/sacloud/gqlp) libsacloud/%s", version.Version, libsacloud.Version),
	})
}

func (r *Resolver) MutationResult(err error) (*model.MutationResult, error) {
	return &model.MutationResult{Success: err == nil}, err
}
