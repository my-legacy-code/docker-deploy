import React, {Component} from 'react';
import {BrowserRouter as Router, Route, Switch} from "react-router-dom";
import './App.css';
import ServicesComponent from "./components/Services.component";

class App extends Component {
    render() {
        return (
            <Router>
                <div className="App">
                    <nav class="ui inverted menu">
                        <div class="ui container">
                            <span class="header logo">Docker Deploy</span>
                        </div>
                    </nav>
                    <Switch>
                        <Route exact path='/' component={ServicesComponent} />
                    </Switch>
                </div>
            </Router>
        );
    }
}

export default App;
