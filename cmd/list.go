/*
Copyright © 2021 dryprojects

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
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"os"
	"sort"
	"text/tabwriter"
	"todo/model"
	"todo/utils"
)

var (
	doneOpt bool
	allOpt  bool
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List todos",
	Run: listRun,
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	listCmd.Flags().BoolVar(
		&doneOpt,
		"done",
		false,
		"show 'Done' Todos",
	)

	listCmd.Flags().BoolVar(
		&allOpt,
		"all",
		false,
		"Show all Todos",
	)
}

func listRun(cmd *cobra.Command, args []string) {
	items, err := model.ReadItems(viper.GetString("datafile"))
	if err != nil {
		log.Printf("%v", err)
	}

	sort.Sort(model.ByPri(items))
	w := tabwriter.NewWriter(os.Stdout, 3, 0, 1, ' ', 0)
	for _, i := range items {
		if allOpt || i.Done == doneOpt {
			out := i.Label() + "\t" + colorTodo(i, i.PrettyDone()) + "\t" + i.PrettyP() + "\t" + colorTodo(i, i.Text+"\t")
			fmt.Fprintln(w, out)
		}
	}
	w.Flush()
}

func colorTodo(item model.Item, text string) string {
	if item.Done {
		return utils.TextGreen(text)
	} else if item.Priority == 1 {
		return utils.TextBlue(text)
	} else if item.Priority == 2 {
		return utils.TextYellow(text)
	} else if item.Priority == 3 {
		return utils.TextRed(text)
	} else {
		return text
	}
}
