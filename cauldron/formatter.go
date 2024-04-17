package cauldron

import (
	"encoding/json"
	"fmt"
	"io"
)

type Formatter interface {
	Format(io.Writer, Processor) error
}

type ConsoleFormatter struct{}

func (c *ConsoleFormatter) Format(w io.Writer, p Processor) error {
	_, err := w.Write([]byte(fmt.Sprintf("%+v", p)))
	return err
}

type JSONFormatter struct {
	Indent string
}

func (j *JSONFormatter) Format(w io.Writer, p Processor) error {
	if j.Indent == "" {
		// default is 2 spaces
		j.Indent = "  "
	}

	bs, err := json.MarshalIndent(p, "", j.Indent)
	if err != nil {
		return fmt.Errorf("error marshalling JSON: %w", err)
	}

	_, err = w.Write(bs)
	return err
}
