// Imports da bilbioteca
import React, { Component } from 'react';
import { withRouter, Route } from 'react-router-dom'
import axios from 'axios'
// Imports de estilos
import "./assets/App.css"
import "./assets/index.css"
// Imports de componentes
import Login from "./componentes/Login"
import Cadastro from "./componentes/Cadastro";
import baseURL from "./assets/api"



class App extends Component{
    logar(usuario, senha){

        console.log("Entrou na tela")
    }

    async cadastrar(nome, tipo, endereco, numero, complemento, usuario, senha){
        const resp = await baseURL.post("customer/signin", {nome, tipo, endereco, numero, complemento, usuario, senha})
        console.log("Entrou na tela")
    }
    render() {
        return (
            <switch >
                <Route exact path="/">
                    <Login logar={this.logar.bind(this)}/>
                </Route>

                <Route path="/cadastro">
                    <Cadastro />
                </Route>
            </switch>
        );
    }
}
export default withRouter(App);