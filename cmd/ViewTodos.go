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
var ViewTodosCmd = &cobra.Command{
	Use:   "ViewTodos",
	Short: "Show's the List od Todos",
	Long: `Shows the list of Todos you have and how much are not completed
	Usage: CLI-TODO ViewTodos`,
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
	rootCmd.AddCommand(ViewTodosCmd)
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
		done := "❌"
		if row.Done {
			done = "✅"
			totalCompleted++
		}
		r := []*simpletable.Cell{
			{Align: simpletable.AlignRight, Text: done},
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
