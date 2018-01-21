// Show details of specific race
function showDetails(response){
    // Shows location on page
    var updateLoc = document.getElementById("location");
    updateLoc.innerText = response["location"];

    // Gets current closing time
    var updateTime = document.getElementById("closeTime");
    var close = response["close"];

    // Gets Unix time of race and converts it to standard form
    var dt=eval(close);
    var myDate = new Date(dt);

    // Shows race close time on page
    var raceRun = document.getElementById("outcome_ts");
    raceRun.innerText = myDate;

    // Gets time remaining until race close in seconds
    close = close - Math.round((new Date()).getTime());
    close = Math.floor(close / 1000);

    var minutes = Math.floor(close / 60);
    var seconds = close % 60;

    updateTime.innerText = " " + minutes + " Minutes " + seconds + " Seconds";

    var len = response["competitors"].length;
    var competitors = response["competitors"];

    // Shows race type
    var updateType = document.getElementById("raceType");
    updateType.innerText = competitors[0]["type"];

    // Creates table elements for each competitor in race
    for (var i = 0; i < len; i++){
        var position = competitors[i]["position"];
        var type = competitors[i]["type"];
        var name = competitors[i]["name"];

        var tr_str = "<tr class='competitor row not-scratched saddle-4'>" +
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
    var table = document.getElementById("closeTime");

    // Starts countdown timer
    var x = setInterval(
        function () {
            var countDown = table.innerText;

            var string = countDown.split(" ");

            // Reduces count down
            var distance = parseInt(string[0]) * 60 + parseInt(string[2]) - 1;

            var minutes = Math.floor(distance / 60);
            var seconds = distance % 60;

            // Display the result
            table.innerText = " " + minutes + " Minutes " + seconds + " Seconds";

            // When the timer reaches 0, set the race status to closed
            if (distance <= 0) {
                clearInterval(x);
                var raceStatus = document.getElementById("Open");
                raceStatus.innerText = "Closed";
                var closed = document.getElementById("closingTime");
                closed.innerText = "CLOSED"
            }
        }, 1000);
}