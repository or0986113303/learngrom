/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	"cobragingorm/api"

	"github.com/spf13/cobra"
)

// ginCmd represents the gin command
var ginCmd = &cobra.Command{
	Use:   "gin",
	Short: "Start http service for grom daemon",
	Long:  `Start the http service for basic grom daemon`,
	Run: func(cmd *cobra.Command, args []string) {
		api.HttpRun()
	},
}

func init() {
	rootCmd.AddCommand(ginCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ginCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ginCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
