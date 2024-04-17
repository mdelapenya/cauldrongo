package cauldron

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/olekukonko/tablewriter"
)

type Formatter interface {
	Format(io.Writer, Printable) error
	FormatHeader(io.Writer) error
}

type ConsoleFormatter struct {
	From string
	To   string
}

func (c *ConsoleFormatter) Format(w io.Writer, p Printable) error {
	table := tablewriter.NewWriter(w)

	table.SetHeader(p.Headers())
	table.SetAutoWrapText(false)
	table.SetColumnAlignment([]int{tablewriter.ALIGN_LEFT, tablewriter.ALIGN_RIGHT})

	for _, v := range p.Data() {
		table.Append(v)
	}

	table.Render()
	return nil
}

func (c *ConsoleFormatter) FormatHeader(w io.Writer) error {
	fmt.Fprintf(w, "From: %s\n", c.From)
	fmt.Fprintf(w, "To: %s\n", c.To)
	return nil
}

type JSONFormatter struct {
	Indent string
	From   string
	To     string
}

func (j *JSONFormatter) Format(w io.Writer, p Printable) error {
	if j.Indent == "" {
		// default is 2 spaces
		j.Indent = "  "
	}

	bs, err := json.MarshalIndent(p, "", j.Indent)
	if err != nil {
		return fmt.Errorf("error marshalling JSON: %w", err)
	}

	bs = append(bs, '\n')

	_, err = w.Write(bs)
	return err
}

func (j *JSONFormatter) FormatHeader(w io.Writer) error {
	if j.Indent == "" {
		// default is 2 spaces
		j.Indent = "  "
	}

	header := map[string]string{
		"from": j.From,
		"to":   j.To,
	}

	bs, err := json.MarshalIndent(header, "", j.Indent)
	if err != nil {
		return fmt.Errorf("error marshalling header: %w", err)
	}

	bs = append(bs, '\n')
	_, err = w.Write(bs)
	return err
}
