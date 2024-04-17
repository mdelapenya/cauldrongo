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

			url := cauldron.NewURL(2296, "2024-04-01", "2024-04-16", tt.tab)
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
