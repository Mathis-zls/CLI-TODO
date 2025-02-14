/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/Mathis-zls/CLI-TODO/utils"
	"github.com/spf13/cobra"
)

// UpdateTodoCmd represents the UpdateTodo command
var UpdateTodoCmd = &cobra.Command{
	Use:   "UpdateTodo",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		description, _ := cmd.Flags().GetString("description")
		dueTime, _ := cmd.Flags().GetString("time")
		id, _ := cmd.Flags().GetInt("id")

		if !isValidIndex(fmt.Sprintf("%d", id)) {
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
	AddTodoCmd.Flags().StringP("name", "n", "", "Name of your task")
	AddTodoCmd.Flags().StringP("description", "d", "", "A description of the task")
	AddTodoCmd.Flags().StringP("time", "t", "", "Deadline for the task")

	AddTodoCmd.Flags().IntP("id", "i", 0, "index that shoud be modified")
	AddTodoCmd.MarkFlagRequired("index")
	AddTodoCmd.MarkFlagsOneRequired("description", "time")
	rootCmd.AddCommand(UpdateTodoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// UpdateTodoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// UpdateTodoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
