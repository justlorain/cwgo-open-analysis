// Copyright 2024 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package graphql

import (
	"context"

	"github.com/cloudwego-contrib/cwgo-open-analysis/config"

	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

var GlobalV4Client *githubv4.Client

// Init githubv4 graphql client
func Init() {
	src := oauth2.StaticTokenSource(
		&oauth2.Token{
			AccessToken: config.GlobalConfig.Backend.Token,
		},
	)
	httpClient := oauth2.NewClient(context.Background(), src)
	GlobalV4Client = githubv4.NewClient(httpClient)
}
