// Show details of specific race
function showDetails(response){
    console.log(response);

    // Shows location on page
    var updateLoc = document.getElementById("location");
    updateLoc.innerText = response["location"];

    // Gets current closing time
    var updateTime = document.getElementById("closeTime");
    var close = response["close"];
    console.log(close);

    // Gets Unix time of race and converts it to standard form
    var dt=eval(close*1000);
    var myDate = new Date(dt);

    // Shows race close time on page
    var raceRun = document.getElementById("outcome_ts");
    raceRun.innerText = myDate;

    // Gets time remaining until race close in seconds
    close = close - Math.round((new Date()).getTime());
    close = Math.floor(close / 1000);
    updateTime.innerText = close;

    var len = response["competitors"].length;
    console.log("Competitors total =" + len);
    var competitors = response["competitors"];

    // Shows race type
    var updateType = document.getElementById("raceType");
    updateType.innerText = competitors[0]["type"];

    // Creates table elements for each competitor in race
    for (var i = 0; i < len; i++){
        var position = competitors[i]["position"];
        console.log(position);
        var type = competitors[i]["type"];
        console.log(type);
        var name = competitors[i]["name"];

        var tr_str = "<tr class='row'>" +
            "<td align='center'>" + name  + "</td>" +
            "<td align='center'>" + position + "</td>" +
            "<td align='center'>" + type + "</td></tr>";

        $("#raceTable").append(tr_str);
    }
    // Sets countdown for race
    setTimers();
}

// Countdown timer
function setTimers() {
    console.log("yo");
    var table = document.getElementById("closeTime");
    console.log("Table length = " + table.innerHTML);

    // Starts countdown timer
    var x = setInterval(
        function () {
            var countDown = table.innerText;
            console.log("countDown = " + countDown);

            // Reduces count down
            var distance = parseInt(countDown) - 1;

            // Display the result
            table.innerText = distance;

            console.log("CD = " + table.innerHTML);

            // When the timer reaches 0, set the race status to closed
            if (distance <= 0) {
                clearInterval(x);
                var raceStatus = document.getElementById("Open");
                raceStatus.innerText = "Closed";
            }
        }, 1000);
}