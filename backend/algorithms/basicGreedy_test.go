package algorithms

import (
	"reflect"
	"testing"

	"github.com/ob-algdatii-ss19/leistungsnachweis-ateam/backend/adjGraph"
)

func TestBasicGreedy(t *testing.T) {
	type args struct {
		graphData adjGraph.Graph
	}
	tests := []struct {
		name string
		args args
		want adjGraph.Graph
	}{
		// TODO: Add test cases.
		{
			"simple first test",
			args{adjGraph.NewGraphAdjMat(2)},
			adjGraph.NewGraphAdjMat(2),
		},
		{
			"simple first test",
			args{adjGraph.NewGraphAdjMat(2)},
			adjGraph.NewGraphAdjMat(2),
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
