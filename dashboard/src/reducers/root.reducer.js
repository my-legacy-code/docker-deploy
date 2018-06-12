import {ADD_SERVICES, UPDATE_SERVICE} from '../actions/service.action';

const initialState = {
    services: []
};

const rootReducer = (state = initialState, action) => {
    switch (action.type) {
        case UPDATE_SERVICE:
            console.log(action.payload);
            console.log({...state.services, [action.payload.repoName]: [action.payload.service]});
            return {...state, services: {...state.services, [action.payload.repoName]: action.payload.service}};
        case ADD_SERVICES:
            return {...state, services: action.payload};
        default:
            return state;
    }
};
export default rootReducer;