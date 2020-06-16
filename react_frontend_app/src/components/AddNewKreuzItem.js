import React, {Component} from "react";

class AddNewKreuzItem extends Component{

    state = {
        entry:{
            cause: '',
            pro: false,
            con: false,
            points: 0
        }

    }

    onChange = (e) => this.setState({ entry:{
            [e.target.name]: e.target.value,
        }


    });
    onSubmit = (e) =>{

        e.preventDefault();
        this.props.addCause(this.state.entry)
        this.setState({cause: '', pro: false, con: false, points: 0})}


    render() {
        return(
            <tr>

                <td>
                    <form onSubmit={this.onSubmit} style={{display: 'flex'}}>
                    <input type="submit" value="Submit" className="btn" style={{flex: '1'}}/>
                    </form>
                </td>
                <td style={{display: 'flex'}}>
                    <div >
                        <input type="text" name="cause" value={this.state.entry.cause} style={{flex: '10',padding: '5px'}} onChange={this.onChange}/>
                    </div>
                </td>
                <td>
                    <input type="checkbox" id="pro" name="pro" value={!this.props.proButtonState.buttonState} checked={this.props.proButtonState.buttonState} onClick={this.props.checkBoxChecked.bind(this, this.props.proButtonState.id)} onChange={this.onChange} />
                </td>
                <td>
                    <input type="checkbox" id="contra" name="con" value={!this.props.conButtonState.buttonState} checked={this.props.conButtonState.buttonState} onClick={this.props.checkBoxChecked.bind(this, this.props.conButtonState.id)} onChange={this.onChange}/>
                </td>
                <td>
                    <input type="number" id="points" name="points" step={1} min={0} max={10}/>
                </td>


            </tr>

        )
    }



}



export default AddNewKreuzItem;
