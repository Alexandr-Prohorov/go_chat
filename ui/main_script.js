const url = 'http://localhost:8080/users'
const usersList = document.querySelector('.chat-list')

async function getUsers () {
    try {
        const resp = await fetch(url, {
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

(async function() {
    let users =  await getUsers();

    users.forEach(elem => {
        const user = document.createElement('div');
        user.classList.add('chat-item');

        const userSpan = document.createElement('span');
        userSpan.classList.add('chat-name');
        userSpan.textContent = elem.Username;
        user.appendChild(userSpan);


        console.log(usersList)
        usersList.appendChild(user);
        console.log('гатова')
    });

    const chat_items = document.querySelectorAll('.chat-item')
    chat_items.forEach(elem => {
        elem.addEventListener('click', function(event) {
            console.log(event)
        })
    })
})();
