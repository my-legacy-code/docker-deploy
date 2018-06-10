window.onload = () => {
    let username = "user";
    let webSocket = new WebSocket(`ws://192.168.1.23:8080/api/connections/${username}`);

    webSocket.addEventListener('open', event => {
        console.log('Established WebSocket connection.');
    });

    webSocket.addEventListener('message', event => {
        let message = JSON.parse(event.data);
        console.log(message);
    });
};