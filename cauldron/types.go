package cauldron

import (
	"fmt"
)

var headers = []string{
	"Metric",
	"Value",
}

/*
	{
	    "commits_activity_overview": 15,
	    "lines_commit_activity_overview": "163.13",
	    "lines_commit_file_activity_overview": "1.05",
	    "issues_created_activity_overview": 6,
	    "issues_closed_activity_overview": 6,
	    "issues_open_activity_overview": 66,
	    "reviews_created_activity_overview": 12,
	    "reviews_closed_activity_overview": 13,
	    "reviews_open_activity_overview": 27,
	}
*/
type Activity struct {
	CommitsActivityOverview         int    `json:"commits_activity_overview"`
	LinesCommitActivityOverview     string `json:"lines_commit_activity_overview"`
	LinesCommitFileActivityOverview string `json:"lines_commit_file_activity_overview"`
	IssuesCreatedActivityOverview   int    `json:"issues_created_activity_overview"`
	IssuesClosedActivityOverview    int    `json:"issues_closed_activity_overview"`
	IssuesOpenActivityOverview      int    `json:"issues_open_activity_overview"`
	ReviewsCreatedActivityOverview  int    `json:"reviews_created_activity_overview"`
	ReviewsClosedActivityOverview   int    `json:"reviews_closed_activity_overview"`
	ReviewsOpenActivityOverview     int    `json:"reviews_open_activity_overview"`
}

func (a *Activity) Data() [][]string {
	return [][]string{
		{"Commits Activity Overview", fmt.Sprintf("%d", a.CommitsActivityOverview)},
		{"Lines Commit Activity Overview", a.LinesCommitActivityOverview},
		{"Lines Commit File Activity Overview", a.LinesCommitFileActivityOverview},
		{"Issues Created Activity Overview", fmt.Sprintf("%d", a.IssuesCreatedActivityOverview)},
		{"Issues Closed Activity Overview", fmt.Sprintf("%d", a.IssuesClosedActivityOverview)},
		{"Issues Open Activity Overview", fmt.Sprintf("%d", a.IssuesOpenActivityOverview)},
		{"Reviews Created Activity Overview", fmt.Sprintf("%d", a.ReviewsCreatedActivityOverview)},
		{"Reviews Closed Activity Overview", fmt.Sprintf("%d", a.ReviewsClosedActivityOverview)},
		{"Reviews Open Activity Overview", fmt.Sprintf("%d", a.ReviewsOpenActivityOverview)},
	}
}

func (a *Activity) Headers() []string {
	return headers
}

/*
	{
		"active_people_git_community_overview": 8,
		"active_people_issues_community_overview": 6,
		"active_people_patches_community_overview": 8,
		"onboardings_git_community_overview": 2,
		"onboardings_issues_community_overview": 4,
		"onboardings_patches_community_overview": 4,
	}
*/
type Community struct {
	ActivePeopleGitCommunityOverview     int `json:"active_people_git_community_overview"`
	ActivePeopleIssuesCommunityOverview  int `json:"active_people_issues_community_overview"`
	ActivePeoplePatchesCommunityOverview int `json:"active_people_patches_community_overview"`
	OnboardingsGitCommunityOverview      int `json:"onboardings_git_community_overview"`
	OnboardingsIssuesCommunityOverview   int `json:"onboardings_issues_community_overview"`
	OnboardingsPatchesCommunityOverview  int `json:"onboardings_patches_community_overview"`
}

func (c *Community) Data() [][]string {
	return [][]string{
		{"Active People Git Community Overview", fmt.Sprintf("%d", c.ActivePeopleGitCommunityOverview)},
		{"Active People Issues Community Overview", fmt.Sprintf("%d", c.ActivePeopleIssuesCommunityOverview)},
		{"Active People Patches Community Overview", fmt.Sprintf("%d", c.ActivePeoplePatchesCommunityOverview)},
		{"Onboardings Git Community Overview", fmt.Sprintf("%d", c.OnboardingsGitCommunityOverview)},
		{"Onboardings Issues Community Overview", fmt.Sprintf("%d", c.OnboardingsIssuesCommunityOverview)},
		{"Onboardings Patches Community Overview", fmt.Sprintf("%d", c.OnboardingsPatchesCommunityOverview)},
	}
}

func (c *Community) Headers() []string {
	return headers
}

/*
	{
	    "commits_overview": 1581,
	    "issues_overview": 386,
	    "reviews_overview": 2057,
	    "questions_overview": "?",
	    "commits_last_year_overview": 661,
	    "issues_last_year_overview": 142,
	    "reviews_last_year_overview": 1258,
	    "questions_last_year_overview": "?",
	    "commits_yoy_overview": 52.66,
	    "issues_yoy_overview": 56.04,
	    "reviews_yoy_overview": 139.62,
	    "questions_yoy_overview": 0,
	    "commit_authors_overview": 227,
	    "issue_submitters_overview": 227,
	    "review_submitters_overview": 245,
	    "question_authors_overview": "?",
	    "commit_authors_last_year_overview": 85,
	    "issue_submitters_last_year_overview": 93,
	    "review_submitters_last_year_overview": 98,
	    "question_authors_last_year_overview": "?",
	    "commit_authors_yoy_overview": 41.67,
	    "issue_submitters_yoy_overview": 69.09,
	    "review_submitters_yoy_overview": 48.48,
	    "question_authors_yoy_overview": 0,
	    "issues_median_time_to_close_overview": 18.14,
	    "reviews_median_time_to_close_overview": 0.91,
	    "issues_median_time_to_close_last_year_overview": 7.27,
	    "reviews_median_time_to_close_last_year_overview": 0.97,
	    "issues_median_time_to_close_yoy_overview": -91.61,
	    "reviews_median_time_to_close_yoy_overview": 61.67,
	}
*/
// Overview is a struct that represents the overview tab of the metrics endpoint.
// The token 'YOY' stands for 'Year-Over-Year'.
type Overview struct {
	CommitsOverview                          int     `json:"commits_overview"`
	IssuesOverview                           int     `json:"issues_overview"`
	ReviewsOverview                          int     `json:"reviews_overview"`
	QuestionsOverview                        string  `json:"questions_overview"`
	CommitsLastYearOverview                  int     `json:"commits_last_year_overview"`
	IssuesLastYearOverview                   int     `json:"issues_last_year_overview"`
	ReviewsLastYearOverview                  int     `json:"reviews_last_year_overview"`
	QuestionsLastYearOverview                string  `json:"questions_last_year_overview"`
	CommitsYoyOverview                       float64 `json:"commits_yoy_overview"`
	IssuesYoyOverview                        float64 `json:"issues_yoy_overview"`
	ReviewsYoyOverview                       float64 `json:"reviews_yoy_overview"`
	QuestionsYoyOverview                     int     `json:"questions_yoy_overview"`
	CommitAuthorsOverview                    int     `json:"commit_authors_overview"`
	IssueSubmittersOverview                  int     `json:"issue_submitters_overview"`
	ReviewSubmittersOverview                 int     `json:"review_submitters_overview"`
	QuestionAuthorsOverview                  string  `json:"question_authors_overview"`
	CommitAuthorsLastYearOverview            int     `json:"commit_authors_last_year_overview"`
	IssueSubmittersLastYearOverview          int     `json:"issue_submitters_last_year_overview"`
	ReviewSubmittersLastYearOverview         int     `json:"review_submitters_last_year_overview"`
	QuestionAuthorsLastYearOverview          string  `json:"question_authors_last_year_overview"`
	CommitAuthorsYoyOverview                 float64 `json:"commit_authors_yoy_overview"`
	IssueSubmittersYoyOverview               float64 `json:"issue_submitters_yoy_overview"`
	ReviewSubmittersYoyOverview              float64 `json:"review_submitters_yoy_overview"`
	QuestionAuthorsYoyOverview               int     `json:"question_authors_yoy_overview"`
	IssuesMedianTimeToCloseOverview          float64 `json:"issues_median_time_to_close_overview"`
	ReviewsMedianTimeToCloseOverview         float64 `json:"reviews_median_time_to_close_overview"`
	IssuesMedianTimeToCloseLastYearOverview  float64 `json:"issues_median_time_to_close_last_year_overview"`
	ReviewsMedianTimeToCloseLastYearOverview float64 `json:"reviews_median_time_to_close_last_year_overview"`
	IssuesMedianTimeToCloseYoyOverview       float64 `json:"issues_median_time_to_close_yoy_overview"`
	ReviewsMedianTimeToCloseYoyOverview      float64 `json:"reviews_median_time_to_close_yoy_overview"`
}

func (o *Overview) Data() [][]string {
	return [][]string{
		{"Commits Overview", fmt.Sprintf("%d", o.CommitsOverview)},
		{"Issues Overview", fmt.Sprintf("%d", o.IssuesOverview)},
		{"Reviews Overview", fmt.Sprintf("%d", o.ReviewsOverview)},
		{"Questions Overview", o.QuestionsOverview},
		{"Commits Last Year Overview", fmt.Sprintf("%d", o.CommitsLastYearOverview)},
		{"Issues Last Year Overview", fmt.Sprintf("%d", o.IssuesLastYearOverview)},
		{"Reviews Last Year Overview", fmt.Sprintf("%d", o.ReviewsLastYearOverview)},
		{"Questions Last Year Overview", o.QuestionsLastYearOverview},
		{"Commits YoY Overview", fmt.Sprintf("%.2f", o.CommitsYoyOverview)},
		{"Issues YoY Overview", fmt.Sprintf("%.2f", o.IssuesYoyOverview)},
		{"Reviews YoY Overview", fmt.Sprintf("%.2f", o.ReviewsYoyOverview)},
		{"Questions YoY Overview", fmt.Sprintf("%d", o.QuestionsYoyOverview)},
		{"Commit Authors Overview", fmt.Sprintf("%d", o.CommitAuthorsOverview)},
		{"Issue Submitters Overview", fmt.Sprintf("%d", o.IssueSubmittersOverview)},
		{"Review Submitters Overview", fmt.Sprintf("%d", o.ReviewSubmittersOverview)},
		{"Question Authors Overview", o.QuestionAuthorsOverview},
		{"Commit Authors Last Year Overview", fmt.Sprintf("%d", o.CommitAuthorsLastYearOverview)},
		{"Issue Submitters Last Year Overview", fmt.Sprintf("%d", o.IssueSubmittersLastYearOverview)},
		{"Review Submitters Last Year Overview", fmt.Sprintf("%d", o.ReviewSubmittersLastYearOverview)},
		{"Question Authors Last Year Overview", o.QuestionAuthorsLastYearOverview},
		{"Commit Authors YoY Overview", fmt.Sprintf("%.2f", o.CommitAuthorsYoyOverview)},
		{"Issue Submitters YoY Overview", fmt.Sprintf("%.2f", o.IssueSubmittersYoyOverview)},
		{"Review Submitters YoY Overview", fmt.Sprintf("%.2f", o.ReviewSubmittersYoyOverview)},
		{"Question Authors YoY Overview", fmt.Sprintf("%d", o.QuestionAuthorsYoyOverview)},
		{"Issues Median Time To Close Overview", fmt.Sprintf("%.2f", o.IssuesMedianTimeToCloseOverview)},
		{"Reviews Median Time To Close Overview", fmt.Sprintf("%.2f", o.ReviewsMedianTimeToCloseOverview)},
		{"Issues Median Time To Close Last Year Overview", fmt.Sprintf("%.2f", o.IssuesMedianTimeToCloseLastYearOverview)},
		{"Reviews Median Time To Close Last Year Overview", fmt.Sprintf("%.2f", o.ReviewsMedianTimeToCloseLastYearOverview)},
		{"Issues Median Time To Close YoY Overview", fmt.Sprintf("%.2f", o.IssuesMedianTimeToCloseYoyOverview)},
		{"Reviews Median Time To Close YoY Overview", fmt.Sprintf("%.2f", o.ReviewsMedianTimeToCloseYoyOverview)},
	}
}

func (o *Overview) Headers() []string {
	return headers
}

/*
	{
	    "issues_time_open_average_performance_overview": 272.41,
	    "issues_time_open_median_performance_overview": 209.19,
	    "open_issues_performance_overview": 66,
	    "reviews_time_open_average_performance_overview": 71.03,
	    "reviews_time_open_median_performance_overview": 47.8,
	    "open_reviews_performance_overview": 27,
	}
*/
type Performance struct {
	IssuesTimeOpenAveragePerformanceOverview  float64 `json:"issues_time_open_average_performance_overview"`
	IssuesTimeOpenMedianPerformanceOverview   float64 `json:"issues_time_open_median_performance_overview"`
	OpenIssuesPerformanceOverview             int     `json:"open_issues_performance_overview"`
	ReviewsTimeOpenAveragePerformanceOverview float64 `json:"reviews_time_open_average_performance_overview"`
	ReviewsTimeOpenMedianPerformanceOverview  float64 `json:"reviews_time_open_median_performance_overview"`
	OpenReviewsPerformanceOverview            int     `json:"open_reviews_performance_overview"`
}

func (p *Performance) Data() [][]string {
	return [][]string{
		{"Issues Time Open Average Performance Overview", fmt.Sprintf("%.2f", p.IssuesTimeOpenAveragePerformanceOverview)},
		{"Issues Time Open Median Performance Overview", fmt.Sprintf("%.2f", p.IssuesTimeOpenMedianPerformanceOverview)},
		{"Open Issues Performance Overview", fmt.Sprintf("%d", p.OpenIssuesPerformanceOverview)},
		{"Reviews Time Open Average Performance Overview", fmt.Sprintf("%.2f", p.ReviewsTimeOpenAveragePerformanceOverview)},
		{"Reviews Time Open Median Performance Overview", fmt.Sprintf("%.2f", p.ReviewsTimeOpenMedianPerformanceOverview)},
		{"Open Reviews Performance Overview", fmt.Sprintf("%d", p.OpenReviewsPerformanceOverview)},
	}
}

func (p *Performance) Headers() []string {
	return headers
}
