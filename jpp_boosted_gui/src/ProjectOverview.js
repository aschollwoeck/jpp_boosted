import React from 'react';
import './ProjectOverview.css';

import c63amg from './c63_amg.jpg';

class ProjectOverview extends React.Component {

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

    getCarImage() {
        var imgUrl = this.props.project.baseModel.imageUrl;
        if (imgUrl !== "") {
            imgUrl = getExternalLink(imgUrl);
        } else {
            imgUrl = c63amg;
        }

        return imgUrl;
    }

    render() {
        var stages = "";
        if (this.props.project.tunings != null) {
            stages = this.props.project.tunings.map(t => {
                const power = <div className="mt-1 row">
                    <div className="ml-3">PS: {t.horsePower}</div>
                    <div className="ml-2">NM: {t.torque}</div>
                    <div className="ml-2">Date: {t.date}</div>
                </div>
                const times = t.times.map(m => {
                    return (
                        <div key={m.id}>{m.speedRange}: {m.time} Sekunden</div>
                    );
                });
                const modParts = t.parts.map(p => {
                    return (
                        <div key={p.id}>{p.name} (<a href={getExternalLink(p.manufacturerUrl)}>{p.manufacturer}</a>)</div>
                    );
                });
                var ytEmbedded = "";
                if (t.youtubeUrl !== "" && t.youtubeUrl !== undefined) {
                    var match = new RegExp("v=([0-9a-zA-Z_]+)", "i").exec(t.youtubeUrl);
                    if (match !== null) {
                        const ytLink = "https://www.youtube.com/embed/" + match[1]
                        ytEmbedded = <iframe title={t.id + match[1]} width="100%" src={ytLink} frameBorder="0" gesture="media" allow="encrypted-media" allowFullScreen></iframe>
                    }
                }
                return (
                    <div key={t.id} className="jpp-search-result-stage">
                        <div className="jpp-search-result-stage-name mt-1">
                            <strong>{t.stage}</strong>
                        </div>
                        <div className="jpp-search-result-stage-content mt-1">
                            <div>{t.description}</div>
                            <div className="mt-1">
                                {power}
                            </div>
                            <div className="mt-1">
                                {times}
                            </div>
                            <div className="mt-1">
                                {modParts}
                            </div>
                        </div>
                        <div className=" mt-1">
                            {ytEmbedded}
                        </div>
                    </div>
                );
            });
        }
        const imgUrl = this.getCarImage();
        return (
            <div className="jpp-search-result col-12">
                <div className="">
                    <div className="jpp-search-result-header">
                        <strong>{this.props.project.title}</strong>
                    </div>
                    <hr />
                    <div className="jpp-search-result-overview">
                        <div className="jpp-search-result-overview-image">
                            <img className="mr-3 jpp-image" src={imgUrl} alt="car" />
                        </div>
                        <div className="jpp-search-result-overview-content">
                            <div>Model: {this.props.project.baseModel.name}</div>
                            <div>Hersteller: {this.props.project.baseModel.manufacturer.name}</div>
                            <div>Baujahr: {this.props.project.baseModel.buildStart}</div>
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
export default ProjectOverview;

function getExternalLink(link) {
    if (link.slice(0, 4) !== "http") {
        link = "http://" + link;
    }
    return link;
}