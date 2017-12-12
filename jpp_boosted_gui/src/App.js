import React, { Component } from 'react';
import './App.css';
import Header from './Header';
import Footer from './Footer';
import Search from './Search';
import CarOverview from './CarOverview';
import Home from './Home';

class App extends Component {
  performSearch(t) {
    window.location = "/search/" + t;
  }

  render() {
    // let docId = 0;
    var match = new RegExp("/search/([0-9a-zA-Z]+)", "i").exec(window.location.href);
    if (match != null) {
      const cars = getVideos().filter(v => v.manufacturer.toLowerCase().indexOf(match[1]) > -1).map(c => {
        return (
          <CarOverview key={c.id} car={c} />
        );
      });
      return (
        <div className="App">
          <Header />
          <div className="jpp-content">
            <div className="container">
              <Search onSubmit={this.performSearch} searchTerm={match[1]} />
              {cars}
            </div>
          </div>
          <Footer />
        </div>
      );
    }

    return (
      <div className="App">
        <Header />
        <div className="jpp-content">
          <Home onSubmit={this.performSearch} />
        </div>
        <Footer />
      </div>
    );
  }
}

export default App;

function getVideos() {
  return [
    {
      id: 1,
      title: "C63 AMG Tuning",
      buildYear: 2016,
      manufacturer: "Mercedes-Benz",
      model: "C63 AMG",
      tuning: [
        {
          id: 1,
          stage: "Stage 1",
          description: "Software",
          measuredTime: [
            {
              id: 1,
              speedRange: "0-200",
              time: "8.7",
              youtubeUrl: ""
            },
            {
              id: 2,
              speedRange: "0-100",
              time: "3.4",
              youtubeUrl: ""
            }
          ]
        }
      ]
    },
    {
      id: 2,
      title: "BMW M2",
      buildYear: 2016,
      manufacturer: "BMW",
      model: "M2",
      tuning: [
        {
          id: 1,
          stage: "Serie",
          description: "Serienkomponenten",
          measuredTime: [
            {
              id: 1,
              speedRange: "0-100",
              time: "4.0",
              youtubeUrl: "",
            }
          ]
        },
        {
          id: 2,
          stage: "Stage 1",
          description: "Software",
          measuredTime: [
            {
              id: 1,
              speedRange: "0-100",
              time: "3.8",
              youtubeUrl: "",
            }
          ]
        }
      ]
    },
    {
      id: 1,
      title: "C63 AMG Tuning",
      buildYear: 2016,
      manufacturer: "Mercedes-Benz",
      model: "C63 AMG",
      tuning: [
        {
          id: 1,
          stage: "Stage 1",
          description: "Software",
          measuredTime: [
            {
              id: 1,
              speedRange: "0-200",
              time: "8.7",
              youtubeUrl: ""
            },
            {
              id: 2,
              speedRange: "0-100",
              time: "3.4",
              youtubeUrl: ""
            }
          ]
        }
      ]
    },
  ];
}