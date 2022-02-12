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
	"path/filepath"

	"github.com/kubetrail/mkpasswd/pkg/flags"
	"github.com/kubetrail/mkpasswd/pkg/run"
	"github.com/spf13/cobra"
)

// genCmd represents the gen command
var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "Generate password",
	Long: `Generate a new password and store in Google
secret manager. If the named password already exists
a new version will be created without destroying
the older version`,
	RunE: run.Gen,
}

func init() {
	rootCmd.AddCommand(genCmd)
	f := genCmd.Flags()
	b := filepath.Base

	f.String(b(flags.Name), "", "Name tag for the password (DNS1123 label format)")
	f.Int(b(flags.Length), 16, "Length of password")
	f.Int(b(flags.NumSymbols), 3, "Number of symbols")
	f.Int(b(flags.NumDigits), 4, "Number of digits")
	f.Bool(b(flags.NoUppercase), false, "No uppercase chars")
	f.Bool(b(flags.AllowRepeat), false, "Allow repeats")
}
