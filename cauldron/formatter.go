package cauldron

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/olekukonko/tablewriter"
)

type Formatter interface {
	Format(io.Writer, Printable) error
}

type ConsoleFormatter struct{}

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

type JSONFormatter struct {
	Indent string
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
