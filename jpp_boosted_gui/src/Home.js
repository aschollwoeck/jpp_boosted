import React from 'react';

import Search from './Search';

import './Home.css';

class Home extends React.Component {
    render() {
        return (
            <div className="home-content">
                <div>
                    <h3>Modifizierte Kraftfahrzeuge der <a target="blank" href="https://www.jp-performance.de">JP Performance GmbH</a></h3>
                </div>
                <div>
                    <h6>Projektautos, gemessene Zeiten und weitere Informationen zu Umbauten</h6>
                </div>
                <Search onSubmit={this.props.onSubmit} />
            </div>
        );
    }
}
export default Home;