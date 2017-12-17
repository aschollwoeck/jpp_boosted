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
      const cars = getVideos().filter(v => v.baseModel.manufacturer.toLowerCase().indexOf(match[1]) > -1).map(c => {
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
      modificationStart: 2016,
      modificationEnd: 2016,
      rating: 4.5,
      baseModel: {
        id: 1,
        manufacturer: "Mercedes-Benz",
        manufacturerUrl: "",
        model: "C63 AMG",
        manufactureDate: 2016,
        bodyType: "Hatchback",
        aspiration: "Turbo",
        transmission: "Manual",
        doors: 3,
        seats: 4,
        weigth: 1700,
        height: 200,
        width: 200,
        length: 473,
        drive: "rear-wheel",
        engine: {
          id: 1,
          vMax: 270,
          capacity: 3400,
          fuel: "petrol",
          cylinder: 8,
          buildTimeStart: 1999,
          buildTimeEnd: 2000,
        },
        tankCapacity: 63,
        wheelbase: 2735,
        maxWeight: 1785,
        emptyWeight: 0,
      },
      tuning: [
        {
          id: 1,
          stage: "Stage 2",
          description: "Software, Krümmer",
          horsePower: 563,
          powerWheels: 258,
          torque: 634,
          modificationDate: 2016,
          modifiedParts: [
            {
              id: 1,
              part: "Krümmer",
              youtubeUrl: "",
              manufacturer: "Hersteller",
              manufacturerUrl: "",
              partUrl: ""
            }
          ],
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