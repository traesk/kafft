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

	"github.com/traesk/kafft/util"

	"github.com/spf13/cobra"
	"github.com/traesk/kafft/crypt"
	"golang.org/x/crypto/ssh/terminal"
)

// openCmd represents the open command
var openCmd = &cobra.Command{
	Use:   "unlock",
	Short: "Decrypt a file",
	Long:  `Decrypts a file, pass in the name of the file as argument.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("No file specified")
		}

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
		fmt.Println("Enter password to unlock file: ")
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
		files, err := util.Unzip(dir+name, dir)
		if err != nil {
			log.Fatal(err)
		}
		for _, f := range files {
			fmt.Println("File(s) decrypted: ", f)
		}

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
