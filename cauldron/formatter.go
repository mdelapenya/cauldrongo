package cauldron

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/olekukonko/tablewriter"
)

type Formatter interface {
	Format(Printable) error
	FormatHeader() error
}

type ConsoleFormatter struct {
	From   string
	To     string
	Writer io.Writer
}

func (c *ConsoleFormatter) Format(p Printable) error {
	table := tablewriter.NewWriter(c.Writer)

	table.SetHeader(p.Headers())
	table.SetAutoWrapText(false)
	table.SetColumnAlignment([]int{tablewriter.ALIGN_LEFT, tablewriter.ALIGN_RIGHT})

	for _, v := range p.Data() {
		table.Append(v)
	}

	table.Render()
	return nil
}

func (c *ConsoleFormatter) FormatHeader() error {
	fmt.Fprintf(c.Writer, "From: %s\n", c.From)
	fmt.Fprintf(c.Writer, "To: %s\n", c.To)
	return nil
}

type JSONFormatter struct {
	Indent string
	From   string
	To     string
	Writer io.Writer
}

func (j *JSONFormatter) Format(p Printable) error {
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

func (j *JSONFormatter) FormatHeader() error {
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
	_, err = j.Writer.Write(bs)
	return err
}
