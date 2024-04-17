package cauldron_test

import (
	"encoding/json"
	"testing"

	"github.com/mdelapenya/cauldrongo/cauldron"
)

type testWriter struct {
	data []byte
}

func (w *testWriter) Write(p []byte) (n int, err error) {
	w.data = p
	return len(p), nil
}

func TestConsoleFormatter(t *testing.T) {
	consoleFormatter := &cauldron.ConsoleFormatter{}

	a := &cauldron.Activity{}
	a.CommitsActivityOverview = 15

	w := &testWriter{}
	err := consoleFormatter.Format(w, a)
	if err != nil {
		t.Fatalf("error formatting: %v", err)
	}

	formatted := string(w.data)

	expected := "&{CommitsActivityOverview:15 LinesCommitActivityOverview: LinesCommitFileActivityOverview: IssuesCreatedActivityOverview:0 IssuesClosedActivityOverview:0 IssuesOpenActivityOverview:0 ReviewsCreatedActivityOverview:0 ReviewsClosedActivityOverview:0 ReviewsOpenActivityOverview:0}"
	if formatted != expected {
		t.Fatalf("expected %s but got %s", expected, formatted)
	}
}

func TestJSONFormatter(t *testing.T) {
	jsonFormatter := &cauldron.JSONFormatter{
		Indent: "	",
	}

	a := &cauldron.Activity{}
	a.CommitsActivityOverview = 15

	w := &testWriter{}

	err := jsonFormatter.Format(w, a)
	if err != nil {
		t.Fatalf("error formatting: %v", err)
	}

	formatted := string(w.data)

	// the indent is 1 tab
	expected := `{
	"commits_activity_overview": 15,
	"lines_commit_activity_overview": "",
	"lines_commit_file_activity_overview": "",
	"issues_created_activity_overview": 0,
	"issues_closed_activity_overview": 0,
	"issues_open_activity_overview": 0,
	"reviews_created_activity_overview": 0,
	"reviews_closed_activity_overview": 0,
	"reviews_open_activity_overview": 0
}`
	if formatted != expected {
		t.Fatalf("expected %s but got %s", expected, formatted)
	}

	var a2 cauldron.Activity
	if err := json.Unmarshal([]byte(formatted), &a2); err != nil {
		t.Fatalf("error unmarshalling: %v", err)
	}

	if a.CommitsActivityOverview != a2.CommitsActivityOverview {
		t.Fatalf("expected CommitsActivityOverview=%d but got %d", a.CommitsActivityOverview, a2.CommitsActivityOverview)
	}
}
