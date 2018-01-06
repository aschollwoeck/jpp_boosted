import React from 'react';
import ProjectOverview from './ProjectOverview';
import Search from './Search';
// import { Redirect } from 'react-router-dom';
// import ReactPlaceholder from 'react-placeholder';
import { MediaBlock, TextBlock } from 'react-placeholder/lib/placeholders';
import 'react-placeholder/lib/reactPlaceholder.css';
import './ProjectSearch.css';

async function getProjects(value) {
    if (value === undefined)
        return [];

    var prom = await fetch("/api/projects/include?search=" + value)
        .then(resp => resp.json())
        .then(data => {
            return data;
        })
        .catch(e => console.log(e));

    if (prom === null)
        return [];
    return prom;
};

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
            this.loadProjects(this.props.match.params.term);
        }
    }

    loadProjects = (value) => {
        getProjects(value).then(data => {
            this.setState({ projects: data, loading: false });
        });
    }

    componentWillReceiveProps(nextProps) {
        this.loadProjects(nextProps.match.params.term);
    }

    performSearch = (searchTerm) => {
        // this.setState({redirect: searchTerm});
        // window.location = '/search/' + searchTerm;
        const { history } = this.props;
        history.push('/search/' + searchTerm);
        this.setState({ loading: true });
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