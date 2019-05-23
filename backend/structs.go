package backend

import "github.com/ob-algdatii-ss19/leistungsnachweis-ateam/backend/adjGraph"

/**
structs for JSON response
*/
type JsonResponse struct {
	ReceivedDataSuccessful bool
	TrafficLightPhases     [][]adjGraph.Node
}

/**
structs for parsing JSON-objects
*/
type GuiRequestData struct {
	Settings     Settings
	Intersection Intersection
}

type Intersection struct {
	Top    Intersection_part
	Right  Intersection_part
	Buttom Intersection_part
	Left   Intersection_part
}

type Intersection_part struct {
	RightLane    bool
	StraightLane bool
	LeftLane     bool
	Pedestrian   Pedestrian
}

type Settings struct {
	Algorithm Algorithm
}

type Algorithm int

const (
	BASIC_GREEDY  Algorithm = 0
	WELSH_POWELL  Algorithm = 1
	BRON_KERBOSCH Algorithm = 2
)

type Pedestrian int

const (
	OFF         Pedestrian = 0
	NORMAL      Pedestrian = 1
	WITH_ISLAND Pedestrian = 2
)
