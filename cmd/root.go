package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use: "ct",
}

func Execute() {
	rootCmd.AddCommand(ThursdayCmd)
	rootCmd.SetHelpCommand(ThursdayCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
