/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"

	"github.com/Mathis-zls/CLI-TODO/utils"
	"github.com/spf13/cobra"
)

// MarkTodoAsCompleteCmd represents the MarkTodoAsComplete command
var MarkTodoAsCompleteCmd = &cobra.Command{
	Use:   "MarkTodoAsComplete",
	Short: "Mark a Todo as Completed",
	Long: `Mark one of your Todos as completed by providig a id of the todo.
	You can only pass valid id (>0 and existing)
	Usage: CLI-TODO MarkTodoAsCompleted -i 1`,
	Run: func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetInt("id")
		if !isValidIndex(id) {
			log.Fatal("Not a valid index")
		}
		err := utils.MarkTodoAsComplete(id)
		if err != nil {
			log.Fatal("Coudn't be marked as Completed:", err)
		}
	},
}

func init() {
	MarkTodoAsCompleteCmd.Flags().IntP("id", "i", 0, "the id from the task that shoud be completed")
	MarkTodoAsCompleteCmd.MarkFlagRequired("id")
	rootCmd.AddCommand(MarkTodoAsCompleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// MarkTodoAsCompleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// MarkTodoAsCompleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
