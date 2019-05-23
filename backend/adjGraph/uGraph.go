package adjGraph

// interface for undirected graphs
type UGraph interface {
	UAddEdge(from Node, to Node)
	UEdges() []Edge
	UAdj(n Node) []Node
}

type UGraphMat struct {
	Graph
}

func NewUGraph(nodes int) UGraphMat {
	return UGraphMat{NewGraphAdjMat(nodes)}
}

func (g UGraphMat) UAddEdge(from Node, to Node) {
	g.Graph.AddEdge(from, to)
	g.Graph.AddEdge(to, from)
}

// all edges of the graph
func (g UGraphMat) UEdges() []Edge {
	return g.Graph.Edges()
}

// all edges which are adjacent to Node n
func (g UGraphMat) UAdj(n Node) []Node {
	return g.Graph.Adj(n)
}
