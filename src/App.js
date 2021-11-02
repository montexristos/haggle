import React from 'react';
import './App.css';
import 'bulma';
import StakeCalculator from './StakeCalculator';
import Arb from './Arb';
import ArbSeries from './ArbSeries';
import EventList from "./EventList";
import SiteEvents from "./SiteEvents";
import Bet from "./Bet";

const host = "http://localhost:8088";

class App extends React.Component {

    constructor(props) {
        super(props);
        this.state = {
            events: [],
            sites: {}
        }
        // fetch(this.getHost() + '/data')
        //     .then(response => response.json())
        //     .then(data => this.setState({
        //         events: data.events,
        //         sites: data.sites
        //     }));
        fetch(this.getHost() + '/all')
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

    scrape = () => {
        fetch(this.getHost() + '/scrape').then(() => {})
    };

    render = () => {
        return <div className="App">
            <header className="App-header">
                <StakeCalculator/>
                <Arb/>
            </header>
            <ArbSeries/>

            <button onClick={this.updateFilters}>Filter</button>
            <button onClick={this.scrape}>Update</button>

            <SiteEvents events={this.state.events}  sites={this.state.sites} />

            {/*<Bet events={this.state.events} sites={this.state.sites}/>*/}
            {/*<EventList events={this.state.events} sites={this.state.sites}/>*/}
        </div>
    };

}

export default App;
