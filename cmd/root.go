// Copyright © 2016 Matthew Holt
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
)

var cfgFile string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "certs",
	Short: "Bulk issuance and maintenance of TLS certificates",
	Long: `Certs is a command line application that makes it easy to
use automated certificate authorities to obtain, renew,
and revoke TLS certificates. In other words, it is an ACME
client designed for bulk usage.`,
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

// func init() {
// 	cobra.OnInitialize(initConfig)

// 	// Here you will define your flags and configuration settings.
// 	// Cobra supports Persistent Flags, which, if defined here,
// 	// will be global for your application.

// 	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.certs.yaml)")
// 	// Cobra also supports local flags, which will only run
// 	// when this action is called directly.
// 	RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
// }

// // initConfig reads in config file and ENV variables if set.
// func initConfig() {
// 	if cfgFile != "" { // enable ability to specify config file via flag
// 		viper.SetConfigFile(cfgFile)
// 	}

// 	viper.SetConfigName(".certs") // name of config file (without extension)
// 	viper.AddConfigPath("$HOME")  // adding home directory as first search path
// 	viper.AutomaticEnv()          // read in environment variables that match

// 	// If a config file is found, read it in.
// 	if err := viper.ReadInConfig(); err == nil {
// 		fmt.Println("Using config file:", viper.ConfigFileUsed())
// 	}
// }
