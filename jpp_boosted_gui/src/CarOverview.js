import React from 'react';
import './CarOverview.css';

import c63amg from './c63_amg.jpg';

class CarOverview extends React.Component {

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

    render() {
        const stages = this.props.car.tuning.map(t => {
            const power =   <div className="row">
                                <div className="col-4">PS: {t.horsePower}</div>
                                <div className="col-4">NM: {t.torque}</div>
                                <div className="col-4">Date: {t.modificationDate}</div>
                            </div>
            const times = t.measuredTime.map(m => {
                return (
                    <div key={m.id}>{m.speedRange}: {m.time} Sekunden</div>
                );
            });
            const modParts = t.modifiedParts.map(p => {
                return (
                    <div key={p.id}>{p.part} (<a href={p.manufacturerUrl}>{p.manufacturer}</a>)</div>
                );
            });
            return (
                <div key={t.id} className="jpp-search-result-stage">
                    <div className="jpp-search-result-stage-name mt-1">
                        <strong>{t.stage}</strong>
                    </div>
                    <div className="jpp-search-result-stage-content mt-1">
                        <div>{t.description}</div>
                        {power}
                        {times}
                        {modParts}
                    </div>
                    <div className=" mt-1">
                        <iframe width="100%" src="https://www.youtube.com/embed/Oe7qUG7ccRI" frameborder="0" gesture="media" allow="encrypted-media" allowfullscreen></iframe>
                    </div>
                </div>
            );
        });

        return (
            <div className="jpp-search-result col-12">
                <div className="">
                    <div className="jpp-search-result-header">
                        <strong>{this.props.car.title}</strong>
                    </div>
                    <hr />
                    <div className="jpp-search-result-overview">
                        <div className="jpp-search-result-overview-image"><img className="mr-3 jpp-image" src={c63amg} alt="car" /></div>
                        <div className="jpp-search-result-overview-content">
                            <div>Model: {this.props.car.baseModel.model}</div>
                            <div>Hersteller: {this.props.car.baseModel.manufacturer}</div>
                            <div>Baujahr: {this.props.car.baseModel.manufactureDate}</div>
                        </div>
                    </div>
                    <div className="jpp-search-result-stages">
                        {stages}
                    </div>
                </div>
            </div>
        );
    }
}
export default CarOverview;