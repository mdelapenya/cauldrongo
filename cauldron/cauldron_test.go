package cauldron_test

import (
	"context"
	"fmt"
	"net/http"
	"path/filepath"
	"testing"

	"github.com/mdelapenya/cauldrongo/cauldron"
	"github.com/testcontainers/testcontainers-go"
	wiremock "github.com/wiremock/wiremock-testcontainers-go"
)

func TestMockHTTPRequests(t *testing.T) {
	absPath, err := filepath.Abs("testdata")
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name       string
		tab        string
		processor  cauldron.Processor
		assertFunc func(t *testing.T, processor cauldron.Processor)
	}{
		{
			name:      "activity",
			tab:       "activity-overview",
			processor: &cauldron.Activity{},
			assertFunc: func(t *testing.T, processor cauldron.Processor) {
				a := processor.(*cauldron.Activity)

				if a.CommitsActivityOverview != 15 {
					t.Fatalf("expected CommitsActivityOverview=15 but got %d", a.CommitsActivityOverview)
				}
			},
		},
		{
			name:      "community",
			tab:       "community-overview",
			processor: &cauldron.Community{},
			assertFunc: func(t *testing.T, processor cauldron.Processor) {
				c := processor.(*cauldron.Community)

				if c.ActivePeopleGitCommunityOverview != 8 {
					t.Fatalf("expected ActivePeopleGitCommunityOverview=8 but got %d", c.ActivePeopleGitCommunityOverview)
				}
			},
		},
		{
			name:      "overview",
			tab:       "overview",
			processor: &cauldron.Overview{},
			assertFunc: func(t *testing.T, processor cauldron.Processor) {
				o := processor.(*cauldron.Overview)

				if o.CommitsOverview != 1581 {
					t.Fatalf("expected CommitsOverview=1581 but got %d", o.CommitsOverview)
				}
			},
		},
		{
			name:      "performance",
			tab:       "performance-overview",
			processor: &cauldron.Performance{},
			assertFunc: func(t *testing.T, processor cauldron.Processor) {
				p := processor.(*cauldron.Performance)

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

			err = tt.processor.Process(body)
			if err != nil {
				t.Fatal(err)
			}

			tt.assertFunc(t, tt.processor)
		})
	}
}
