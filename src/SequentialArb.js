/**


 Parlay arb was an amazing tool to camouflage your action and boost your profits.
 Maybe you know, maybe you don't know; do you know you can change the odds in your bet ticket?

 For example,

 Dundee United  2,50
 Chelsea 1,30
 Liverpool 1,50

 You can change odds in bet ticket how you want,
 You can change odds like this for example;

 Dundee United Reserves 2,00
 Chelsea 1,30
 Liverpool 1,875

 Chelsea 1.50 (Opponent +0,50 is 3.50)
 Porto 1.90 (Opponent +0,50 is 2.25)
 Liverpool 1.40 (Opponent +0,50 is 3.50)

 This parlay pays 3,99

 You can change Chelsea odds.

 You make Chelsea 1,40. Since you bet other side at 3.50, you don't make any profit in first leg at all.

 Parlay was paying 3,99 with Chelsea. Well, you decreased Chelsea to 1,40.

 So the remaining odds total is 2,85! (3,99 / 1,40)

 You can keep Porto at 1,90. This way, you will have Liverpool at 1,50.
 Alternatively, you can decrease Porto's price to 1,80. This way, you won't make any profit in second leg as well but if Porto wins, you will have Liverpol at 1,5833!

 HalfKelly
 https://en.m.wikipedia.org/wiki/Kelly_criterion

 */
import React from 'react';
import ArbEvent from './ArbEvent';

class SequentialArb extends React.Component {

    constructor(props) {
        super(props);
        const storedArbs = localStorage.getItem("seqarbs");
        let arbs = JSON.parse(storedArbs);
        if(!Array.isArray(arbs)){
            arbs = [{
                events: [],
                stake: 10,
                bonus: 0
            }];
        }
        this.state = {
            arbs: arbs || []
        };
    }

    calculatePerc = (sel1, sel2) => {
        const perc1=(1/sel1)*100;
        const perc2=  (1/sel2) * 100;
        return Math.round(100 * (perc1+perc2)) / 100;
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
        arb.events[idx].arbStake = parseFloat(event.target.value);
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
        console.log(arbIndex);
        const arb = arbs[arbIndex];
        arb.events.push({
            name: "",
            odd: 2.0,
            overrideOdd: 2.0,
            arb: 2.0,
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
        localStorage.setItem("seqarbs", JSON.stringify(arbs));
        let events = null;
        const tables = arbs.map((arb, arbIndex) => {
            let totalStake = arb.stake;
            let adjustedMultiplier = 1;
            let parlayLost = false;
            let winnings = 0;
            let adjustedWinnings = 0;
            let arbWinnings = 0;
            //arb winnings should be the stake plus all arb stakes
            if (arb.events && arb.events.length > 0) {
                events = arb.events.map((event, index) => {
                    const name = event.name;
                    const odd = event.odd;
                    // let shouldWin = arb.stake;

                    if (index === 0) {
                        // event.arbStake = Math.round(100 * arb.stake / (event.arb -1)) / 100;
                    } else {
                        // for (let i=0;i<index;i++) {
                        //     if (arb.events[i].status === "lost") {
                        //         parlayLost = true;
                        //     }
                        //     if (arb.events[i].status !== "lost") {
                        //         shouldWin += arb.events[i].arbStake;
                        //     }
                        // }
                        //count all stakes until now
                        // event.arbStake = Math.round(100 * shouldWin / (event.arb -1)) / 100;
                        // if (parlayLost) {
                        //     event.arbStake = 0;
                        // }
                    }
                    // if (event.arb === 0) {
                    //     event.arbStake = 0;
                    // }
                    totalStake += parseFloat(event.arbStake);
                    const arbWinnings = event.arbStake * event.arb;
                    return <ArbEvent name={name}
                                     odd={odd}
                                     overrideOdd={event.overrideOdd}
                                     overrideOddCb={this.overrideOddCb}
                                     changeOdd={this.changeOdd}
                                     setName={this.setName}
                                     setArb={this.setArb}
                                     overrideArbStake={this.overrideArbStake}
                                     arb={event.arb}
                                     arbWinnings={arbWinnings}
                                     arbStake={event.arbStake}
                                     key={arbIndex + "-" + index+"-"+event.name}
                                     index={index}
                                     changeStatus={this.changeStatus}
                                     arbIndex={arbIndex}
                                     status={event.status}
                                     deleteEvent={this.deleteEvent}
                    />;
                });
                winnings = arb.stake;
                adjustedWinnings = arb.stake;
                arb.events.forEach((event, i) => {
                    winnings = winnings * event.odd;
                    adjustedWinnings = adjustedWinnings * event.overrideOdd
                    if (event.status === "lost") {
                        arbWinnings += Math.round(event.arbStake * event.arb * 100) / 100;
                    }
                });
                winnings = winnings + arb.bonus;
            } else {
                events = null;
            }
            totalStake = Math.round(totalStake * 100) / 100;
            winnings = Math.round(winnings * 100) / 100;
            adjustedWinnings = Math.round(adjustedWinnings * 100) / 100;
            return <table className="table" key={arbIndex}>
                <thead>
                <tr>
                    <td>Event</td>
                    <td>Odd</td>
                    <td>Adjusted Odd</td>
                    <td>Arb</td>
                    <td>Arb Perc</td>
                    <td>Arb Stake</td>
                    <td>Arb Winnings</td>
                    <td colSpan="2">Actions</td>
                </tr>
                </thead>
                <tbody>
                { events }
                <tr>
                    <td><input type="number" defaultValue={arb.stake} onChange={(event) => this.changeStake(event, arbIndex)}/></td>
                    <td><input type="number" defaultValue={arb.bonus} onChange={(event) => this.changeBonus(event, arbIndex)}/></td>
                    <td>{adjustedMultiplier}</td>
                    <td></td>
                </tr>
                </tbody>
                <tfoot>
                <tr>
                    <td>
                        <input type="button" onClick={() => {
                            this.addEvent(arbIndex)
                        }} value="Add Event" />

                        <input type="button" onClick={() => {
                            this.deleteArb(arbIndex)
                        }} value="Delete Arb" />
                    </td>
                    <td>Parlay Winnings: {winnings}</td>
                    <td>Adj Winnings: {adjustedWinnings}</td>
                    <td>Total Stake: {totalStake}</td>
                    <td>Arb Winnings: {arbWinnings}</td>
                </tr>
                </tfoot>
            </table>
        });

        return <React.Fragment>
            <h1>Sequential Arb</h1>
            {tables}
            <input type="button" onClick={this.addArb} value="Add Arb" />
        </React.Fragment>;
    }
};



export default SequentialArb;