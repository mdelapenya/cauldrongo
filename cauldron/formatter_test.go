package cauldron_test

import (
	"testing"

	"github.com/mdelapenya/cauldrongo/cauldron"
	"github.com/mdelapenya/cauldrongo/project"
)

var testProject = project.Project{ID: 1, Name: "Test Project", RepoURL: []string{"http://example.com/repo", "http://example.com/repo.git"}}

type testWriter struct {
	data []byte
}

func (w *testWriter) Write(p []byte) (n int, err error) {
	w.data = append(w.data, p...)
	return len(p), nil
}

func TestConsoleFormatter(t *testing.T) {
	tests := []struct {
		name      string
		printable cauldron.Printable
		writer    *testWriter
		expected  string
	}{
		{
			name:      "activity",
			printable: &cauldron.Activity{CommitsActivityOverview: 15},
			expected: `+-------------------------------------+-------+
|        METRIC (TEST PROJECT)        | VALUE |
+-------------------------------------+-------+
| Commits Activity Overview           |    15 |
| Lines Commit Activity Overview      |       |
| Lines Commit File Activity Overview |       |
| Issues Created Activity Overview    |     0 |
| Issues Closed Activity Overview     |     0 |
| Issues Open Activity Overview       |     0 |
| Reviews Created Activity Overview   |     0 |
| Reviews Closed Activity Overview    |     0 |
| Reviews Open Activity Overview      |     0 |
+-------------------------------------+-------+
`,
			writer: &testWriter{},
		},
		{
			name:      "community",
			printable: &cauldron.Community{ActivePeopleGitCommunityOverview: 87},
			expected: `+------------------------------------------+-------+
|          METRIC (TEST PROJECT)           | VALUE |
+------------------------------------------+-------+
| Active People Git Community Overview     |    87 |
| Active People Issues Community Overview  |     0 |
| Active People Patches Community Overview |     0 |
| Onboardings Git Community Overview       |     0 |
| Onboardings Issues Community Overview    |     0 |
| Onboardings Patches Community Overview   |     0 |
+------------------------------------------+-------+
`,
			writer: &testWriter{},
		},
		{
			name:      "overview",
			printable: &cauldron.Overview{CommitsOverview: 1587},
			expected: `+-------------------------------------------------+-------+
|              METRIC (TEST PROJECT)              | VALUE |
+-------------------------------------------------+-------+
| Commits Overview                                |  1587 |
| Issues Overview                                 |     0 |
| Reviews Overview                                |     0 |
| Commits Last Year Overview                      |     0 |
| Issues Last Year Overview                       |     0 |
| Reviews Last Year Overview                      |     0 |
| Commits YoY Overview                            |  0.00 |
| Issues YoY Overview                             |  0.00 |
| Reviews YoY Overview                            |  0.00 |
| Commit Authors Overview                         |     0 |
| Issue Submitters Overview                       |     0 |
| Review Submitters Overview                      |     0 |
| Commit Authors Last Year Overview               |     0 |
| Issue Submitters Last Year Overview             |     0 |
| Review Submitters Last Year Overview            |     0 |
| Commit Authors YoY Overview                     |  0.00 |
| Issue Submitters YoY Overview                   |  0.00 |
| Review Submitters YoY Overview                  |  0.00 |
| Issues Median Time To Close Overview            |  0.00 |
| Reviews Median Time To Close Overview           |  0.00 |
| Issues Median Time To Close Last Year Overview  |  0.00 |
| Reviews Median Time To Close Last Year Overview |  0.00 |
| Issues Median Time To Close YoY Overview        |  0.00 |
| Reviews Median Time To Close YoY Overview       |  0.00 |
+-------------------------------------------------+-------+
`,
			writer: &testWriter{},
		},
		{
			name:      "performance",
			printable: &cauldron.Performance{IssuesTimeOpenAveragePerformanceOverview: 272.41},
			expected: `+------------------------------------------------+--------+
|             METRIC (TEST PROJECT)              | VALUE  |
+------------------------------------------------+--------+
| Issues Time Open Average Performance Overview  | 272.41 |
| Issues Time Open Median Performance Overview   |   0.00 |
| Open Issues Performance Overview               |      0 |
| Reviews Time Open Average Performance Overview |   0.00 |
| Reviews Time Open Median Performance Overview  |   0.00 |
| Open Reviews Performance Overview              |      0 |
+------------------------------------------------+--------+
`,
			writer: &testWriter{},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(innerT *testing.T) {
			innerT.Parallel()

			consoleFormatter := cauldron.NewConsoleFormatter(testProject, "2021-01-01", "2021-12-31", tt.writer)

			err := consoleFormatter.Format(tt.printable)
			if err != nil {
				innerT.Fatalf("error formatting: %v", err)
			}

			formatted := string(tt.writer.data)

			expected := `Project: Test Project (1)
Repo URLs: [http://example.com/repo http://example.com/repo.git]
From: 2021-01-01
To: 2021-12-31
`
			expected += tt.expected

			if formatted != expected {
				innerT.Fatalf("expected \n%s but got \n%s", expected, formatted)
			}
		})
	}
}

func TestJSONFormatter(t *testing.T) {
	w := &testWriter{}
	// using tab as indent
	jsonFormatter := cauldron.NewJSONFormatter(testProject, "2021-01-01", "2021-12-31", "	", w)

	a := &cauldron.Activity{}
	a.CommitsActivityOverview = 15

	err := jsonFormatter.Format(a)
	if err != nil {
		t.Fatalf("error formatting: %v", err)
	}

	formatted := string(w.data)

	// the indent is 2 ep

	expected := `{
	"project": {
		"id": 1,
		"name": "Test Project",
		"RepoURL": [
			"http://example.com/repo",
			"http://example.com/repo.git"
		]
	},
	"from": "2021-01-01",
	"to": "2021-12-31",
	"response": {
		"commits_activity_overview": 15,
		"lines_commit_activity_overview": "",
		"lines_commit_file_activity_overview": "",
		"issues_created_activity_overview": 0,
		"issues_closed_activity_overview": 0,
		"issues_open_activity_overview": 0,
		"reviews_created_activity_overview": 0,
		"reviews_closed_activity_overview": 0,
		"reviews_open_activity_overview": 0
	}
}
`
	if formatted != expected {
		t.Fatalf("expected %s but got %s", expected, formatted)
	}
}
