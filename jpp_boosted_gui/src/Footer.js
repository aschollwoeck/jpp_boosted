import React from 'react';
import './Footer.css';

class Footer extends React.Component {
    render() {
        const currentDate = new Date().getFullYear();
        return (
            <div className="jpp-footer">
                <div className="mt-1">
                    <ul className="footerLinks">
                        <li>
                            <a href="/about">Über JPP-Boosted</a>
                        </li>
                        <li className="bar">|</li>
                        <li>
                            <a href="/submit">Beitragen</a>
                        </li>
                        <li className="bar">|</li>
                        <li>
                            <a href="/datenschutz">Datenschutz</a>
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
                    <a href="/">Steuer bei zu Kaffee oder helft bei Serverkosten</a>
                    </div>
                <div className="mt-1">
                    &copy; {currentDate} jpp-boosted.de
                            </div>
            </div>
        );
    }
}
export default Footer;