chrome.action.onClicked.addListener(async (tab) => {
    const opt = await chrome.storage.sync.get('save_dir')
    const save_dir = opt.save_dir ?? '/tmp'


    const res = await fetch('http://localhost:5906/request-download', {
        method: 'POST',
        body: JSON.stringify({
            url: tab.url,
            filename: tab.title,
            save_dir
        })
    })

    console.log(res)
})
