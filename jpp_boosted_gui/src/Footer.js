import React from 'react';
import './Footer.css';

class Footer extends React.Component {
    render() {
        const currentDate = new Date().getFullYear();
        return (
            <div className="jpp-footer">
                <div className="container">
                    <div className="row">
                        <div className="col-sm-4">
                            <strong>
                                <div>
                                    <a href="/">About</a>
                                </div>
                                <div>
                                    <a href="/">Submit</a>
                                </div>
                            </strong>
                        </div>
                        <div className="col-sm-4">
                            <div>
                                Wünsche, Anregungen, Ideen?
                            </div>
                            <div>
                                Schreib eine <a href="mailto:contact@jpp-boosted.de">Mail</a>!
                            </div>
                            <div>
                                &copy; {currentDate} jpp-boosted.de
                            </div>
                        </div>
                        <div className="col-sm-4">
                            <strong>
                                <div>
                                    <a href="/">Terms</a>
                                </div>
                                <div>
                                    <a href="/">Privacy</a>
                                </div>
                            </strong>
                        </div>

                    </div>
                </div>
            </div>
        );
    }
}
export default Footer;