import React, {Component} from 'react';
import './Services.component.css';

class ServicesComponent extends Component {

    constructor(props) {
        super(props);
        this.state = {
            services: []
        }
    }

    _renderServices() {
        this.state.services.map(service =>
            <tr>
                <td>
                    {service.serviceConfig.Name}
                </td>
                <td>
                    <div className="status status-{{.Status}}">
                    </div>
                </td>
                <td>
                    {service.deployed_at}
                </td>
                <td>
                    <a href="https://hub.docker.com/r/{{.ServiceConfig.Namespace}}">
                        {service.serviceConfig.namespace}
                    </a>
                    /
                    <a href={`https://hub.docker.com/r/${service.serviceConfig.namespace}/${service.serviceConfig.name}`}>
                        {service.Name}
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