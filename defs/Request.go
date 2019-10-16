package defs

type Request struct {
	Url string
	Parse func([]byte, string)
}
