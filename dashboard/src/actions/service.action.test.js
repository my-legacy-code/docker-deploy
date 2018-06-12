import * as actions from './service.action';
import {UPDATE_SERVICE} from './service.action';

describe('service actions', () => {
    it('should create a action to update service state', () => {
        const service = {
            deployed_at: "2018-06-11T18:13:17.108431754-07:00",
            service_config: {
                name: `teamy_api`,
                namespace: `teamyapp`,
                docker_run_args: [`-p`, `8001:3000`]
            },
            status: "deploying"
        };

        const expectedAction = {
            type: UPDATE_SERVICE,
            payload: {
                repoName: 'teamyapp/teamy_api',
                service: {
                    deployed_at: "2018-06-11T18:13:17.108431754-07:00",
                    service_config: {
                        name: `teamy_api`,
                        namespace: `teamyapp`,
                        docker_run_args: [`-p`, `8001:3000`]
                    },
                    status: "deploying"
                }
            }
        };

        expect(actions.updateService(service)).toEqual(expectedAction);

    });
});