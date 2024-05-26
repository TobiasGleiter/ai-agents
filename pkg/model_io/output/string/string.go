package output

import (
	"strings"
)

type CommaSeparatedListOutputParser struct {}
type StringOutputParser struct {}

func (p *CommaSeparatedListOutputParser) Parse(output string) interface{} {
    return strings.Split(output, ",")
}

func (p *StringOutputParser) Parse(output string) interface{} {
    return output
}



