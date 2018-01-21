$( document ).ready(function(){
    //Perform Ajax request.
    $.ajax({
        url: '/receive',
        type: 'GET',
        dataType: 'JSON',
        success: function (response){
            getRaces(response)
        }
    })});

function getRaces(response) {
    var len = response["races"].length;
    var races = response["races"];

    for (var i = 0; i < len; i++){
        var id = i;

        // Gets location of race
        var location = response["location"];

        // Gets Unix closing time of race
        var close = races[i]["close"];

        var closingTime =  races[i]["closeDate"];

        // Gets full data value for race
        var dt=eval(close);
        var myDate = new Date(dt);
        var dateString = myDate.toUTCString();

        // Gets the type of race
        var raceType = races[i]["racetype"];

        // Gets time in seconds until race closes
        close = close - Math.round((new Date()).getTime());
        close = Math.floor(close / 1000);

        var minutes = Math.floor(close / 60);
        var seconds = close % 60;

        // Sets up table rows for race
        var tr_str = "<tr class='row'>" +
            "<td align='center' class='odds' id='race_" + id + "'><a href='/races/" + races[i]["id"] + "'>" + id + "</a></td>" +
            "<td align='center'>" + location + "</td>" +
            "<td align='center'>" + raceType + "</td>" + "" +
            "<td align='center'>" + dateString + "</td>" +
            "<td align='center' class='timer'> <span id='closingMinutes'>" + minutes + "</span>" + " Minutes " + "<span id='closingSeconds'>" + seconds + "</span>" + " Seconds" + "</td>" + "</tr>";

        $("#raceTable").append(tr_str);
    }
    // Starts countdown timer for race
    setTimers();
}

function setTimers() {
    // Gets the race table that will be updated with the timers
    var table = document.getElementById("raceTable");

    // Countdown timer
    var x = setInterval(
        function () {

            if (table.rows.length == 1) {
                clearInterval(x);
            }

            // Goes through each row, gets the column for the timer to be set in, and sets the countdown timer values
            for (var i = 1, row; row = table.rows[i], i < table.rows.length; i++) {
                var endDate = row.cells[4];
                var countDown = row.cells[4].textContent;

                var string = countDown.split(" ");

                // Find the distance between now and the race finish
                var distance = parseInt(string[1]) * 60 + parseInt(string[3]) - 1;

                var minutes = Math.floor(distance / 60);
                var seconds = distance % 60;

                // Display the result in the counter row
                endDate.innerText = " " + minutes + " Minutes " + seconds + " Seconds";

                // When the countdown is finished, remove the row so no bets can be placed
                if (distance < 0) {
                    row.remove();
                }
            }}, 1000);
}
