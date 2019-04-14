package graph

type Node int

type Edge struct {
	from Node
	to   Node
}

type Graph interface {
	AddEdge(from Node, to Node)
	Edges() []Edge
}
