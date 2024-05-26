package output

import (
	"strings"
)

type OutputParser interface {
	Parse(output string) (interface{}, error)
}

type CommaSeparatedListOutputParser struct {}

func (p *CommaSeparatedListOutputParser) Parse(output string) ([]string, error) {
    return strings.Split(output, ","), nil
}



