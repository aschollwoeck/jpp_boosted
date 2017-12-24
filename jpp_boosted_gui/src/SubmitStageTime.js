import React from 'react';

class SubmitStageTime extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            index: this.props.index,
            measuredTime: this.props.measuredTime,
        }
    }

    handleUserInput = (e) => {
        const measuredTime = this.state.measuredTime;
        measuredTime[e.target.name] = e.target.value;

        this.setState({ measuredTime: measuredTime });
        this.props.onChange(this.props.index, this.state.measuredTime);
    }

    render() {
        return (
            <div>
                <div className="form-group">
                    <div>
                        <label className="keyName" htmlFor="speedRange">KM/H</label>
                        <input type="text" name="speedRange" placeholder="0-100" 
                            value={this.state.measuredTime.speedRange}
                            onChange={this.handleUserInput} />
                    </div>
                    <div>
                        <label className="keyName" htmlFor="time">Zeit (Sek.)</label>
                        <input type="text" name="time" placeholder="3.8" 
                            value={this.state.measuredTime.time}
                            onChange={this.handleUserInput} />
                    </div>
                </div>
                <div className="form-group">
                    <div>
                        <label className="keyName" htmlFor="youtubeUrl">YouTube</label>
                        <input type="text" name="youtubeUrl" placeholder="youtube.com/watch?v=v1kFDOnueug" 
                            value={this.state.measuredTime.youtubeUrl}
                            onChange={this.handleUserInput} />
                    </div>
                </div>
            </div>
        );
    }
}
export default SubmitStageTime;