import React, {Component} from 'react';
import {BrowserRouter as Router, Route, Switch} from 'react-router-dom';
import './App.css';
import ServicesComponent from './components/Services.component';
import {initWebSocket} from './ws';
import store from './store';

class App extends Component {
    constructor(props) {
        super(props);
        initWebSocket(store);
    }

    render() {
        return (
            <Router>
                <div className="App">
                    <nav className="ui inverted menu">
                        <div className="ui container">
                            <span className="header logo">Docker Deploy</span>
                        </div>
                    </nav>
                    <Switch>
                        <Route exact path='/' render={(props) => <ServicesComponent />}/>
                    </Switch>
                </div>
            </Router>
        );
    }
}

export default App;
