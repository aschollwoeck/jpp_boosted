import React from 'react';
import { Switch, Route } from 'react-router-dom';
import Home from './Home';
import Privacy from './Privacy';
import Submit from './Submit';
import About from './About';
import ProjectSearch from './ProjectSearch';
import Projects from './Projects';

class Main extends React.Component {
    render() {
        return (
            <Switch>
                <Route exact path='/' component={Home} />
                <Route exact path='/search' component={ProjectSearch} />
                <Route path='/search/:term' component={ProjectSearch} />
                <Route exact path='/submit' component={Submit} />
                <Route path='/submit/:id' component={Submit} />
                <Route path='/privacy' component={Privacy} />
                <Route path='/about' component={About} />
                <Route path='/datenschutz' component={Privacy} />
                <Route path='/projects' component={Projects} />
            </Switch>
        );
    }
}
export default Main;