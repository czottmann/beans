package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "beans",
	Short: "Beans is a sample CLI application",
	Long:  `Beans is a sample CLI application to demonstrate cobra.`,
}

func Execute() error {
	return rootCmd.Execute()
}
