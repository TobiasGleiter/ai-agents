package model

type SimpleModel struct{}

func (m *SimpleModel) Process(input string) string {
	generated := input
	return generated
}