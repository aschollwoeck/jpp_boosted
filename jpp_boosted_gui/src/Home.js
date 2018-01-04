import React from 'react';
import { Redirect } from 'react-router-dom';

import Search from './Search';

import './Home.css';

class Home extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            redirect: undefined,
        };
    }

    performSearch = (t) => {
        // window.location = "/search/" + t;
        this.setState({redirect: t});
    }

    render() {
        if(this.state.redirect !== undefined) {
            return <Redirect to={{
                pathname: '/search/' + this.state.redirect, 
                state: {from: "test"}}} />
        }

        return (
            <div className="home-content">
                <div>
                    <h3>Modifizierte Kraftfahrzeuge der <a target="blank" href="https://www.jp-performance.de">JP Performance GmbH</a></h3>
                </div>
                <div>
                    <h6>Projektautos, gemessene Zeiten und weitere Informationen zu Umbauten</h6>
                </div>
                <Search onSubmit={this.performSearch} />
            </div>
        );
    }
}
export default Home;