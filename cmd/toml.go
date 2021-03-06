// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
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
	"io/ioutil"
	"log"
	"os"
	"regexp"

	plog "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// tomlCmd represents the toml command
var tomlCmd = &cobra.Command{
	Use:   "toml",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		dependency := args[0]

		dir, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		file, err := os.Open(dir + "/Gopkg.toml")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		plog.Info("File " + file.Name())

		bts, err := ioutil.ReadFile(file.Name())

		re := regexp.MustCompile(`(?m)` + dependency + `"\n  version = "(.*)"`)

		text := string(bts)
		if re.Match([]byte(text)) {
			fmt.Println(re.FindStringSubmatch(text)[1])
		}

	},
}

func init() {
	rootCmd.AddCommand(tomlCmd)
}
