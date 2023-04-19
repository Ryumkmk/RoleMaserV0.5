// const fileInput = document.getElementById('file-input');
// const fileInputLabel = document.getElementById('file-input-label');
// fileInput.addEventListener('change', () => {
//     const files = fileInput.files;
//     if (files.length === 0) {
//         fileInputLabel.textContent = 'ファイルを選択してください';
//     } else {
//         fileInputLabel.textContent = files[0].name;
//     }
// });

// HTMLファイルに<input id="file-input">と<label id="file-input-label">があることを確認してください。

// const fileInput = document.getElementById('file-input');
// const fileInputLabel = document.getElementById('file-input-label');

// // fileInputおよびfileInputLabelがnullでないことを確認してから、addEventListener()メソッドを呼び出す
// if (fileInput && fileInputLabel) {
//   fileInput.addEventListener('change', () => {
//     const files = fileInput.files;
//     if (files.length === 0) {
//         fileInputLabel.textContent = 'ファイルを選択してください';
//     } else {
//         fileInputLabel.textContent = files[0].name;
//     }
//   });
// } else {
//   console.log('fileInputまたはfileInputLabelが見つかりませんでした');
// }

// function copyValue(input, targetName, copyName) {
//     var targetInput = input.closest(".AM").querySelectorAll("[name='" + targetName + "']")[0];
//     var copyInput = input.closest(".AM").querySelectorAll("[name='" + copyName + "']")[0];
//     targetInput.value = input.value;
//     copyInput.value = input.value;
// }

// function copyValue2(input, targetName, copyName) {
//     var targetInput = input.closest(".PM").querySelectorAll("[name='" + targetName + "']")[0];
//     var copyInput = input.closest(".PM").querySelectorAll("[name='" + copyName + "']")[0];
//     targetInput.value = input.value;
//     copyInput.value = input.value;
// }

// function confirmDelete(event) {
//     event.preventDefault();
//     if (confirm("本当に削除しますか？")) {
//         // 削除処理
//         document.getElementById("delete-form").submit();
//     }
// }

// input要素を取得
const input = document.querySelector('input[name="drinkmain"]');

// pjs-item要素を取得
const items = document.querySelectorAll('.pjs-item');

// inputの入力値が変更されたら
input.addEventListener('input', function (event) {
    // 入力値を取得
    const inputValue = event.target.value;
    console.log(inputValue);
    // pjs-item要素をループ
    for (const item of items) {
        // itemの名前を取得
        const itemName = item.querySelector('.Names').textContent;
        // itemNameがinputValueを含むかどうかを判定
        if (itemName.includes(inputValue)) {
            // 背景色を変更
            item.style.backgroundColor = 'yellow';
        } else {
            // 背景色を初期化
            item.style.backgroundColor = 'red';
        }
    }
});
