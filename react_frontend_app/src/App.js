import React, {Component} from 'react';
import './App.css';
import AdenauerKreuz from "./components/AdenauerKreuz";
import AddNewKreuzItem from "./components/AddNewKreuzItem";
import Header from "./components/Header";

class App extends Component{

    state = {
        aKreuz: [
            {
                id: 1,
                cause: 'Langweilig',
                points: 10,
                proOrCon: false
             },
            {
                id: 2,
                cause: 'Wichtig',
                points: 90,
                proOrCon: true
            },
            {
                id: 3,
                cause: 'Schlechtes Wetter',
                points: 30,
                proOrCon: false
            },
            {
                id: 4,
                cause: 'na Echt ned',
                points: 99,
                proOrCon: false
            }
        ],
        proButtonState: {id: 1, buttonState: false},
        conButtonState: {id: 2, buttonState: false}

    }




    checkBoxChecked = (id) => {


        if (id===1)
        {
            let tmp = this.state.proButtonState.buttonState;
            tmp=!tmp;
            this.setState({proButtonState:{id: 1,buttonState: tmp}});
            this.setState({conButtonState:{id: 2,buttonState: false}});

        }
        if (id===2)
        {
            let tmp = this.state.conButtonState.buttonState;
            tmp=!tmp;
            this.setState({proButtonState:{id: 1,buttonState: false}});
            this.setState({conButtonState:{id: 2,buttonState: tmp}});
        }

    }


    addCause = (entry) =>{
        console.log(entry)
    }


    render() {
        return (
            <div className="App">
                <div className="container">
                    <Header topicName="Heinz"/>

                    <table className="kreuzTable">
                        <tbody>
                        <tr>
                            <th>    </th>
                            <th>Argument</th>
                            <th>Pro</th>
                            <th>Contra</th>
                            <th>Wert</th>
                        </tr>
                        <AddNewKreuzItem addCause={this.addCause} proButtonState={this.state.proButtonState} conButtonState={this.state.conButtonState} checkBoxChecked={this.checkBoxChecked}/>
                        <AdenauerKreuz kreuz={this.state.aKreuz}/>
                        </tbody>

                    </table>
                </div>


            </div>
        );
    }
}



export default App;
