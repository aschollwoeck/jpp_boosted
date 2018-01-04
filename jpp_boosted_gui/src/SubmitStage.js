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

    getNewTime() {
        return { id: 0, speedRange: "", time: "", youtubeUrl: "" };
    }

    addTime = (e) => {
        e.preventDefault();
        const tuning = this.state.tuning;
        if(tuning.times === null) {
            tuning.times = [];
        }
        tuning.times.push(this.getNewTime());

        this.setState({ tuning: tuning });
    }

    getNewPart() {
        return { id: 0, name: "", youtubeUrl: "", manufacturer: "", manufacturerUrl: "", url: "" };
    }

    addPart = (e) => {
        e.preventDefault();
        const tuning = this.state.tuning;
        if(tuning.parts === null) {
            tuning.parts = [];
        }
        tuning.parts.push(this.getNewPart());

        this.setState({ tuning: tuning });
    }

    timeChanged = (index, value) => {
        const tuning = this.state.tuning;
        tuning.times[index] = value;
        this.setState({ tuning: tuning });
        this.props.onChange(this.props.index, this.state.tuning);
    }

    partChanged = (index, value) => {
        const tuning = this.state.tuning;
        tuning.parts[index] = value;
        this.setState({ tuning: tuning });
        this.props.onChange(this.props.index, this.state.tuning);
    }

    render() {
        var times = "";
        if (this.state.tuning.times !== null) {
            times = this.state.tuning.times.map((m, i) => {
                return <SubmitStageTime key={"time" + i} index={i} time={m} onChange={this.timeChanged} />
            });
        }

        var parts = "";
        if (this.state.tuning.parts !== null) {
            parts = this.state.tuning.parts.map((p, i) => {
                return <SubmitStagePart key={"part" + i} index={i} part={p} onChange={this.partChanged} />
            });
        }

        return (
            <div>
                <div className="form-group">
                    <div>
                        <label className="keyName" htmlFor="stage">Name</label>
                        <input type="text" name="stage" placeholder="Serie, Stage 1..."
                            value={this.state.tuning.stage}
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
                        <label className="keyName" htmlFor="horsePower">PS</label>
                        <input type="text" name="horsePower" placeholder="414"
                            value={this.state.tuning.horsePower}
                            onChange={this.handleUserInput} />
                    </div>
                    <div>
                        <label className="keyName" htmlFor="torque">NM</label>
                        <input type="text" name="torque" placeholder="510"
                            value={this.state.tuning.torque}
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