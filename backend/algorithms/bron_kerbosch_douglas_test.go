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
	//Output: {ABC EFG true}{EFG IJK true}{IJK MNO true}{MNO ABC true}
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
			fmt.Print(element)
		}
		fmt.Println("")
	}
	//Output: {ABC MNO true}{IJK EFG true}
	//{EFG ABC true}{MNO IJK true}
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

	//Output: {ABC EFG true}{EFG IJK true}{IJK MNO true}{MNO ABC true}
	//{ABC P1 true}{EFG P1 true}{IJK P1 true}{MNO P1 true}

	toTest := GetMaxCliques(compGraph)
	for _, index := range toTest {
		for _, element := range index {
			fmt.Print(element)
		}
		fmt.Println("")
	}
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
			fmt.Print(element)
		}
		fmt.Println("")
	}
	//Output: {ABC MNO true}
	//{MNO EFG true}
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
			fmt.Print(element)
		}
		fmt.Println("")
	}
	//Output:
	//{ABC IJK true}{IJK ABC true}
	//{EFG MNO true}{MNO EFG true}
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
			fmt.Print(element)
		}
		fmt.Println("")
	}
	//Output:
	//{ABC P1 true}{ABC P2 true}{EFG P1 true}{EFG P2 true}{IJK P1 true}{IJK P2 true}{MNO P1 true}{MNO P2 true}
	//{ABC EFG true}{EFG IJK true}{IJK MNO true}{MNO ABC true}
}

func ExampleGetMaxClique7() {
	//Test2 : Nur Linksabbieger, Fußgänger Ohne Insel
	//Input : M,A,E,I

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
			fmt.Print(element)
		}
		fmt.Println("")
	}
	//Output: {ABC MNO true}{IJK EFG true}
	//{EFG ABC true}{MNO IJK true}
	//{ABC P1 true}{EFG P1 true}{IJK P1 true}{MNO P1 true}
}
