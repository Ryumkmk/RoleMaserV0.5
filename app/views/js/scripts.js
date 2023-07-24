
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

function submitForm(obj, date, ampm) {
    var form = document.getElementById("typing-form");
    var formData = $(form).serialize(); // フォームデータのシリアライズ
    $.ajax({
        url: "/uploadPj",
        type: "POST",
        data: formData,
        success: function () {
            makeresttypingForm(obj, date);
            if (ampm == "ダブル") {
                ispjinputeddouble(date);
            } else {
                ispjinputed(date, ampm);
            }
        },
        error: function (error) {
            console.log(error);
        }
    });
}

function ispjinputeddouble(date) {
    $.ajax({
        url: "/ispjinputeddouble",
        type: "GET",
        data: {
            date: date,
        },
        success: function (response) {
            displaynotinputedpj(response.AMNames, "");
            displaynotinputedpj(response.PMNames, "P");
        },
        error: function (error) {
            console.log(error);
        }
    });
}

function ispjinputed(date, ampm) {
    $.ajax({
        url: "/ispjinputed",
        type: "GET",
        data: {
            date: date,
            ampm: ampm,
        },
        success: function (response) {
            displaynotinputedpj(response.Names, "");
        },
        error: function (error) {
            console.log(error);
        }
    });
}

function displaynotinputedpj(names, obj) {
    var notInputedpjnameElement = document.getElementById("notInputedpjname" + obj);
    notInputedpjnameElement.innerHTML = ""; // 一旦中身をクリア
    var list = document.createElement("div");
    if (obj == "P") {
        list.textContent += "PM の出勤PJ："
    } else {
        list.textContent += "AM の出勤PJ："
    }

    for (var i = 0; i < names.length; i++) {
        var name = names[i];
        list.textContent += " ※ " + name;
    }
    list.textContent += "は未入力ですが宜しいですか？"
    notInputedpjnameElement.appendChild(list);

}

function makeresttypingForm(obj, date) {
    $.ajax({
        url: "/makeresttypingpage",
        type: "GET",
        data: {
            date: date
        },
        success: function (response) {
            displayRestTypingPage(obj, response.RICPs);
        },
        error: function (error) {
            console.log(error);
        }
    });
}

function displayRestTypingPage(obj, RICPs) {
    if (obj.endsWith("D")) {
        displayRestTypingPageByAmpm("P", RICPs)
        displayRestTypingPageByAmpm("", RICPs)
    } else if (obj.endsWith("P")) {
        displayRestTypingPageByAmpm("P", RICPs)
    } else {
        displayRestTypingPageByAmpm("", RICPs)
    }
}
function displayRestTypingPageByAmpm(obj, RICPs) {

    var restTypingElement = document.getElementById("rest" + obj);
    if (restTypingElement) {
        var inputs = restTypingElement.querySelectorAll("input[type='text']");
        for (var i = 0; i < inputs.length; i++) {
            inputs[i].value = "";
        }
    } else {
        restTypingElement = document.getElementById("rest-typing-display-none" + obj);
        restTypingElement.setAttribute("id", "rest" + obj);
        var inputs = restTypingElement.querySelectorAll("input[type='text']");
        for (var i = 0; i < inputs.length; i++) {
            inputs[i].value = "";
        }
    }
    for (var i = 0; i < RICPs.length; i++) {
        var rICP = RICPs[i];
        switch (rICP.RoleName) {
            case "ドリカン＆作" + obj:
            case "ドリカン＆聞" + obj:
                var restDrink = document.getElementById("rest-drink" + obj);
                restDrink.value += rICP.PjName + " ";
                break;
            case "コーヒー" + obj:
            case "洗い場洗い" + obj:
            case "洗い場拭き" + obj:
                var restCoffee = document.getElementById("rest-coffee" + obj);
                restCoffee.value += rICP.PjName + " ";
                break;
            case "門番＆第二" + obj:
            case "門番＆第一" + obj:
                var restGatekeeper = document.getElementById("rest-gatekeeper" + obj);
                restGatekeeper.value += rICP.PjName + " ";
                break;
            case "クロークサブ" + obj:
            case "リターンリーダー" + obj:
            case "クローク" + obj:
                var restCloak = document.getElementById("rest-cloak" + obj);
                restCloak.value += rICP.PjName + " ";
                break;
            case "アペ＆水" + obj:
            case "トイレタバコ男" + obj:
                var restApe = document.getElementById("rest-ape" + obj);
                restApe.value += rICP.PjName + " ";
                break;
            case "シルバー＆ワイン" + obj:
                var restSilver = document.getElementById("rest-silver" + obj);
                restSilver.value += rICP.PjName + " ";
                break;
            case "シャンパン＆ワイン" + obj:
            case "トイレタバコ女" + obj:
                var restChampagne = document.getElementById("rest-champagne" + obj);
                restChampagne.value += rICP.PjName + " ";
                break;
            case "リーダー" + obj:
                var restLeader = document.getElementById("rest-leader" + obj);
                restLeader.value += rICP.PjName + " ";
                break;
            case "サブPJ" + obj:
                var restSubPj = document.getElementById("rest-subpj" + obj);
                restSubPj.value += rICP.PjName + " ";
                break;
        }
    }
}

function getRoleCount(obj, pjname) {
    $.ajax({
        url: "/getRoleCount",
        type: "GET",
        data: {
            pjname: pjname
        },
        success: function (response) {
            displayRoleCounts(obj, response.RoleCounts, pjname);
        },
        error: function (error) {
            console.log(error);
        }
    });
}

function displayRoleCounts(obj, roleCounts, pjname) {

    if (obj.endsWith("P")) {
        var roleCountsElement = document.getElementById("role-counts-display-noneP");
        roleCountsElement.setAttribute("id", "role-counts-displayP");
    } else {
        var roleCountsElement = document.getElementById("role-counts-display-none");
        roleCountsElement.setAttribute("id", "role-counts-display");
    }
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

function displayRoleCountsNone(obj) {
    if (obj.endsWith("P")) {
        var roleCountsElement = document.getElementById(obj);
        roleCountsElement.setAttribute("id", "role-counts-display-noneP");
    } else {
        var roleCountsElement = document.getElementById("role-counts-display");
        roleCountsElement.setAttribute("id", "role-counts-display-none");
    }

}





