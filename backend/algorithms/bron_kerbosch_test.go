package algorithms

import (
	"fmt"
	"github.com/ob-algdatii-ss19/leistungsnachweis-ateam/backend/adjGraph"
	"testing"
)

func TestGetAllMaxClique(t *testing.T) {
	//Test Number 1
	//Nur Rechtsabbieger, keine Fussgaenger, eine Phase

	args := []adjGraph.Edge{}
	graph := adjGraph.NewUGraph(20)
	for _, edge := range args {
		graph.UAddEdge(edge.From, edge.To)
	}
	argsForEntries := []adjGraph.Edge{{4, 1},
		{1, 2},
		{2, 3},
		{3, 4}}
	g := adjGraph.NewGraphAdjMat(6)
	for _, edge := range argsForEntries {
		g.AddEdge(edge.From, edge.To)
	}
	entryList := adjGraph.MakeList(g)
	packageC := adjGraph.ConflictGraphPackage{entryList, graph}
	compGraph := adjGraph.MakeCompatibilityGraph(packageC)

	//Output: {ABC EFG true}{EFG IJK true}{IJK MNO true}{MNO ABC true}

	toTest := getAllMaxCliques(compGraph)
	for _, index := range toTest {
		for _, element := range index {
			fmt.Print(element)
		}
		fmt.Println("")
	}
	fmt.Print(len(toTest))
}

func TestGetAllMaxClique2(t *testing.T) {
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

	//Output: {ABC MNO true}{IJK EFG true}
	//{EFG ABC true}{MNO IJK true}

	toTest := getAllMaxCliques(compGraph)
	for _, index := range toTest {
		for _, element := range index {
			fmt.Print(element)
		}
		fmt.Println("")
	}
	fmt.Print(len(toTest))
}

func TestGetAllMaxClique3(t *testing.T) {
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
	//{ABC EFG true}{EFG IJK true}{MNO P1 true}
	//{ABC EFG true}{IJK P1 true}{MNO ABC true}
	//{ABC EFG true}{IJK P1 true}{MNO P1 true}
	//{ABC P1 true}{EFG IJK true}{IJK MNO true}
	//{ABC P1 true}{EFG IJK true}{MNO P1 true}
	//{ABC P1 true}{EFG P1 true}{IJK MNO true}
	//{ABC P1 true}{EFG P1 true}{IJK P1 true}{MNO P1 true}
	//{EFG P1 true}{IJK MNO true}{MNO ABC true}
	//{EFG P1 true}{IJK P1 true}{MNO ABC true}

	toTest := getAllMaxCliques(compGraph)
	for _, index := range toTest {
		for _, element := range index {
			fmt.Print(element)
		}
		fmt.Println("")
	}
	fmt.Print(len(toTest))
}

func TestGetAllMaxClique4(t *testing.T) {
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

	//Output: {ABC EFG true}{EFG IJK true}{IJK MNO true}{MNO ABC true}
	//{ABC EFG true}{IJK P1 true}{MNO ABC true}
	//{ABC P1 true}{EFG IJK true}{IJK MNO true}
	//{ABC P1 true}{EFG P1 true}{IJK MNO true}
	//{ABC P1 true}{EFG P1 true}{IJK P1 true}{MNO P1 true}
	//{EFG P1 true}{IJK MNO true}{MNO ABC true}
	//{EFG P1 true}{IJK P1 true}{MNO ABC true}

	toTest := getAllMaxCliques(compGraph)
	for _, index := range toTest {
		for _, element := range index {
			fmt.Print(element)
		}
		fmt.Println("")
	}
	fmt.Print(len(toTest))
}
