package backend

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/ob-algdatii-ss19/leistungsnachweis-ateam/backend/adjGraph"
)

func Test_buildGraphObjectFromJSON(t *testing.T) {

	sut01 := GuiRequestData{
		Settings{Algorithm: BASIC_GREEDY},
		Intersection{
			Top: Intersection_part{
				RightLane:    true,
				StraightLane: true,
				LeftLane:     true,
				Pedestrian:   NORMAL,
			},
			Left: Intersection_part{
				RightLane:    true,
				StraightLane: true,
				LeftLane:     true,
				Pedestrian:   NORMAL,
			},
			Buttom: Intersection_part{
				RightLane:    true,
				StraightLane: true,
				LeftLane:     true,
				Pedestrian:   NORMAL,
			},
			Right: Intersection_part{
				RightLane:    true,
				StraightLane: true,
				LeftLane:     true,
				Pedestrian:   NORMAL,
			},
		},
	}

	sut02 := GuiRequestData{
		Settings{Algorithm: BASIC_GREEDY},
		Intersection{
			Top: Intersection_part{
				RightLane:    true,
				StraightLane: true,
				LeftLane:     true,
				Pedestrian:   WITH_ISLAND,
			},
			Left: Intersection_part{
				RightLane:    true,
				StraightLane: true,
				LeftLane:     true,
				Pedestrian:   NORMAL,
			},
			Buttom: Intersection_part{
				RightLane:    true,
				StraightLane: true,
				LeftLane:     true,
				Pedestrian:   NORMAL,
			},
			Right: Intersection_part{
				RightLane:    true,
				StraightLane: true,
				LeftLane:     true,
				Pedestrian:   NORMAL,
			},
		},
	}

	//starts with 1, not 0
	sutResult01 := adjGraph.NewGraphAdjMat(4)

	//top
	sutResult01.AddEdge(1, 2)
	sutResult01.AddEdge(1, 3)
	sutResult01.AddEdge(1, 4)

	//left
	sutResult01.AddEdge(2, 1)
	sutResult01.AddEdge(2, 3)
	sutResult01.AddEdge(2, 4)

	//bottom
	sutResult01.AddEdge(3, 1)
	sutResult01.AddEdge(3, 2)
	sutResult01.AddEdge(3, 4)

	//right
	sutResult01.AddEdge(4, 1)
	sutResult01.AddEdge(4, 2)
	sutResult01.AddEdge(4, 3)

	fmt.Println("[DEBUG] graph in Test", sutResult01)

	type args struct {
		data GuiRequestData
	}
	tests := []struct {
		name string
		args args
		want adjGraph.Graph
	}{
		{
			"basic test for build graph object",
			args{sut01},
			sutResult01,
		},
		{
			"second basic test for build graph object",
			args{sut02},
			sutResult01,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildGraphObjectFromJSON(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildGraphObjectFromJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}
