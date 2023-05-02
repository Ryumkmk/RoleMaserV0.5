const btn = document.querySelector('.btn-box button');
if (btn) {
    btn.addEventListener('click', () => {
        const more = document.querySelector('.more');
        more.classList.toggle('appear');

        if (btn.textContent == "アプリをホームに追加") {
            btn.textContent = "やり方";
        } else {
            btn.textContent = "アプリをホームに追加";
        }
    });
}