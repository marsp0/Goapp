function toggler(value) {
	"use strict";
    if (value === "ranked") {
        var live_1 = document.getElementById("live");
        var ranked_1 = document.getElementById("ranked");
        live_1.style.display = "none";
        ranked_1.style.display = "block";
        $("#live-li").removeClass("active");
        $("#ranked-li").addClass("active");

    } else {
        var ranked_3 = document.getElementById("ranked");
        var live_3 = document.getElementById("live");
        ranked_3.style.display = "none";
        live_3.style.display = "block";
        $("#live-li").addClass("active");
        $("#ranked-li").removeClass("active");
    }
}