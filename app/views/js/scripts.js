

if (window.location.pathname === '/typingpage') {
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
    function fadeballon() {
        // console.log(balloonArray)
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
        // console.log(balloonArray)
    }
}


$(window).on('load', function () {
    setTimeout(function () {
        $('#loading').fadeOut();
    }, 1000);
});


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


window.addEventListener("load", function () {
    if (window.location.pathname === '/') {
        checkInstalled();
    }
});




