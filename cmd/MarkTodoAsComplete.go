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

// MarkTodoAsCompleteCmd represents the MarkTodoAsComplete command
var MarkTodoAsCompleteCmd = &cobra.Command{
	Use:   "MarkTodoAsComplete",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetInt("id")
		if !isValidIndex(fmt.Sprintf("%d", id)) {
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
