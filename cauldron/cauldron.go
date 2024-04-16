package cauldron

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const (
	baseScheme               = "https"
	baseURL                  = "cauldron.io"
	metricsURLFormat         = "/project/%d/metrics"
	metricsQueryStringFormat = "from=%s&to=%s&tab=%s"
)

func NewURL(projectID int, from, to, tab string) url.URL {
	return url.URL{
		Scheme:   baseScheme,
		Host:     baseURL,
		Path:     fmt.Sprintf(metricsURLFormat, projectID),
		RawQuery: fmt.Sprintf(metricsQueryStringFormat, from, to, tab),
	}
}

func HttpRequest(url *url.URL) (io.Reader, error) {
	httpCli := http.Client{}

	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("error creating HTTP request: %v", err)
	}

	req.Header.Add("Authority", "cauldron.io")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("User-Agent", "CauldronGo")

	resp, err := httpCli.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making HTTP request: %v", err)
	}
	// we are intentionally not closing the body here, as it will be read by the caller

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return resp.Body, nil
}
