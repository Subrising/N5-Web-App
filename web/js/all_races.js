$( document ).ready(function(){
    //Perform Ajax request.
    $.ajax({
        url: '/receive',
        type: 'GET',
        dataType: 'JSON',
        success: function (response){
            console.log(response);
            var len = response["races"].length;
            console.log("race length =" + len);
            var races = response["races"];

            for (var i = 0; i < len; i++){
                var id = i;

                // Gets location of race
                var location = response["location"];
                console.log(location);

                // Gets Unix closing time of race
                var close = races[i]["close"];
                console.log(close);

                // Gets full data value for race
                var dt=eval(close*1000);
                var myDate = new Date(dt);

                // Gets the type of race
                var raceType = races[i]["racetype"];
                console.log(raceType);

                // Gets time in seconds until race closes
                close = close - Math.round((new Date()).getTime());
                close = Math.floor(close / 1000);

                // Sets up table rows for race
                var tr_str = "<tr id='faded'>" +
                    "<td align='center' id='race_" + id + "'><a href='/races/" + races[i]["id"] + "'>" + id + "</a></td>" +
                    "<td align='center'>" + location + "</td>" +
                    "<td align='center'>" + raceType + "</td>" + "" +
                    "<td align='center'>" + myDate + "</td>" +
                    "<td align='center' class='timer'>" + close + "</td>" + "</tr>";

                $("#raceTable").append(tr_str);
            }
            // Starts countdown timer for race
            setTimers();
        }
    })});

function setTimers() {
    // Gets the race table that will be updated with the timers
    var table = document.getElementById("raceTable");

    // Countdown timer
    var x = setInterval(
        function () {
            // Gets amount of rows of timers that need to be set
            console.log("rows = ", table.rows.length);

            if (table.rows.length == 1) {
                clearInterval(x);
            }

            // Goes through each row, gets the column for the timer to be set in, and sets the countdown timer values
            for (var i = 1, row; row = table.rows[i], i < table.rows.length; i++) {
                var endDate = row.cells[4];
                var countDown = row.cells[4].textContent;
                console.log("countDown = " + countDown);

                // Find the distance between now and the race finish
                var distance = parseInt(countDown) - 1;

                // Display the result in the counter row
                endDate.innerText = distance;

                console.log("CD = " + endDate.innerHTML);

                // When the countdown is finished, remove the row so no bets can be placed
                if (distance < 0) {
                    row.remove();
                }
            }}, 1000);
}
