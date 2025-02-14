/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"

	"github.com/Mathis-zls/CLI-TODO/utils"
	"github.com/spf13/cobra"
)

// NewTodoListCmd represents the New command
var NewTodoListCmd = &cobra.Command{
	Use:   "NewTodoList",
	Short: "Creates a new TodoList,with a given name",
	Long: `Creates a new TodoList. You need to provied a name for the List. This TodoList is saved as a json and can be loaded with LoadTodoList
	Usage: CLI-TODO NewTodoList -n "Homework"`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		err := utils.NewSave(name)
		if err != nil {
			log.Fatal("Error while creating new TodoList:", err)
		}

	},
}

func init() {
	NewTodoListCmd.Flags().StringP("name", "n", "", "Name of your new Todolist")
	NewTodoListCmd.MarkFlagRequired("name")
	rootCmd.AddCommand(NewTodoListCmd)
}
