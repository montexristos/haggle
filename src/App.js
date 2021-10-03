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
            events: [],
            sites: {}
        }
        fetch(this.getHost() + '/data')
            .then(response => response.json())
            .then(data => this.setState({
                events: data.events,
                sites: data.sites
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
        const sites = this.state.sites;
        let rows = [];
        for (const eventId in this.state.events) {
            const events = this.state.events[eventId];
            if (!events) {
                continue;
            }
            if (events.length) {
                const matchResult = {};
                const overUnder = {};
                const btts = {};
                for (let siteId=0; siteId<events.length;siteId++) {
                    matchResult[siteId] = events[siteId].Markets[0].MarketType ? events[siteId].Markets[0] : "";
                    overUnder[siteId] = events[siteId].Markets[1].MarketType ? events[siteId].Markets[1] : "";
                    btts[siteId] = events[siteId].Markets[2].MarketType ? events[siteId].Markets[2] : "";
                    const site = sites[events[siteId].SiteID];
                    let nameTag = "";
                    if (siteId === 0) {
                        nameTag = <td rowSpan={events.length}>
                            {events[0].CanonicalName} <br/>
                            {events[0].Date}
                        </td>;
                    }
                    rows.push(<tr>
                        {nameTag}
                        <td>{site}</td>
                        <td>{matchResult[siteId] && matchResult[siteId].Selections ? matchResult[siteId].Selections[0].Price : ""}</td>
                        <td>{matchResult[siteId] && matchResult[siteId].Selections ? matchResult[siteId].Selections[1].Price : ""}</td>
                        <td>{matchResult[siteId] && matchResult[siteId].Selections ? matchResult[siteId].Selections[2].Price : ""}</td>
                        <td>{overUnder[siteId] && overUnder[siteId].Selections ? overUnder[siteId].Selections[0].Price : ""}</td>
                        <td>{overUnder[siteId] && overUnder[siteId].Selections ? overUnder[siteId].Selections[1].Price : ""}</td>
                        <td>{btts[siteId] && btts[siteId].Selections ? btts[siteId].Selections[0].Price : ""}</td>
                        <td>{btts[siteId] && btts[siteId].Selections ? btts[siteId].Selections[1].Price : ""}</td>
                    </tr>);
                }
            }
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
