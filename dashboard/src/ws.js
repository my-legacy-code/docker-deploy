import {updateServices, updateService} from './actions/service.action';

const hostname = window.location.hostname;
const port = hostname === 'localhost' ? 8000 : window.location.port;

const baseURL = `${hostname}:${port}`;

const ws = new WebSocket(`ws://${baseURL}/api/connect`);

export const initWebSocket = (store) => {
    const UPDATE_SERVICE_STATES = 'update_service_states';
    const UPDATE_SERVICE_STATE = 'update_service_state';

    ws.addEventListener('open', () => {
        console.log('Connection established.');
    });

    ws.addEventListener('message', event => {
        let message = JSON.parse(event.data);

        switch(message.type) {
            case UPDATE_SERVICE_STATES:
                store.dispatch(updateServices(message.body));
                break;
            case UPDATE_SERVICE_STATE:
                store.dispatch(updateService(message.body));
                break;
        }
    });

    ws.addEventListener('close', () => {
        console.log('Connection closed.');
    });
};

export const emit = (type, data) => ws.send(data);