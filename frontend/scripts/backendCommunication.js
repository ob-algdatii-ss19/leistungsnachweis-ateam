export {sendIntersectionJSONDataToBackend};

/**
 * Sending and Receiving JSON-Data to/from the backend
 */
function sendIntersectionJSONDataToBackend() {

    //send JSON
    let xhr = new XMLHttpRequest();
    let url = "/json";
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
            generateNewGUI(response);
        }
    };

    //prepare data for sending
    let data = collectDataFromIntersection();

    console.log("[INFO] call go-backend with:", data);
    xhr.send(JSON.stringify(data));
}

/**
 * Function will be called after receiving JSON-Data from the backend.
 * It generates a new html-page with the data from the backend.
 */
function generateNewGUI(responseData) {

    let redirectUrl = "result.html";

    //AJAX Request for async loading of html-pages
    let xhttp = new XMLHttpRequest();
    xhttp.onreadystatechange = function() {
        if (this.readyState == 4 && this.status == 200) {

            console.log("[INFO] redirect to " + redirectUrl);

            //replace html-body of intersection.html with content from result.html
            document.getElementById("intersection-settings").innerHTML = this.responseText;

            //insert JSON-Data
            let jsonResult = document.createTextNode(JSON.stringify(responseData.TrafficLightPhases));
            document.getElementById("json-result").appendChild(jsonResult);
        }
    };
    xhttp.open("GET", redirectUrl, true);
    xhttp.send();
}

/**
 * Collect all relevant Data from the intersection and build JSON object
 */
function collectDataFromIntersection() {

    let settings = {Algorithm: parseInt(document.getElementById("algorithm-type").value, 10)};

    let intersection_part_top = {
        RightLane: document.getElementById("node-O").checked,
        StraightLane: document.getElementById("node-N").checked,
        LeftLane: document.getElementById("node-M").checked,
        Pedestrian: pedestrianType("top")};
    let intersection_part_right = {
        RightLane: document.getElementById("node-K").checked,
        StraightLane: document.getElementById("node-J").checked,
        LeftLane: document.getElementById("node-I").checked,
        Pedestrian: pedestrianType("right")};
    let intersection_part_buttom = {
        RightLane: document.getElementById("node-G").checked,
        StraightLane: document.getElementById("node-F").checked,
        LeftLane: document.getElementById("node-E").checked,
        Pedestrian: pedestrianType("bottom")};
    let intersection_part_left = {
        RightLane: document.getElementById("node-C").checked,
        StraightLane: document.getElementById("node-B").checked,
        LeftLane: document.getElementById("node-A").checked,
        Pedestrian: pedestrianType("left")};

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

/**
 * Evaluate the the type of the pedestrian crossing
 */
function pedestrianType(intersectionPart) {

    const PEDESTRIAN_OFF = 0;
    const PEDESTRIAN_NORMAL = 1;
    const PEDESTRIAN_WITH_ISLAND = 2;

    let pedestrianType;
    let pedestrianId = "pedestrians-" + intersectionPart;
    let pedestrianIslandId = "pedestrians-island-" + intersectionPart;

    console.log("[DEBUG] evaluate pedestrian type for " + intersectionPart);

    if(document.getElementById(pedestrianId).checked) {

        if(document.getElementById(pedestrianIslandId).checked){
            pedestrianType = PEDESTRIAN_WITH_ISLAND;
        } else {
            pedestrianType = PEDESTRIAN_NORMAL;
        }
    } else {
        pedestrianType = PEDESTRIAN_OFF;
    }

    return pedestrianType;
}
