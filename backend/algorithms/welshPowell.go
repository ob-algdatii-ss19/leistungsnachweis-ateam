package algorithms

import (
	"fmt"
	"github.com/ob-algdatii-ss19/leistungsnachweis-ateam/backend/adjGraph"
)

/*
Calculate the optimization of the traffic lights with WelshPowell algorithm
*/

var numberOfNodes int =20 //20 only important for test!
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


		//in wich streets are cars drive in?
		var streetsWithInsertions []bool
		for k := 0; k < numberOfNodes/countNodesPerStreet; k++ { //maximal count of traffic light phases is cound of streets
			streetsWithInsertions = append(streetsWithInsertions, false) //actul no insertions
		}
		//[0]=ABC, [1]=EFG, [2]=KJI, [3]=MNO

		if(len(nodeGroupArray[i])==0){
			/*for iTest := i; iTest < numberOfNodes/countNodesPerStreet; iTest++ { //fulltest, maybe another index has values
				if(len(nodeGroupArray[iTest])>0){
					i=iTest;
					break
				}
			}*/
			fmt.Println("all 'bigger' indices already could drive");
			continue
		}
		//calculate actIndex
		var actIndex int =0;
		for iTest := 1; iTest <= numberOfNodes/countNodesPerStreet; iTest++ {
			if ((nodeGroupArray[i][0] <= countNodesPerStreet*iTest) &&(nodeGroupArray[i][0] > countNodesPerStreet*(iTest-1))) {
				actIndex=iTest-1; //-1 because starts with 1
			}
		}

		fmt.Println("actIndex", actIndex)

			//actArr :=nodeGroupArray[i];
		//l :=len(nodeGroupArray[i])
		for j:= 0; j < len(nodeGroupArray[i]); j++ {



			actVal :=nodeGroupArray[i][j]

			//right
			if(actVal==1||actVal==7||actVal==13||actVal==16){
				var ind =giveIndex(4, actIndex, 1) //right is adding 1
				if(streetsWithInsertions[ind]==false){
					streetsWithInsertions[ind]=true;
					innerArray=append(innerArray, adjGraph.Node(nodeGroupArray[i][j]))
					nodeGroupArray[i]=remove(nodeGroupArray[i], j)
					j--;
				}

			}
			//straight
			if(actVal==2||actVal==8||actVal==11||actVal==17){
				var ind =giveIndex(4, actIndex, 2) //straight is adding 2
				if(streetsWithInsertions[ind]==false){
					streetsWithInsertions[ind]=true;
					innerArray=append(innerArray, adjGraph.Node(nodeGroupArray[i][j]))
					nodeGroupArray[i]=remove(nodeGroupArray[i], j)
					j--;
				}
			}
			//left
			if(actVal==3||actVal==6||actVal==12||actVal==18){
				var ind =giveIndex(4, actIndex, 3) //left is adding 3
				if(streetsWithInsertions[ind]==false){
					streetsWithInsertions[ind]=true;
					innerArray=append(innerArray, adjGraph.Node(nodeGroupArray[i][j]))
					nodeGroupArray[i]=remove(nodeGroupArray[i], j)
					j--;
				}
			}
		}
		//all for this street are finish now
		fmt.Println("nach erster schleife",nodeGroupArray , streetsWithInsertions)


		//are there other streets which can drive simultaneously?
		//right: 1, 7, 13, 16
		//straight: 2,8,11,17
		//left: 3,6,12,18


		//iterate over all, when find possible street--> add this
		for inner := i+1; inner < numberOfNodes/countNodesPerStreet; inner++ { //i+1 because all streets with smaller index already deleted
			actIndex =inner;
			if(len(nodeGroupArray[inner])>0){
				for iTest := 1; iTest <= numberOfNodes/countNodesPerStreet; iTest++ {
					if ((nodeGroupArray[inner][0] <= countNodesPerStreet*iTest) &&(nodeGroupArray[inner][0] > countNodesPerStreet*(iTest-1))) {
						actIndex=iTest-1; //-1 because starts with 1
					}
				}
			}




			for j:= 0; j < len(nodeGroupArray[inner]); j++ {



				actVal :=nodeGroupArray[inner][j]

				//right
				if(actVal==1||actVal==7||actVal==13||actVal==16){
					var ind =giveIndex(4, actIndex, 1) //right is adding 1
					fmt.Println("RIGHT inner, j, ind", nodeGroupArray[inner][j], ind, streetsWithInsertions[ind])
					if(streetsWithInsertions[ind]==false){
						streetsWithInsertions[ind]=true;
						innerArray=append(innerArray, adjGraph.Node(nodeGroupArray[inner][j]))
						nodeGroupArray[inner]=remove(nodeGroupArray[inner], j)
						j--;
					}

				}
				//straight
				if(actVal==2||actVal==8||actVal==11||actVal==17){
					var ind =giveIndex(4, actIndex, 2) //straight is adding 2
					fmt.Println("STRIAGHT inner, j, ind", nodeGroupArray[inner][j], ind, streetsWithInsertions[ind])

					if(streetsWithInsertions[ind]==false){
						streetsWithInsertions[ind]=true;
						innerArray=append(innerArray, adjGraph.Node(nodeGroupArray[inner][j]))
						nodeGroupArray[inner]=remove(nodeGroupArray[inner], j)
						j--;
					}
				}
				//left
				if(actVal==3||actVal==6||actVal==12||actVal==18){
					var ind =giveIndex(4, actIndex, 3) //left is adding 3
					fmt.Println("LEFT inner, j, ind", nodeGroupArray[inner][j], ind, streetsWithInsertions[ind])
					if(streetsWithInsertions[ind]==false){
						streetsWithInsertions[ind]=true
						innerArray=append(innerArray, adjGraph.Node(nodeGroupArray[inner][j]))
						nodeGroupArray[inner]=remove(nodeGroupArray[inner], j)
						j--;
					}
				}
				fmt.Println("index array",actVal,nodeGroupArray )

			}
		}
		fmt.Println("nach erstem Komplettdurchlauf",nodeGroupArray )



		result = append(result, innerArray)
	}
	fmt.Println("node group array nach delete",nodeGroupArray )


	return result

}

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

//returns the index, this edge has in the actual index (ABC, EFG, IJK, MNO)
func giveIndex(arrayLenght int, actIndex int, adding int) int{
	var retInt int=actIndex+adding;
	if(retInt>=arrayLenght){
		retInt=retInt-arrayLenght
	}
	return retInt
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


//siehe http://mrsleblancsmath.pbworks.com/w/file/fetch/46119304/vertex%20coloring%20algorithm.pdf]
//Konfliktgraph
//drüber nicht erforderlich, siehe ugraph::UAdj
//Knoten mit meisten Konflikten fahren zuerst
//keine Gruppen bilden!!!
