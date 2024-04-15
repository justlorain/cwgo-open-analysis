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

package cmd

import (
	"context"
	"github.com/cloudwego-contrib/cwgo-open-analysis/api"
	"github.com/cloudwego-contrib/cwgo-open-analysis/util"

	"github.com/spf13/cobra"
)

// restartCmd represents the restart command
var restartCmd = &cobra.Command{
	Use:   "restart",
	Short: "restart cwgo-open-analysis service",
	Long: `restart cwgo-open-analysis service 
e.g. cwgo-open-analysis restart -c "cron-spec" path2config.yaml`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		configPath := ""
		if !util.IsEmptySlice(args) {
			configPath = args[0]
		}
		if err := api.ReadInConfig(configPath); err != nil {
			cobra.CheckErr(err)
		}
		if TokenF != "" {
			api.SetToken(TokenF)
		}
		if CronF != "" {
			api.SetCron(CronF)
		}
		if RetryF != -1 {
			api.SetRetry(RetryF)
		}
		if err := api.Init(); err != nil {
			cobra.CheckErr(err)
		}
		if err := api.Restart(context.Background()); err != nil {
			cobra.CheckErr(err)
		}
	},
}

func init() {
	setupCommand(restartCmd)
}
