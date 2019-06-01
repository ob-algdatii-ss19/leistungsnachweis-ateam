package adjGraph

type nodeName string
type ConflictGraph UGraph

type trafficEntry struct {
	from   nodeName
	to     nodeName
	isTrue bool
}

type ReturnType struct {
	Entries []trafficEntry
	UGraph  ConflictGraph
}

func MakeConflictGraphOutOfConnectionGraph(connectionGraph AdjMat) ReturnType {
	var Entries []trafficEntry = makeList(connectionGraph)
	var conflictGraph UGraph = NewUGraph(20)
	for i := 0; i < 20; i++ {
		for j := i + 1; j < 20; j++ {
			if Entries[i].isTrue && Entries[j].isTrue {
				if i%5 == 3 { //Pruefung auf Fu-g'nger bei i
					if j%5 != 3 && j%5 != 4 { //Pruefung auf weiteren Fussgaenger- zwei Fussgaenger koennen zueinander nicht im Konflikt stehen
						if Entries[i+1].isTrue { //mit Fussgaengerinsel
							if Entries[i].from == Entries[j].from {
								conflictGraph.UAddEdge(Node(i+1), Node(j+1))
							}
						} else { //ohne Fu-g'nger Insel
							if Entries[i].from == Entries[j].to {
								conflictGraph.UAddEdge(Node(i+1), Node(j+1))
							}
							if Entries[i].from == Entries[j].from {
								conflictGraph.UAddEdge(Node(i+1), Node(j+1))
							}
						}
					}
				} else if j%5 == 3 { //Prufung der Fussg'nger bei j
					if i%5 != 3 && i%5 != 4 { //Pruefung auf weiteren Fussgaenger- zwei Fussgaenger koennen zueinander nicht im Konflikt stehen
						if Entries[j+1].isTrue { //mit Fussgaengerinsel
							if Entries[i].from == Entries[j].from {
								conflictGraph.UAddEdge(Node(i+1), Node(j+1))
							}
						} else { //Ohne Fussgaengerinsel
							if Entries[j].from == Entries[i].to {
								conflictGraph.UAddEdge(Node(j+1), Node(i+1))
							}
							if Entries[j].from == Entries[i].from {
								conflictGraph.UAddEdge(Node(i+1), Node(j+1))
							}

						}
					}
				} else if i%5 == 4 { //Falls Fussgaengerinsel ueberhaupt vorhanden
					if j%5 != 3 && j%5 != 4 { //Pruefung auf weiteren Fussgaenger- zwei Fussgaenger koennen zueinander nicht im Konflikt stehen
						if Entries[i].isTrue { //Falls Fussgaengerinsel ueberhaupt vorhanden
							if Entries[i].from == Entries[j].to {
								conflictGraph.UAddEdge(Node(i+1), Node(j+1))
							}
						}
					}
				} else if j%5 == 4 { //Falls Fussgaengerinsel ueberhaupt vorhanden
					if i%5 != 3 && i%5 != 4 { //Pruefung auf weiteren Fussgaenger- zwei Fussgaenger koennen zueinander nicht im Konflikt stehen
						if Entries[j].isTrue {
							if Entries[j].from == Entries[i].to {
								conflictGraph.UAddEdge(Node(i+1), Node(j+1))
							}
						}
					}
				} else {
					if i != 0 && j != 0 && i != 6 && j != 6 && i != 12 && j != 12 && i != 15 && j != 15 { //Abfangen der Rechtsabbieger
						//if i != 1 && j != 1 && i != 7 && j != 7 && i != 13 && j != 13 && i != 16 && j != 16 { //Abfangen der Rechtsabbieger
						if !((i == 2 && j == 11) || (i == 5 && j == 17)) { //kein Linksabbieger Paar 2&11; 5&17
							if Entries[i].from != Entries[j].from {
								if Entries[i].to != Entries[j].to {
									conflictGraph.UAddEdge(Node(i+1), Node(j+1))
									//conflictGraph.AddEdge(Node(j),Node(i))  //der Graph ist ja Symmetrisch
								}
							}
						}
					}
				}
			}
		}
	}
	tmp := ReturnType{Entries, conflictGraph}
	return tmp
}

func makeList(matrix AdjMat) []trafficEntry {
	//var Entries []trafficEntry
	var Entries = make([]trafficEntry, 20)
	var everyNode = [6]string{"ABC", "EFG", "IJK", "MNO", "P1", "P2"}
	//Entries[0] = trafficEntry{"ABC","EFG",false}
	var counter = 0
	for i := 1; i <= len(everyNode)-2; i++ {
		for j := 1; j <= len(everyNode); j++ {
			if i != j {
				//fmt.Println("i= ",i," ,j= ",j," ,counter= ",counter)
				var isEdge = matrix[i][j]
				//fmt.Println("inserted in Matrix")
				Entries[counter] = trafficEntry{nodeName(everyNode[i-1]), nodeName(everyNode[j-1]), isEdge}
				//fmt.Println("inserted in Entriy list")
				counter++
			}
		}
	}
	return Entries
}
