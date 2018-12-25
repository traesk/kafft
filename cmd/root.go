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

	"golang.org/x/crypto/ssh/terminal"

	"github.com/spf13/cobra"
	"github.com/traesk/kafft/crypt"
)

var delete bool
var uniquePassword bool
var dir string
var printPassword bool
var save bool
var folder bool

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.Flags().BoolVarP(&delete, "delete", "d", false, "Delete original file after encryption.")
	rootCmd.Flags().BoolVarP(&uniquePassword, "uniquepassword", "u", false, "Make it possible to set a unique password for each file.")
	rootCmd.Flags().BoolVarP(&printPassword, "printpassword", "p", false, "Print the password set for each file.")
	rootCmd.Flags().BoolVarP(&save, "save", "s", false, "Write filename and password to file. Unsafe!")
	rootCmd.Flags().BoolVarP(&folder, "folder", "f", false, "Zips and encrypts a folder, with all its files. ")
}
func initConfig() {

}

var rootCmd = &cobra.Command{
	Use:   "kafft",
	Short: "Encrypt a file.",
	Long:  "Encrypts a file, pass in the file name as argument.",
	Args: func(cmd *cobra.Command, args []string) error {
		for _, arg := range args {

			if len(args) < 1 {
				return errors.New("No file specified")
			}

			dir, err := os.Getwd()
			if err != nil {
				log.Fatal(err)
				os.Exit(1)
			}
			if _, err := os.Open(dir + arg); !os.IsNotExist(err) {
				return fmt.Errorf("Could not find file: %s", args[0])

			}
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		var password []byte
		var passwordEntered bool
		for _, arg := range args {

			if uniquePassword || !passwordEntered {
				password = inputPassword()
				passwordEntered = true
			}
			outputName, err := crypt.Encrypt(dir, arg, password, delete)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			// Tell the user it's done
			if printPassword {
				fmt.Print(printInfoPassword(outputName, password))
			} else if save {
				util.SaveInfo(outputName, string(password))
				if err != nil {
					log.Fatal(err)
					os.Exit(1)
				}
			} else {
				printInfo(outputName)

			}
		}
		fmt.Println()
	},
}

// Execute the application
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
func printInfo(outputName string) string {
	return fmt.Sprintf(`
	File encrypted, please remember the password.
	New name: %s`, outputName)

}
func printInfoPassword(outputName string, password []byte) string {
	return fmt.Sprintf(`
	File encrypted, please remember the password.
	New name: %s
	Password: %s
`, outputName, string(password))

}
func inputPassword() []byte {
	fmt.Println("\nEnter password to lock file: ")
	password, err := terminal.ReadPassword(0)
	if err != nil {
		fmt.Println("Erroneous password")
		os.Exit(1)
	}
	fmt.Println(`
Password entered. Encrypting...`)
	return password
}
