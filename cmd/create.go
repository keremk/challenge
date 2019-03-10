// Copyright © 2019 Kerem Karatal
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
	"fmt"

	challenge "github.com/keremk/challenge/lib"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates a coding challenge for a candidate",
	Long: `This command creates a coding challenge for a candidate in the configured 
repository. You need to supply the github name of the candidate and the discipline of 
the coding challenge to be created. Discipline needs to match one that is in the configuration 
file. This command also creates the first task, adds the coding challenge as an issue in the
tracking repository and invites the candidate the coding challenge.`,
	Run: func(cmd *cobra.Command, args []string) {
		user, _ := cmd.Flags().GetString("user")
		if user == "" {
			user = "no user specified"
		}
		discipline, _ := cmd.Flags().GetString("discipline")
		if discipline == "" {
			discipline = "no discipline specified"
		}
		fmt.Println("create called for", user, discipline)

		challenge.CreateChallenge(user, discipline)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	createCmd.Flags().StringP("user", "u", "", "Provide Github user name for the candidate")
	createCmd.Flags().StringP("discipline", "d", "", "Provide the discipline (e.g. ios, android, backend ...) for the challenge")
}
