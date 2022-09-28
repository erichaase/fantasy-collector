package cmd

import (
	"github.com/erichaase/fantasy-collector/internal/espn"
	"github.com/spf13/cobra"
)

// storeCmd represents the store command
var storeCmd = &cobra.Command{
	Use:   "store",
	Short: "Collects and stores NBA stats from ESPN",
	Long:  `Collects and stores NBA stats from ESPN`,
	Run: func(cmd *cobra.Command, args []string) {
		espn.GetGameLines()
	},
}

func init() {
	rootCmd.AddCommand(storeCmd)
	// TODO: add format flag: html, csv, etc.
}
