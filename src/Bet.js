import React, {Fragment} from 'react';
import './App.css';
import 'bulma';
import Competitions from './Competitions';

const host = "http://localhost:8088";

class Bet extends React.Component {

    constructor(props) {
        super(props);
        this.state = {
            competitions: [],
            tournaments: [],
            weeks: [],
            start: "",
            end: "",
            selectedDate: "all",
            correct: 0,
            over: 0,
            wrong: 0,
            threshold: 60,
            threshold1: 65,
            threshold2: 70,
            threshold3: 70,
            threshold4: 80,
            sort: "tournament",
            hideCompleted: false,
            showO0: false,
            showO1: false,
            showO2: false,
            showO3: false,
            showO4: false,
            showU0: false,
            showU1: false,
            showU2: false,
            showU3: false,
            showU4: false,
            homeScoredWeight: 10,
            homeAwayScoredWeight: 8,
            homeConceidedWeight: 10,
            homeAwayConceidedWeight: 3,
            awayScoredWeight: 10,
            awayHomeScoredWeight: 0,
            awayConceidedWeight: 10,
            awayHomeConceidedWeight: 0,
            homeOverWeight: 0,
            awayOverWeight: 0,
            homeOpWeight: 0,
            awayOpWeight: 0,
            cardThreshold: 5,
            cornerThreshold: 13,
            ggs: [],
            overs: []
        }
        const queryParams = "?threshold0=" + this.state.threshold +
            "&threshold1=" + this.state.threshold1 +
            "&threshold2=" + this.state.threshold2 +
            "&threshold3=" + this.state.threshold3 +
            "&threshold4=" + this.state.threshold4;
        fetch(this.getHost() + '/home' + queryParams)
            .then(response => response.json())
            .then(data => this.updateState(
                data.tournaments,
                data.weeks,
                data.start,
                data.end,
                "all",
                data.stats
            ));

    }

    getHost = () => {
        if (!process.env.NODE_ENV || process.env.NODE_ENV === 'development') {
            return host;
        } else {
            return ""
        }
    }

    changeWeek = (event) => {
        if (event.target.value === 'all') {
            fetch(this.getHost() + '/all')
                .then(response => response.json())
                .then(data => this.updateState(
                    data.tournaments,
                    data.weeks,
                    data.start,
                    data.end,
                    "all",
                    data.stats
                ));
        }
        else {
            const queryParams = "?threshold0=" + this.state.threshold +
                "&threshold1=" + this.state.threshold1 +
                "&threshold2=" + this.state.threshold2 +
                "&threshold3=" + this.state.threshold3 +
                "&threshold4=" + this.state.threshold4;
            fetch(this.getHost() + '/week/' + event.target.value + queryParams)
                .then(response => response.json())
                .then(data => this.updateState(
                    data.tournaments,
                    data.weeks,
                    data.start,
                    data.end,
                    "all",
                    data.stats
                ));
        }
    };

    updateState = (competitions, weeks, start, end, selectedDate, stats) => {
        this.setState({
            tournaments: competitions,
            weeks: weeks,
            start: start,
            end: end,
            selectedDate: selectedDate,
            stats: stats,
            threshold: stats.threshold0,
            threshold1: stats.threshold1,
            threshold2: stats.threshold2,
            threshold3: stats.threshold3,
            threshold4: stats.threshold4,
        });
    };

    filterDay = (event) => {
        const dt = new Date(event.target.value);
        this.setState({
            selectedDate: dt.getDay()
        })
    };
    sort = (event) => {
        this.setState({
            sort: event.target.value
        })
    };

    update = () => {
        fetch(this.getHost() + '/cache').then(() => {})
    };

    ml = () => {
        fetch(this.getHost() + '/ml').then((resp) => {
            resp.json().then((json) => {
                this.setState({
                    ggs: json.ggs,
                    overs: json.overs
                })
            });

        })
    };

    updateFilters = () => {
        fetch(this.getHost() + '/week/' + this.state.end)
            .then(response => response.json())
            .then(data => this.updateState(
                data.tournaments,
                data.weeks,
                data.start,
                data.end,
                "all",
                data.stats
            ));
    };

    render = () => {
        if (this.state.tournaments === undefined) {
            return null;
        }
        return <Fragment>
                <table>
                    <tbody>

                    </tbody>
                </table>
                <button onClick={this.updateFilters}>Filter</button>
                <button onClick={this.update}>Update</button>
                <button onClick={this.ml}>Calculate</button>
                <select onChange={this.sort} defaultValue="tournament">
                    <option value="tournament">Tournament</option>
                    <option value="date">Date</option>
                </select>
                <Competitions
                    changeWeek={this.changeWeek}
                    filterDay={this.filterDay}
                    weeks={this.state.weeks}
                    start={this.state.start}
                    end={this.state.end}
                    date={this.state.selectedDate}
                    tournaments={this.state.tournaments}
                    sort={this.state.sort}
                    ggs={this.state.ggs}
                    overs={this.state.overs}
                    events={this.props.events}
                    sites={this.props.sites}
                />
        </Fragment>
    };

}

export default Bet;
