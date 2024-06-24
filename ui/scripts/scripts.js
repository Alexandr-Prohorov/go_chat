document.getElementById('loginForm').addEventListener('submit', function(event) {
    event.preventDefault();
    const username = document.getElementById('username').value;
    const password = document.getElementById('password').value;
    const url = 'http://192.168.137.149:8080/auth'
    const body = {
        Login: username,
        Password: password
    }

    // Implement your login logic here
    fetch(url, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(body),
        credentials: 'include' // Обязательно включите этот параметр для отправки и получения куки
    }).then(resp => {
        if (resp.ok)  window.location.href = '/main';
        console.log(resp)
    }).catch(error => console.log('Error:', error));
});
