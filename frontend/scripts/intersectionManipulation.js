import * as backendCommunication from "./backendCommunication.js";

var svgDoc;

window.onload=function() {

    // Get the SVG document from inside the Object tag
    svgDoc = document.getElementById("intersection-svg").contentDocument;

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
        backendCommunication.sendIntersectionJSONDataToBackend();
    });
}


/*
    Checkbox-Listener to display and hide traffic-lights
 */
function addCheckboxListener(){
    let nodeList = ['node-A', 'node-B','node-C', 'node-E', 'node-F', 'node-G','node-I','node-J','node-K','node-M','node-N','node-O'];
    let elementList = ["A-left-left","B-left-straight","C-left-right","E-bottom-left","F-bottom-straight","G-bottom-right","I-right-left","J-right-straight","K-right-right","M-top-left","N-top-straight","O-top-right"];
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
    let nodeListPedestrians =['pedestrians-right','pedestrians-bottom','pedestrians-left','pedestrians-top'];
    let nodeListIslandPedestrians =['pedestrians-island-right','pedestrians-island-bottom','pedestrians-island-left','pedestrians-island-top'];
    let elementListPedestriansOutterLights = ["L-outter","H-outter","D-outter","P-outter"];
    let elementListPedestriansInnerLights = ["L-inner","H-inner","D-inner","P-inner"];
    let elementListPedestriansIsland = ["pi-right","pi-bottom","pi-left","pi-top"];

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

    let trafficLight = svgDoc.getElementById(trafficLightID);
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
