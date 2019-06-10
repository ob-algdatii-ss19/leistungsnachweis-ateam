package algorithms

import (
	"reflect"
	"testing"

	"github.com/ob-algdatii-ss19/leistungsnachweis-ateam/backend/adjGraph"
)

func TestWelshPowell(t *testing.T) {

	sutUGraphMat := adjGraph.NewUGraph(12)
	sutUGraphMat.UAddEdge(1, 2)
	sutUGraphMat.UAddEdge(1, 3)
	sutUGraphMat.UAddEdge(2, 3)
	sutUGraphMat.UAddEdge(2, 4)
	sutUGraphMat.UAddEdge(3, 4)
	sutUGraphMat.UAddEdge(4, 5)
	sutTrafficEntries := []adjGraph.TrafficEntry{
		{adjGraph.ABC, adjGraph.MNO, true},
		{adjGraph.ABC, adjGraph.IJK, true},
		{adjGraph.ABC, adjGraph.EFG, true},
		{adjGraph.EFG, adjGraph.ABC, true},
		{adjGraph.EFG, adjGraph.MNO, true},
		{adjGraph.EFG, adjGraph.IJK, false},
		{adjGraph.IJK, adjGraph.EFG, false},
		{adjGraph.IJK, adjGraph.ABC, false},
		{adjGraph.IJK, adjGraph.MNO, false},
		{adjGraph.MNO, adjGraph.IJK, false},
		{adjGraph.MNO, adjGraph.EFG, false},
		{adjGraph.MNO, adjGraph.ABC, false},
	}

	sutUGraphMat01 := adjGraph.NewUGraph(12)
	sutTrafficEntries01 := []adjGraph.TrafficEntry{
		{adjGraph.ABC, adjGraph.MNO, false},
		{adjGraph.ABC, adjGraph.IJK, false},
		{adjGraph.ABC, adjGraph.EFG, true},
		{adjGraph.EFG, adjGraph.ABC, false},
		{adjGraph.EFG, adjGraph.MNO, false},
		{adjGraph.EFG, adjGraph.IJK, true},
		{adjGraph.IJK, adjGraph.EFG, false},
		{adjGraph.IJK, adjGraph.ABC, false},
		{adjGraph.IJK, adjGraph.MNO, true},
		{adjGraph.MNO, adjGraph.IJK, false},
		{adjGraph.MNO, adjGraph.EFG, false},
		{adjGraph.MNO, adjGraph.ABC, true},
	}

	sutUGraphMat02 := adjGraph.NewUGraph(0)
	sutTrafficEntries02 := []adjGraph.TrafficEntry{{}}

	sutUGraphMat03 := adjGraph.NewUGraph(4)
	sutUGraphMat03.UAddEdge(1, 2)
	sutUGraphMat03.UAddEdge(1, 3)
	sutUGraphMat03.UAddEdge(1, 4)
	sutUGraphMat03.UAddEdge(2, 3)
	sutUGraphMat03.UAddEdge(2, 4)
	sutUGraphMat03.UAddEdge(4, 3)

	sutTrafficEntries03 := []adjGraph.TrafficEntry{
		{adjGraph.ABC, adjGraph.MNO, true},
		{adjGraph.ABC, adjGraph.IJK, true},
		{adjGraph.ABC, adjGraph.EFG, true},
		{adjGraph.EFG, adjGraph.ABC, true},
	}

	type args struct {
		argsData adjGraph.ReturnType
	}

	tests := []struct {
		name string
		args args
		want [][]adjGraph.Node
	}{
		{
			"coloring simple graph",
			args{adjGraph.ReturnType{sutTrafficEntries, sutUGraphMat}},
			[][]adjGraph.Node{
				{adjGraph.Node(1), adjGraph.Node(4)},
				{adjGraph.Node(2), adjGraph.Node(5)},
				{adjGraph.Node(3)},
			},
		},
		{
			"coloring graph without edges",
			args{adjGraph.ReturnType{sutTrafficEntries01, sutUGraphMat01}},
			[][]adjGraph.Node{
				{adjGraph.Node(3), adjGraph.Node(6), adjGraph.Node(9), adjGraph.Node(12)},
			},
		},
		{
			"coloring empty graph",
			args{adjGraph.ReturnType{sutTrafficEntries02, sutUGraphMat02}},
			[][]adjGraph.Node{},
		},
		{
			"handling nil objects",
			args{adjGraph.ReturnType{}},
			[][]adjGraph.Node{},
		},
		{
			"coloring graph with only different colors",
			args{adjGraph.ReturnType{sutTrafficEntries03, sutUGraphMat03}},
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
			if got := BasicGreedy(tt.args.argsData); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BasicGreedy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSortNodesDescending(t *testing.T) {
	result := make([][]int, 0)
	var innerArray1 []int;
	innerArray1 = append(innerArray1, 1)
	innerArray1 = append(innerArray1, -12)
	result = append(result, innerArray1)

	var innerArray2 []int;
	innerArray2 = append(innerArray2, 1)
	innerArray2 = append(innerArray2, 1)
	innerArray2 = append(innerArray2, 0)
	result = append(result, innerArray2)

	var innerArray3 []int;
	innerArray3 = append(innerArray3, 1)
	result = append(result, innerArray3)


	wanted := make([][]int, 0)

	wanted = append(wanted, innerArray2)
	wanted = append(wanted, innerArray1)
	wanted = append(wanted, innerArray3)



	tests := []struct {
		name string
		args [][]int
		want [][]int
	}{
		{
			"first test:",
			result,
			wanted,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortNodesDescending(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildGraphObjectFromJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}
