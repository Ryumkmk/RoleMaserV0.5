
const tapDeleteBalloon = document.getElementById("tap-delete-balloon");
var balloonArray = [];

if (tapDeleteBalloon) {
    tapDeleteBalloon.addEventListener('click', function (event) {
        // クリックされた要素が"balloonoya"クラスを持つ場合、fadeballon()関数を呼び出さない
        if (!event.target.closest(".balloonoya")) {
            fadeballon();
        }
    });
}
$(window).on('load', function () {
    setTimeout(function() {
        $('#loading').fadeOut();
    }, 1000);
});

function fadeballon() {
    // console.log(balloonArray)
    for (let i = 0; i <= balloonArray.length; i++) {
        var wObjballoon = document.getElementById(balloonArray[i]);
        if (wObjballoon.className == "balloon") {
            wObjballoon.className = "balloonnone";
        }
    }
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
        // 削除処理
        document.getElementById("delete-form").submit();
    }
}

function checkInstalled() {
    if (window.matchMedia('(display-mode: standalone)').matches || window.navigator.standalone == true) {
        // ホームに追加済み
        document.getElementById("addToHomeScreenButton").style.display = "none";
    } else {
        // まだ追加されていない
        document.getElementById("addToHomeScreenButton").style.display = "block";
    }
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
    // console.log(balloonArray)
}



window.addEventListener("load", function () {
    checkInstalled();
});




