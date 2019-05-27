package adjGraph

import (
	"testing"
)

func TestAdjMat(t *testing.T) {
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

	g := NewGraphAdjMat(9)
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

func BenchmarkAdjMatEdges(b *testing.B) {
	g := NewGraphAdjMat(9)
	benchmarkHelper(b, g)
}

func benchmarkHelper(b *testing.B, g Graph) {
	edges := []Edge{
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

	for _, edge := range edges {
		g.AddEdge(edge.from, edge.to)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		g.Edges()
	}
}
