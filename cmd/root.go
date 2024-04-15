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
	"os"

	"github.com/spf13/cobra"
)

const Name = "cwgo-open-analysis"

var rootCmd = &cobra.Command{
	Use:     Name,
	Short:   "cwgo-open-analysis",
	Long:    `a service for analyzing cloudwego open-source community github data`,
	Version: "v0.1.0",
}

func init() {
	rootCmd.SetVersionTemplate("{{ .Version }}")
	rootCmd.CompletionOptions.DisableDefaultCmd = true
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var (
	TokenF string
	CronF  string
	RetryF int
)

var (
	defaultTokenF = ""
	defaultCronF  = ""
	defaultRetry  = -1
)

func setupCommand(cmd *cobra.Command) {
	rootCmd.AddCommand(cmd)

	cmd.Flags().StringVarP(&TokenF, "token", "t", defaultTokenF, "your github token")
	cmd.Flags().StringVarP(&CronF, "cron", "c", defaultCronF, "your cron spec")
	cmd.Flags().IntVarP(&RetryF, "retry", "r", defaultRetry, "retry times")
}
