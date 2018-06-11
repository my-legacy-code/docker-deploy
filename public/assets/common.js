window.onload = () => {
    let username = "user";
    let webSocket = new WebSocket(`ws://${window.location.host}/api/connections/${username}`);

    webSocket.addEventListener('open', event => {
        console.log('Established WebSocket connection.');
    });

    webSocket.addEventListener('message', event => {
        let message = JSON.parse(event.data);
        console.log(message);
    });
};