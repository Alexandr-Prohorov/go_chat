function sendMessage() {

}

async function getOneChat () {
    try {
        const resp = await fetch(url + `${userId}`, {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json'
            },
            credentials: 'include'
        })
        const data = resp.json()
        console.log(data)
        return data
    } catch (e) {
        console.log('Error:', e)
    }
}

(async function () {
    const chatData = await getOneChat()
    console.log(chatData)
})()