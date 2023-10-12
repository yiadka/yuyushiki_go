package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var yukariCmd = &cobra.Command{
	Use:   "yukari",
	Short: "Display Yukari's ASCII art",
	Long:  `Display Yukari's ASCII art`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		b, err := os.ReadFile("aa/yukari.txt")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(b))
	},
}

func init() {
	aaCmd.AddCommand(yukariCmd)
}
