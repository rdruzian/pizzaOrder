import React from 'react';
import ReactDOM from 'react-dom';
import { BrowserRouter as Router, Route, Switch } from 'react-router-dom'
import './assets/index.css';
import App from './App';
import Login from "./componentes/Login"
import Cadastro from "./componentes/Cadastro"
// import Usuario from "./componentes/Usuario"

ReactDOM.render(
    // command:() => this.props.history.push('/')
    <Router>
        <App>
            <Switch>
                <Route exact path="/" component={Login}/>
                <Route path="/cadastro" component={Cadastro}/>
                {/*<Route path="/usuario" component={Usuario}/>*/}
            </Switch>
        </App>
    </Router>,
    document.getElementById('root')
);