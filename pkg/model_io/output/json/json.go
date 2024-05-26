package output 

import (
	"encoding/json"
)

type JsonOutputParser struct {}

func (p *JsonOutputParser) Parse(output string) interface{} {
	var result interface{}
	return json.Unmarshal([]byte(output), &result)
}