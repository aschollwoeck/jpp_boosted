import React from 'react';
import './Search.css';

class Search extends React.Component {
    constructor(props) {
        super(props);
        this.state = { searchTerm: this.props.searchTerm };

        this.handleSearchTermChange = this.handleSearchTermChange.bind(this);
        this.submit = this.submit.bind(this);
    }

    handleSearchTermChange(e) {
        this.setState({
            searchTerm: e.target.value
        });
    }

    submit(e) {
        e.preventDefault();
        this.props.onSubmit(this.state.searchTerm);
    }

    render() {
        return (
            <div>
                <form onSubmit={this.submit} id="search">
                    <div className="Search">
                        <input type="search" placeholder="Nick, Hersteller, Modell..." value={this.state.searchTerm} onChange={this.handleSearchTermChange} />
                        <span className="input-group-btn">
                            <button className="btn">
                                <i className="fas fa-search"></i>
                            </button>
                        </span>
                    </div>
                </form>
            </div>
        );
    }
}
export default Search;