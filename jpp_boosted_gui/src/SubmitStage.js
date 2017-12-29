import React from 'react';
import SubmitStageTime from './SubmitStageTime';
import SubmitStagePart from './SubmitStagePart';

class SubmitStage extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            tuning: this.props.tuning,
        };
    }

    handleUserInput = (e) => {
        const tuning = this.state.tuning;
        tuning[e.target.name] = e.target.value;

        this.setState({ tuning: tuning });
        this.props.onChange(this.props.index, this.state.tuning);
    }

    //             speedRange: "0-200",
    //             time: "10.3",
    //             youtubeUrl: ""
    getNewTime() {
        return { speedRange: "", time: "", youtubeUrl: "" };
    }

    addTime = (e) => {
        e.preventDefault();
        const tuning = this.state.tuning;
        tuning.measuredTime.push(this.getNewTime());

        this.setState({ tuning: tuning });
    }

    // {
    //     id: 1,
    //     part: "Krümmer",
    //     youtubeUrl: "",
    //     manufacturer: "Hersteller",
    //     manufacturerUrl: "",
    //     partUrl: ""
    //   }
    getNewPart() {
        return { part: "", youtubeUrl: "", manufacturer: "", manufacturerUrl: "", partUrl: "" };
    }

    addPart = (e) => {
        e.preventDefault();
        const tuning = this.state.tuning;
        tuning.modifiedParts.push(this.getNewPart());

        this.setState({ tuning: tuning });
    }

    timeChanged = (index, value) => {
        const tuning = this.state.tuning;
        tuning.measuredTime[index] = value;
        this.setState({tuning: tuning});
        this.props.onChange(this.props.index, this.state.tuning);
    }

    partChanged = (index, value) => {
        const tuning = this.state.tuning;
        tuning.modifiedParts[index] = value;
        this.setState({tuning: tuning});
        this.props.onChange(this.props.index, this.state.tuning);
    }

    render() {
        const times = this.state.tuning.measuredTime.map((m, i) => {
            return <SubmitStageTime key={"time" + i} index={i} measuredTime={m} onChange={this.timeChanged} />
        });
        const parts = this.state.tuning.modifiedParts.map((p, i) => {
            return <SubmitStagePart key={"part" + i} index={i} modifiedPart={p} onChange={this.partChanged} />
        });

        return (
            <div>
                <div className="form-group">
                    <div>
                        <label className="keyName" htmlFor="name">Name</label>
                        <input type="text" name="name" placeholder="Serie, Stage 1..."
                            value={this.state.tuning.name}
                            onChange={this.handleUserInput} />
                    </div>
                </div>
                <div className="form-group">
                    <div>
                        <label className="keyName" htmlFor="description">Beschreibung</label>
                        <textarea name="description" placeholder="Software, Abgasanlage"
                            value={this.state.tuning.description}
                            onChange={this.handleUserInput} />
                    </div>
                </div>
                <div className="form-group">
                    <div>
                        <label className="keyName" htmlFor="youtubeUrl">YouTube</label>
                        <textarea name="youtubeUrl" placeholder="wwww.youtube.com/..."
                            value={this.state.tuning.youtubeUrl}
                            onChange={this.handleUserInput} />
                    </div>
                </div>
                <h5>Messungen</h5>
                <div className="form-group">
                    <div>
                        <label className="keyName" htmlFor="ps">PS</label>
                        <input type="text" name="ps" placeholder="414"
                            value={this.state.tuning.ps}
                            onChange={this.handleUserInput} />
                    </div>
                    <div>
                        <label className="keyName" htmlFor="nm">NM</label>
                        <input type="text" name="nm" placeholder="510"
                            value={this.state.tuning.nm}
                            onChange={this.handleUserInput} />
                    </div>
                </div>
                <div className="form-group">
                    <div>
                        <label className="keyName" htmlFor="date">Datum</label>
                        <input type="date" name="date"
                            value={this.state.tuning.date}
                            onChange={this.handleUserInput} />
                    </div>
                </div>
                <button className="btn btn-primary" onClick={this.addTime}>
                    + Zeit hinzufügen
                </button>
                {times}
                <h5>Teile</h5>
                <button className="btn btn-primary" onClick={this.addPart}>
                    + Teil hinzufügen
                </button>
                {parts}
            </div>
        );
    }
}
export default SubmitStage;