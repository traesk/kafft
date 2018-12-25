// Copyright © 2018 Joel Kratz joel@kratz.nu
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
	"os"

	"github.com/spf13/cobra"
	"github.com/traesk/kafft/crypt"
	"golang.org/x/crypto/ssh/terminal"
)

// openCmd represents the open command
var openCmd = &cobra.Command{
	Use:   "open",
	Short: "Decrypt a file",
	Long:  `Decrypts a file, pass in the name of the file as argument.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Enter password: ")
		password, err := terminal.ReadPassword(0)

		if err != nil {
			fmt.Println("Erroneous password")
			os.Exit(1)
		}
		name, err := crypt.Decrypt(dir, args[0], password)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		// Tell the user it's done
		fmt.Println("File decrypted: ", name)

	},
}

func init() {
	rootCmd.AddCommand(openCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// openCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// openCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
