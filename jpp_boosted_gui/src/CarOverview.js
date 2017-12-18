import React from 'react';
import './CarOverview.css';

import c63amg from './c63_amg.jpg';

class CarOverview extends React.Component {

    // {
    //     id: 1,
    //     title: "C63 AMG Tuning",
    //     buildYear: 2016,
    //     manufacturer: "Mercedes-Benz",
    //     model: "C63 AMG",
    //     tuning: [
    //       {
    //         id: 1,
    //         stage: "1",
    //         description: "Software",
    //         measuredTime: [
    //           {
    //             id: 1,
    //             speedRange: "0-200",
    //             time: "8.7",
    //             youtubeUrl: ""
    //           },
    //           {
    //             id: 2,
    //             speedRange: "0-100",
    //             time: "3.4",
    //             youtubeUrl: ""
    //           }
    //         ]
    //       }
    //     ]
    //   }

    render() {
        const stages = this.props.car.tuning.map(t => {
            const times = t.measuredTime.map(m => {
                return (
                    <div key={m.id}>{m.speedRange}: {m.time} Sekunden</div>
                );
            });
            return (
                <div key={t.id} className="jpp-search-result-stage row">
                    <div className="col-3">
                        <strong>{t.stage}</strong>
                    </div>
                    <div className="col-5">
                        <div>{t.description}</div>
                        {times}  
                    </div>
                    <div className="col-4">
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
                    <div className="jpp-search-result-manufacturer media">
                        <img className="mr-3 carOverviewImg" src={c63amg} alt="car" />
                        <div className="media-body">
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