export const UPDATE_SERVICE = 'UPDATE_SERVICE';
export const INIT_SERVICES = 'INIT_SERVICES';

export const updateService = (service) => ({
    type: UPDATE_SERVICE,
    payload: {
        repoName: `${service.service_config.namespace}/${service.service_config.name}`,
        service
    }
});

export const initServices = (services) => {
   return ({
    type: INIT_SERVICES,
    payload: services
})};