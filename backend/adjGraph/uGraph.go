package adjGraph

// interface for undirected graphs
type UGraph interface {
	UAddEdge(from Node, to Node)
	UEdges() []Edge
	UAdj(n Node) []Node
	UNumberOfNodes() int
	GetMatrixEntryAtIndex(i int, j int) bool
}

type UGraphMat struct {
	GraphObject AdjMat
}

func NewUGraph(nodes int) UGraphMat {
	return UGraphMat{NewGraphAdjMat(nodes)}
}

func (g UGraphMat) UAddEdge(from Node, to Node) {
	g.GraphObject.AddEdge(from, to)
	g.GraphObject.AddEdge(to, from)
}

// all edges of the graph
func (g UGraphMat) UEdges() []Edge {
	return g.GraphObject.Edges()
}

// all edges which are adjacent To Node n
func (g UGraphMat) UAdj(n Node) []Node {
	return g.GraphObject.Adj(n)
}

func (g UGraphMat) UNumberOfNodes() int {
	return g.GraphObject.NumberOfNodes()
}

func (g UGraphMat) GetMatrixEntryAtIndex(i int, j int) bool {
	return g.GraphObject[i][j]
}
