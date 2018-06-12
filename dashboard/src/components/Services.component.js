import React, {Component} from 'react';
import './Services.component.css';
import store from "../store";

class ServicesComponent extends Component {

    constructor(props) {
        super(props);
        this.state = {
            services: []
        };

        store.subscribe(() => {
            this.setState({...this.state, services: store.getState().services});
        });
    }

    _renderServices() {
        return this.state.services.map(service =>
            <tr>
                <td>
                    {service.service_config.name}
                </td>
                <td>
                    <div className={`status status-${service.status}`}>
                    </div>
                </td>
                <td>
                    {service.deployed_at}
                </td>
                <td>
                    <a href={`https://hub.docker.com/r/${service.service_config.namespace}`} >
                    {service.service_config.namespace}
                        </a>
                        /
                        <a href={`https://hub.docker.com/r/${service.service_config.namespace}/${service.service_config.name}`}>
                        {service.service_config.name}
                        </a>
                        </td>
                        </tr>
                        )
                        }

                       render() {
                    return (
                    <div className="Services ui main container">
                    <h1 className="ui header">Services</h1>
                    <table className="ui very basic collapsing celled table">
                    <thead>
                    <tr>
                    <th>
                    Name
                    </th>
                    <th>
                    Status
                    </th>
                    <th>
                    Deployed At
                    </th>
                    <th>Repository</th>
                    </tr>
                    </thead>
                    <tbody>
                    {this._renderServices()}
                    </tbody>
                    </table>
                    </div>
                    );
                }
                    }

                    export default ServicesComponent;