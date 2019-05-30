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

	//starts with 1, not 0
	sutResult01 := adjGraph.NewGraphAdjMat(6)

	//top
	sutResult01.AddEdge(1, 2)
	sutResult01.AddEdge(1, 3)
	sutResult01.AddEdge(1, 4)
	sutResult01.AddEdge(1, 5) //because Pedestrains


	//left
	sutResult01.AddEdge(2, 1)
	sutResult01.AddEdge(2, 3)
	sutResult01.AddEdge(2, 4)
	sutResult01.AddEdge(2, 5) //because Pedestrains


	//bottom
	sutResult01.AddEdge(3, 1)
	sutResult01.AddEdge(3, 2)
	sutResult01.AddEdge(3, 4)
	sutResult01.AddEdge(3, 5) //because Pedestrains


	//right
	sutResult01.AddEdge(4, 1)
	sutResult01.AddEdge(4, 2)
	sutResult01.AddEdge(4, 3)
	sutResult01.AddEdge(4, 5) //because Pedestrains


	//fmt.Println("[DEBUG] graph in Test", sutResult01)



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
				Pedestrian:   WITH_ISLAND,
			},
			Buttom: Intersection_part{
				RightLane:    true,
				StraightLane: true,
				LeftLane:     true,
				Pedestrian:   WITH_ISLAND,
			},
			Right: Intersection_part{
				RightLane:    true,
				StraightLane: true,
				LeftLane:     true,
				Pedestrian:   WITH_ISLAND,
			},
		},
	}
	//starts with 1, not 0
	sutResult02 := adjGraph.NewGraphAdjMat(6)

	//top
	sutResult02.AddEdge(1, 2)
	sutResult02.AddEdge(1, 3)
	sutResult02.AddEdge(1, 4)
	sutResult02.AddEdge(1, 5) //because Pedestrains
	sutResult02.AddEdge(1, 6) //because Pedestrain Island


	//left
	sutResult02.AddEdge(2, 1)
	sutResult02.AddEdge(2, 3)
	sutResult02.AddEdge(2, 4)
	sutResult02.AddEdge(2, 5) //because Pedestrains
	sutResult02.AddEdge(2, 6) //because Pedestrain Island


	//bottom
	sutResult02.AddEdge(3, 1)
	sutResult02.AddEdge(3, 2)
	sutResult02.AddEdge(3, 4)
	sutResult02.AddEdge(3, 5) //because Pedestrains
	sutResult02.AddEdge(3, 6) //because Pedestrain Island


	//right
	sutResult02.AddEdge(4, 1)
	sutResult02.AddEdge(4, 2)
	sutResult02.AddEdge(4, 3)
	sutResult02.AddEdge(4, 5) //because Pedestrains
	sutResult02.AddEdge(4, 6) //because Pedestrain Island


	//fmt.Println("[DEBUG] graph in Test", sutResult02)


	sut03 := GuiRequestData{
		Settings{Algorithm: BASIC_GREEDY},
		Intersection{
			Left: Intersection_part{
				LeftLane:    false,
				StraightLane: true ,
				RightLane:     false,
				Pedestrian:   OFF,
			},
			Buttom: Intersection_part{
				LeftLane:    false,
				StraightLane: false,
				RightLane:     true,
				Pedestrian:   OFF,
			},
			Right: Intersection_part{
				LeftLane:    false,
				StraightLane: false,
				RightLane:     false,
				Pedestrian:   NORMAL,
			},
			Top: Intersection_part{
				LeftLane:    true,
				StraightLane: false,
				RightLane:     false,
				Pedestrian:    OFF,
			},
		},
	}
	//starts with 1, not 0
	sutResult03 := adjGraph.NewGraphAdjMat(6)

	//left
	//sutResult03.AddEdge(1, 2)
	sutResult03.AddEdge(1, 3)
	//sutResult03.AddEdge(1, 4)
	//sutResult03.AddEdge(1, 5)
	//sutResult03.AddEdge(1, 6)

	//bottom
	//sutResult03.AddEdge(2, 1)
	sutResult03.AddEdge(2, 3)
	//sutResult03.AddEdge(2, 4)
	//sutResult03.AddEdge(2, 5)
	//sutResult03.AddEdge(2, 6)

	//right
	//sutResult03.AddEdge(3, 1)
	//sutResult03.AddEdge(3, 2)
	//sutResult03.AddEdge(3, 4)
	sutResult03.AddEdge(3, 5)
	//sutResult03.AddEdge(3, 6)

	//top
	//sutResult03.AddEdge(4, 1)
	//sutResult03.AddEdge(4, 2)
	sutResult03.AddEdge(4, 3)
	//sutResult03.AddEdge(4, 5)
	//sutResult03.AddEdge(4, 6)

	fmt.Println("[DEBUG] guidata3", buildGraphObjectFromJSON(sut03))
	fmt.Println("[DEBUG] grap03h in Test", sutResult03)

	type args struct {
		data GuiRequestData
	}
	tests := []struct {
		name string
		args args
		want adjGraph.Graph
	}{
		{
			"first test: all ways, normal pedestrians",
			args{sut01},
			sutResult01,
		},
		{
			"second test: all ways, pedestrians with islands",
			args{sut02},
			sutResult02,
		},
		{
			"third test: some ways, 2 normal pedestrians, 2 pedestrians with island",
			args{sut03},
			sutResult03,
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
