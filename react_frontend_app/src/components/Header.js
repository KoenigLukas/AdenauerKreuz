import React, {Component} from "react";
import PropTypes from 'prop-types';
import KreuzItem from "./KreuzItem";


class Header extends Component{
    render() {
        return(
            <header style={headerStyle}>
                <h1>{this.props.topicName}</h1>
            </header>
        )
    }



}

const headerStyle = {
    background: '#333',
    color: '#fff',
    padding: '10px'
}

export default Header;

