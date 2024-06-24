const userId = localStorage.getItem('userId')
const messagesList = document.querySelector('.chat-messages')
const url = 'http://192.168.137.149:8080/'
let socket;

async function getOneChat () {
    try {
        const resp = await fetch(`${url}chat-room/${userId}/chat`, {
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

async function getMessages (chatId) {
    try {
        const resp = await fetch(`${url}chat-room/${userId}/messages/${chatId}`, {
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

function sendMessage() {
    const messageInput = document.getElementById('message-input');
    const message = messageInput.value;
    socket.send(JSON.stringify({ content: message }));
    messageInput.value = '';
}

(async function () {
    const chatData = await getOneChat()
    console.log(chatData)
    const messages = await getMessages(chatData[0].ChatId)
    console.log(chatData)
    console.log(messages)
    messages.forEach(message => {
        const messageContainer = document.createElement('div');
        messageContainer.classList.add('chat-message');

        const messageContent = document.createElement('div');
        messageContent.classList.add('message-content');
        messageContent.textContent = message.Content;
        console.log(message)

        messageContainer.appendChild(messageContent);
        messagesList.appendChild(messageContainer)

    })

    socket = new WebSocket(`ws://${window.location.host}/ws/chat-room/${chatData[0].ChatId}`);

    socket.onmessage = (event) => {
        const message = JSON.parse(event.data);
        const messageElement = document.createElement('div');
        messageElement.classList.add('message-content');
        messageElement.textContent = `${message.content}`;
        messagesList.appendChild(messageElement);
    };

    socket.onopen = () => {
        console.log('WebSocket connection established');
    };

    socket.onclose = () => {
        console.log('WebSocket connection closed');
    };

    socket.onerror = (error) => {
        console.error('WebSocket error:', error);
    };
})()