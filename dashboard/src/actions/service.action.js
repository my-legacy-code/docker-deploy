export const UPDATE_SERVICE = 'UPDATE_SERVICE';
export const ADD_SERVICES = 'ADD_SERVICES';

export const updateService = (service) => ({
    type: UPDATE_SERVICE,
    payload: {
        repoName: `${service.service_config.namespace}/${service.service_config.name}`,
        service
    }
});

export const addServices = (services) => {
   return ({
    type: ADD_SERVICES,
    payload: services
})};