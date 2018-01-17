import React from 'react';
import ProjectOverview from './ProjectOverview';
import Search from './Search';
import { MediaBlock, TextBlock } from 'react-placeholder/lib/placeholders';
import 'react-placeholder/lib/reactPlaceholder.css';

async function getProjects() {
    var prom = await fetch("/api/projects")
        .then(resp => resp.json())
        .then(data => {
            return data;
        })
        .catch(e => console.log(e));

    if (prom === null)
        return [];
    return prom;
};

class Projects extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            projects: [],
            loading: true,
            searchTerm: "",
        };

        this.loadProjects();
    }

    loadProjects = () => {
        getProjects().then(data => {
            this.setState({ projects: data, loading: false });
        });
    }

    performSearch = (searchTerm) => {
        const { history } = this.props;
        history.push('/search/' + searchTerm);
    }

    render() {
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
export default Projects;