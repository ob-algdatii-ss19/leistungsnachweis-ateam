
window.onload=function() {

    //register button-listener
    document.getElementById("myBtn").addEventListener("click", function(){
        console.log("[INFO] button clicked, execute EventListener");

        //get values from the html-form
        var firstname = document.getElementById("firstname").value;
        var lastname = document.getElementById("lastname").value;

        callBackend(firstname, lastname);
    });

};

/**
 * Sending and Receiving JSON-Data
 */
function callBackend(firstname, lastname) {

    console.log("[INFO] call go-backend with: " + firstname + ", " + lastname);

    //send JSON
    var xhr = new XMLHttpRequest();
    var url = "/json";
    xhr.open("POST", url, true);
    xhr.setRequestHeader("Content-Type", "application/json");

    //function will be called when receiving JSON
    xhr.onreadystatechange = function () {
        if (xhr.readyState === 4 && xhr.status === 200) {
            var json = JSON.parse(xhr.responseText);

            //Read data from JSON-Object which was generated in go
            //names of the variables like "Fristname" should be the same as in the go struct
            console.log("[INFO] received json response from go-backend: " + json.Firstname + ", " + json.Lastname);

            //generate new html-site with received data
            generateNewGUI(json.Firstname, json.Lastname);
        }
    };

    //prepare data for sending
    var data = JSON.stringify({"Firstname": firstname, "Lastname": lastname});
    xhr.send(data);
}


/**
 * Function will be called after receiving JSON-Data from the backend.
 * It generates a new html-page with the data from the backend.
 */
function generateNewGUI(firstname, lastname) {

    var redirectUrl = "/result.html";

    //AJAX Request for async loading of html-pages
    var xhttp = new XMLHttpRequest();
    xhttp.onreadystatechange = function() {
        if (this.readyState == 4 && this.status == 200) {

            console.log("redirect to " + redirectUrl);

            document.body.innerHTML = this.responseText; //replace html-body with content from result.html

            //insert JSON-Data
            var jsonResult = document.createTextNode("JSON-Data from the backend: " + firstname + ", " + lastname);
            document.getElementById("json-result").appendChild(jsonResult);
        }
    };
    xhttp.open("GET", "result.html", true);
    xhttp.send();
}
