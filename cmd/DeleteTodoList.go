/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"

	"github.com/Mathis-zls/CLI-TODO/utils"
	"github.com/spf13/cobra"
)

// DeleteTodoListCmd represents the DeleteTodoList command
var DeleteTodoListCmd = &cobra.Command{
	Use:   "DeleteTodoList",
	Short: "Delete a Todolist",
	Long: `Delete a Todolist. You need to provide a name (-n)
	Usage: CLI-TODO DeleteTodoList -n "exampleList"`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		err := utils.DeleteSave(name)
		if err != nil {
			log.Fatal("Couldn't delete Todolist:", err)
		}
	},
}

func init() {
	DeleteTodoListCmd.Flags().StringP("name", "n", "", "name of the list")
	DeleteTodoListCmd.MarkFlagRequired("name")
	rootCmd.AddCommand(DeleteTodoListCmd)
}
