package adjGraph

// interface for directed graphs

type Node int

type Edge struct {
	From Node
	To   Node
}

type Graph interface {
	AddEdge(from Node, to Node)
	Edges() []Edge
	Adj(n Node) []Node
	NumberOfNodes() int
}
