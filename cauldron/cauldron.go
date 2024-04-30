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

func NewURL(projectID int, from, to, tab string, repoURLs []string) url.URL {
	rawQuery := fmt.Sprintf(metricsQueryStringFormat, from, to, tab)
	if len(repoURLs) > 0 {
		queryRepos := ""
		for _, repoURL := range repoURLs {
			queryRepos += "&" + url.QueryEscape("repo_url[]") + "=" + url.QueryEscape(repoURL)
		}

		rawQuery = rawQuery + queryRepos
	}

	return url.URL{
		Scheme:   baseScheme,
		Host:     baseURL,
		Path:     fmt.Sprintf(metricsURLFormat, projectID),
		RawQuery: rawQuery,
	}
}

func HttpRequest(url url.URL) (io.ReadCloser, int, error) {
	httpCli := http.Client{}

	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("error creating HTTP request: %v", err)
	}

	req.Header.Add("Authority", "cauldron.io")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("User-Agent", "CauldronGo")

	resp, err := httpCli.Do(req)
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("error making HTTP request: %v", err)
	}
	// we are intentionally not closing the body here, as it will be read by the caller

	return resp.Body, resp.StatusCode, nil
}
