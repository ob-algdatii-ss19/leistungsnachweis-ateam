package adjGraph

// interface for directed graphs

type Node int

type Edge struct {
	from Node
	to   Node
}

type Graph interface {
	AddEdge(from Node, to Node)
	Edges() []Edge
	Adj(n Node) []Node
}
