import React, {Component} from "react";
import PropTypes from 'prop-types';
import AdenauerKreuz from "./AdenauerKreuz";

class KreuzItem extends Component{


    getStyle = () =>{

        return{
            textAlign: 'center',
            padding: '10px',
            backgroundColor: this.props.kreuz.proOrCon ?
                '#07da63' : '#ff2400'
        }
    }

    render() {
        return(
            <tr>
                <td style={this.getStyle()}>
                    {'\t\t\t'}
                </td>
                <td>
                    <div >
                        {this.props.kreuz.cause}
                    </div>
                </td>
                <td>
                    <input type="checkbox" checked={this.props.kreuz.proOrCon}/>
                </td>
                <td>
                    <input type="checkbox" checked={!(this.props.kreuz.proOrCon)}/>
                </td>
                <td>
                    {this.props.kreuz.points}
                </td>

            </tr>

        )
    }
}

KreuzItem.propTypes = {
    kreuz: PropTypes.object.isRequired
}



export default KreuzItem;

