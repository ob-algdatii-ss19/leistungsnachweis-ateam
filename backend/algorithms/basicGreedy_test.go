package algorithms

import (
	"reflect"
	"testing"

	"github.com/ob-algdatii-ss19/leistungsnachweis-ateam/backend/adjGraph"
)

func TestBasicGreedy(t *testing.T) {
	type args struct {
		graphData adjGraph.UGraph
	}
	testAdjMat01 := adjGraph.NewGraphAdjMat(9)
	testAdjMat01.AddEdge(1, 2)
	testAdjMat01.AddEdge(3, 2)
	testAdjMat01.AddEdge(4, 3)

	tests := []struct {
		name string
		args args
		want [][]adjGraph.Node
	}{
		{
			"simple first test",
			args{adjGraph.NewUGraph(3)},
			[][]adjGraph.Node{
				{adjGraph.Node(1)},
				{adjGraph.Node(2), adjGraph.Node(4)},
				{adjGraph.Node(3)},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BasicGreedy(tt.args.graphData); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BasicGreedy() = %v, want %v", got, tt.want)
			}
		})
	}
}
