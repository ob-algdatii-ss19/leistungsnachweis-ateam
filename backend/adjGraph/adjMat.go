package adjGraph

type AdjMat [][]bool

func NewGraphAdjMat(nodes int) AdjMat {
	g := make([][]bool, nodes+1)
	for i := range g {
		if i != 0 {
			g[i] = make([]bool, nodes+1)
		}
	}
	return g
}

func (g AdjMat) AddEdge(from Node, to Node) {
	g[from][to] = true
}

func (g AdjMat) RemoveEdge(from Node, to Node) {
	g[from][to] = false
}

func (g AdjMat) NumberOfNodes() int {
	return len(g) - 1
}

func (g AdjMat) Edges() []Edge {
	es := make([]Edge, 0)
	for i, row := range g {
		if i == 0 {
			continue
		}
		for j, edge := range row {
			if j == 0 {
				continue
			}
			if edge {
				es = append(es, Edge{Node(i), Node(j)})
			}
		}
	}
	return es
}

func (g AdjMat) Adj(n Node) []Node {
	adjNodes := make([]Node, 0)

	for i, edge := range g[n] {
		if edge {
			adjNodes = append(adjNodes, Node(i))
		}
	}

	return adjNodes
}
