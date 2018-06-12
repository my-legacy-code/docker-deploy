import {addServices} from './actions/service.action';

const baseURL = `localhost:8000`;

const ws = new WebSocket(`ws://${baseURL}/api/connect`);

export const initWebSoket = (store) => {
    const INITIAL_SERVICE_STATES = 'initial_service_states';

    ws.addEventListener('open', () => {
        console.log('Connection established.');
    });

    ws.addEventListener('message', event => {
        let message = JSON.parse(event.data);
        console.log(message);

        switch(message.type) {
            case INITIAL_SERVICE_STATES:
                store.dispatch(addServices(message.body));
                break;
        }
    });

    ws.addEventListener('close', () => {
        console.log('Connection closed.');
    });
};

export const emit = (type, data) => ws.send(data);