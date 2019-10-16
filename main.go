package main

import (
	"spiders/defs"
	"spiders/engine"
)

func main() {
	var seeds []defs.Request
	seeds = []defs.Request{
		{
			Url:   "http://www.zhenai.com/zhenghun/beijing",
			Parse: engine.NewFuncParser(parser.ParseCity, "ParseCity"),
		},
	}
	engine.Run()
}
