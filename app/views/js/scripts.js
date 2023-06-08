
if (window.location.pathname === '/typingpage') {
    const tapDeleteBalloon = document.getElementById("tap-delete-balloon");
    var balloonArray = [];
    if (tapDeleteBalloon) {
        tapDeleteBalloon.addEventListener('click', function (event) {
            if (!event.target.closest(".balloonoya")) {
                fadeballon();
            }
        });
    }
    function fadeballon() {
        for (let i = 0; i <= balloonArray.length - 1; i++) {
            var wObjballoon = document.getElementById(balloonArray[i]);
            if (wObjballoon.className == "balloon") {
                wObjballoon.className = "balloonnone";

            }
        }
        balloonArray.splice(0, balloonArray.length); // balloonArrayを空の配列にする
    }
    function showBalloon(id) {
        var wObjballoon = document.getElementById(id);
        if (wObjballoon.className == "balloonnone") {
            wObjballoon.className = "balloon";
            balloonArray.push(id);
        } else {
            wObjballoon.className = "balloonnone";
            var index = balloonArray.indexOf(id);
            if (index !== -1) {
                balloonArray.splice(index, 1); // idをballoonArrayから削除
            }
        }
    }

    window.onload = function () {
        var pmForm = document.getElementsByName("pm-form");

        if (pmForm.length > 0 && pmForm[0].value === "試食会") {
            var tastingElements = document.getElementsByClassName("tasting");
            for (var i = 0; i < tastingElements.length; i++) {
                tastingElements[i].classList.add("tasting-color");
            }
        }
    };



}
var rowNum = 20;
$(function () {
    var rowNum = 20;
    $("#add-btn").on("click", function () {
        addRow(rowNum);
        rowNum++;
    });

    $("#delete-btn").on("click", function () {
        $("#trainer-trainee-form-tbody tr:last").remove();
        rowNum--;
    });
});

function addRow(rowNum) {
    var newRow = $("<tr>");
    newRow.append("<td class='trainer-trainee-form-tbody-td'><input name='trainer" + rowNum + "'></td>");
    newRow.append("<td class='trainer-trainee-form-tbody-td'><input name='trainee" + rowNum + "'></td>");
    $("#trainer-trainee-form-tbody").append(newRow);
}

$(window).on('load', function () {
    setTimeout(function () {
        $('#loading').fadeOut();
    }, 1000);
});

function getRoleCount(pjname) {
    $.ajax({
        url: "/getRoleCount",
        type: "GET",
        data: {
            pjname: pjname
        },
        success: function (response) {
            displayRoleCounts(response.RoleCounts, pjname);
        },
        error: function (error) {
            console.log(error);
        }
    });
}
function getRoleCountP(pjname) {
    $.ajax({
        url: "/getRoleCount",
        type: "GET",
        data: {
            pjname: pjname
        },
        success: function (response) {
            displayRoleCountsP(response.RoleCounts, pjname);
        },
        error: function (error) {
            console.log(error);
        }
    });
}
function displayRoleCounts(roleCounts, pjname) {
    var roleCountsElement = document.getElementById("role-counts-display-none");
    roleCountsElement.setAttribute("id", "role-counts-display");
    roleCountsElement.innerHTML = ""; // 一旦中身をクリア

    var title = document.createElement("div");
    title.setAttribute("class", "rolecount-title");

    title.textContent = pjname;
    roleCountsElement.appendChild(title);

    var table = document.createElement("table");
    table.setAttribute("class", "rolecount-table");
    var tbody = document.createElement("tbody");
    tbody.setAttribute("class", "rolecount-tbody");

    var trHeader = document.createElement("tr");
    trHeader.setAttribute("class", "rolecount-tr-header");
    var thRole = document.createElement("th");
    thRole.setAttribute("class", "rolecount-header-th");
    var thCount3Mon = document.createElement("th");
    thCount3Mon.setAttribute("class", "rolecount-header-th");
    var thCountAll = document.createElement("th");
    thCountAll.setAttribute("class", "rolecount-header-th");

    thRole.textContent = "役割";
    thCount3Mon.textContent = "過去3ヶ月";
    thCountAll.textContent = "全期間";

    trHeader.appendChild(thRole);
    trHeader.appendChild(thCount3Mon);
    trHeader.appendChild(thCountAll);
    tbody.appendChild(trHeader);

    for (var i = 0; i < roleCounts.length; i++) {
        var roleCount = roleCounts[i];

        var trData = document.createElement("tr");
        trData.setAttribute("class", "rolecount-tr-data");
        var tdRole = document.createElement("td");
        tdRole.setAttribute("class", "rolecount-td-data");
        var tdCount3Mon = document.createElement("td");
        tdCount3Mon.setAttribute("class", "rolecount-td-data");
        var tdCountAll = document.createElement("td");
        tdCountAll.setAttribute("class", "rolecount-td-data");


        tdRole.textContent = roleCount.name;
        tdCount3Mon.textContent = roleCount.count3mon + " 回";
        tdCountAll.textContent = roleCount.countall + " 回";

        trData.appendChild(tdRole);
        trData.appendChild(tdCount3Mon);
        trData.appendChild(tdCountAll);
        tbody.appendChild(trData);
    }
    table.appendChild(tbody);
    roleCountsElement.appendChild(table);

    var uptitle = document.createElement("div");
    uptitle.setAttribute("class", "to-updatepj");

    var form = document.createElement("form");
    form.setAttribute("action", "/updatepj");
    form.setAttribute("method", "POST");

    var inputname = document.createElement("input");
    inputname.setAttribute("type", "hidden");
    inputname.setAttribute("name", "pjname");
    inputname.setAttribute("value", pjname);

    var input = document.createElement("input");
    input.setAttribute("type", "submit");
    input.setAttribute("value", "一人前チェックシートを見る");

    var btndesign = document.createElement("div");
    btndesign.setAttribute("class", "btn-design-1");

    form.appendChild(inputname);
    form.appendChild(btndesign);
    btndesign.appendChild(input);
    uptitle.appendChild(form);
    roleCountsElement.appendChild(uptitle);
}

function displayRoleCountsP(roleCounts, pjname) {
    var roleCountsElement = document.getElementById("role-counts-display-noneP");
    roleCountsElement.setAttribute("id", "role-counts-displayP");
    roleCountsElement.innerHTML = ""; // 一旦中身をクリア

    var title = document.createElement("div");
    title.setAttribute("class", "rolecount-title");

    title.textContent = pjname;
    roleCountsElement.appendChild(title);

    var table = document.createElement("table");
    table.setAttribute("class", "rolecount-table");
    var tbody = document.createElement("tbody");
    tbody.setAttribute("class", "rolecount-tbody");

    var trHeader = document.createElement("tr");
    trHeader.setAttribute("class", "rolecount-tr-header");
    var thRole = document.createElement("th");
    thRole.setAttribute("class", "rolecount-header-th");
    var thCount3Mon = document.createElement("th");
    thCount3Mon.setAttribute("class", "rolecount-header-th");
    var thCountAll = document.createElement("th");
    thCountAll.setAttribute("class", "rolecount-header-th");

    thRole.textContent = "役割";
    thCount3Mon.textContent = "過去3ヶ月";
    thCountAll.textContent = "全期間";

    trHeader.appendChild(thRole);
    trHeader.appendChild(thCount3Mon);
    trHeader.appendChild(thCountAll);
    tbody.appendChild(trHeader);

    for (var i = 0; i < roleCounts.length; i++) {
        var roleCount = roleCounts[i];

        var trData = document.createElement("tr");
        trData.setAttribute("class", "rolecount-tr-data");
        var tdRole = document.createElement("td");
        tdRole.setAttribute("class", "rolecount-td-data");
        var tdCount3Mon = document.createElement("td");
        tdCount3Mon.setAttribute("class", "rolecount-td-data");
        var tdCountAll = document.createElement("td");
        tdCountAll.setAttribute("class", "rolecount-td-data");


        tdRole.textContent = roleCount.name;
        tdCount3Mon.textContent = roleCount.count3mon + " 回";
        tdCountAll.textContent = roleCount.countall + " 回";

        trData.appendChild(tdRole);
        trData.appendChild(tdCount3Mon);
        trData.appendChild(tdCountAll);
        tbody.appendChild(trData);
    }
    table.appendChild(tbody);
    roleCountsElement.appendChild(table);

    var uptitle = document.createElement("div");
    uptitle.setAttribute("class", "to-updatepj");

    var form = document.createElement("form");
    form.setAttribute("action", "/updatepj");
    form.setAttribute("method", "POST");

    var inputname = document.createElement("input");
    inputname.setAttribute("type", "hidden");
    inputname.setAttribute("name", "pjname");
    inputname.setAttribute("value", pjname);

    var input = document.createElement("input");
    input.setAttribute("type", "submit");
    input.setAttribute("value", "一人前チェックシートを見る");

    var btndesign = document.createElement("div");
    btndesign.setAttribute("class", "btn-design-1");

    form.appendChild(inputname);
    form.appendChild(btndesign);
    btndesign.appendChild(input);
    uptitle.appendChild(form);
    roleCountsElement.appendChild(uptitle);

}

function displayRoleCountsNone() {
    var roleCountsElement = document.getElementById("role-counts-display");
    roleCountsElement.setAttribute("id", "role-counts-display-none");
}
function displayRoleCountsNoneP() {
    var roleCountsElement = document.getElementById("role-counts-displayP");
    roleCountsElement.setAttribute("id", "role-counts-display-noneP");
}




