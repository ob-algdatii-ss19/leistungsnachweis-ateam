package algorithms

import (
	"fmt"
	"github.com/ob-algdatii-ss19/leistungsnachweis-ateam/backend/adjGraph"
)

/*
Calculate the optimization of the traffic lights with basic Greedy algorithm
*/

var numberOfNodes int
var countNodesPerStreet int =5;
func WelshPowell(returnType adjGraph.ReturnType) [][]adjGraph.Node {
	fmt.Println("[INFO] Called WelshPowell Algorithm")

	graphData := returnType.UGraph
	trafficEntries := returnType.Entries

	if graphData == nil || trafficEntries == nil {
		return [][]adjGraph.Node{}
	}

	numberOfNodes = graphData.UNumberOfNodes()
	var usedNodes []int


	//go through all nodes (vertices)
	for i := 1; i <= numberOfNodes; i++ {

		//user selected the node via checkbox in the gui
		if trafficEntries[i-1].IsTrue {
			usedNodes=append(usedNodes, i)
		}
	}

	fmt.Println("end of welch powell", usedNodes)
	return giveSameColor(usedNodes)
}



//make graph to []
//call sortNodesDescending

//loop over all nodes to find all nodes where the cars can drive simultaneously
	//starts with node with most edges


//returns a [][]
	//columns: all nodes
	//rows: nodes where cars can drive simultaneously
func giveSameColor(usedNodes []int) [][]adjGraph.Node{
	nodeGroupArray :=toNodeGroups(usedNodes)
	fmt.Println("node group array",nodeGroupArray )

	nodeGroupArray=sortNodesDescending(nodeGroupArray)
	fmt.Println("node group array SORTED",nodeGroupArray )


	//now we have a 2dimensional Array, order by the first inner array hase the most nodes
	result := make([][]adjGraph.Node, 0)

	for i := 0; i < numberOfNodes/countNodesPerStreet; i++ { //maximal count of traffic light phases is cound of streets
		var innerArray []adjGraph.Node;
		actArr :=nodeGroupArray[i];
		for j:= 0; j < len(actArr); j++ {

			innerArray=append(innerArray, adjGraph.Node(actArr[j]))
		}
		result = append(result, innerArray)
	}

	return result

}



//returns a 2 disemsional array with groups of node (out street)
func toNodeGroups(usedNodes []int) [][]int{

	result := make([][]int, 0)

	for i := 1; i <= numberOfNodes/countNodesPerStreet; i++ { //4 because we have 4 streets
		var innerArray []int;
		for j:= 0; j < len(usedNodes); j++ {
			if ((usedNodes[j] <= countNodesPerStreet*i) &&(usedNodes[j] > countNodesPerStreet*(i-1))) {
				innerArray = append(innerArray, usedNodes[j])
			}
		}
		result = append(result, innerArray)
	}
	return result
}



//this function make the first step and order the nodes descending
//node with most edges is first, with less edges is last
func sortNodesDescending (nodeGroupArray [][]int) [][]int{
	var result = nodeGroupArray

	for i := 0; i < len(result); i++{
		//actElementOuter :=result[i];
		for j := (i+1); j <  len(result); j++{
		//actElementInner :=result[j];
			if(len(result[j])>len(result[i])){ //swap
				var tmp=result[j]
				result[j]=result[i];
				result[i]=tmp;

			}
		}
		fmt.Println("sorting",result )
	}

	return result
}


//returns the count aof edges in a node
func countEdgesPerNode(node adjGraph.Node) int{
	return 0

}
