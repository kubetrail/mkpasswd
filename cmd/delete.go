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

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete all versions of named passwd",
	Long: `This command will delete all versions of
the named password.`,
	RunE:    run.Delete,
	Args:    cobra.ExactArgs(1),
	Example: fmt.Sprintf("  %s delete my-passwd", run.AppName),
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	f := deleteCmd.Flags()
	b := filepath.Base

	f.Bool(b(flags.Force), false, "Force delete without asking confirmation")
}
