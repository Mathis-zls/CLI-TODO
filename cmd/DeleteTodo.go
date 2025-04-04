/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"

	"github.com/Mathis-zls/CLI-TODO/utils"
	"github.com/spf13/cobra"
)

// DeleteTodoCmd represents the DeleteTodo command
var DeleteTodoCmd = &cobra.Command{
	Use:   "DeleteTodo",
	Short: "Delete an existing Todo by id",
	Long: `Delete a Todo.You neeed to provide the id of the Todo
	Only valid ids can be used
	Usage: CLI-TODO DeleteTodo -i`,
	Run: func(cmd *cobra.Command, args []string) {

		id, _ := cmd.Flags().GetInt("id")
		if !isValidIndex(id) {
			log.Fatal("Error not a valid id:")
		}
		err := utils.DeleteTodo(id)
		if err != nil {
			log.Fatal("Error while deleting Todo:", err)
		}
	},
}

func init() {
	DeleteTodoCmd.Flags().IntP("id", "i", 0, "id of the Todo")
	DeleteTodoCmd.MarkFlagRequired("id")
	rootCmd.AddCommand(DeleteTodoCmd)

}

func isValidIndex(index int) bool {
	if id, err := utils.GetMaxID(); err != nil || index > id || index < 1 {
		return false
	}
	return true
}
