import React from 'react';
import './App.css';
import 'bulma';
import StakeCalculator from './StakeCalculator';
import Arb from './Arb';
import ParlayArb from './ParlayArb';
import ArbSeries from './ArbSeries';
import ArbSeriesNew from './ArbSeriesNew';
import SequentialArb from "./SequentialArb";
import EventList from "./EventList";

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

        return <div className="App">
            <header className="App-header">
                <StakeCalculator/>
            </header>
            <Arb/>
            <SequentialArb/>
            <ArbSeries/>
            <ArbSeriesNew/>
            <ParlayArb/>

            <button onClick={this.updateFilters}>Filter</button>
            <button onClick={this.update}>Update</button>
            <EventList events={this.state.events} sites={this.state.sites}/>
        </div>
    };

}

export default App;
