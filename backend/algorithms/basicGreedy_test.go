package algorithms

import (
	"reflect"
	"testing"

	"github.com/ob-algdatii-ss19/leistungsnachweis-ateam/backend/adjGraph"
)

func TestBasicGreedy(t *testing.T) {

	sutUGraphMat := adjGraph.NewUGraph(5)
	sutUGraphMat.UAddEdge(1, 2)
	sutUGraphMat.UAddEdge(1, 3)
	sutUGraphMat.UAddEdge(2, 3)
	sutUGraphMat.UAddEdge(2, 4)
	sutUGraphMat.UAddEdge(3, 4)
	sutUGraphMat.UAddEdge(4, 5)

	sutUGraphMat01 := adjGraph.NewUGraph(4)

	sutUGraphMat02 := adjGraph.NewUGraph(0)

	sutUGraphMat03 := adjGraph.NewUGraph(4)
	sutUGraphMat03.UAddEdge(1, 2)
	sutUGraphMat03.UAddEdge(1, 3)
	sutUGraphMat03.UAddEdge(1, 4)
	sutUGraphMat03.UAddEdge(2, 3)
	sutUGraphMat03.UAddEdge(2, 4)
	sutUGraphMat03.UAddEdge(4, 3)

	type args struct {
		graphData adjGraph.UGraph
	}

	tests := []struct {
		name string
		args args
		want [][]adjGraph.Node
	}{
		{
			"coloring simple graph",
			args{sutUGraphMat},
			[][]adjGraph.Node{
				{adjGraph.Node(1), adjGraph.Node(4)},
				{adjGraph.Node(2), adjGraph.Node(5)},
				{adjGraph.Node(3)},
			},
		},
		{
			"coloring graph without edges",
			args{sutUGraphMat01},
			[][]adjGraph.Node{
				{adjGraph.Node(1), adjGraph.Node(2), adjGraph.Node(3), adjGraph.Node(4)},
			},
		},
		{
			"coloring empty graph",
			args{sutUGraphMat02},
			[][]adjGraph.Node{},
		},
		{
			"coloring graph with only different colors",
			args{sutUGraphMat03},
			[][]adjGraph.Node{
				{adjGraph.Node(1)},
				{adjGraph.Node(2)},
				{adjGraph.Node(3)},
				{adjGraph.Node(4)},
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
