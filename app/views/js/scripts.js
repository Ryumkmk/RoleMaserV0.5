
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
    var isSafari = /^((?!chrome|android).)*safari/i.test(navigator.userAgent);
    var isiOS = /iPad|iPhone|iPod/.test(navigator.userAgent) && !window.MSStream;
    var isStandalone = window.matchMedia('(display-mode: standalone)').matches;
    
    if (isiOS && isSafari && !isStandalone) {
      var manifest = {
        "name": "RoleMaster",
        "short_name": "RoleMaster",
        "icons": [
          {
            "src": "RoleMaster.png",
            "sizes": "500x500"
          }
        ],
        "start_url": "/",
        "display": "standalone",
        "background_color": "#ffffff",
        "theme_color": "#ffffff"
      };
      
      // iOS Safariでホーム画面に追加するための処理
      var addToHomeScreen = window.addToHomeScreen || {};
      addToHomeScreen.show(true, manifest);
    } else {
      alert('ホーム画面に追加できるのは、iOS Safariのみです。');
    }
  }


