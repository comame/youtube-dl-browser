/** @type {HTMLInputElement} saveDirInput */
const saveDirInput = document.getElementById('save_dir')

saveDirInput.addEventListener('input', () => {
    chrome.storage.sync.set({
        save_dir: saveDirInput.value
    })
})

chrome.storage.sync.get('save_dir', ({ save_dir }) => {
    saveDirInput.value = save_dir ?? '/tmp'
})
