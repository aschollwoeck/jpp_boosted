import React from 'react';
import ProjectOverview from './ProjectOverview';
import Search from './Search';
import { Redirect } from 'react-router-dom';
// import ReactPlaceholder from 'react-placeholder';
import {MediaBlock,TextBlock} from 'react-placeholder/lib/placeholders';
import 'react-placeholder/lib/reactPlaceholder.css';
import './ProjectSearch.css';

class ProjectSearch extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            projects: [],
            loading: true,
            searchTerm: this.props.match.params.term,
            redirect: undefined,
        };

        if (this.props.match.params.term != null) {
            fetch("/api/projects/include?search=" + this.props.match.params.term)
                .then(resp => resp.json())
                .then(data => {
                    // setTimeout(() => this.setState({ projects: data, loading: false }), 8000);
                    this.setState({ projects: data, loading: false });
                })
                .catch(e => console.log(e));
        }
    }

    performSearch = (searchTerm) => {
        // this.setState({redirect: searchTerm});
        window.location = '/search/' + searchTerm;
        // const { history } = this.props;
        // history.push('/search/' + searchTerm);
        // this.setState({loading: true });

        // fetch("/api/projects/include?search=" + searchTerm)
        // .then(resp => resp.json())
        // .then(data => {
        //     setTimeout(() => this.setState({ projects: data, loading: false }), 2000);
        //     // this.setState({ projects: data, loading: false });
        // })
        // .catch(e => {
        //     console.log(e);
        //     this.setState({ loading: false });
        // });
    }

    render() {


        // const redirect = this.state.redirect;
        // if(redirect !== undefined && '/search/' + redirect !== this.props.location.pathname) {
        //     return <Redirect to={{
        //         pathname: '/search/' + redirect, 
        //         state: {from: this.props.location}}} />
        // }

        var projects = "";
        if (this.state.loading) {
            projects = <div>
            <div className="jpp-placeholder">
                <MediaBlock color='#E0E0E0' rows={4} />
                <TextBlock color='#E0E0E0' rows={4} />
            </div>
            <div className="jpp-placeholder">
                <MediaBlock color='#E0E0E0' rows={4} />
                <TextBlock color='#E0E0E0' rows={4} />
            </div>
                    </div>
        } else {
            if (this.state.projects !== null) {
                projects = this.state.projects.map(p => {
                    return (
                        <ProjectOverview key={"project" + p.id} project={p} />
                    );
                });
            }
        }
        return (
            <div>
                <Search onSubmit={this.performSearch} searchTerm={this.state.searchTerm} />
                {projects}
            </div>
        );
    }
}
export default ProjectSearch;