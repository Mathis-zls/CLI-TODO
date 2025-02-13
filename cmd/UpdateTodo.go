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

var id int

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

		if !isValidIndex(fmt.Sprintf("%d", id)) {
			log.Fatal("Not a valid Index")
		}
		utils.UpdateTodo(id, utils.TodoParams{
			Name:        name,
			Description: description,
			Due_time:    due_too,
		})
	},
}

func init() {
	rootCmd.AddCommand(UpdateTodoCmd)
	AddTodoCmd.Flags().StringVarP(&name, "name", "n", "", "name of you task")
	AddTodoCmd.Flags().StringVarP(&description, "description", "d", "", "a description off the task")
	AddTodoCmd.Flags().StringVarP(&due_too, "time", "t", "", "when shoud the task be completed")
	AddTodoCmd.Flags().IntVarP(&id, "index", "i", 0, "index that shoud be modified")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// UpdateTodoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// UpdateTodoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
