import React from 'react'
import logo from './jpp_small.png';
import boost from './boost_out_black.svg';
import './Header.css';

class Header extends React.Component {
    render() {
        return (
            <header>
                <div className="header">
                    <div className="header-name">
                        <a className="header-home-link" href="/">JPP-Boosted!</a>
                    </div>
                    <div className="header-logo">
                        <a target="blank" href="https://jp-performance.de">
                            <img className="logo" src={logo} alt="logo" />
                        </a>
                    </div>
                </div>
            </header>
        );
    }
}
export default Header;