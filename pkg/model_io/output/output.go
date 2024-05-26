package output

import (
	"strings"
)

type OutputParser interface {
	Parse(output string) []string
}

type CommaSeparatedListOutputParser struct {}

func (p *CommaSeparatedListOutputParser) Parse(output string) []string {
    return strings.Split(output, ",")
}