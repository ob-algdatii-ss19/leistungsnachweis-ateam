/* function should be in a different js-file, which performs the algorithms for the different traffic-light-phases
function callBackend() {
    alert("backend call");

    var firstname = document.getElementById("firstname").value;
    var lastname = document.getElementById("lastname").value;


    //send JSON
    var xhr = new XMLHttpRequest();
    var url = "url";
    xhr.open("POST", url, true);
    xhr.setRequestHeader("Content-Type", "application/json");
    xhr.onreadystatechange = function () {
        if (xhr.readyState === 4 && xhr.status === 200) {
            var json = JSON.parse(xhr.responseText);
            console.log("RESPONSE FROM GO" + json.response);
        }
    };
    var data = JSON.stringify({"Firstname": firstname, "Lastname": lastname});
    xhr.send(data);
}*/
/*
function foo() {
    alert("The form was submitted");
}
*/
