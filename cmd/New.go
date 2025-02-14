/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"

	"github.com/Mathis-zls/CLI-TODO/utils"
	"github.com/spf13/cobra"
)

// NewCmd represents the New command
var NewCmd = &cobra.Command{
	Use:   "New",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		err := utils.NewSave(name)
		if err != nil {
			log.Fatal("Error while creating new TodoList:", err)
		}

	},
}

func init() {
	NewCmd.Flags().StringP("name", "n", "", "Name of your new Todolist")
	NewCmd.MarkFlagRequired("name")
	rootCmd.AddCommand(NewCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// NewCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// NewCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
