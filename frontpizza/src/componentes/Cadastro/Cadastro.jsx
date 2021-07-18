import React, {Component} from "react";
import "./cadastro.css"

class Cadastro extends Component {
    constructor(props) {
        super(props);
        this.nome = "";
        this.tipo = "";
        this.endereco = "";
        this.numero = "";
        this.complemento = "";
        this.usuario = "";
        this.senha = "";
        this.confirmaSenha = "";
    }

    _handleMudancaNome(evento){
        evento.stopPropagation();
        this.nome = evento.target.value;
    }
    _handleMudancaTipo(evento){
        evento.stopPropagation();
        this.tipo = evento.target.value;
    }
    _handleMudancaEndereco(evento){
        evento.stopPropagation();
        this.endereco = evento.target.value;
    }
    _handleMudancaNumero(evento){
        evento.stopPropagation();
        this.numero = evento.target.value;
    }
    _handleMudancaComplemento(evento){
        evento.stopPropagation();
        this.complemento = evento.target.value;
    }
    _handleMudancaUsuario(evento){
        evento.stopPropagation();
        this.usuario = evento.target.value;
    }
    _handleMudancaSenha(evento){
        evento.stopPropagation();
        this.senha = evento.target.value;
    }
    _handleMudancaConfirmaSenha(evento){
        evento.stopPropagation();
        this.confirmaSenha = evento.target.value;
    }

    _cadastrar(evento){
        evento.preventDefault();
        evento.stopPropagation();
        if(this.senha === this.confirmaSenha) {
            this.props.cadastrar(this.nome, this.tipo, this.endereco, this.numero, this.complemento, this.usuario, this.senha);
        }
        else {
            console.log("Senhas devem ser iguais")
        }
    }
    render() {
        return (
            <form className="form-cadastro" onSubmit={this._cadastrar.bind(this)}>
                <input
                    type="text"
                    placeholder="Nome"
                    className="form-cadastro_input"
                    onChange={this._handleMudancaNome.bind(this)}/>
                <select name="tipo" id="tipo" className="form-cadastro_input" onChange={this._handleMudancaTipo.bind(this)}>
                    <option value=""></option>
                    <option value="Rua">Rua</option>
                    <option value="Avenida">Avenida</option>
                    <option value="Estrada">Estrada</option>
                </select>
                <input
                    type="text"
                    placeholder="Endereço"
                    className="form-cadastro_input"
                    onChange={this._handleMudancaEndereco.bind(this)}/>
                <input
                    type="text"
                    placeholder="Nro"
                    className="form-cadastro_input"
                    onChange={this._handleMudancaNumero.bind(this)}/>
                <input
                    type="text"
                    placeholder="Complemento"
                    className="form-cadastro_input"
                    onChange={this._handleMudancaComplemento.bind(this)}/>
                <input
                    type="text"
                    placeholder="Usuario"
                    className="form-cadastro_input"
                    onChange={this._handleMudancaUsuario.bind(this)}/>
                <input
                    type="password"
                    placeholder="Senha"
                    className="form-cadastro_input"
                    onChange={this._handleMudancaSenha.bind(this)}/>
                <input
                    type="password"
                    placeholder="Confirma senha"
                    className="form-cadastro_input"
                    onChange={this._handleMudancaConfirmaSenha.bind(this)}/>
                <button className="form-cadastro_input form-cadastro_submit">Cadastrar</button>
            </form>
        )
    }
}

export default Cadastro