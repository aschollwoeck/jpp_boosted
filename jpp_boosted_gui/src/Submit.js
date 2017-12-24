import React from 'react';

import SubmitStage from './SubmitStage';

import './Submit.css';

class Submit extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            title: "",
            manufacturer: "",
            model: "",
            tuning: [this.getNewTuning()],
        }
    }

    // {
    //     id: 1,
    //     title: "C63 AMG Tuning",
    //     modificationStart: 2016,
    //     modificationEnd: 2016,
    //     rating: 4.5,
    //     baseModel: {
    //       id: 1,
    //       manufacturer: "Mercedes-Benz",
    //       manufacturerUrl: "",
    //       model: "C63 AMG",
    //       manufactureDate: 2016,
    //       bodyType: "Hatchback",
    //       aspiration: "Turbo",
    //       transmission: "Manual",
    //       doors: 3,
    //       seats: 4,
    //       weigth: 1700,
    //       height: 200,
    //       width: 200,
    //       length: 473,
    //       drive: "rear-wheel",
    //       engine: {
    //         id: 1,
    //         vMax: 270,
    //         capacity: 3400,
    //         fuel: "petrol",
    //         cylinder: 8,
    //         buildTimeStart: 1999,
    //         buildTimeEnd: 2000,
    //       },
    //       tankCapacity: 63,
    //       wheelbase: 2735,
    //       maxWeight: 1785,
    //       emptyWeight: 0,
    //     },
    //     tuning: [
    //       {
    //         id: 1,
    //         stage: "Serie",
    //         description: "Herstellerangaben",
    //         horsePower: 563,
    //         powerWheels: 258,
    //         torque: 634,
    //         modificationDate: 2016,
    //         modifiedParts: [
        // {
        //     id: 1,
        //     part: "Krümmer",
        //     youtubeUrl: "",
        //     manufacturer: "Hersteller",
        //     manufacturerUrl: "",
        //     partUrl: ""
        //   }
    //         ],
    //         measuredTime: [
    //           {
    //             id: 1,
    //             speedRange: "0-200",
    //             time: "10.3",
    //             youtubeUrl: ""
    //           },
    //           {
    //             id: 2,
    //             speedRange: "0-100",
    //             time: "4.2",
    //             youtubeUrl: ""
    //           }
    //         ]
    //       },

    getNewTuning() {
        return {name: "", description: "", ps: "", nm: "", date: "", measuredTime: [], modifiedParts: []};
    }

    handleUserInput = (e) => {
        console.log(e.target.name + " " + e.target.value);
        const name = e.target.name;
        const value = e.target.value;
        this.setState({[name]: value});
    }

    addStage = (e) => {
        e.preventDefault();
        console.log("Neuer Umbau");
        const tuning = this.state.tuning;
        tuning.push(this.getNewTuning());

        this.setState({tuning: tuning});
    }

    tuningChanged = (index, value) => {
        console.log(value);

        const tunings = this.state.tuning.slice();
        tunings[index] = value;
        this.setState({tuning: tunings});
    }

    render() {
        const stages = this.state.tuning.map((t, i) => {
            return (
                <SubmitStage key={"tuning"+i} index={i} tuning={t} onChange={this.tuningChanged} />
            );
        });

        return (
            <form>
                <h2>Neues Fahrzeug</h2>
                <div className="form-group">
                    <label className="keyName" htmlFor="title">Titel</label>
                    <input type="text" name="title" placeholder="BMW V8"
                        value={this.state.title}
                        onChange={this.handleUserInput} />
                </div>
                <div className="form-group">
                    <div>
                        <label className="keyName" htmlFor="manufacturer">Hersteller</label>
                        <input type="text" name="manufacturer" placeholder="BMW"
                            value={this.state.manufacturer}
                            onChange={this.handleUserInput} />
                    </div>
                    <div>
                        <label className="keyName" htmlFor="model">Model</label>
                        <input type="text" name="model" placeholder="z.B M3"
                            value={this.state.model}
                            onChange={this.handleUserInput} />
                    </div>
                </div>
                <h4>Umbauten</h4>
                <button className="btn btn-primary" onClick={this.addStage}>
                    + Umbau hinzufügen
                </button>
                <hr />
                {stages}
                <button type="submit" className="btn btn-primary">
                    Absenden
                </button>
            </form>
        );
    }
}
export default Submit;