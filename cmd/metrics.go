package cmd

import (
	"fmt"
	"io"
	"net/url"
	"os"
	"time"

	"github.com/spf13/cobra"
	"golang.org/x/sync/errgroup"

	"github.com/mdelapenya/cauldrongo/cauldron"
)

var projectID int
var from string
var to string
var tab string

func init() {
	now := time.Now()
	yearAgo := now.AddDate(-1, 0, 0)

	// format now in the format YYYY-MM-DD
	formattedNow := now.Format("2006-01-02")
	formattedYearAgo := yearAgo.Format("2006-01-02")

	cmdMetrics.Flags().IntVarP(&projectID, "project-id", "p", 0, "The project ID to fetch metrics")
	cmdMetrics.Flags().StringVarP(&from, "from", "f", formattedYearAgo, "The start date to fetch metrics")
	cmdMetrics.Flags().StringVarP(&to, "to", "t", formattedNow, "The end date to fetch metrics")
	cmdMetrics.Flags().StringVarP(&tab, "tab", "b", "", "The tab to fetch metrics. Possible values are: overview, activity-overview, community-overview, performance-overview. If empty, it will fetch the overview tab.")

	rootCmd.AddCommand(cmdMetrics)
}

var cmdMetrics = &cobra.Command{
	Use:   "metrics",
	Short: "Fetch metrics for a given project",
	Long: `Fetch metrics for a given project. It will return the metrics for the
				  project in the requested format.`,
	Run: func(cmd *cobra.Command, args []string) {
		overviewURL := cauldron.NewURL(projectID, from, to, tab)
		urls := []url.URL{overviewURL}
		if tab == "" {
			urls = append(urls, cauldron.NewURL(projectID, from, to, "activity-overview"))
			urls = append(urls, cauldron.NewURL(projectID, from, to, "community-overview"))
			urls = append(urls, cauldron.NewURL(projectID, from, to, "erformance-overview"))
		}

		// execute all requests concurrently, waiting for the last one to finish, capturing errors
		// and printing them

		responses := make(chan io.Reader, len(urls))

		errorGroup := errgroup.Group{}
		for _, u := range urls {
			u := u
			errorGroup.Go(func() error {
				reader, err := cauldron.HttpRequest(&u)
				if err != nil {
					fmt.Printf("Error fetching metrics: %v. URL: %s\n", err, u.String())
					return err
				}

				responses <- reader
				return nil
			})
		}

		if err := errorGroup.Wait(); err != nil {
			fmt.Println("Error fetching metrics:", err)
			os.Exit(1)
		}

		// process all responses
		for i := 0; i < len(urls); i++ {
			reader := <-responses

			var processor cauldron.Processor
			u := urls[i]
			switch u.Query().Get("tab") {
			case "activity-overview":
				processor = &cauldron.Activity{}
			case "community-overview":
				processor = &cauldron.Community{}
			case "performance-overview":
				processor = &cauldron.Performance{}
			default:
				processor = &cauldron.Overview{}
			}

			err := processor.Process(reader)
			if err != nil {
				fmt.Printf("Error processing metrics: %v\n", err)
				os.Exit(1)
			}

			fmt.Printf("Metrics processed successfully: %+v\n", processor)
		}
	},
}
