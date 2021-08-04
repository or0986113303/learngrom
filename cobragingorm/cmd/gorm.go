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
	"cobragingorm/internal/pkg/control"
	"cobragingorm/internal/pkg/model"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/tidwall/gjson"
	"github.com/tidwall/pretty"
)

// gormCmd represents the gorm command
var (
	Name    string
	Number  uint
	user    model.User
	log     = logrus.New()
	dbh     *control.ORMEngine
	gormCmd = &cobra.Command{
		Use:   "gorm",
		Short: "A brief description of your command",
		Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("gorm entry point test")
		},
	}
	gormSelectLastCmd = &cobra.Command{
		Use:   "gormselectlast",
		Short: "Select last one by condition",
		Long: `Select the last one from db by input 
		condition at parameters`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(Name) > 0 && Number <= 0 {
				res := dbh.SelectConditionLast("Name", Name)
				log.Info(res)
			} else if Number > 0 && len(Name) == 0 {
				res := dbh.SelectConditionLast("Number", Number)
				log.Info(res)
			} else {
				log.Info("Wrong parameters")
			}
		},
	}
	gormSelectFirstCmd = &cobra.Command{
		Use:   "gormselectfirst",
		Short: "Select first one by condition",
		Long: `Select the first one from db by input 
		condition at parameters`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(Name) > 0 && Number <= 0 {
				res := dbh.SelectConditionFirst("Name", Name)
				log.Info(res)
			} else if Number > 0 && len(Name) == 0 {
				res := dbh.SelectConditionFirst("Number", Number)
				log.Info(res)
			} else {
				log.Info("Wrong parameters")
			}
		},
	}
	gormInsertCmd = &cobra.Command{
		Use:   "gorminsert",
		Short: "Insert",
		Long:  `Insert simple data by input parameters which are name, number`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(Name) > 0 && Number > 0 {
				src := model.User{Name: Name, Number: Number}
				res := dbh.Insert(src)
				log.Info(res)
			} else {
				for _, value := range args {
					reducejson := pretty.Pretty([]byte(value))
					if len(reducejson) > 3 {
						src := model.User{Name: gjson.Get(value, "Name").String(), Number: uint(gjson.Get(value, "Number").Uint())}
						res := dbh.Insert(src)
						log.Info(res)
					} else {
						log.Info("Wrong parameters")
					}
				}
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(gormCmd)
	rootCmd.AddCommand(gormSelectLastCmd)
	rootCmd.AddCommand(gormSelectFirstCmd)
	rootCmd.AddCommand(gormInsertCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// gormCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// gormCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	log.Out = os.Stdout
	control.NewEngine()
	dbh = control.GetEngine()

	gormSelectLastCmd.Flags().StringVarP(&Name, "name", "N", "", "select by name")
	gormSelectLastCmd.Flags().UintVarP(&Number, "number", "n", 0, "value of number")
	gormSelectLastCmd.MarkFlagRequired("name")

	gormSelectFirstCmd.Flags().StringVarP(&Name, "name", "N", "", "select by name")
	gormSelectFirstCmd.Flags().UintVarP(&Number, "number", "n", 0, "value of number")
	gormSelectFirstCmd.MarkFlagRequired("name")

	gormInsertCmd.Flags().StringVarP(&Name, "name", "N", "", "value of name")
	gormInsertCmd.Flags().UintVarP(&Number, "number", "n", 0, "value of number")
}
