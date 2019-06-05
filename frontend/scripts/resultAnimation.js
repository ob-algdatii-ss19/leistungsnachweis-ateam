let svgDoc;
let currentVisibleLanes;

function displayPhasesOnIntersection(trafficLightPhasesString) {

    svgDoc = document.getElementById("intersection-svg").contentDocument;

    if (currentVisibleLanes !== undefined) {
        switchVisibilityOfTrafficLaneArray(currentVisibleLanes, false);

    }

    console.log("[DEBUG] hide current visible lanes ", currentVisibleLanes);

    let trafficLightPhasesArray = trafficLightPhasesString.split(",");

    console.log("[DEBUG] display all lanes of the array ", trafficLightPhasesArray);
    switchVisibilityOfTrafficLaneArray(trafficLightPhasesArray, true);

    currentVisibleLanes = trafficLightPhasesArray;
}

function switchVisibilityOfTrafficLaneArray(trafficLightPhasesArray,switchOn){
    let styleValue;

    if(switchOn === true) {
        styleValue = "visibility:visible";
    } else {
        styleValue = "visibility:hidden";
    }

    for (let i=0; i<trafficLightPhasesArray.length; i++) {
        //console.log("[DEBUG] set style " + styleValue + " for lane " + trafficLightPhasesArray[i]);
        switchVisibilityOfTrafficLane(trafficLightPhasesArray[i], styleValue);
    }
}

let trafficLightLetterToIdMap = new Map();
trafficLightLetterToIdMap.set("A", "left-left");
trafficLightLetterToIdMap.set("B", "left-straight");
trafficLightLetterToIdMap.set("C", "left-right");
trafficLightLetterToIdMap.set("D", "tl-pedestrian-left-complete");
trafficLightLetterToIdMap.set("D1", "tl-pedestrian-left-long");
trafficLightLetterToIdMap.set("D2", "tl-pedestrian-left-short");
trafficLightLetterToIdMap.set("E", "bottom-left");
trafficLightLetterToIdMap.set("F", "bottom-straight");
trafficLightLetterToIdMap.set("G", "bottom-right");
trafficLightLetterToIdMap.set("H", "tl-pedestrian-bottom-complete");
trafficLightLetterToIdMap.set("H1", "tl-pedestrian-bottom-long");
trafficLightLetterToIdMap.set("H2", "tl-pedestrian-bottom-short");
trafficLightLetterToIdMap.set("I", "right-left");
trafficLightLetterToIdMap.set("J", "right-straight");
trafficLightLetterToIdMap.set("K", "right-right");
trafficLightLetterToIdMap.set("L", "tl-pedestrian-right-complete");
trafficLightLetterToIdMap.set("L1", "tl-pedestrian-right-long");
trafficLightLetterToIdMap.set("L2", "tl-pedestrian-right-short");
trafficLightLetterToIdMap.set("M", "top-left");
trafficLightLetterToIdMap.set("N", "top-straight");
trafficLightLetterToIdMap.set("O", "top-right");
trafficLightLetterToIdMap.set("P", "tl-pedestrian-top-complete");
trafficLightLetterToIdMap.set("P1", "tl-pedestrian-top-long");
trafficLightLetterToIdMap.set("P2", "tl-pedestrian-top-short");


function switchVisibilityOfTrafficLane(trafficLightLetter, styleValue) {

    let laneId = trafficLightLetterToIdMap.get(trafficLightLetter);
    svgDoc.getElementById(laneId).setAttribute("style", styleValue);

}
