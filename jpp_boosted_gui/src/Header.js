import React from 'react'
import logo from './jpp_small.png';
import './Header.css';

class Header extends React.Component {
    render() {
        return (
            <header>
                <div className="Header">
                    <a href="/">
                        <img src={logo} className="App-logo" alt="logo" />
                        Boosted!
                    </a>
                </div>
            </header>
        );
    }
}
export default Header;