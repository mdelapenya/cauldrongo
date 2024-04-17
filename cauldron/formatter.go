package cauldron

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/mdelapenya/cauldrongo/project"
	"github.com/olekukonko/tablewriter"
)

type Formatter interface {
	Format(Printable) error
	FormatHeader() error
}

func NewConsoleFormatter(p project.Project, from string, to string, w io.Writer) *consoleFormatter {
	return &consoleFormatter{
		Project: p,
		From:    from,
		To:      to,
		Writer:  w,
	}
}

type consoleFormatter struct {
	From    string
	To      string
	Writer  io.Writer
	Project project.Project
}

func (c *consoleFormatter) Format(p Printable) error {
	table := tablewriter.NewWriter(c.Writer)

	var headers = []string{
		"Metric (" + c.Project.Name + ")",
		"Value",
	}

	table.SetHeader(headers)
	table.SetAutoWrapText(false)
	table.SetColumnAlignment([]int{tablewriter.ALIGN_LEFT, tablewriter.ALIGN_RIGHT})

	for _, v := range p.Data() {
		table.Append(v)
	}

	table.Render()
	return nil
}

func (c *consoleFormatter) FormatHeader() error {
	fmt.Fprintf(c.Writer, "Project: %s (%d)\n", c.Project.Name, c.Project.ID)
	fmt.Fprintf(c.Writer, "From: %s\n", c.From)
	fmt.Fprintf(c.Writer, "To: %s\n", c.To)
	return nil
}

func NewJSONFormatter(p project.Project, from string, to string, indent string, w io.Writer) *jsonFormatter {
	if len(indent) == 0 {
		indent = "  "
	}

	return &jsonFormatter{
		Project: p,
		From:    from,
		To:      to,
		Writer:  w,
		Indent:  indent,
	}
}

type jsonFormatter struct {
	Indent  string
	From    string
	To      string
	Writer  io.Writer
	Project project.Project
}

func (j *jsonFormatter) Format(p Printable) error {
	if j.Indent == "" {
		// default is 2 spaces
		j.Indent = "  "
	}

	bs, err := json.MarshalIndent(p, "", j.Indent)
	if err != nil {
		return fmt.Errorf("error marshalling JSON: %w", err)
	}

	bs = append(bs, '\n')

	_, err = j.Writer.Write(bs)
	return err
}

func (j *jsonFormatter) FormatHeader() error {
	if j.Indent == "" {
		// default is 2 spaces
		j.Indent = "  "
	}

	type h struct {
		Project project.Project `json:"project"`
		From    string          `json:"from"`
		To      string          `json:"to"`
	}

	header := h{
		Project: j.Project,
		From:    j.From,
		To:      j.To,
	}

	bs, err := json.MarshalIndent(header, "", j.Indent)
	if err != nil {
		return fmt.Errorf("error marshalling header: %w", err)
	}

	bs = append(bs, '\n')
	_, err = j.Writer.Write(bs)
	return err
}
