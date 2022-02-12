/*
Copyright © 2022 kubetrail.io authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/kubetrail/mkpasswd/pkg/flags"
	"github.com/kubetrail/mkpasswd/pkg/run"
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get password payload value",
	Long: `Retrieve latest version of the named password
or get a specific version`,
	RunE:    run.Get,
	Example: fmt.Sprintf("  %s get my-passwd\n  %s get my-passwd --version=7", run.AppName, run.AppName),
	Args:    cobra.ExactArgs(1),
}

func init() {
	rootCmd.AddCommand(getCmd)
	f := getCmd.Flags()
	b := filepath.Base

	f.String(b(flags.Version), "latest", "Get specific version")
}
