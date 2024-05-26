package output 

import (
	"encoding/json"
)

type OutputParser interface {
	Parse(output string, target interface{}) error
}

type JsonOutputParser struct {}

func (p *JsonOutputParser) Parse(output string, target interface{}) error {
	return json.Unmarshal([]byte(output), target)
}