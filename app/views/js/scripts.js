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

