package cmd

import (
	"embed"
	"fmt"

	"github.com/spf13/cobra"
)

//go:embed *.txt
var files embed.FS

var yuiCmd = &cobra.Command{
	Use:   "yukari",
	Short: "Display Yui's ASCII art",
	Long:  `Display Yui's ASCII art`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		yui, _ := files.ReadFile("aa/yui.txt")
		fmt.Println(yui)
	},
}

func init() {
	aaCmd.AddCommand(yuiCmd)
}
