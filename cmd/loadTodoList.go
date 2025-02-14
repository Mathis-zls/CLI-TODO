/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"

	"github.com/Mathis-zls/CLI-TODO/utils"
	"github.com/spf13/cobra"
)

// LoadTodoListCmd represents the load command
var LoadTodoListCmd = &cobra.Command{
	Use:   "LoadTodoList",
	Short: "Load a existing todolist by name",
	Long: `Load an existing todolist while providing the name of the list.
	It will return an error if the name is wrong
	Usage: CLI-TODO LoadTodoList -n "Shopping"`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")

		err := utils.LoadSave(name)
		if err != nil {
			log.Fatal("Error while loading file:", err)
		}
	},
}

func init() {
	LoadTodoListCmd.Flags().StringP("name", "n", "", "the name of the file that should be loaded")
	LoadTodoListCmd.MarkFlagRequired("name")
	rootCmd.AddCommand(LoadTodoListCmd)
}
