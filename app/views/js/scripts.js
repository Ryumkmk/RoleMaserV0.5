
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
    checkInstalled();
});

function addToHomeScreen() {
    var userAgent = navigator.userAgent.toLowerCase();
    var isSafari = userAgent.indexOf('safari') !== -1 && userAgent.indexOf('chrome') === -1;
    if (isSafari) {
        var appleTouchIcons = [
            { href: '/RoleMaster.png', sizes: '500x500' }
        ];
        window.addEventListener('load', function () {
            if (window.navigator.standalone === true) {
                // Already installed
                document.getElementById('addToHomeScreenButton').style.display = 'none';
            } else {
                var lastTouchIcon = appleTouchIcons[appleTouchIcons.length - 1];
                var link = document.createElement('link');
                link.setAttribute('rel', 'apple-touch-icon');
                link.setAttribute('sizes', lastTouchIcon.sizes);
                link.setAttribute('href', lastTouchIcon.href);
                document.head.appendChild(link);
            }
        });
    } else {
        alert('ホーム画面に追加できません。');
    }
}



