import {initServices, updateService} from './actions/service.action';

const baseURL = `localhost:8000`;

const ws = new WebSocket(`ws://${baseURL}/api/connect`);

export const initWebSoket = (store) => {
    const INITIAL_SERVICE_STATES = 'initial_service_states';
    const UPDATE_SERVICE_STATE = 'update_service_state';

    ws.addEventListener('open', () => {
        console.log('Connection established.');
    });

    ws.addEventListener('message', event => {
        let message = JSON.parse(event.data);

        switch(message.type) {
            case INITIAL_SERVICE_STATES:
                store.dispatch(initServices(message.body));
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