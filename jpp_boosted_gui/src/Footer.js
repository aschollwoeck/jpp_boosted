import React from 'react';
import { Link } from 'react-router-dom'

import './Footer.css';

class Footer extends React.Component {
    render() {
        const currentDate = new Date().getFullYear();
        return (
            <div className="jpp-footer">
                <div className="mt-1">
                    <ul className="footerLinks">
                        <li>
                            <Link to="/about">Über JPP-Boosted</Link>
                        </li>
                        <li className="bar">|</li>
                        <li>
                            <Link to="/submit">Beitragen</Link>
                        </li>
                        <li className="bar">|</li>
                        <li>
                            <Link to="/datenschutz">Datenschutz</Link>
                        </li>
                    </ul>
                </div>
                <div className="mt-1">
                    <div>
                        Wünsche, Anregungen, Ideen?
                            </div>
                    <div>
                        Schreib eine <a href="mailto:contact@jpp-boosted.de">Mail</a>!
                            </div>
                </div>
                
                <div className="mt-1">
                    &copy; {currentDate} jpp-boosted.de
                            </div>
            </div>
        );
    }
}
// <div className="mt-1">
//                     <a href="/">Steuer bei zu Kaffee oder helft bei Serverkosten</a>
//                     </div>
export default Footer;