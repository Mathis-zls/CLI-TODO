/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/Mathis-zls/CLI-TODO/utils"
	"github.com/alexeyco/simpletable"
	"github.com/spf13/cobra"
)

// ViewTodosCmd represents the ViewTodos command
var view = &cobra.Command{
	Use:   "ViewTodos",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		todos, err := utils.GetTodos()
		if err != nil {
			log.Fatal("Coudn't load the Todos:", err)
			return
		}
		todoView(todos)
	},
}

func init() {
	rootCmd.AddCommand(view)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ViewTodosCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ViewTodosCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func todoView(todos []utils.Todo) {
	table := simpletable.New()
	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "Completed"},
			{Align: simpletable.AlignCenter, Text: "#"},
			{Align: simpletable.AlignCenter, Text: "Todo"},
			{Align: simpletable.AlignCenter, Text: "Description"},
			{Align: simpletable.AlignCenter, Text: "Due too"},
		},
	}
	totalCompleted := 0
	for _, row := range todos {
		done := '❌'
		if row.Done {
			done = '✅'
			totalCompleted++
		}
		r := []*simpletable.Cell{
			{Align: simpletable.AlignRight, Text: fmt.Sprintf("%v", done)},
			{Align: simpletable.AlignRight, Text: fmt.Sprintf("%v", row.Id)},
			{Align: simpletable.AlignRight, Text: fmt.Sprintf("%v", row.Name)},
			{Align: simpletable.AlignRight, Text: fmt.Sprintf("%v", row.Description)},
			{Align: simpletable.AlignRight, Text: fmt.Sprintf("%v", row.Due_time)},
		}
		table.Body.Cells = append(table.Body.Cells, r)
	}

	table.Footer = &simpletable.Footer{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: fmt.Sprintf("Total Task's completed: %d", totalCompleted), Span: 5},
		},
	}

	fmt.Println(table.String())

}
