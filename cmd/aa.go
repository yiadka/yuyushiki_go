package cmd

import (
	"errors"

	"github.com/spf13/cobra"
)

var aaCmd = &cobra.Command{
	Use:   "aa",
	Short: "Display ASCII art",
	Long:  `Display ASCII art according to the command entered`,
	RunE:  func(cmd *cobra.Command, args []string) error { return errors.New("not found command") },
}

func init() {
	rootCmd.AddCommand(aaCmd)
}
