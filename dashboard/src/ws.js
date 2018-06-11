const baseURL = `localhost:8000`;

const ws = new WebSocket(`ws://${baseURL}/api/connect`);

export const initWebSoket = (store) => {
    ws.addEventListener('open', () => {
        console.log('Connection established.');
    });

    ws.addEventListener('message', event => {
        let message = JSON.parse(event.data);
        console.log(message);
    });

    ws.addEventListener('close', () => {
        console.log('Connection closed.');
    });
};

export const emit = (type, data) => ws.send(data);