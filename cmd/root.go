/*
Copyright © 2024 Mattia Pun <mattia@punjwani.be>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "PicPorto",
	Short: "Easily create a portfolio  with your photos.",
	Long:  "PicPorto is a CLI tool that allows you to easily create a portfolio with your photos.",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}

}

func init() {
	rootCmd.AddCommand(createCmd)
}
