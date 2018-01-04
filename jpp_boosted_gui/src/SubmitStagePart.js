import React from 'react';

class SubmitStagePart extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            index: this.props.index,
            part: this.props.part,
        };
    }

    handleUserInput = (e) => {
        const part = this.state.part;
        part[e.target.name] = e.target.value;

        this.setState({ part: part });
        this.props.onChange(this.props.index, this.state.part);
    }
    
    render() {
        return (
            <div>
                <div className="form-group">
                    <div>
                        <label className="keyName" htmlFor="name">Teilename</label>
                        <input type="text" name="name" placeholder="Krümmer"
                            value={this.state.part.name}
                            onChange={this.handleUserInput} />
                    </div>
                    <div>
                        <label className="keyName" htmlFor="url">URL</label>
                        <input type="text" name="url" placeholder=""
                            value={this.state.part.url}
                            onChange={this.handleUserInput} />
                    </div>
                </div>
                <div className="form-group">
                    <div>
                        <label className="keyName" htmlFor="manufacturer">Teilehersteller</label>
                        <input type="text" name="manufacturer" placeholder=""
                            value={this.state.part.manufacturer}
                            onChange={this.handleUserInput} />
                    </div>
                    <div>
                        <label className="keyName" htmlFor="manufacturerUrl">URL</label>
                        <input type="text" name="manufacturerUrl" placeholder="" 
                            value={this.state.part.manufacturerUrl}
                            onChange={this.handleUserInput} />
                    </div>
                </div>
            </div>
        );
    }
}
export default SubmitStagePart;