import React from 'react';
import './Submit.css';

class Submit extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            title: "",
            manufacturer: "",
            model: "",
            tuning: [{ps: "", nm: ""}],
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

    handleUserInput = (e) => {
        console.log(e.target.name + " " + e.target.value);
        const name = e.target.name;
        const value = e.target.value;
        this.setState({[name]: value});
    }

    render() {
        return (
            <form>
                <h2>Neues Fahrzeug</h2>
                <div className="form-group">
                    <label htmlFor="title">Titel</label>
                    <input type="text" name="title" placeholder="BMW V8"
                        value={this.state.title}
                        onChange={this.handleUserInput} />
                </div>
                <div className="form-group">
                    <label htmlFor="manufacturer">Hersteller</label>
                    <input type="text" name="manufacturer" placeholder="BMW"
                        value={this.state.manufacturer}
                        onChange={this.handleUserInput} />
                    <label htmlFor="model">Model</label>
                    <input type="text" name="model" placeholder="z.B M3"
                        value={this.state.model}
                        onChange={this.handleUserInput} />
                </div>
                <h4>Umbauten</h4>
                <div className="form-group">
                    <label htmlFor="stage-name">Name</label>
                    <input type="text" name="stage-name" placeholder="Serie, Stage 1..." />
                </div>
                <div className="form-group">
                    <label htmlFor="stage-description">Beschreibung</label>
                    <textarea name="stage-description" placeholder="Software, Abgasanlage" />
                </div>
                <h5>Messungen</h5>
                <div className="form-group">
                    <label htmlFor="stage-ps">PS</label>
                    <input type="text" name="stage-ps" placeholder="414"
                        value={this.state.tuning[0].ps}
                        onChange={this.handleUserInput} />
                    <label htmlFor="tuning[0].nm">NM</label>
                    <input type="text" name="tuning[0].nm" placeholder="510"
                        value={this.state.tuning[0].nm}
                        onChange={this.handleUserInput} />
                </div>
                <div className="form-group">
                    <label htmlFor="stage-date">Datum</label>
                    <input type="date" name="stage-date" />
                </div>
                <div className="form-group">
                    <label htmlFor="stage-speedRange">KM/H</label>
                    <input type="text" name="stage-speedRange" placeholder="0-100" />
                    <label htmlFor="stage-speedTime">Zeit (Sek.)</label>
                    <input type="text" name="stage-speedTime" placeholder="3.8" />
                </div>
                <div className="form-group">
                    <label htmlFor="stage-speedYoutube">YouTube</label>
                    <input type="text" name="stage-speedYoutube" placeholder="youtube.com/watch?v=v1kFDOnueug" />
                </div>
                <h5>Teile</h5>
                <div className="form-group">
                    <label htmlFor="stage-part-name">Teilename</label>
                    <input type="text" name="stage-part-name" placeholder="Krümmer" />
                    <label htmlFor="stage-part-url">URL</label>
                    <input type="text" name="stage-part-url" placeholder="" />
                </div>
                <div className="form-group">
                    <label htmlFor="stage-part-manufacturer">Teilehersteller</label>
                    <input type="text" name="stage-part-manufacturer" placeholder="" />
                    <label htmlFor="stage-part-manufacturer-url">URL</label>
                    <input type="text" name="stage-part-manufacturer-url" placeholder="" />
                </div>
                <button type="submit" className="btn btn-primary">
                    Absenden
                </button>
            </form>
        );
    }
}
        //     id: 1,
        //     part: "Krümmer",
        //     youtubeUrl: "",
        //     manufacturer: "Hersteller",
        //     manufacturerUrl: "",
        //     partUrl: ""
export default Submit;