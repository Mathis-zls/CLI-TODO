/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"strconv"

	"github.com/Mathis-zls/CLI-TODO/utils"
	"github.com/spf13/cobra"
)

// DeleteTodoCmd represents the DeleteTodo command
var DeleteTodoCmd = &cobra.Command{
	Use:   "DeleteTodo",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if err := cobra.ExactArgs(1); err != nil {
			return errors.New("not one Arg")
		}
		if isValidIndex(args[0]) {
			return errors.New("not a Valid ID")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		id, _ := strconv.Atoi(args[0])
		utils.DeleteTodo(id)
	},
}

func init() {
	rootCmd.AddCommand(DeleteTodoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// DeleteTodoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// DeleteTodoCmd.Flags().Int("index", 0, "Index that shoud be deleted")
}

func isValidIndex(index string) bool {

	indx, err := strconv.Atoi(index)
	if err != nil || indx < 1 {
		return false
	}
	if id, err := utils.GetMaxID(); err != nil || indx > id {
		return false
	}
	return true

}
