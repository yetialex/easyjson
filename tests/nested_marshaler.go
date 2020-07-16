package tests

import (
	"github.com/yetialex/easyjson"
	"github.com/yetialex/easyjson/jlexer"
	"github.com/yetialex/easyjson/jwriter"
)

type MarshalerUnmarshaler interface {
	easyjson.Marshaler
	easyjson.Unmarshaler
}

//easyjson:json
type NestedMarshaler struct {
	Value  MarshalerUnmarshaler
	Value2 int
}

type StructWithMarshaler struct {
	Value int
}

func (s *StructWithMarshaler) UnmarshalEasyJSON(w *jlexer.Lexer) {
	s.Value = w.Int()
}

func (s *StructWithMarshaler) MarshalEasyJSON(w *jwriter.Writer) {
	w.Int(s.Value)
}
