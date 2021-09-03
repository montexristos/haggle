/**


 Parlay arb was an amazing tool to camouflage your action and boost your profits.
 Maybe you know, maybe you don't know; do you know you can change the odds in your bet ticket?

For example,

Dundee United Reserves 2,50
Chelsea 1,30
Liverpool 1,50

You can change odds in bet ticket how you want,
    You can change odds like this for example;

                                      Dundee United Reserves 2,00
Chelsea 1,30
Liverpool 1,875


Last leg of parlay (Liverpool) becomes 1,875 from 1,50! If opposite side is 3,20 which creates a %2.13 arb; Well, it becomes a Massive %18.23 ARB!

It's specially useful when arbing lower leagues where there's parlay requirement by some bookmakers.


You don't change them at the bookie, you virtually change them yourself as it suits you...the end result is the same mathematically.
If you have made a ticket for 100€ betting 3 games, each odd 2.0 you have a parlay paying 800€ if it wins.  You can "arb" the first game as if the odd was 1.5 with say no profit, then "arb" the 2nd game as if it was 2 and then the last game as if it was 2.667.
Your bet is made but you can adjust the counter bets by the odds you choose to set yourself, mathematically the end result is the same for your ticket but you can greatly increase profits by virtually creating a bigger arb in the last leg if you "decrease" the first odd and increase the last odd this way.  It's your choice.



When you parlay two events scheduled at same or similiar time. You have chance to get x25 of your arb profit. Yes, x25!

However, you reduce the overall profit little bit..but you keep your chance to hit the jackpot.

Allright, here's the parlay:

Man City 1,35     Saturday 17:00 (Lay is 1,25)
Liverpool 1,35    Saturday 17:15 (Lay is 1,25)

Stake: 3000€

The key is..We lay each bet seperate.

    Risk for City is 1120,39€ Lay stake is 4481,56€.
    Risk for Liverpool is 1120,39€. Lay stake is 4481,56€.

    Let's have a look at scenarios now..

Both City and Liverpool win +226€ Profit
City wins, Lİverpool not winner  +226€ Profit
Liverpool wins, City not winner +226€ Profit

What if both Liverpool and City does not win?

    Well, you HIT THE JACKPOT in that case!

    +5739€ profit if both Liverpool and City fail to win.

    Parlay bet lost -3000€
City lay win +4481,56€ (+4369,52€ Net)
Liverpool lay win +4481,56€ (+4369,52€ Net)

Net Profit: +5739€

PS: Betfair comission is %2.5

PS2: This strategy works best with short odds. Doesn't work very well on bigger odds.

PS3: If you think winning 2 lay bets is very unlikely, Well, think again! It just happened this weekend! Liverpool lost 7-2 while Leeds-Man City match ended 1-1.

PS4: If two events were scheduled at different times, and you were arbing classic way. You would make +223€ profit if first leg lost and +524€ profit if first leg was winner.


    HalfKelly
https://en.m.wikipedia.org/wiki/Kelly_criterion

*/
import React from 'react';

class Arb extends React.Component {

    constructor(props) {
        super(props);
        const storedArbs = localStorage.getItem("simplearb");
        let arbs = JSON.parse(storedArbs);
        if (!Array.isArray(arbs)) {
            arbs = [{
                sel1: 2.00,
                sel2: 2.00,
                name1: "",
                name2: "",
                stake1: 10,
                stake2: 10
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

    updateSel1 = (event, arbIndex) => {
        event.persist();
        const arbs = this.state.arbs;
        const arb = arbs[arbIndex];
        arb.sel1 = event.target.value
        this.setState({
            arbs: arbs
        });
    };
    updateSel2 = (event, arbIndex) => {
        event.persist();
        const arbs = this.state.arbs;
        const arb = arbs[arbIndex];
        arb.sel2 = event.target.value
        this.setState({
            arbs: arbs
        });
    };

    setName1 = (event, arbIndex) => {
        const arbs = this.state.arbs;
        const arb = arbs[arbIndex];
        arb.name1 = event.target.value;
        this.setState({
            arbs: arbs
        });
    }

    setName2 = (event, arbIndex) => {
        const arbs = this.state.arbs;
        const arb = arbs[arbIndex];
        arb.name2 = event.target.value;
        this.setState({
            arbs: arbs
        });
    }

    updateStake1 = (event, arbIndex) => {
        event.persist();
        const arbs = this.state.arbs;
        const arb = arbs[arbIndex];
        arb.stake1 = parseFloat(event.target.value);
        this.setState({
            arbs: arbs
        });
    }
    render = () => {
        const arbs = this.state.arbs.filter(function (el) {
            return el != null;
        });
        localStorage.setItem("simplearb", JSON.stringify(arbs));

        const tables = arbs.map((arb, arbIndex) => {
            const perc = this.calculatePerc(arb.sel1, arb.sel2);
            let percColor;
            if (perc < 99) {
                percColor = "green";
            } else {
                percColor = "transparent";
            }
            let winnings = 0;
            if (arb.stake1 > 0) {
                winnings = Math.round(100 * arb.sel1 * arb.stake1) / 100;
            }
            const stake2 = Math.round(100 * winnings / arb.sel2) / 100;
            const inputStyle = {
                width: '200px',
            };
            return <table className="table" key={"simple"+arbIndex}>
                <thead>
                <tr>
                    <th>Selection 1</th>
                    <th>Odd</th>
                    <th>Stake</th>
                    <th>Selection 2</th>
                    <th>Odd</th>
                    <th>Stake</th>
                    <th>Perc</th>
                </tr>
                </thead>
                <tbody>
                <tr>
                    <td><input type="text" defaultValue={arb.name1} style={inputStyle} maxLength="40" onBlur={(event) => {
                        this.setName1(event, arbIndex)
                    }} /></td>
                    <td><input name="selection1" type="number" step=".01" defaultValue="2.00" onChange={(event) => {
                        this.updateSel1(event, arbIndex)
                    }} /></td>
                    <td><input name="stake1" type="number" step="1" defaultValue="10.00" onChange={(event) => {
                        this.updateStake1(event, arbIndex)
                    }} /></td>
                    <td><input type="text" defaultValue={arb.name2} style={inputStyle} maxLength="40" onBlur={(event) => {
                        this.setName2(event, arbIndex)
                    }} /></td>
                    <td><input name="selection2" type="number" step=".01" defaultValue="2.00" onChange={(event) => {
                        this.updateSel2(event, arbIndex)
                    }} /></td>
                    <td>{stake2}</td>
                    <td id="arbPerc" style={{background: percColor}}>{perc} %</td>
                </tr>
                </tbody>
                <tfoot>
                <tr>
                    <td>total Stake: {arb.stake1 + stake2}</td>
                    <td>total winnings: {winnings}</td>
                </tr>
                </tfoot>
            </table>;
        })

        return <React.Fragment>
            <h1>Arb</h1>
            { tables }
        </React.Fragment>;
    }
};



export default Arb;