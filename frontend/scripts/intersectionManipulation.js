var svgDoc;

window.onload=function() {

    svgDoc = document.getElementById("intersection-svg").contentDocument; // Get the SVG document inside the Object tag

    // hide all traffic-lanes
    var trafficLanes = svgDoc.getElementById("traffic-lanes");
    trafficLanes.setAttribute("style", "visibility:hidden");

    // hide all traffic-lights
    var trafficLights = svgDoc.getElementById("traffic-lights");
    trafficLights.setAttribute("style", "visibility:hidden");

    addCheckboxListener();
    addButtonListener();
};

/*
    Register Button-Listener for evaluating the intersection
*/
function addButtonListener(){
    document.getElementById("Evaluate").addEventListener("click", function(){
        console.log("[INFO] Evaluate button clicked, execute EventListener");
        sendIntersectionJSONDataToBackend();
    });
}


/*
    Checkbox-Listener to display and hide traffic-lights
    TODO change visibility of all traffic lights and refactor method
 */
function addCheckboxListener(){
    nodeList = ['node-A', 'node-B','node-C', 'node-E', 'node-F', 'node-G','node-I','node-J','node-K','node-M','node-N','node-O'];
    elementList = ["A-left-left","B-left-straight","C-left-right","E-bottom-left","F-bottom-straight","G-bottom-right","I-right-left","J-right-straight","K-right-right","M-top-left","N-top-straight","O-top-right"];
    for(let i= 0; i<nodeList.length;i++) {
        //alert(nodeList.length);
        document.getElementById(nodeList[i]).addEventListener('change', e => {

            if (e.target.checked) {
                svgDoc.getElementById(elementList[i]).setAttribute("style", "visibility:visible");
            } else {
                svgDoc.getElementById(elementList[i]).setAttribute("style", "visibility:hidden");
            }
        });
    }
    nodeListPedestrians =['pedestrians-right','pedestrians-bottom','pedestrians-left','pedestrians-top'];
    nodeListIslandPedestrians =['pedestrians-island-right','pedestrians-island-bottom','pedestrians-island-left','pedestrians-island-top'];
    elementListPedestriansOutterLights = ["L-outter","H-outter","D-outter","P-outter"];
    elementListPedestriansInnerLights = ["L-inner","H-inner","D-inner","P-inner"];
    elementListPedestriansIsland = ["pi-right","pi-bottom","pi-left","pi-top"];

    //svgDoc.getElementById("pi-left").setAttribute("style", "visibility:hidden");
    //svgDoc.getElementById(elementListPedestriansIsland[0]+"-short").setAttribute("style", "visibility:hidden");

    for(let i= 0; i<nodeListPedestrians.length;i++) {
        //alert(nodeListPedestrians[i]+" "+elementListPedestrians[i]);
        svgDoc.getElementById(elementListPedestriansIsland[i]).setAttribute("style", "visibility:hidden");
        document.getElementById(nodeListPedestrians[i]).addEventListener('change', e => {
            if (e.target.checked) {
                svgDoc.getElementById(elementListPedestriansOutterLights[i]).setAttribute("style", "visibility:visible");

                //show checkbox and text for pedestrian islands
                document.getElementById(nodeListIslandPedestrians[i]).disabled = false;
                document.getElementById(nodeListIslandPedestrians[i]).parentElement.classList.remove("inactive-checkbox-text");

                document.getElementById(nodeListIslandPedestrians[i]).addEventListener('change', e=>{
                    if(e.target.checked){
                        svgDoc.getElementById(elementListPedestriansIsland[i]).setAttribute("style", "visibility:visible");
                        svgDoc.getElementById(elementListPedestriansIsland[i]).setAttribute("style", "visibility:visible");
                        svgDoc.getElementById(elementListPedestriansInnerLights[i]).setAttribute("style", "visibility:visible");
                    }else {
                        svgDoc.getElementById(elementListPedestriansIsland[i]).setAttribute("style", "visibility:hidden");
                        svgDoc.getElementById(elementListPedestriansIsland[i]).setAttribute("style", "visibility:hidden");
                        svgDoc.getElementById(elementListPedestriansInnerLights[i]).setAttribute("style", "visibility:hidden");
                    }
                });
            } else {
                //hide pedestrian-traffic-lights
                svgDoc.getElementById(elementListPedestriansOutterLights[i]).setAttribute("style", "visibility:hidden");
                svgDoc.getElementById(elementListPedestriansInnerLights[i]).setAttribute("style", "visibility:hidden");

                //hide checkbox for pedestrian islands and gray out
                var checkbox = document.getElementById(nodeListIslandPedestrians[i]);
                checkbox.disabled = true;
                checkbox.checked = false;
                document.getElementById(nodeListIslandPedestrians[i]).parentElement.classList.add("inactive-checkbox-text");

                //hide pedestrian islands
                svgDoc.getElementById(elementListPedestriansIsland[i]).setAttribute("style", "visibility:hidden");
                svgDoc.getElementById(elementListPedestriansIsland[i]).setAttribute("style", "visibility:hidden");
            }
        });
    }

}
function switchLightTo( trafficLightID,  color){

    var trafficLight = svgDoc.getElementById(trafficLightID);
    switch(color){
        case "red":
            trafficLight.children[0].children[1].setAttribute("style","visibility:visible");
            trafficLight.children[0].children[2].setAttribute("style","visibility:hidden");
            trafficLight.children[0].children[3].setAttribute("style","visibility:hidden");

            break;
        case "yellow":
            trafficLight.children[0].children[1].setAttribute("style","visibility:hidden");
            trafficLight.children[0].children[2].setAttribute("style","visibility:visible");
            trafficLight.children[0].children[3].setAttribute("style","visibility:hidden");
            break;
        case "green":
            trafficLight.children[0].children[1]  .setAttribute("style","visibility:hidden");
            trafficLight.children[0].children[2]  .setAttribute("style","visibility:hidden");
            trafficLight.children[0].children[3]  .setAttribute("style","visibility:visible");
    }
}

function blendInGreenLines(trafficLightLetter,on_or_off){
    var trafficLight = svgDoc.getElementById(trafficLightLetter);
    switch(trafficLightLetter){
        case "A":
            svgDoc.getElementById("left-left").setAttribute("style", "visibility:visible");
            break;
        case "B":
            svgDoc.getElementById("left-straight").setAttribute("style", "visibility:visible");
            break;
        case "C":
            svgDoc.getElementById("left-right").setAttribute("style", "visibility:visible");
            break;
        case "E":
            svgDoc.getElementById("bottom-left").setAttribute("style", "visibility:visible");
            break;
        case "F":
            svgDoc.getElementById("bottom-straight").setAttribute("style", "visibility:visible");
            break;
        case "G":
            svgDoc.getElementById("bottom-left").setAttribute("style", "visibility:visible");
            break;
        case "K":
            svgDoc.getElementById("right-right").setAttribute("style", "visibility:visible");
            break;
        case "J":
            svgDoc.getElementById("right-straight").setAttribute("style", "visibility:visible");
            break;
        case "I":
            svgDoc.getElementById("right-left").setAttribute("style","visibility:visible");
            break;
        case "M":
            svgDoc.getElementById("top-left").setAttribute("style", "visibility:visible");
            break;
        case "N":
            svgDoc.getElementById("top-straight").setAttribute("style", "visibility:visible");
            break;
        case "O":
            svgDoc.getElementById("top-right").setAttribute("style", "visibility:visible");
            break;
        case "pedestrians-top":
            svgDoc.getElementById("top_1_").setAttribute("style", "visibility:visible");
            break;
        case "pedestrians-bottom":
            svgDoc.getElementById("bottom_1_").setAttribute("style", "visibility:visible");
            break;
        case "pedestrians-right":
            svgDoc.getElementById("right_1_").setAttribute("style", "visibility:visible");
            break;
        case "pedestrians-left":
            svgDoc.getElementById("left_1_").setAttribute("style", "visibility:visible");
            break;
    }
}

/**
    write function to switch between green, yellow and red of a specific traffic-light

    You can do this with:
    var trafficLight = svgDoc.getElementById("A-left-left");

    var trafficLightRed = trafficLight.children[0].children[1];
    var trafficLightYellow = trafficLight.children[0].children[2];
    var trafficLightGreen = trafficLight.children[0].children[3];

    trafficLightYellow.setAttribute("style", "visibility:hidden");
 */

/**
 * Sending and Receiving JSON-Data to/from the backend
 */
function sendIntersectionJSONDataToBackend() {

    //send JSON
    var xhr = new XMLHttpRequest();
    var url = "/json";
    xhr.open("POST", url, true);
    xhr.setRequestHeader("Content-Type", "application/json");

    //function will be called when receiving JSON
    xhr.onreadystatechange = function () {
        if (xhr.readyState === 4 && xhr.status === 200) {
            let response = JSON.parse(xhr.responseText);

            //Read data from JSON-Object which was generated in go
            //names of the variables like "Fristname" should be the same as in the go struct
            console.log("[INFO] received json response from go-backend: ", response);

            //generate new html-site with received data
            //generateNewGUI(json.Firstname, json.Lastname);    //TODO display response
        }
    };

    //prepare data for sending
    let data = collectDataFromIntersection();

    console.log("[INFO] call go-backend with:", data);
    xhr.send(JSON.stringify(data));
}

/**
 * Collect all relevant Data from the intersection and build JSON object
 */
function collectDataFromIntersection() {
    const ALGORITHM_BASIC_GREEDY = 0;
    const ALGORITHM_WELSH_POWELL = 1;
    const ALGORITHM_BRON_KERBOSCH = 2;
    const PEDESTRIAN_NORMAL = 0;
    const PEDESTRIAN_WITH_ISLAND = 1;

    let settings = {Algorithm: ALGORITHM_BASIC_GREEDY};

    let intersection_part_top = {RightLane: true, StraightLane: true, LeftLane: true, Pedestrian: PEDESTRIAN_NORMAL};
    let intersection_part_right = {RightLane: true, StraightLane: true, LeftLane: true, Pedestrian: PEDESTRIAN_NORMAL};
    let intersection_part_buttom = {RightLane: true, StraightLane: true, LeftLane: true, Pedestrian: PEDESTRIAN_NORMAL};
    let intersection_part_left = {RightLane: true, StraightLane: true, LeftLane: true, Pedestrian: PEDESTRIAN_NORMAL};

    let intersection = {
        Top: intersection_part_top,
        Right: intersection_part_right,
        Buttom: intersection_part_buttom,
        Left: intersection_part_left
    };

    let jsonData = {};
    jsonData.Settings = settings;
    jsonData.Intersection = intersection;

    return jsonData;
}

