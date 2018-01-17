import React from 'react'
import { Link } from 'react-router-dom'

import logo from './jpp_small.png';
import './Header.css';

class Header extends React.Component {
    render() {
        return (
            <header>
                <div className="header">
                    <div className="header-name">
                        <Link className="header-home-link" to="/">JPP-Boosted!</Link>
                        <Link className="ml-2 header-link" to="/projects">Alle Projekte</Link>
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