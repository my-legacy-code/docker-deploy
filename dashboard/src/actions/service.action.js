export const ADD_SERVICE = 'ADD_SERVICE';
export const ADD_SERVICES = 'ADD_SERVICES';

export const addService = (service) => ({
    type: ADD_SERVICE,
    service
});

export const addServices = (services) => {
   return ({
    type: ADD_SERVICES,
    payload: Object.keys(services).map(serviceName => services[serviceName])
})};