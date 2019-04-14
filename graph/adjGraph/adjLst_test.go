package adjGraph

import (
	"fmt"
	"testing"
)

func TestAdjLst(t *testing.T) {
	g := NewGraphAdjLst(9)
	testGraph(t, g)
}

func testGraph(t *testing.T, g Graph) {
	wantedEdges := []Edge{
		{1, 2},
		{1, 3},
		{1, 7},
		{4, 6},
		{5, 4},
		{6, 1},
		{6, 5},
		{6, 6},
		{7, 5},
		{9, 8},
	}

	for _, edge := range wantedEdges {
		g.AddEdge(edge.from, edge.to)
	}

	gotEdges := g.Edges()

outerWanted:
	for _, wanted := range wantedEdges {
		for _, got := range gotEdges {
			if wanted.from == got.from && wanted.to == got.to {
				continue outerWanted
			}
		}
		t.Errorf("Edge %v wanted, but not found", wanted)
	}

outerGot:
	for _, got := range gotEdges {
		for _, wanted := range wantedEdges {
			if wanted.from == got.from && wanted.to == got.to {
				continue outerGot
			}
		}
		t.Errorf("Edge %v found, but not wanted", got)
	}
}

func ExampleAdjLst() {
	g := NewGraphAdjLst(9)
	mkExampleGraph(&g)
	fmt.Printf("%v\n", g.Edges())
	// Output:
	// [{1 2} {1 3} {1 7} {4 6} {5 4} {6 1} {6 5} {6 6} {7 5} {9 8}]
}

func BenchmarkAdjLstEdges(b *testing.B) {
	g := NewGraphAdjLst(9)
	benchmarkHelper(b, g)
}
