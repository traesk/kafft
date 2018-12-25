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
	"errors"
	"fmt"
	"log"
	"os"

	"golang.org/x/crypto/ssh/terminal"

	"github.com/spf13/cobra"
	"github.com/traesk/kafft/crypt"
)

var delete bool
var dir string

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.Flags().BoolVarP(&delete, "delete", "d", false, "delete original file after encryption")
}
func initConfig() {

}

var rootCmd = &cobra.Command{
	Use:   "kafft",
	Short: "Encrypt a file.",
	Long:  "Encrypts a file, pass in the file name as argument.",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("No file specified")
		}
		// ../
		dir, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		if _, err := os.Open(dir + args[0]); os.IsNotExist(err) {
			return nil
		}
		return fmt.Errorf("Could not find file: %s", args[0])
	},
	Run: func(cmd *cobra.Command, args []string) {
		// pew pew
		fmt.Println("Enter password: ")
		password, err := terminal.ReadPassword(0)

		if err != nil {
			fmt.Println("Erroneous password")
			os.Exit(1)
		}

		outputName, err := crypt.Encrypt(dir, args[0], password, delete)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		// Tell the user it's done
		fmt.Println("File encrypted, please remember the password")
		fmt.Println("New name: ", outputName)

	},
}

// Execute gets rid of squigglies
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
