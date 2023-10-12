package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var yuzukoCmd = &cobra.Command{
	Use:   "yuzuko",
	Short: "Display Yuzuko's ASCII art",
	Long:  `Display Yuzuko's ASCII art`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		b, err := os.ReadFile("aa/yuzuko.txt")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(b))
	},
}

func init() {
	aaCmd.AddCommand(yuzukoCmd)
}
