package cauldron

type Printable interface {
	Data() [][]string
	Headers() []string
}
