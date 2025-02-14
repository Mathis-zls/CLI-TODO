/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"

	"github.com/Mathis-zls/CLI-TODO/utils"
	"github.com/spf13/cobra"
)

// UpdateTodoCmd represents the UpdateTodo command
var UpdateTodoCmd = &cobra.Command{
	Use:   "UpdateTodo",
	Short: "Update a existing Todo",
	Long: `A command for Updating a Todo. 
	Requierment a Todo is exisiting with a id
	You can Update name,description,time
	Use: CLI-TODO UpdateTodo -i 2 -n "Eat" -d "Go out to eat..." -t "Monday"
	You can update 1 ore more param (n,d,t) but you need to provide an id`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		description, _ := cmd.Flags().GetString("description")
		dueTime, _ := cmd.Flags().GetString("time")
		id, _ := cmd.Flags().GetInt("id")

		if !isValidIndex(id) {
			log.Fatal("Not a valid Index")
		}
		err := utils.UpdateTodo(id, utils.TodoParams{
			Name:        name,
			Description: description,
			Due_time:    dueTime,
		})
		if err != nil {
			log.Fatal("Error while Updating Todo:", err)
		}
	},
}

func init() {
	UpdateTodoCmd.Flags().StringP("name", "n", "", "Name of your task")
	UpdateTodoCmd.Flags().StringP("description", "d", "", "A description of the task")
	UpdateTodoCmd.Flags().StringP("time", "t", "", "Deadline for the task")

	UpdateTodoCmd.Flags().IntP("id", "i", 0, "index that shoud be modified")
	UpdateTodoCmd.MarkFlagRequired("index")
	UpdateTodoCmd.MarkFlagsOneRequired("description", "time", "name")
	rootCmd.AddCommand(UpdateTodoCmd)

}
