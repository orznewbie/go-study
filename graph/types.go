package graph

import "encoding/json"

type Uid = string

type (
	KV struct {
		Key string
		Value Value
	}
	Value json.RawMessage
)

type Vertex struct {
	ID Uid
	Properties map[string]interface{}
	Labels []string
}

type Edge struct {
	ID Uid
	Source Uid
	Kind string
	Target Uid
	Properties map[string]string
}

type Mutator interface {
	AddVertex(...*Vertex)
	AddEdge(...*Edge)
}

type Queryer interface {

}

type Traverser interface {
	V(...interface{}) Traverser
	Out() Traverser
	OutE() Traverser
	In() Traverser
	InE() Traverser
	Values()
	HasLabel(string) Traverser
}