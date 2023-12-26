package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ThursdayCmd = &cobra.Command{
	Use:   "Thursday [文案]",
	Short: "crazy Thursday",
	Long:  `crazy Thursday`,
	Run: func(cmd *cobra.Command, args []string) {
		res := Fetch()
		Print(res)
		fmt.Println(res)
	},
}
