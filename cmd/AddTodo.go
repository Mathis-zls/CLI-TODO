/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"

	"github.com/Mathis-zls/CLI-TODO/utils"
	"github.com/spf13/cobra"
)

var (
	name        string
	description string
	due_too     string
)

// AddTodoCmd represents the AddTodo command
var AddTodoCmd = &cobra.Command{
	Use:   "AddTodo",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := utils.AddTodo(utils.TodoParams{
			Name:        name,
			Description: description,
			Due_time:    due_too,
		})
		if err != nil {
			log.Fatal("A Problem with adding the Todo")
		}
	},
}

func init() {
	rootCmd.AddCommand(AddTodoCmd)

	AddTodoCmd.Flags().StringVarP(&name, "name", "n", "", "name of you task")
	AddTodoCmd.Flags().StringVarP(&description, "description", "d", "", "a description off the task")
	AddTodoCmd.Flags().StringVarP(&due_too, "time", "t", "", "when shoud the task be completed")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// AddTodoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// AddTodoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
