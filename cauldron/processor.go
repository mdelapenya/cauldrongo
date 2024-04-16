package cauldron

import "io"

type Processor interface {
	Process(io.Reader) error
}
