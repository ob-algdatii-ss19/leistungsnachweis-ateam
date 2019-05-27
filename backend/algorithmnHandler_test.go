package backend

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"reflect"
	"testing"

	"github.com/ob-algdatii-ss19/leistungsnachweis-ateam/backend/adjGraph"
)

func TestAdjMat(t *testing.T) {

	//read json-data from request
	//<nil> <nil> false true {0 0} false false false
	//{{0} {{false true true 0} {false false false 0} {false true false 0} {false false false 0}}}
	var jsonString = "{{0} {{false false false 0} {false false false 0} {true false false 0} {true false true 0}}}"
	var r = ioutil.NopCloser(bytes.NewReader([]byte(jsonString))) // r type is io.ReadCloser

	decoder := json.NewDecoder(r)
	var receivedData GuiRequestData
	decoder.Decode(&receivedData)
	graphObject := buildGraphObjectFromJSON(receivedData)

	//graphObject="[[] [false false false false false] [false false false true false] [false true true false false] [false false false false false]]";

	fmt.Println("test output SHOULD BE", "[[] [false false false false false] [false false false true false] [false true true false false] [false false false false false]]")
	fmt.Println("test output in TestAdjMat", graphObject)
}

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

	sutResult01 := adjGraph.NewGraphAdjMat(3)
	sutResult01.AddEdge(0, 1)
	sutResult01.AddEdge(1, 3)

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
			args{sut01},
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
