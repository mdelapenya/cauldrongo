package cauldron

import (
	"encoding/json"
	"fmt"
)

type Formatter interface {
	Format(Processor) (string, error)
}

type ConsoleFormatter struct{}

func (c *ConsoleFormatter) Format(p Processor) (string, error) {
	return fmt.Sprintf("%+v", p), nil
}

type JSONFormatter struct {
	Indent string
}

func (j *JSONFormatter) Format(p Processor) (string, error) {
	if j.Indent == "" {
		// default is 2 spaces
		j.Indent = "  "
	}

	bs, err := json.MarshalIndent(p, "", j.Indent)
	if err != nil {
		return "", fmt.Errorf("error marshalling JSON: %w", err)
	}

	return string(bs), nil
}
