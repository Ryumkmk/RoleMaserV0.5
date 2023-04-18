const fileInput = document.getElementById('file-input');
const fileInputLabel = document.getElementById('file-input-label');
fileInput.addEventListener('change', () => {
    const files = fileInput.files;
    if (files.length === 0) {
        fileInputLabel.textContent = 'ファイルを選択してください';
    } else {
        fileInputLabel.textContent = files[0].name;
    }
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