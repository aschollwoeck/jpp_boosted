import React from 'react';
import moment from 'moment';

import './ProjectOverview.css';
import c63amg from './c63_amg.jpg';
import { Link } from 'react-router-dom';

class ProjectOverview extends React.Component {
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
                stages = this.props.project.tunings.sort(t => t.date).map(t => {
                const ps = t.horsePower !== 0 ? <div className="ml-3">PS: {t.horsePower}</div> : ""
                const nm = t.torque !== 0 ? <div className="ml-2">NM: {t.torque}</div> : ""
                const power = <div className="mt-1">
                    <div className="row">
                        {ps}
                        {nm}
                    </div>
                    <div>Date: {moment(t.date).format("MMM. YY")}</div>
                </div>
                var times = "";
                if (t.times !== null && t.times.length > 0) {
                    times = t.times.map(m => {
                        return (
                            <div key={m.id}><span className="speedRangeKey">{m.speedRange}</span>: {m.time} Sekunden</div>
                        );
                    });
                    
                    times = <div><div><strong>Gemessene Zeiten (km/h)</strong></div>{times}</div>
                }
                var modParts = "";
                if (t.parts !== null && t.parts.length > 0) {
                    modParts = t.parts.map(p => {
                        const partName = p.url === "" ? p.name : <a target="blank" href={getExternalLink(p.url)}>{p.name}</a>
                        const manName = p.manufacturerUrl === "" ? p.manufacturer : <a target="blank" href={getExternalLink(p.manufacturerUrl)}>{p.manufacturer}</a>
                        return (
                            <div key={p.id}>{partName} ({manName})</div>
                        );
                    });
                    modParts = <div><div><strong>Verbaute Teile</strong></div>{modParts}</div>
                }
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
                            <div className="mt-2">
                                {times}
                            </div>
                            <div className="mt-2">
                                {modParts}
                            </div>
                        </div>
                        <div className="mt-1">
                            {ytEmbedded}
                        </div>
                    </div>
                );
            });
        }

        var modTimespan = "";
        if (this.props.project.tunings !== null && this.props.project.tunings.length > 0) {
            var dates = this.props.project.tunings.map(t => new Date(t.date));
            var maxDate = new Date(Math.max.apply(null, dates));
            var minDate = new Date(Math.min.apply(null, dates));
            if (moment(maxDate).format("MMM. YY") === moment(minDate).format("MMM. YY")) {
                modTimespan = moment(minDate).format("MMM. YY")
            } else {
                modTimespan = moment(minDate).format("MMM. YY") + " - " + moment(maxDate).format("MMM. YY")
            }
        }

        const imgUrl = this.getCarImage();

        var editButton = "";
        if (this.props.project !== null) {
            editButton = (
                <Link  className="jpp-link-edit btn btn-primary" to={"/submit/" + this.props.project.id}>Bearbeiten</Link>
            );
        }

        return (
            <div className="jpp-search-result col-12">
                <div className="">
                    <div className="jpp-search-result-header">
                        <strong>{this.props.project.title}</strong>
                        {editButton}
                    </div>
                    <hr />
                    <div className="jpp-search-result-overview">
                        <div className="jpp-search-result-overview-image">
                            <img className="mr-3 jpp-image" src={imgUrl} alt="car" />
                        </div>
                        <div className="jpp-search-result-overview-content">
                            <div>Model: {this.props.project.baseModel.name} ({this.props.project.baseModel.seriesCode})</div>
                            <div>Hersteller: {this.props.project.baseModel.manufacturer.name}</div>
                            <div>Baujahr: {this.props.project.buildYear}</div>
                            <div>Bauzeit: {modTimespan}</div>
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