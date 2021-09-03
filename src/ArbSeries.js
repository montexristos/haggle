import React from 'react';
import ArbSeriesEvent from './ArbSeriesEvent';

class ArbSeries extends React.Component {

    constructor(props) {
        super(props);
        const storedArbs = localStorage.getItem("arbseriesconcurrent");
        let arbs = JSON.parse(storedArbs);
        if (!Array.isArray(arbs)) {
            arbs = [{
                events: [],
                bonus: 0,
                toWin: 10,
                stake: 10
            }];
        }
        this.state = {
            arbs: arbs || []
        };
    }

    calculatePerc = (sel1, sel2) => {
        const perc1 = (1 / sel1) * 100;
        const perc2 = (1 / sel2) * 100;
        return Math.round(100 * (perc1 + perc2)) / 100;
    }

    overrideOddCb = (event, idx, arbIndex) => {
        event.persist();
        const arbs = this.state.arbs;
        const arb = arbs[arbIndex];
        let overrideOdd = parseFloat(event.target.value);
        let remainingMultiplier = this.getMultiplier(arb);
        arb.events.forEach((event, index) => {
            if (index !== arb.events.length - 1) {
                remainingMultiplier = remainingMultiplier / event.overrideOdd;
            }
        });
        arb.events[arb.events.length - 1].overrideOdd = Math.round(remainingMultiplier * 100) / 100;
        arb.events[idx].overrideOdd = overrideOdd;
        this.setState({
            arbs: arbs
        });
    }
    changeStake = (event, arbIndex) => {
        event.persist();
        const arbs = this.state.arbs;
        const arb = arbs[arbIndex];
        arb.stake = parseFloat(event.target.value);
        this.setState({
            arbs: arbs
        });
    }
    changeBonus = (event, arbIndex) => {
        event.persist();
        const arbs = this.state.arbs;
        const arb = arbs[arbIndex];
        arb.bonus = parseFloat(event.target.value);
        this.setState({
            arbs: arbs
        });
    }
    changeToWin = (event, arbIndex) => {
        event.persist();
        const arbs = this.state.arbs;
        const arb = arbs[arbIndex];
        arb.toWin = parseFloat(event.target.value);
        this.setState({
            arbs: arbs
        });
    }
    changeOdd = (event, idx, arbIndex) => {
        event.persist();
        const arbs = this.state.arbs;
        const arb = arbs[arbIndex];
        arb.events[idx].odd = parseFloat(event.target.value);
        this.setState({
            arbs: arbs
        });
    }
    changeStatus = (event, idx, arbIndex) => {
        event.persist();
        const arbs = this.state.arbs;
        const arb = arbs[arbIndex];
        console.log(event.target.value);
        arb.events[idx].status = event.target.value;
        this.setState({
            arbs: arbs
        });
    }
    overrideArbStake = (event, idx, arbIndex) => {
        event.persist();
        const arbs = this.state.arbs;
        const arb = arbs[arbIndex];
        arb.events[idx].overrideArbStake = parseFloat(event.target.value);
        this.setState({
            arbs: arbs
        });
    }
    setArb = (event, idx, arbIndex) => {
        event.persist();
        const arbs = this.state.arbs;
        const arb = arbs[arbIndex];
        arb.events[idx].arb = parseFloat(event.target.value);
        this.setState({
            arbs: arbs
        });
    }
    setName = (event, idx, arbIndex) => {
        const arbs = this.state.arbs;
        const arb = arbs[arbIndex];
        arb.events[idx].name = event.target.value;
        this.setState({
            arbs: arbs
        });
    }
    addArb = () => {
        let arbs = this.state.arbs;
        arbs.push({
            events: [],
            stake: 10,
            overrides: {},

        });
        this.setState({
            arbs: arbs
        });
    }
    addEvent = (arbIndex) => {
        const arbs = this.state.arbs;
        const arb = arbs[arbIndex];
        arb.events.push({
            name: "",
            odd: 1.9,
            arb: 1.9,
            arbStake: 10,
            status: "pending"
        });
        this.setState({
            arbs: arbs
        });
    }

    deleteEvent = (event, idx, arbIndex) => {
        const arbs = this.state.arbs;
        const arb = arbs[arbIndex];
        arb.events.splice(idx, 1);
        this.setState({
            arbs: arbs
        });
    }
    deleteArb = (arbIndex) => {
        const arbs = this.state.arbs;
        arbs.splice(arbIndex, 1);
        this.setState({
            arbs: arbs
        });
    }

    getMultiplier = (arb) => {
        let originalMultiplier = 1;
        if (arb.events.length) {
            arb.events.forEach((event, index) => {
                if (event.odd > 1) {
                    originalMultiplier = originalMultiplier * event.odd;
                }
            });
        }
        return originalMultiplier;
    }

    combine = (a, min) => {
        var fn = function(n, src, got, all) {
            if (n === 0) {
                if (got.length > 0) {
                    all[all.length] = got;
                }
                return;
            }
            for (var j = 0; j < src.length; j++) {
                fn(n - 1, src.slice(j + 1), got.concat([src[j]]), all);
            }
            return;
        }
        var all = [];
        for (var i = min; i < a.length; i++) {
            fn(i, a, [], all);
        }
        all.push(a);
        return all;
    }

    render = () => {
        const arbs = this.state.arbs.filter(function (el) {
            return el != null;
        });
        localStorage.setItem("arbseriesconcurrent", JSON.stringify(arbs));
        let events = null;
        const tables = arbs.map((arb, arbIndex) => {
            let headers;
            let multiplierArray = {};
            let seriesWinnings;
            let stakes;
            let winnings = 0;
            let totalStake = 0;
            const originalMultiplier = this.getMultiplier(arb);
            const parlayStake = arb.stake;
            // let parlayLost = false;
            if (arb.events && arb.events.length > 1) {
                const combinations = this.combine(arb.events, arb.events.length);
                console.log(combinations);
                let series = [];
                for (let comb in combinations) {
                    console.log(comb);
                    // series.push(comb.forEach((evt, idx) => {
                    //     const key = evt.name + idx;
                    //
                    //     return <td key={key}>a</td>;
                    // }));
                }

                events = arb.events.map((event, index) => {

                    return <ArbSeriesEvent name={event.name}
                                           odd={event.odd}
                                           overrideOdd={event.overrideOdd}
                                           overrideOddCb={this.overrideOddCb}
                                           changeOdd={this.changeOdd}
                                           setName={this.setName}
                                           setArb={this.setArb}
                                           overrideArbStake={this.overrideArbStake}
                                           arb={event.arb}
                                           arbStake={event.arbStake}
                                           key={arbIndex + "-" + index + "-" + event.name}
                                           index={index}
                                           changeStatus={this.changeStatus}
                                           arbIndex={arbIndex}
                                           status={event.status}
                                           series={series}
                                           deleteEvent={this.deleteEvent}
                    />;
                });

                const stakeArray = {};
                stakes = arb.events.map((r, j) => {
                    //for the first selections we want to break event
                    if (multiplierArray[j]) {
                        if (j === arb.events.length - 1) {
                            let toWin = parlayStake * originalMultiplier;
                            stakeArray[j] = Math.round(100 * (toWin) / r.arb ) / 100;
                        } else {
                            let toWin = parlayStake;
                            let mult = r.arb;
                            for (let o=0;o<j;o++) {
                                toWin += stakeArray[o];
                            }
                            for (let q=j+1;q<arb.events.length;q++){
                                mult = mult * arb.events[q].odd;
                            }
                            stakeArray[j] = Math.round(100 * (toWin) / (mult - 1)) / 100;
                        }
                    } else {
                        stakeArray[j] = 0;
                    }
                    totalStake += parseFloat(stakeArray[j]);
                    return <td key={"stake-" + j}>{stakeArray[j]}</td>;
                })
                headers = arb.events.map((e, i) => {
                    return <td key={"td-" + i}>Series {i}</td>;
                })
                seriesWinnings = arb.events.map((v, k) => {
                    return <td key={"win-" + k}>Winnings {Math.round(100 * stakeArray[k] * multiplierArray[k]) / 100}</td>;
                })
            }
            totalStake += parlayStake;
            totalStake = Math.round(totalStake * 100) / 100;
            winnings = Math.round(parlayStake * originalMultiplier * 100) / 100;
            return <table className="table" key={arbIndex}>
                <thead>
                <tr>
                    <td key={"td--1"}>Event</td>
                    <td key={"td--3"}>Odd</td>
                    <td key={"td--4"}>Arb</td>
                    {headers}
                    <td key={"td--5"} colSpan="2">Actions</td>
                </tr>
                </thead>
                <tbody>
                {events}
                <tr>
                    <td>Parlay Stake: <input type="number" defaultValue={arb.stake}
                                             onChange={(event) => this.changeStake(event, arbIndex)}/></td>
                    <td>Parlay Bonus: <input type="number" defaultValue={arb.bonus}
                                             onChange={(event) => this.changeBonus(event, arbIndex)}/></td>
                    <td>To Win: <input type="number" defaultValue={arb.toWin}
                                       onChange={(event) => this.changeToWin(event, arbIndex)}/></td>
                    {stakes}
                </tr>
                </tbody>
                <tfoot>
                <tr>
                    <td>
                        <input type="button" onClick={() => {
                            this.addEvent(arbIndex)
                        }} value="Add Event"/>

                        <input type="button" onClick={() => {
                            this.deleteArb(arbIndex)
                        }} value="Delete Arb" />
                    </td>
                    <td>Parlay Winnings: {winnings}</td>
                    <td>Total Stake: {totalStake}</td>
                    {seriesWinnings}
                </tr>
                </tfoot>
            </table>
        });

        return <React.Fragment>
            <h1>Arb Series</h1>
            {tables}
            <input type="button" onClick={this.addArb} value="Add Arb"/>
        </React.Fragment>;
    }
};


export default ArbSeries;