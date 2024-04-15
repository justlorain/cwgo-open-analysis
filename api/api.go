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

package api

import (
	"context"

	"github.com/cloudwego-contrib/cwgo-open-analysis/client/graphql"
	"github.com/cloudwego-contrib/cwgo-open-analysis/client/rest"
	"github.com/cloudwego-contrib/cwgo-open-analysis/config"
	"github.com/cloudwego-contrib/cwgo-open-analysis/cron"
	"github.com/cloudwego-contrib/cwgo-open-analysis/storage"
)

func Start(ctx context.Context) error {
	return cron.Start(ctx)
}

func Restart(ctx context.Context) error {
	return cron.Restart(ctx)
}

func ReadInConfig(path string) error {
	if path == "" {
		path = "./default.yaml"
	}
	if err := config.GlobalConfig.ReadInConfig(path); err != nil {
		return err
	}
	return nil
}

func Init() error {
	if err := storage.Init(); err != nil {
		return err
	}
	// NOTE: graphql client MUST initialize before rest client due to dependency
	graphql.Init()
	rest.Init()
	return nil
}

func AddGroups(groups ...config.Group) {
	config.GlobalConfig.Groups = append(config.GlobalConfig.Groups, groups...)
}

func SetDataSource(ds config.DataSource) {
	config.GlobalConfig.DataSource = ds
}

func SetBackend(be config.Backend) {
	config.GlobalConfig.Backend = be
}

func SetCron(spec string) {
	config.GlobalConfig.Backend.Cron = spec
}

func SetToken(token string) {
	config.GlobalConfig.Backend.Token = token
}

func SetRetry(times int) {
	config.GlobalConfig.Backend.Retry = times
}
