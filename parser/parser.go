package parser

type Parser interface {
	Parse(content []byte, id int8) any
}

type NilParser struct {
}

func (n NilParser) Parse(content []byte, id int8) any {
	return nil
}
