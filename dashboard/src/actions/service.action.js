export const UPDATE_SERVICE = 'UPDATE_SERVICE';
export const UPDATE_SERVICES = 'UPDATE_SERVICES';

export const updateService = (service) => ({
    type: UPDATE_SERVICE,
    payload: {
        repoName: `${service.service_config.namespace}/${service.service_config.name}`,
        service
    }
});

export const updateServices = (services) => {
   return ({
    type: UPDATE_SERVICES,
    payload: services
})};