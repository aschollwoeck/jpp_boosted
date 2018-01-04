import React, { Component } from 'react';

import './App.css';
import Header from './Header';
import Footer from './Footer';
import Main from './Main';

class App extends Component {

  render() {
    return (
      <div className="App">
        <Header />
        <div className="jpp-content">
          <div className="container">
            <Main />
          </div>
        </div>
        <Footer />
      </div>
    );

    // var match = new RegExp("/search/(.+)", "i").exec(window.location.href);
    // if (match != null) {
    //   var projects = "";

    //   if (this.state.projects !== null) {
    //     projects = this.state.projects.map(p => {
    //       return (
    //         <ProjectOverview key={"project" + p.id} project={p} />
    //       );
    //     });
    //   }
    //   return (
    //     <div className="App">
    //       <Header />
    //       <div className="jpp-content">
    //         <div className="container">
    //           <Search onSubmit={this.performSearch} searchTerm={match[1]} />
    //           {projects}
    //         </div>
    //       </div>
    //       <Footer />
    //     </div>
    //   );

    // }

    // if (window.location.href.indexOf("datenschutz") > 0) {
    //   return (
    //     <div className="App">
    //       <Header />
    //       <div className="jpp-content">
    //         <div className="container">
    //           <Privacy />
    //         </div>
    //       </div>
    //       <Footer />
    //     </div>
    //   );
    // }

    // if (window.location.href.indexOf("submit") > 0) {
    //   var match = new RegExp("/submit/(.+)", "i").exec(window.location.href);
    //   if (match != null) {
    //     return (
    //       <div className="App">
    //         <Header />
    //         <div className="jpp-content">
    //           <div className="container">
    //             <Submit ProjectID={match[1]} />
    //           </div>
    //         </div>
    //         <Footer />
    //       </div>
    //     );
    //   }
    //   else {
    //     return (
    //       <div className="App">
    //         <Header />
    //         <div className="jpp-content">
    //           <div className="container">
    //             <Submit />
    //           </div>
    //         </div>
    //         <Footer />
    //       </div>
    //     );
    //   }
    // }

    // if (window.location.href.indexOf("about") > 0) {
    //   return (
    //     <div className="App">
    //       <Header />
    //       <div className="jpp-content">
    //         <div className="container">
    //           <About />
    //         </div>
    //       </div>
    //       <Footer />
    //     </div>
    //   );
    // }

    // return (
    //   <div className="App">
    //     <Header />
    //     <div className="jpp-content">
    //       <Home onSubmit={this.performSearch} />
    //     </div>
    //     <Footer />
    //   </div>
    // );
  }
}

export default App;