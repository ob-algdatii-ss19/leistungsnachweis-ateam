package adjGraph

import (
	"fmt"
	"testing"
)

func TestMakeConflictGraphOutOfConnectionGraph(t *testing.T) {
	//Test1 : Nur Rechtsabbieger, keine Fußgänger
	//Input : O,C,G,K

	wantedEdges1 := []Edge{
		{4, 1},
		{1, 2},
		{2, 3},
		{3, 4},
	}

	g := NewGraphAdjMat(6)
	for _, edge := range wantedEdges1 {
		g.AddEdge(edge.from, edge.to)
	}

	//	gotEdges := g.Edges()
	m := MakeConflictGraphOutOfConnectionGraph(g)
	gotEdges := m.UEdges()

	if len(gotEdges) != 0 {
		t.Errorf("Rechtsabbieger sollten keinen Konflikt aufrufen")
	}
}

func TestMakeConflictGraphOutOfConnectionGraph2(t *testing.T) {
	//Test2 : Nur Linksabbieger, keine Fußgänger
	//Input : M,A,E,I

	wantedEdges1 := []Edge{
		{4, 3},
		{1, 4},
		{2, 1},
		{3, 2},
	}

	g := NewGraphAdjMat(6)
	for _, edge := range wantedEdges1 {
		g.AddEdge(edge.from, edge.to)
	}

	//	gotEdges := g.Edges()
	m := MakeConflictGraphOutOfConnectionGraph(g)
	gotEdges := m.UEdges()

	expected := []Edge{{3, 6},
		{3, 18},
		{6, 3},
		{6, 12},
		{12, 6},
		{12, 18},
		{18, 3},
		{18, 12}}
	for index, element := range gotEdges {
		if expected[index] != element {
			t.Errorf("Unexpected Node")
		}
	}
}

func TestMakeConflictGraphOutOfConnectionGraph3(t *testing.T) {
	//Test3 : keine Fußgänger, zwei Rechtsabbieger,  1 Gerade aus Fahrer, 2 Linksabbieger
	//Input : E,G,I,B,C

	wantedEdges1 := []Edge{
		{1, 2},
		{1, 3},
		{2, 1},
		{2, 3},
		{3, 2},
	}

	g := NewGraphAdjMat(6)
	for _, edge := range wantedEdges1 {
		g.AddEdge(edge.from, edge.to)
	}

	//	gotEdges := g.Edges()
	m := MakeConflictGraphOutOfConnectionGraph(g)
	gotEdges := m.UEdges()

	expected := []Edge{{2, 6}, {2, 12}, {6, 2}, {6, 12}, {12, 2}, {12, 6}}

	for index, element := range gotEdges {
		if expected[index] != element {
			t.Errorf("Wrong Node")
		}
	}
}

/*2 ABC IJK
6 EFG ABC
12 IJK EFG
*/

func TestMakeConflictGraphOutOfConnectionGraph5(t *testing.T) {
	//Test5 : keine Fußgänger, nur geradeausfahrer
	//Input : B,F,J,N

	wantedEdges1 := []Edge{
		{1, 3},
		{2, 4},
		{3, 1},
		{4, 2},
	}

	g := NewGraphAdjMat(6)
	for _, edge := range wantedEdges1 {
		g.AddEdge(edge.from, edge.to)
	}

	//	gotEdges := g.Edges()
	m := MakeConflictGraphOutOfConnectionGraph(g)
	gotEdges := m.UEdges()

	expected := []Edge{{2, 8},
		{2, 11},
		{2, 17},
		{8, 2},
		{8, 11},
		{8, 17},
		{11, 2},
		{11, 8},
		{11, 17},
		{17, 2},
		{17, 8},
		{17, 11}}

	for index, element := range gotEdges {
		if expected[index] != element {
			t.Errorf("Wrong Node")
		}
	}
}

func TestMakeConflictGraphOutOfConnectionGraph6(t *testing.T) {
	//Test5 : Fußgänger ohne insel,  rechtsabbieger
	//Input : B,F,J,N

	wantedEdges1 := []Edge{
		{4, 1},
		{1, 2},
		{2, 3},
		{3, 4},
		{1, 5},
		{2, 5},
		{3, 5},
		{4, 5},
	}

	g := NewGraphAdjMat(6)
	for _, edge := range wantedEdges1 {
		g.AddEdge(edge.from, edge.to)
	}

	//	gotEdges := g.Edges()
	m := MakeConflictGraphOutOfConnectionGraph(g)
	gotEdges := m.UEdges()

	/*	expected := []Edge{{1, 9},
		{4, 16},
		{7, 14},
		{9, 1},
		{13, 19},
		{14, 7},
		{16, 4},
		{19, 13}}
	*/
	/*for index, element := range gotEdges {
		if expected[index] != element {
			t.Errorf("Wrong Node")
		}
	}*/

	for _, element := range gotEdges {
		fmt.Println(element)
	}
}
