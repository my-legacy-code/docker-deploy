import {ADD_SERVICE, ADD_SERVICES} from './actions/service.action';

const initialState = {
  services: []
};

const rootReducer = (state = initialState, action) => {
    switch (action.type) {
        case UPDATE_SERVICE:
            return {...state, services: [...state.services, action.payload]};
        case ADD_SERVICES:
            return {...state, services: action.payload};
        default:
            return state;
    }
};
export default rootReducer;