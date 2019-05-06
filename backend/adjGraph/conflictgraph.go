package adjGraph

import _ "Strings"


//Annotation: Input: Verbindungsgraph + Liste mit Knotennamen des Verbindungsgraph

func MakeConflictGraphOutOfConnectionGraph(connectionGraph AdjMat) AdjMat{
	//var name_List = [4]string{"ABC","EFG","KIJ","MNO"}
	//var turn_righties = [1]Edge{Edge(Node(4),Node(1))}

	//Alle trues [ber Edges herausfinden
	var connections = connectionGraph.Edges()
	//remove edge f[r rechtsabbieger

	var conflictGraph = NewGraphAdjMat(len(connections))
	//var conflictGraphNodeName =
	for i:=0; i<len(conflictGraph);i++ {
		for j:=i+1; j<len(conflictGraph);j++ {
			//Rechtsabbieger abfangen
			var diff1 = connections[j].from - connections[i].from
			var diff2 = connections[j].to - connections[i].to
			if diff1!= 1 &&diff2 != 1 {
				//Linksabbieger abfangen:
				//if(linksabbieger()) {
					if (connections[i].from != connections[j].from) && (connections[i].to != connections[j].to) {
						conflictGraph.AddEdge(Node(i), Node(j))
					}
				//}
			}
		}
	}
	//remove edge f[r linksabbieger

	return conflictGraph

}

func linksabbieger(edge1 Edge,edge2 Edge) bool{
	var isLefty bool = false

	if((edge1.from==1)&&(edge1.to==0))&&((edge2.from==3)&&(edge2.to==2)){
		isLefty = true
	}else if((edge1.from==2)&&(edge1.to==1))&&((edge2.from==0)&&(edge2.to==3)){
		isLefty = true
	}
	return isLefty
}



//Input: JSon File
//OutPut: Adjazensmatrix mit

func MakeConnectionGraphOutOfJson()AdjMat{
	var newMatrix = NewGraphAdjMat(4)
	return newMatrix
}
