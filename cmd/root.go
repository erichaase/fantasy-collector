package cmd

import (
	"fmt"
	"os"

	"github.com/erichaase/fantasy-collector/internal/espn"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "fantasy-collector",
	Short: "Collects and stores NBA stats from ESPN",
	Long:  `Collects and stores NBA stats from ESPN`,
	Run: func(cmd *cobra.Command, args []string) {
		lines, err := espn.GetGameLines()
		if err != nil {
			// TODO: add a logger
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		fmt.Println(lines)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// TODO: add --format flag: html, csv, etc.
	// TODO: add --dest flag: s3, stdout, etc.
}
