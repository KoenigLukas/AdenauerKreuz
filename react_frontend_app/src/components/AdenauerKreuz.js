import React, {Component} from "react";
import KreuzItem from "./KreuzItem";
import PropTypes from 'prop-types';

class AdenauerKreuz extends Component{
    render() {
        return this.props.kreuz.map((kreuz)=>(
            <KreuzItem key={kreuz.id} kreuz ={kreuz}/>
        ));
    }
}

AdenauerKreuz.propTypes = {
    aKreuz: PropTypes.array.isRequired
}

export default AdenauerKreuz;

