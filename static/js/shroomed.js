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

function GetMatchInfo(matchId,index,server) {
	"use strict";
	var x = $("#" + index).attr("aria-expanded");
    var y = document.getElementsByClassName("info-present-"+matchId);
	if (x === "false" && y.length == 0) {
		$.ajax({
		url: "get_match_info",
		type:"get",
		data: {"MatchId":matchId, "Server":server},
        success: function(data) {$("#" + index).append("<p>" + data + "</p>");}
		});
	}
}