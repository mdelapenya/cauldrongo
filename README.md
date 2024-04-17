# Cauldron Go

This is a simple Go program that fetches the OSS metrics from the Cauldron API.

## Cauldron APIs

### Metrics

The metrics endpoint has the following format:

```
https://cauldron.io/project/${PROJECT_ID}/metrics?from=${FROM}&to=${TO}&tab=${TAB}
```

There is one URL parameter: `PROJECT_ID`. This is the ID of the project in Cauldron. The querystring receives three parameters:

- `from`: the start date of the metrics, in the format `YYYY-MM-DD`
- `to`: the end date of the metrics, in the format `YYYY-MM-DD`
- `tab`: the tab of the metrics

The `tab` parameter can be one of the following values:

- `overview`: the overview tab
- `activity-overview`: the activity overview tab
- `community-overview`: the community overview tab
- `performance-overview`: the performance overview tab

It's **important** to note that the repositories, and their datasources, must be refreshed before fetching the metrics. So please do it manually before querying the API.

## Usage

The CLI has one subcommand: `metrics`. It has the following flags:

- `--project | -p`: the project ID. Required.
- `--from | -f`: the start date of the metrics, in the format `YYYY-MM-DD`. Default is one year ago.
- `--to | -t`: the end date of the metrics, in the format `YYYY-MM-DD`. Default is today.
- `--tab | -T`: the tab of the metrics. Default is `overview`.
- `--format | -F`: the output format, can be `console` or `json`. Default is `console`.

### Examples

```sh
# Fetch the metrics for the project 1, from one year ago to today, using the overview tab
cauldron-go metrics --project 1
```

## Not implemented (yet)

- Refresh the repositories and datasources before fetching the metrics, using the refresh endpoint, but we need to deal with credentials, so I'm postponing it for a second iteration.
