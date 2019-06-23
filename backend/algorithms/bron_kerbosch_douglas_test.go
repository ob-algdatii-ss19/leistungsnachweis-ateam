package algorithms

import (
	"fmt"
	"github.com/ob-algdatii-ss19/leistungsnachweis-ateam/backend/adjGraph"
	"testing"
)

func ExampleGetMaxClique() {
	//Test Number 1
	//Nur Rechtsabbieger, keine Fussgaenger, eine Phase

	args := []adjGraph.Edge{
		{4, 1},
		{1, 2},
		{2, 3},
		{3, 4}}
	g := adjGraph.NewGraphAdjMat(6)
	for _, edge := range args {
		g.AddEdge(edge.From, edge.To)
	}
	graph := adjGraph.MakeConflictGraphOutOfConnectionGraph(g)
	compGraph := adjGraph.MakeCompatibilityGraph(graph)

	toTest := GetMaxCliques(compGraph)
	for _, index := range toTest {
		for _, element := range index {
			fmt.Println(element)
		}
		fmt.Println("")
	}
	//Output: 1
	//7
	//13
	//16
}

func ExampleGetMaxClique2() {
	//Test2 : Nur Linksabbieger, keine Fußgänger
	//Input : M,A,E,I

	args := []adjGraph.Edge{
		{4, 3},
		{1, 4},
		{2, 1},
		{3, 2},
	}
	g := adjGraph.NewGraphAdjMat(6)
	for _, edge := range args {
		g.AddEdge(edge.From, edge.To)
	}
	graph := adjGraph.MakeConflictGraphOutOfConnectionGraph(g)
	compGraph := adjGraph.MakeCompatibilityGraph(graph)

	toTest := GetMaxCliques(compGraph)
	for _, index := range toTest {
		for _, element := range index {
			fmt.Println(element)
		}
		fmt.Println("")
	}
	//Output: 3
	//12
	//
	//6
	//18
}

func ExampleGetMaxClique3() {
	//Test3 : Fußgänger ohne insel,  rechtsabbieger
	//Input : B,F,J,N

	args := []adjGraph.Edge{
		{4, 1},
		{1, 2},
		{2, 3},
		{3, 4},
		{1, 5},
		{2, 5},
		{3, 5},
		{4, 5},
	}
	g := adjGraph.NewGraphAdjMat(6)
	for _, edge := range args {
		g.AddEdge(edge.From, edge.To)
	}
	graph := adjGraph.MakeConflictGraphOutOfConnectionGraph(g)
	compGraph := adjGraph.MakeCompatibilityGraph(graph)

	toTest := GetMaxCliques(compGraph)
	for _, index := range toTest {
		for _, element := range index {
			fmt.Println(element)
		}
		fmt.Println("")
	}
	//Output: 1
	//7
	//13
	//16
	//
	//4
	//9
	//14
	//19
}

func TestGetMaxClique4(t *testing.T) {
	//Ein Linksabbieger und ein Geradeausfahrer -  die Spuren Kreuzen sich
	// A, N

	args := []adjGraph.Edge{
		{1, 4},
		{4, 2},
	}
	g := adjGraph.NewGraphAdjMat(6)
	for _, edge := range args {
		g.AddEdge(edge.From, edge.To)
	}
	graph := adjGraph.MakeConflictGraphOutOfConnectionGraph(g)
	compGraph := adjGraph.MakeCompatibilityGraph(graph)

	toTest := GetMaxCliques(compGraph)
	for _, index := range toTest {
		for _, element := range index {
			fmt.Println(element)
		}
		fmt.Println("")
	}
	//Output: 3
	//
	//17
}

func ExampleGetMaxClique5() {
	//Test5 : keine Fußgänger, nur geradeausfahrer
	//Input : B,F,J,N

	args := []adjGraph.Edge{
		{1, 3},
		{2, 4},
		{3, 1},
		{4, 2},
	}
	g := adjGraph.NewGraphAdjMat(6)
	for _, edge := range args {
		g.AddEdge(edge.From, edge.To)
	}
	graph := adjGraph.MakeConflictGraphOutOfConnectionGraph(g)
	compGraph := adjGraph.MakeCompatibilityGraph(graph)

	toTest := GetMaxCliques(compGraph)
	for _, index := range toTest {
		for _, element := range index {
			fmt.Println(element)
		}
		fmt.Println("")
	}
	//Output:
	//2
	//11
	//
	//8
	//17
}

func ExampleGetMaxClique6() {
	//Test6 : Fußgänger mit Insel, Rechtsabbieger
	//Input : B,F,J,N

	args := []adjGraph.Edge{
		{4, 1},
		{1, 2},
		{2, 3},
		{3, 4},
		{1, 5},
		{2, 5},
		{3, 5},
		{4, 5},
		{1, 6},
		{2, 6},
		{3, 6},
		{4, 6},
	}
	g := adjGraph.NewGraphAdjMat(6)
	for _, edge := range args {
		g.AddEdge(edge.From, edge.To)
	}
	graph := adjGraph.MakeConflictGraphOutOfConnectionGraph(g)
	compGraph := adjGraph.MakeCompatibilityGraph(graph)

	toTest := GetMaxCliques(compGraph)
	for _, index := range toTest {
		for _, element := range index {
			fmt.Println(element)
		}
		fmt.Println("")
	}
	//Output:
	//4
	//5
	//9
	//10
	//14
	//15
	//19
	//20
	//
	//1
	//7
	//13
	//16
}

func ExampleGetMaxClique7() {
	//Test7 : Linksabbieger, Fußgänger Ohne Insel

	args := []adjGraph.Edge{
		{4, 3},
		{1, 4},
		{2, 1},
		{3, 2},
		{1, 5},
		{2, 5},
		{3, 5},
		{4, 5},
	}
	g := adjGraph.NewGraphAdjMat(6)
	for _, edge := range args {
		g.AddEdge(edge.From, edge.To)
	}
	graph := adjGraph.MakeConflictGraphOutOfConnectionGraph(g)
	compGraph := adjGraph.MakeCompatibilityGraph(graph)

	toTest := GetMaxCliques(compGraph)
	for _, index := range toTest {
		for _, element := range index {
			fmt.Println(element)
		}
		fmt.Println("")
	}
	//Output:4
	//9
	//14
	//19
	//
	//3
	//12
	//
	//6
	//18
}

func ExampleGetMaxClique8() {
	//Test7 : alles

	args := []adjGraph.Edge{
		{1, 2},
		{1, 3},
		{1, 4},
		{1, 5},
		{1, 6},
		{2, 1},
		{2, 3},
		{2, 4},
		{2, 5},
		{2, 6},
		{3, 1},
		{3, 2},
		{3, 4},
		{3, 5},
		{3, 6},
		{4, 1},
		{4, 2},
		{4, 3},
		{4, 5},
		{4, 6},
	}
	g := adjGraph.NewGraphAdjMat(6)
	for _, edge := range args {
		g.AddEdge(edge.From, edge.To)
	}
	graph := adjGraph.MakeConflictGraphOutOfConnectionGraph(g)
	compGraph := adjGraph.MakeCompatibilityGraph(graph)

	toTest := GetMaxCliques(compGraph)
	for _, index := range toTest {
		for _, element := range index {
			fmt.Println(element)
		}
		fmt.Println("")
	}
	//Output:4
	//5
	//9
	//10
	//14
	//15
	//19
	//20
	//
	//1
	//2
	//3
	//16
	//
	//7
	//11
	//12
	//13
	//
	//6
	//8
	//
	//17
	//18
}

func ExampleGetMaxClique9() {
	//Test9

	args := []adjGraph.Edge{
		{1, 2},
		{1, 3},
		{1, 5},
		{1, 6},
		{2, 1},
		{2, 3},
		{2, 5},
		{2, 6},
		{3, 1},
		{3, 2},
		{3, 5},
		{3, 6},
	}
	g := adjGraph.NewGraphAdjMat(6)
	for _, edge := range args {
		g.AddEdge(edge.From, edge.To)
	}
	graph := adjGraph.MakeConflictGraphOutOfConnectionGraph(g)
	compGraph := adjGraph.MakeCompatibilityGraph(graph)

	toTest := GetMaxCliques(compGraph)
	for _, index := range toTest {
		for _, element := range index {
			fmt.Println(element)
		}
		fmt.Println("")
	}
	//Output:4
	//5
	//9
	//10
	//14
	//15
	//
	//1
	//2
	//11
	//
	//6
	//7
	//
	//12
}
