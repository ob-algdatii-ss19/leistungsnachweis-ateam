package algorithms

import (
	"github.com/ob-algdatii-ss19/leistungsnachweis-ateam/backend/adjGraph"
	"reflect"
	"testing"
)

func TestWelshPowell(t *testing.T) {

	sutUGraphMat := adjGraph.NewUGraph(18)
	sutTrafficEntries := []adjGraph.TrafficEntry{
		{adjGraph.ABC, adjGraph.EFG, true},
		{adjGraph.ABC, adjGraph.IJK, false},
		{adjGraph.ABC, adjGraph.MNO, false},
		{adjGraph.P1, adjGraph.P1, false},
		{adjGraph.P2, adjGraph.P2, false},

		{adjGraph.EFG, adjGraph.ABC, false},
		{adjGraph.EFG, adjGraph.IJK, true},
		{adjGraph.EFG, adjGraph.MNO, false},
		{adjGraph.P1, adjGraph.P1, false},
		{adjGraph.P2, adjGraph.P2, false},

		{adjGraph.IJK, adjGraph.ABC, false},
		{adjGraph.IJK, adjGraph.EFG, false},
		{adjGraph.IJK, adjGraph.MNO, true},
		{adjGraph.P1, adjGraph.P1, false},
		{adjGraph.P2, adjGraph.P2, false},

		{adjGraph.MNO, adjGraph.ABC, true},
		{adjGraph.MNO, adjGraph.EFG, false},
		{adjGraph.MNO, adjGraph.IJK, false},
		{adjGraph.P1, adjGraph.P1, false},
		{adjGraph.P2, adjGraph.P2, false},
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
				{adjGraph.Node(1), adjGraph.Node(7), adjGraph.Node(13), adjGraph.Node(16)},
			},
		},
		{
			"handling nil objects",
			args{adjGraph.ReturnType{}},
			[][]adjGraph.Node{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WelshPowell(tt.args.argsData); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WelshPowell() = %v, want %v", got, tt.want)
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


	//wanted order
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
				t.Errorf("sortNodesDescending() = %v, want %v", got, tt.want)
			}
		})
	}
}
