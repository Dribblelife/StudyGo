package parser

import (
	"regexp"
	"LearGoProject/crawler/engine"
	"fmt"
)

var re = regexp.MustCompile(
	`<script>window.__INITIAL_STATE__=(.+);\(function`)

func ParseProfile1(contents []byte) engine.ParseResult  {
	match:= re.FindSubmatch(contents)
	if len(match)>=2 {
		json:=match[1]
		fmt.Printf("json:%s\n",json)
	}
	return engine.ParseResult{}
}

