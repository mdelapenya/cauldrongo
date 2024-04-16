package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var projectID int

func init() {
	cmdMetrics.Flags().IntVarP(&projectID, "project-id", "p", 0, "The project ID to fetch metrics")

	rootCmd.AddCommand(cmdMetrics)
}

var cmdMetrics = &cobra.Command{
	Use:   "metrics",
	Short: "Fetch metrics for a given project",
	Long: `Fetch metrics for a given project. It will return the metrics for the
				  project in the format requested.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Print: %d\n", projectID)
	},
}
