import React from 'react';

class SubmitStagePart extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            index: this.props.index,
            modifiedPart: this.props.modifiedPart,
        };
    }

    handleUserInput = (e) => {
        const modifiedPart = this.state.modifiedPart;
        modifiedPart[e.target.name] = e.target.value;

        this.setState({ modifiedPart: modifiedPart });
        this.props.onChange(this.props.index, this.state.modifiedPart);
    }
    //     part: "Krümmer",
    //     youtubeUrl: "",
    //     manufacturer: "Hersteller",
    //     manufacturerUrl: "",
    //     partUrl: ""
    render() {
        return (
            <div>
                <div className="form-group">
                    <div>
                        <label className="keyName" htmlFor="part">Teilename</label>
                        <input type="text" name="part" placeholder="Krümmer"
                            value={this.state.modifiedPart.part}
                            onChange={this.handleUserInput} />
                    </div>
                    <div>
                        <label className="keyName" htmlFor="partUrl">URL</label>
                        <input type="text" name="partUrl" placeholder=""
                            value={this.state.modifiedPart.partUrl}
                            onChange={this.handleUserInput} />
                    </div>
                </div>
                <div className="form-group">
                    <div>
                        <label className="keyName" htmlFor="manufacturer">Teilehersteller</label>
                        <input type="text" name="manufacturer" placeholder=""
                            value={this.state.modifiedPart.manufacturer}
                            onChange={this.handleUserInput} />
                    </div>
                    <div>
                        <label className="keyName" htmlFor="manufacturerUrl">URL</label>
                        <input type="text" name="manufacturerUrl" placeholder="" 
                            value={this.state.modifiedPart.manufacturerUrl}
                            onChange={this.handleUserInput} />
                    </div>
                </div>
            </div>
        );
    }
}
export default SubmitStagePart;