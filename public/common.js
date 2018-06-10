window.onload = () => {
    let username = "user";
    let webSocket = new WebSocket(`ws://192.168.1.23:8080/api/connections/${username}`);

    webSocket.addEventListener('open', event => {
        console.log('Hello, server!');
    });

    webSocket.addEventListener('message', event => {
        console.log(`message from server ${event.data}`);
    });
};