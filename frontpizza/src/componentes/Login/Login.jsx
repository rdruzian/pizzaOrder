import React, {Component} from "react";
import "./login.css"

class Login extends Component {
    constructor(props) {
        super(props);
        this.usuario = "";
        this.senha = "";
    }

    _handleMudancaUsuario(evento) {
        evento.stopPropagation();
        this.usuario = evento.target.value;
    }

    _handleMudancaSenha(evento) {
        evento.stopPropagation();
        this.senha = evento.target.value;
    }

    _login(evento) {
        evento.preventDefault();
        evento.stopPropagation();
        this.props.logar(this.usuario, this.senha);
    }

    render() {
        return (
            <form className="form-cadastro" onSubmit={this._login.bind(this)}>
                <input
                    type="text"
                    placeholder="usuário"
                    className="form-cadastro_input"
                    onChange={this._handleMudancaUsuario.bind(this)}
                />
                <input
                    type="password"
                    placeholder="senha"
                    className="form-cadastro_input"
                    onChange={this._handleMudancaSenha.bind(this)}
                />
                <div className="form-cadastro_button">
                    <button className="form-cadastro_input form-cadastro_submit">Login</button>
                    <button
                        className="form-cadastro_input form-cadastro_cadastro" onSubmit={() => this.props.history.push('/cadastro')}>Cadastre-se
                    </button>
                </div>
            </form>
        )
    }
}

export default Login