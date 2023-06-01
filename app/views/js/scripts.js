

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

var rowNum = 1;
$(function () {
    var rowNum = getCookie("rowNum") || 1;
    for (var i = 1; i < rowNum; i++) {
        addRow(i);
    }

    $("#add-btn").on("click", function () {
        addRow(rowNum);
        rowNum++;
        setCookie("rowNum", rowNum);
    });

    $("#delete-btn").on("click", function () {
        $("#table-body tr:last").remove();
        rowNum--;
        setCookie("rowNum", rowNum);
    });
});

function addRow(rowNum) {
    var newRow = $("<tr>");
    newRow.append("<td><input name='trainer" + rowNum + "'></td>");
    newRow.append("<td><input name='trainee" + rowNum + "'></td>");
    $("#table-body").append(newRow);
}

function setCookie(name, value) {
    document.cookie = name + "=" + value + "; path=/";
}

function getCookie(name) {
    var value = "; " + document.cookie;
    var parts = value.split("; " + name + "=");
    if (parts.length == 2) {
        return parts.pop().split(";").shift();
    }
    return null;
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
    title.textContent = pjname;
    title.setAttribute("class", "rolecount-title");
    roleCountsElement.appendChild(title);

    for (var i = 0; i < roleCounts.length; i++) {
        var roleCount = roleCounts[i];
        var div = document.createElement("div");
        div.textContent = roleCount.name + " : " + roleCount.count + " 回";
        div.setAttribute("class", "rolecount-container");
        roleCountsElement.appendChild(div);
    }
}
function displayRoleCountsP(roleCounts, pjname) {
    var roleCountsElement = document.getElementById("role-counts-display-noneP");
    roleCountsElement.setAttribute("id", "role-counts-displayP");
    roleCountsElement.innerHTML = ""; // 一旦中身をクリア
    var title = document.createElement("div");
    title.textContent = pjname;
    title.setAttribute("class", "rolecount-titleP");
    roleCountsElement.appendChild(title);

    for (var i = 0; i < roleCounts.length; i++) {
        var roleCount = roleCounts[i];
        var div = document.createElement("div");
        div.textContent = roleCount.name + " : " + roleCount.count + " 回";
        div.setAttribute("class", "rolecount-containerP");
        roleCountsElement.appendChild(div);
    }
}

function displayRoleCountsNone() {
    var roleCountsElement = document.getElementById("role-counts-display");
    roleCountsElement.setAttribute("id", "role-counts-display-none");
}
function displayRoleCountsNoneP() {
    var roleCountsElement = document.getElementById("role-counts-displayP");
    roleCountsElement.setAttribute("id", "role-counts-display-noneP");
}
function copyValue(input, targetName, copyName) {
    var targetInput = input.closest(".AM").querySelectorAll("[name='" + targetName + "']")[0];
    var copyInput = input.closest(".AM").querySelectorAll("[name='" + copyName + "']")[0];
    targetInput.value = input.value;
    copyInput.value = input.value;
}

function copyValue2(input, targetName, copyName) {
    var targetInput = input.closest(".PM").querySelectorAll("[name='" + targetName + "']")[0];
    var copyInput = input.closest(".PM").querySelectorAll("[name='" + copyName + "']")[0];
    targetInput.value = input.value;
    copyInput.value = input.value;
}

function confirmDelete(event) {
    event.preventDefault();
    if (confirm("本当に削除しますか？")) {
        document.getElementById("delete-form").submit();
    }
}

function checkInstalled() {
    if (window.matchMedia('(display-mode: standalone)').matches || window.navigator.standalone == true) {
        document.getElementById("addToHomeScreenButton").style.display = "none";
    } else {
        document.getElementById("addToHomeScreenButton").style.display = "block";
    }
}


window.addEventListener("load", function () {
    if (window.location.pathname === '/top') {
        checkInstalled();
    }
});




