import React from 'react';

class StakeCalculator extends React.Component {
    constructor(props) {
        super(props);
        const bank = parseInt(localStorage.getItem("bank"));
        if (bank > 0) {

            this.state = {
                stake: 0,
                bank: bank,
                probability: 0.5
            };
        } else {
            this.state = {
                stake: 0,
                bank: 0,
                probability: 0.5
            };
        }
    }

    render = () => {
        localStorage.setItem("bank", parseInt(this.state.bank));
        const stake = Math.round(this.calculateStake(this.state.odd, this.state.probability) * this.state.bank * 100) / 100;
        const stakeView = isNaN(stake) ? "" : stake;
        return <div>
            <label>Odd</label>
            <input name="odd" onChange={this.updateOdd} />
            <label>Bank</label>
            <input name="bank" defaultValue={this.state.bank} onChange={this.updateBank} />
            <span>U: {this.state.bank * 0.02} - {this.state.bank * 0.06} &euro;</span>
            <input name="probability" defaultValue=".5" onChange={this.updateProbability} />
            <span>{stakeView}</span>
        </div>
    }

    updateOdd = (event) => {
        this.setState({
            odd: event.target.value
        });
    }
    updateBank = (event) => {
        this.setState({
            bank: event.target.value
        });
    }
    updateProbability = (event) => {
        this.setState({
            probability: event.target.value
        });
    }

    /**
     * For simple bets with two outcomes, one involving losing the entire amount bet, and the other involving winning the bet amount multiplied by the payoff odds, the Kelly bet is:

{\displaystyle f^{*}=p-{\frac {q}{b}}={\frac {bp-q}{b}}={\frac {bp-(1-p)}{b}}={\frac {p(b+1)-1}{b}}}{\displaystyle f^{*}=p-{\frac {q}{b}}={\frac {bp-q}{b}}={\frac {bp-(1-p)}{b}}={\frac {p(b+1)-1}{b}}}
where:

{\displaystyle f^{*}}f^{*} is the fraction of the current bankroll to wager; (i.e. how much to bet, expressed in fraction)
{\displaystyle b}b is the net fractional odds received on the wager; (e.g. betting $10, on win, rewards $4 plus wager; then {\displaystyle b=0.4}{\displaystyle b=0.4})
{\displaystyle p}p is the probability of a win;
{\displaystyle q=1-p}{\displaystyle q=1-p} is the probability of a loss.
     */
    calculateStake = (odd, probability) => {
        if (odd <= 1) {
            return 0;
        } 
        const b = odd - 1;
        const p = probability;
        const stake = (p * (b+1) - 1) / b
        return stake;
    }

}

export default StakeCalculator;