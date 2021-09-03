import React from 'react';
import ArbSeriesEvent from './ArbSeriesEvent';

class ArbSeriesNew extends React.Component {

    constructor(props) {
        super(props);
        const storedArbs = localStorage.getItem("arbseries");
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

    render = () => {
        const arbs = this.state.arbs.filter(function (el) {
            return el != null;
        });
        localStorage.setItem("arbseries", JSON.stringify(arbs));
        let events = null;
        const tables = arbs.map((arb, arbIndex) => {
            let headers;
            let multiplierArray = {};
            let seriesWinnings;
            let stakes;
            let mults;
            let winnings = 0;
            let totalStake = 0;
            const originalMultiplier = this.getMultiplier(arb);
            const parlayStake = arb.stake;
            // let parlayLost = false;
            if (arb.events && arb.events.length > 0) {
                events = arb.events.map((event, index) => {
                    multiplierArray[index] = 1;
                    const series = arb.events.map((evt, idx) => {
                        const key = evt.name + idx;
                        if (idx === index) {
                            if (arb.events[index].arb) {
                                multiplierArray[idx] = multiplierArray[idx] * arb.events[index].arb;
                            } else {
                                multiplierArray[idx] = 0;
                            }
                            return <td key={key}>{arb.events[index].arb}</td>;
                        }
                        if (idx < index) {
                            multiplierArray[idx] = multiplierArray[idx] * arb.events[index].odd;
                            return <td key={key}>{arb.events[index].odd}</td>;
                        }
                        return <td key={key}></td>;
                    });

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
                                           key={arbIndex + "-" + index+"-"+event.name}
                                           index={index}
                                           changeStatus={this.changeStatus}
                                           arbIndex={arbIndex}
                                           status={event.status}
                                           series={series}
                                           deleteEvent={this.deleteEvent}
                    />;
                });
                const stakeArray = {};
                const multArray = {};
                let toWin = parlayStake;
                let parlayMult = 1;
                arb.events.forEach((r, j)  => {
                    parlayMult *= arb.events[j].odd;
                    let seriesMult = 1;
                    // console.log(j)
                    for (let q=j;q<arb.events.length;q++){
                        let odd;
                        if (q === j) {
                            odd = arb.events[q].arb;
                        } else {
                            odd = arb.events[q].odd;
                        }
                        seriesMult *= odd
                    }
                    console.log(seriesMult);

                    multArray[j] = seriesMult;
                });
                toWin = parlayStake * parlayMult;
                // console.log(multArray);
                stakes = arb.events.map((r, j) => {
                    if (multArray[j]) {
                        //for the first selections we want to break even
                        stakeArray[j] = Math.round(100 * (toWin) / (multArray[j])) / 100;
                    } else {
                        stakeArray[j] = 0;
                    }
                    totalStake += parseFloat(stakeArray[j]);
                    return <td key={"stake-" + j}>{stakeArray[j]}</td>;
                })
                mults = arb.events.map((r, j) => {
                    return <td key={"mult-" + j}>{multArray[j]}</td>;
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
                    <td></td>
                    <td></td>
                    <td></td>
                    {mults}
                </tr>
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
            <h1>Arb Series New</h1>
            {tables}
            <input type="button" onClick={this.addArb} value="Add Arb"/>
        </React.Fragment>;
    }
};


export default ArbSeriesNew;