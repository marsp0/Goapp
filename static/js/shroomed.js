function toggler(value) {
	"use strict";
    if (value === "ranked") {
        var live_1 = document.getElementById("live");
        var recent_1 = document.getElementById("recent");
        var ranked_1 = document.getElementById("ranked");
        live_1.style.display = "none";
        recent_1.style.display = "none";
        ranked_1.style.display = "block";
        $("#live-li").removeClass("active");
        $("#recent-li").removeClass("active");
        $("#ranked-li").addClass("active");

    } else if (value === "recent") {
        var ranked_2 = document.getElementById("ranked");
        var live_2 = document.getElementById("live");
        var recent_2 = document.getElementById("recent");
        live_2.style.display = "none";
        ranked_2.style.display = "none";
        recent_2.style.display = "block";
        $("#live-li").removeClass("active");
        $("#recent-li").addClass("active");
        $("#ranked-li").removeClass("active");
    } else {
        var ranked_3 = document.getElementById("ranked");
        var live_3 = document.getElementById("live");
        var recent_3 = document.getElementById("recent");
        ranked_3.style.display = "none";
        recent_3.style.display = "none";
        live_3.style.display = "block";
        $("#live-li").addClass("active");
        $("#recent-li").removeClass("active");
        $("#ranked-li").removeClass("active");
    }
}

function GetMatchInfo(matchId,index) {
	"use strict";
	// need to add check if the game was already requested
	// Do not wanna make multiple calls for the same game
	var x = $("#" + index).attr("aria-expanded");
	if (x === "false") {
		$.ajax({
		url: "get_match_info",
		type:"get",
		data: {"MatchId":matchId},

		})
	}
}