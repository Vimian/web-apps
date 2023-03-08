package main

// index is the index page
const index string = `
<!DOCTYPE html>
<html>
	<head>
		<title>iFrame Timer</title>
	</head>
	<body>
		<script>
			function goto(params) {
				var prefix = "/";
				if (window.location.pathname.endsWith("/")) {
					prefix = "";
				}
				const url =
					window.location.origin +
					window.location.pathname +
					prefix +
					"timer?site=" +
					document.getElementById("url").value;
				window.location.replace(url);
			}
		</script>
		<input type="text" id="url" placeholder="URL..." />
		<button type="submit" onClick="goto()">Goto</button>
	</body>
</html>
`

// timer is the timer page
const timer string = `
<!DOCTYPE html>
<html>
        <head>
                <title>iFrame Timer</title>
                <link
                        rel="stylesheet"
                        href="https://cdn.jsdelivr.net/npm/bootstrap@4.3.1/dist/css/bootstrap.min.css"
                        integrity="sha384-ggOyR0iXCbMQv3Xipma34MD+dH/1fQ784/j6cY/iJTQUOhcWr7x9JvoRxT2MZw1T"
                        crossorigin="anonymous"
                />

                <style>
                        body {
                                margin: 0;
                                padding: 0;
                        }
                </style>
        </head>
        <body>
                <div class="row no-gutters">
                        <div class="col-9" style="display: flex; height: 100vh">
                                <iframe
                                        id="iframe"
                                        width="100%"
                                        height="100%"
                                        style="border: none"
                                ></iframe>

                                <!-- If no query string is present, redirect to the parent directory -->
                                <script>
                                        if (window.location.search == "") {
                                                window.location.replace(
                                                        window.location.origin +
                                                                window.location.pathname
                                                                        .split(
                                                                                "/"
                                                                        )
                                                                        .slice(
                                                                                0,
                                                                                -1
                                                                        )
                                                                        .join(
                                                                                "/"
                                                                        )
                                                );
                                        }

                                        document.getElementById("iframe").src =
                                                window.location.search
                                                        .split("&")[0]
                                                        .split("=")
                                                        .slice(1)
                                                        .join("=");
                                </script>
                        </div>

                        <div class="col-3">
                                <script>
                                        var time = 120;
                                        var set = 0;
                                        var timer = 0;
                                        setInterval(function () {
                                                if (timer == 1) {
                                                        document.getElementById(
                                                                "audio"
                                                        ).play();
                                                        set = set + 1;
                                                        document.getElementById(
                                                                "set"
                                                        ).innerHTML = set;
                                                }

                                                if (timer > 0) {
                                                        timer = timer - 1;
                                                        document.getElementById(
                                                                "timer"
                                                        ).innerHTML = timer;
                                                }
                                        }, 1000);

                                        function start() {
                                                timer = time;
                                                document.getElementById(
                                                        "set"
                                                ).innerHTML = set;
                                        }

                                        function reset() {
                                                timer = time;
                                                set = 0;
                                                document.getElementById(
                                                        "set"
                                                ).innerHTML = set;
                                        }
                                </script>

                                <!-- interval buttons -->
                                <div id="intervals">
                                        <input
                                                type="number"
                                                class="form-control"
                                                id="inputTime"
                                                placeholder="Time"
                                                value="120"
                                                onchange="time = this.value"
                                        />
                                        <script>
                                                var intervals = [
                                                        30, 60, 90, 120, 150,
                                                        180,
                                                ];

                                                for (
                                                        var i = 0;
                                                        i < intervals.length;
                                                        i++
                                                ) {
                                                        // add buttons with onclick to set time to interval, display interval and set inputTime to interval
                                                        document.getElementById(
                                                                "intervals"
                                                        ).innerHTML +=
                                                                '<button type="button" class="btn btn-primary" onclick="time = ' +
                                                                intervals[i] +
                                                                "; document.getElementById('inputTime').value = " +
                                                                intervals[i] +
                                                                '">' +
                                                                intervals[i] +
                                                                "</button>";
                                                }
                                        </script>
                                </div>

                                <!-- timer and set display -->
                                <h1>T: <span id="timer"></span></h1>
                                <h1>Set: <span id="set"></span></h1>

                                <!-- Pause and Next/Start buttons -->
                                <button
                                        type="button"
                                        class="btn btn-primary"
                                        onclick="start()"
                                >
                                        Pause
                                </button>
                                <button
                                        type="button"
                                        class="btn btn-primary"
                                        onclick="reset()"
                                >
                                        Next/Start
                                </button>

                                <!-- + and - set buttons -->
                                <button
                                        type="button"
                                        class="btn btn-primary"
                                        onclick="set = set + 1; document.getElementById('set').innerHTML = set"
                                >
                                        +
                                </button>
                                <button
                                        type="button"
                                        class="btn btn-primary"
                                        onclick="set = set - 1; document.getElementById('set').innerHTML = set"
                                >
                                        -
                                </button>

                                <!-- audio -->
                                <audio
                                        id="audio"
                                        src="https://cdn.pixabay.com/download/audio/2022/03/10/audio_2cd9a13a55.mp3?filename=buzzer-dog-39284.mp3"
                                ></audio>
                        </div>
                </div>

                <script
                        src="https://code.jquery.com/jquery-3.3.1.slim.min.js"
                        integrity="sha384-q8i/X+965DzO0rT7abK41JStQIAqVgRVzpbzo5smXKp4YfRvH+8abtTE1Pi6jizo"
                        crossorigin="anonymous"
                ></script>
                <script
                        src="https://cdn.jsdelivr.net/npm/popper.js@1.14.7/dist/umd/popper.min.js"
                        integrity="sha384-UO2eT0CpHqdSJQ6hJty5KVphtPhzWj9WO1clHTMGa3JDZwrnQq4sF86dIHNDz0W1"
                        crossorigin="anonymous"
                ></script>
                <script
                        src="https://cdn.jsdelivr.net/npm/bootstrap@4.3.1/dist/js/bootstrap.min.js"
                        integrity="sha384-JjSmVgyd0p3pXB1rRibZUAYoIIy6OrQ6VrjIEaFf/nJGzIxFDsf4x0xIM+B07jRM"
                        crossorigin="anonymous"
                ></script>
        </body>
</html>
`
