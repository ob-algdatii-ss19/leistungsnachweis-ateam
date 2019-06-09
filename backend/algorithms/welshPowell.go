package algorithms

import (
	"fmt"
	"github.com/ob-algdatii-ss19/leistungsnachweis-ateam/backend/adjGraph"
)

/*
Calculate the optimization of the traffic lights with basic Greedy algorithm
*/
func WelshPowell(graphData adjGraph.UGraph) [][]adjGraph.Node {
	fmt.Println("[INFO] Called welshPowell Algorithm")



	return nil
}



//make graph to []
//call sortNodesDescending

//loop over all nodes to find all nodes where the cars can drive simultaneously
	//starts with node with most edges


//returns a [][]
	//columns: all nodes
	//rows: nodes where cars can drive simultaneously
func giveSameColor(graphData adjGraph.UGraph) [][]adjGraph.Node{


	return nil

}


//this function make the first step and order the nodes descending
//node with most edges is first, with less edges is last
func sortNodesDescending (graphArray []adjGraph.Node) []adjGraph.Node{
	returnGraphArray := graphArray; //copy

	for i := 0; i < len(graphArray); i++{
		actElementOuter :=graphArray[i];
		for j := (i+1); j <  len(graphArray); j++{
		actElementInner :=graphArray[j];
			if(actElementInner>actElementOuter){ //swap
				returnGraphArray[i]=actElementInner;
				returnGraphArray[j]=actElementOuter;
			}
		}
	}

	return returnGraphArray
}
