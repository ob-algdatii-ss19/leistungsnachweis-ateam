package graph

type Elem struct {
	value Node
	next  *Elem
}

type AdjLst []*Elem

func NewGraphAdjLst(nodes int) AdjLst {
	g := make([]*Elem, nodes+1)
	return g
}

func (g AdjLst) AddEdge(from Node, to Node) {
	if g[from] == nil {
		g[from] = &Elem{to, nil}
	} else {
		elem := g[from]
		for elem.next != nil {
			elem = elem.next
		}
		elem.next = &Elem{to, nil}
	}
}

func (g AdjLst) Edges() []Edge {
	es := make([]Edge, 0)
	for i, elem := range g {
		for elem != nil {
			es = append(es, Edge{Node(i), elem.value})
			elem = elem.next
		}
	}
	return es
}
