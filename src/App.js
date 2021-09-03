import React from 'react';
import './App.css';
import 'bulma';
import Competitions from './Competitions';
import StakeCalculator from './StakeCalculator';
import Arb from './Arb';
import ParlayArb from './ParlayArb';
import ArbSeries from './ArbSeries';
import ArbSeriesNew from './ArbSeriesNew';
import SequentialArb from "./SequentialArb";
import Event from "./Event";

const host = "http://localhost:8088";

class App extends React.Component {

    constructor(props) {
        super(props);
        this.state = {
            events: []
        }
        fetch(this.getHost() + '/data')
            .then(response => response.json())
            .then(data => this.setState({
                events: data.events
            }));
    }

    getHost = () => {
        if (!process.env.NODE_ENV || process.env.NODE_ENV === 'development') {
            return host;
        } else {
            return ""
        }
    }

    updateState = (events) => {
        this.setState({
            events: events
        });
    };

    update = () => {
        fetch(this.getHost() + '/scrape').then(() => {})
    };

    render = () => {
        if (this.state.events === undefined || !this.state.events) {
            return null;
        }
        let rows = [];
        for (const eventId in this.state.events) {
            const event = this.state.events[eventId];
            console.log(event);
            if (!event) {
                continue;
            }
            console.log(event[0].Markets[0]);
            const matchResult1 = event[0].Markets[0].MarketType ? event[0].Markets[0] : "";
            const matchResult2 = event[1].Markets[0].MarketType ? event[1].Markets[0] : "";
            const overUnder1 = event[0].Markets[1].MarketType ? event[0].Markets[1] : "";
            const overUnder2 = event[1].Markets[1].MarketType ? event[1].Markets[1] : "";
            const btts1 = event[0].Markets[2].MarketType ? event[0].Markets[2] : "";
            const btts2 = event[1].Markets[2].MarketType ? event[1].Markets[2] : "";

            rows.push(<tr>
                <td rowSpan="2">
                    {event[0].Name}
                </td>
                <td>{matchResult1 ? matchResult1.Selections[0].Price : ""}</td>
                <td>{matchResult1 ? matchResult1.Selections[1].Price : ""}</td>
                <td>{matchResult1 ? matchResult1.Selections[2].Price : ""}</td>
                <td>{overUnder1 ? overUnder1.Selections[0].Price : ""}</td>
                <td>{overUnder1 ? overUnder1.Selections[1].Price : ""}</td>
                <td>{btts1 ? btts1.Selections[0].Price : ""}</td>
                <td>{btts1 ? btts1.Selections[1].Price : ""}</td>
            </tr>,
            <tr>
                <td>{matchResult2 ? matchResult2.Selections[0].Price : ""}</td>
                <td>{matchResult2 ? matchResult2.Selections[1].Price : ""}</td>
                <td>{matchResult2 ? matchResult2.Selections[2].Price : ""}</td>
                <td>{overUnder2 ? overUnder2.Selections[0].Price : ""}</td>
                <td>{overUnder2 ? overUnder2.Selections[1].Price : ""}</td>
                <td>{btts2 ? btts2.Selections[0].Price : ""}</td>
                <td>{btts2 ? btts2.Selections[1].Price : ""}</td>
            </tr>
            );
        }
        return <div className="App">
            <header className="App-header">
                <StakeCalculator />
            </header>
            <Arb />
            <SequentialArb />
            <ArbSeries />
            <ArbSeriesNew />
            <ParlayArb />

            <button onClick={this.updateFilters}>Filter</button>
            <button onClick={this.update}>Update</button>
            <table  className="table">
                <thead>
                    <th>Event name</th>
                    <th colspan="3">match result</th>
                    <th colspan="2">under/over</th>
                    <th colspan="2">BTTS</th>
                </thead>
                <tbody>
                { rows }
                </tbody>
            </table>
        </div>
    };

}

export default App;
