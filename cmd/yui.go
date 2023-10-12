package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var yuiCmd = &cobra.Command{
	Use:   "yukari",
	Short: "Display Yui's ASCII art",
	Long:  `Display Yui's ASCII art`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		b, err := os.ReadFile("aa/yui.txt")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(b))
	},
}

func init() {
	aaCmd.AddCommand(yuiCmd)
}
