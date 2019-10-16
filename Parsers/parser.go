package Parsers

import "spiders/defs"

type ParseFunc func(body []byte, url string) defs.ParseResult

type Parser interface {
	Parse(contents []byte, url string) defs.ParseResult
	Serialize() (name string, args interface{})
}

func NilParseFunc(body []byte, url string) defs.ParseResult {
	return defs.ParseResult{}
}
