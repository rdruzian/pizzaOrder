import React, { Component } from 'react';
import Login from "./componentes/Login"
import {withRouter} from 'react-router-dom'
import "./assets/App.css"
import "./assets/index.css"

class App extends Component{
    logar(usuario, senha){
    }
    render() {
        return (
            <Login logar={this.logar.bind(this)}/>
        );
    }
}
export default withRouter(App);