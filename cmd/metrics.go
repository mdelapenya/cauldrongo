package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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
var format string

func init() {
	now := time.Now()
	yearAgo := now.AddDate(-1, 0, 0)

	// format now in the format YYYY-MM-DD
	formattedNow := now.Format("2006-01-02")
	formattedYearAgo := yearAgo.Format("2006-01-02")

	cmdMetrics.Flags().IntVarP(&projectID, "project-id", "p", 0, "The project ID to fetch metrics. Required.")
	cmdMetrics.Flags().StringVarP(&from, "from", "f", formattedYearAgo, "The start date to fetch metrics. Default is one year ago.")
	cmdMetrics.Flags().StringVarP(&to, "to", "t", formattedNow, "The end date to fetch metrics. Default is today.")
	cmdMetrics.Flags().StringVarP(&tab, "tab", "T", "", "The tab to fetch metrics. Possible values are: overview, activity-overview, community-overview, performance-overview. Default is overview.")
	cmdMetrics.Flags().StringVarP(&format, "format", "F", "console", "The format to output the metrics. Possible values are: console and json. Default is console.")

	rootCmd.AddCommand(cmdMetrics)
}

var cmdMetrics = &cobra.Command{
	Use:   "metrics",
	Short: "Fetch metrics for a given project",
	Long: `Fetch metrics for a given project. It will return the metrics for the
				  project in the requested format.`,
	Run: func(cmd *cobra.Command, args []string) {
		projectIDs := []int{projectID}
		if len(projects) > 0 {
			// if the configuration file contains projects, we will ignore the projectID flag
			projectIDs = []int{}
			for _, p := range projects {
				projectIDs = append(projectIDs, p.ID)
			}
		}

		var formatter cauldron.Formatter
		switch format {
		case "json":
			formatter = &cauldron.JSONFormatter{
				From:   from,
				To:     to,
				Writer: os.Stdout,
			}
		default:
			formatter = &cauldron.ConsoleFormatter{
				From:   from,
				To:     to,
				Writer: os.Stdout,
			}
		}

		formatter.FormatHeader()

		if err := metricsRun(projectIDs, formatter, from, to, tab); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func metricsRun(projectIDs []int, formatter cauldron.Formatter, from string, to string, tab string) error {
	formatter.FormatHeader()

	for _, projectID := range projectIDs {
		overviewURL := cauldron.NewURL(projectID, from, to, tab)
		urls := []url.URL{overviewURL}
		if tab == "" {
			urls = append(urls, cauldron.NewURL(projectID, from, to, "activity-overview"))
			urls = append(urls, cauldron.NewURL(projectID, from, to, "community-overview"))
			urls = append(urls, cauldron.NewURL(projectID, from, to, "erformance-overview"))
		}

		// execute all requests concurrently, waiting for the last one to finish, capturing errors
		// and printing them

		responses := make(chan io.ReadCloser, len(urls))

		errorGroup := errgroup.Group{}
		for _, u := range urls {
			u := u
			errorGroup.Go(func() error {
				reader, code, err := cauldron.HttpRequest(u)
				if err != nil {
					return fmt.Errorf("error fetching metrics: %w. URL: %s", err, u.String())
				}

				if code != http.StatusOK {
					return fmt.Errorf("error fetching metrics: HTTP status code %d. URL: %s", code, u.String())
				}

				responses <- reader
				return nil
			})
		}

		if err := errorGroup.Wait(); err != nil {
			return err
		}

		// process all responses
		for i := 0; i < len(urls); i++ {
			reader := <-responses
			defer reader.Close()

			var printable cauldron.Printable
			u := urls[i]
			switch u.Query().Get("tab") {
			case "activity-overview":
				printable = &cauldron.Activity{}
			case "community-overview":
				printable = &cauldron.Community{}
			case "performance-overview":
				printable = &cauldron.Performance{}
			default:
				printable = &cauldron.Overview{}
			}

			bs, err := io.ReadAll(reader)
			if err != nil {
				return fmt.Errorf("error reading metrics: %w", err)
			}

			if err := json.Unmarshal(bs, printable); err != nil {
				return fmt.Errorf("error unmarshalling metrics: %w", err)
			}

			if err := formatter.Format(printable); err != nil {
				return fmt.Errorf("error formatting metrics: %w", err)
			}
		}
	}

	return nil
}
