/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"

	"github.com/Mathis-zls/CLI-TODO/utils"
	"github.com/spf13/cobra"
)

// AddTodoCmd represents the AddTodo command
var AddTodoCmd = &cobra.Command{
	Use:   "AddTodo",
	Short: "Add a new Todo to your TodoList",
	Long: `This command is for adding new Todos. 
	Requierement you already loaded/created a TodoList with load/New.
	Usage CLI-TODO AdddTodo -n "Einkaufen" -d "Einkaufen bei Rewe" -t "Dienstag"
	You only need to specify the name -n`,
	Run: func(cmd *cobra.Command, args []string) {

		name, _ := cmd.Flags().GetString("name")
		description, _ := cmd.Flags().GetString("description")
		dueTime, _ := cmd.Flags().GetString("time")

		err := utils.AddTodo(utils.TodoParams{
			Name:        name,
			Description: description,
			Due_time:    dueTime,
		})
		if err != nil {
			log.Fatal("A Problem with adding the Todo", err)
		}
	},
}

func init() {
	AddTodoCmd.Flags().StringP("name", "n", "", "Name of your task")
	AddTodoCmd.Flags().StringP("description", "d", "", "A description of the task")
	AddTodoCmd.Flags().StringP("time", "t", "", "Deadline for the task")
	AddTodoCmd.MarkFlagRequired("name")
	rootCmd.AddCommand(AddTodoCmd)
}
