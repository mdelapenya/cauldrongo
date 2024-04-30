package cauldron_test

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"path/filepath"
	"testing"

	"github.com/mdelapenya/cauldrongo/cauldron"
	"github.com/testcontainers/testcontainers-go"
	wiremock "github.com/wiremock/wiremock-testcontainers-go"
)

func TestMockHTTPRequests(t *testing.T) {
	absPath, err := filepath.Abs(filepath.Join("..", "testdata"))
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name       string
		tab        string
		printable  cauldron.Printable
		assertFunc func(t *testing.T, printable cauldron.Printable)
	}{
		{
			name:      "activity",
			tab:       "activity-overview",
			printable: &cauldron.Activity{},
			assertFunc: func(t *testing.T, printable cauldron.Printable) {
				a := printable.(*cauldron.Activity)

				if a.CommitsActivityOverview != 15 {
					t.Fatalf("expected CommitsActivityOverview=15 but got %d", a.CommitsActivityOverview)
				}
			},
		},
		{
			name:      "community",
			tab:       "community-overview",
			printable: &cauldron.Community{},
			assertFunc: func(t *testing.T, printable cauldron.Printable) {
				c := printable.(*cauldron.Community)

				if c.ActivePeopleGitCommunityOverview != 8 {
					t.Fatalf("expected ActivePeopleGitCommunityOverview=8 but got %d", c.ActivePeopleGitCommunityOverview)
				}
			},
		},
		{
			name:      "overview",
			tab:       "overview",
			printable: &cauldron.Overview{},
			assertFunc: func(t *testing.T, printable cauldron.Printable) {
				o := printable.(*cauldron.Overview)

				if o.CommitsOverview != 1581 {
					t.Fatalf("expected CommitsOverview=1581 but got %d", o.CommitsOverview)
				}
			},
		},
		{
			name:      "performance",
			tab:       "performance-overview",
			printable: &cauldron.Performance{},
			assertFunc: func(t *testing.T, printable cauldron.Printable) {
				p := printable.(*cauldron.Performance)

				if p.IssuesTimeOpenAveragePerformanceOverview != 272.41 {
					t.Fatalf("expected IssuesTimeOpenAveragePerformanceOverview=272.41 but got %f", p.IssuesTimeOpenAveragePerformanceOverview)
				}
			},
		},
	}

	opts := []testcontainers.ContainerCustomizer{}
	for _, tt := range tests {
		opts = append(opts, wiremock.WithMappingFile(tt.name, filepath.Join(absPath, tt.name+".json")))
	}

	ctx := context.Background()
	container, err := wiremock.RunContainerAndStopOnCleanup(ctx, t, opts...)
	if err != nil {
		t.Fatal(err)
	}

	host, err := container.Host(ctx)
	if err != nil {
		t.Fatal(err)
	}

	port, err := container.MappedPort(ctx, "8080")
	if err != nil {
		t.Fatal(err)
	}

	baseURL := fmt.Sprintf("%s:%s", host, port.Port())

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(innerT *testing.T) {
			innerT.Parallel()

			// setting repo_urls to an empty slice on purpose,
			// as the test data does not include any repoURLs
			repoURLs := []string{}

			url := cauldron.NewURL(2296, "2024-04-01", "2024-04-16", tt.tab, repoURLs)
			url.Scheme = "http"
			url.Host = baseURL

			body, statusCode, err := cauldron.HttpRequest(url)
			if err != nil {
				innerT.Fatal(err, "Failed to get a response")
			}

			if statusCode != http.StatusOK {
				innerT.Fatalf("expected HTTP-200 but got %d", statusCode)
			}

			bs, err := io.ReadAll(body)
			if err != nil {
				innerT.Fatalf("Failed to read the response %v", err)
			}

			if err := json.Unmarshal(bs, tt.printable); err != nil {
				innerT.Fatalf("Failed to unmarshal the response %v", err)
			}

			tt.assertFunc(t, tt.printable)
		})
	}
}

func TestNewURL(t *testing.T) {
	testCases := []struct {
		name      string
		projectID int
		from      string
		to        string
		tab       string
		repoURLs  []string
		expected  string
	}{
		{
			name:      "no-repo-urls",
			projectID: 2296,
			from:      "2024-04-01",
			to:        "2024-04-16",
			tab:       "activity-overview",
			repoURLs:  []string{},
			expected:  "https://cauldron.io/project/2296/metrics?from=2024-04-01&to=2024-04-16&tab=activity-overview",
		},
		{
			name:      "with-repo-url",
			projectID: 2296,
			from:      "2024-04-01",
			to:        "2024-04-16",
			tab:       "activity-overview",
			repoURLs: []string{
				"https://github.com/testcontainers/testcontainers-go",
			},
			expected: "https://cauldron.io/project/2296/metrics?from=2024-04-01&to=2024-04-16&tab=activity-overview&repo_url%5B%5D=https%3A%2F%2Fgithub.com%2Ftestcontainers%2Ftestcontainers-go",
		},
		{
			name:      "with-repo-urls",
			projectID: 2296,
			from:      "2024-04-01",
			to:        "2024-04-16",
			tab:       "activity-overview",
			repoURLs: []string{
				"https://github.com/testcontainers/testcontainers-go",
				"https://github.com/testcontainers/testcontainers-go.git",
			},
			expected: "https://cauldron.io/project/2296/metrics?from=2024-04-01&to=2024-04-16&tab=activity-overview&repo_url%5B%5D=https%3A%2F%2Fgithub.com%2Ftestcontainers%2Ftestcontainers-go&repo_url%5B%5D=https%3A%2F%2Fgithub.com%2Ftestcontainers%2Ftestcontainers-go.git",
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(tt *testing.T) {
			tt.Parallel()

			url := cauldron.NewURL(testCase.projectID, testCase.from, testCase.to, testCase.tab, testCase.repoURLs)
			if url.String() != testCase.expected {
				tt.Fatalf("expected %s but got %s", testCase.expected, url.String())
			}
		})
	}
}
