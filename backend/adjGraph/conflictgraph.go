package adjGraph

type nodeName string

type TrafficEntry struct {
	From         nodeName
	To           nodeName
	ChosenByUser bool
}

type ConflictGraphPackage struct {
	Entries       []TrafficEntry
	ConflictGraph UGraph
}

func GetIndexInConflictGraph(graphPackage ConflictGraphPackage, entry TrafficEntry) int {
	rightIndex := -1
	for index, element := range graphPackage.Entries {
		if entry.To == element.To && entry.From == element.From {
			rightIndex = index
			break
		}
	}
	return rightIndex + 1
}

const ABC = "ABC"
const EFG = "EFG"
const IJK = "IJK"
const MNO = "MNO"
const P = "P"
const P1 = "P1"
const P2 = "P2"

func MakeConflictGraphOutOfConnectionGraph(connectionGraph AdjMat) ConflictGraphPackage {
	var Entries []TrafficEntry = makeList(connectionGraph)
	var conflictGraph UGraph = NewUGraph(20)
	for i := 0; i < 20; i++ {
		for j := i + 1; j < 20; j++ {
			if Entries[i].ChosenByUser && Entries[j].ChosenByUser {
				if i%5 == 3 { //Pruefung auf Fu-g'nger bei i
					if j%5 != 3 && j%5 != 4 { //Pruefung auf weiteren Fussgaenger- zwei Fussgaenger koennen zueinander nicht im Konflikt stehen
						if Entries[i+1].ChosenByUser { //mit Fussgaengerinsel
							if Entries[i].From == Entries[j].From {
								conflictGraph.UAddEdge(Node(i+1), Node(j+1))
							}
						} else { //ohne Fu-g'nger Insel
							if Entries[i].From == Entries[j].To {
								conflictGraph.UAddEdge(Node(i+1), Node(j+1))
							}
							if Entries[i].From == Entries[j].From {
								conflictGraph.UAddEdge(Node(i+1), Node(j+1))
							}
						}
					}
				} else if j%5 == 3 { //Prufung der Fussg'nger bei j
					if i%5 != 3 && i%5 != 4 { //Pruefung auf weiteren Fussgaenger- zwei Fussgaenger koennen zueinander nicht im Konflikt stehen
						if Entries[j+1].ChosenByUser { //mit Fussgaengerinsel
							if Entries[i].From == Entries[j].From {
								conflictGraph.UAddEdge(Node(i+1), Node(j+1))
							}
						} else { //Ohne Fussgaengerinsel
							if Entries[j].From == Entries[i].To {
								conflictGraph.UAddEdge(Node(j+1), Node(i+1))
							}
							if Entries[j].From == Entries[i].From {
								conflictGraph.UAddEdge(Node(i+1), Node(j+1))
							}

						}
					}
				} else if i%5 == 4 { //Falls Fussgaengerinsel ueberhaupt vorhanden
					if j%5 != 3 && j%5 != 4 { //Pruefung auf weiteren Fussgaenger- zwei Fussgaenger koennen zueinander nicht im Konflikt stehen
						if Entries[i].ChosenByUser { //Falls Fussgaengerinsel ueberhaupt vorhanden
							if Entries[i].From == Entries[j].To {
								conflictGraph.UAddEdge(Node(i+1), Node(j+1))
							}
						}
					}
				} else if j%5 == 4 { //Falls Fussgaengerinsel ueberhaupt vorhanden
					if i%5 != 3 && i%5 != 4 { //Pruefung auf weiteren Fussgaenger- zwei Fussgaenger koennen zueinander nicht im Konflikt stehen
						if Entries[j].ChosenByUser {
							if Entries[j].From == Entries[i].To {
								conflictGraph.UAddEdge(Node(i+1), Node(j+1))
							}
						}
					}
				} else {

					if i != 0 && j != 0 && i != 6 && j != 6 && i != 12 && j != 12 && i != 15 && j != 15 { //Abfangen der Rechtsabbieger
						//if !((i == 0 || i == 6 || i == 12 || i == 15) && (j == 0 || j == 6 || j == 12 || j == 15)) { //Zwei Rechtsabbieger stehen nicht im Konflikt
						if !((i == 2 && j == 11) || (i == 5 && j == 17)) { //kein Linksabbieger Paar 2&11; 5&17
							if Entries[i].From != Entries[j].From {
								if !((Entries[i].From == Entries[j].To) && (Entries[i].To == Entries[j].From)) {
									conflictGraph.UAddEdge(Node(i+1), Node(j+1))
								}
							}
						}
					}
					if Entries[i].To == Entries[j].To {
						conflictGraph.UAddEdge(Node(i+1), Node(j+1))
					}
				}
			}
		}
	}
	tmp := ConflictGraphPackage{Entries, conflictGraph}
	return tmp
}

func makeList(matrix AdjMat) []TrafficEntry {
	//var Entries []trafficEntry
	var Entries = make([]TrafficEntry, 20)
	var everyNode = [6]string{ABC, EFG, IJK, MNO, P1, P2}
	//Entries[0] = trafficEntry{"ABC","EFG",false}
	var counter = 0
	for i := 1; i <= len(everyNode)-2; i++ {
		for j := 1; j <= len(everyNode); j++ {
			if i != j {
				//fmt.Println("i= ",i," ,j= ",j," ,counter= ",counter)
				var isEdge = matrix[i][j]
				//fmt.Println("inserted in Matrix")
				Entries[counter] = TrafficEntry{nodeName(everyNode[i-1]), nodeName(everyNode[j-1]), isEdge}
				counter++
			}
		}
	}
	return Entries
}

func MakeCompatibilityGraph(conflictReturn ConflictGraphPackage) ConflictGraphPackage {
	var compGraph UGraphMat = NewUGraph(20)
	for i := 0; i < len(conflictReturn.Entries); i++ {
		for j := i + 1; j < len(conflictReturn.Entries); j++ {
			if conflictReturn.Entries[i].ChosenByUser && conflictReturn.Entries[j].ChosenByUser {
				if !conflictReturn.ConflictGraph.GetMatrixEntryAtIndex(i+1, j+1) { //wenn es eine kante hat -> Kriegt keine Kante
					compGraph.UAddEdge(Node(i+1), Node(j+1))
				}
			}

		}
	}
	tmp := ConflictGraphPackage{conflictReturn.Entries, compGraph}
	return tmp
}
