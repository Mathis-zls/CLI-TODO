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
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// NewCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// NewCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
