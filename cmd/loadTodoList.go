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
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loadCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loadCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
