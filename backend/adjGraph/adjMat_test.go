package adjGraph

import (
	"fmt"
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

func ExampleAdjMat() {
	g := NewGraphAdjMat(9)
	mkExampleGraph(&g)
	fmt.Printf("%v\n", g.Edges())
	// Output:
	// [{1 2} {1 3} {1 7} {4 6} {5 4} {6 1} {6 5} {6 6} {7 5} {9 8}]
}

func mkExampleGraph(g Graph) {
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
}

func ExampleGraphAdjMat_Adj() {
	g := NewGraphAdjMat(9)
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
	fmt.Printf("%v\n", g.Adj(1))
	// Output:
	// [2 3 7]
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

func TestAdjMat_NumberOfNodes(t *testing.T) {

	sut := NewGraphAdjMat(4)
	edges := []Edge{
		{1, 2},
		{1, 4},
	}
	for _, edge := range edges {
		sut.AddEdge(edge.from, edge.to)
	}

	tests := []struct {
		name string
		g    AdjMat
		want int
	}{
		{
			name: "test number of nodes",
			g:    sut,
			want: 4,
		},
		{
			name: "test number of nodes: with empty matrix",
			g:    NewGraphAdjMat(0),
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.g.NumberOfNodes(); got != tt.want {
				t.Errorf("AdjMat.NumberOfNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}
