import React from 'react';

class ArbEvent extends React.Component {


    calculatePerc = (sel1, sel2) => {
        const perc1=(1/sel1)*100;
        const perc2=  (1/sel2) * 100;
        return Math.round(100 - (100 * (perc1+perc2)) / 100);
    }

    render = () => {
        let overridedOdd = this.props.odd
        if (this.props.overrideOdd) {
            overridedOdd = this.props.overrideOdd;
        }
        let arbStake = this.props.arbStake
        if (this.props.overrideOdd) {
            arbStake = this.props.arbStake;
        }
        const arbWinnings = Math.round(arbStake * this.props.arb * 100) / 100;
        const arbPerc = this.calculatePerc(overridedOdd, this.props.arb);
        const origArbPerc = this.calculatePerc(this.props.odd, this.props.arb);
        const inputStyle = {
            width: '200px',
        };
        return <tr>
            <td><input type="text" defaultValue={this.props.name} style={inputStyle} maxLength="40" onBlur={(event) => {
                this.props.setName(event, this.props.index, this.props.arbIndex)
            }} /></td>
            <td><input type="number" step="0.01" value={this.props.odd} onChange={(event) => {
                this.props.changeOdd(event, this.props.index, this.props.arbIndex)
            }} /></td>
            <td><input type="number" step="0.01" value={overridedOdd} onChange={
                (event) => {
                    this.props.overrideOddCb(event, this.props.index, this.props.arbIndex)
                }}/></td>
            <td><input type="number" step="0.01" value={this.props.arb} onChange={(event) => {
                this.props.setArb(event, this.props.index, this.props.arbIndex)
            }} /></td>
            <td>{arbPerc}% ({origArbPerc}%)</td>
            <td><input type="number" step="0.01" value={arbStake} onChange={
                (event) => {
                    this.props.overrideArbStake(event, this.props.index, this.props.arbIndex)}
                }/></td>
            <td>{arbWinnings}</td>
            <td>
                <input type="button" onClick={(event) =>
                    this.props.deleteEvent(event, this.props.index, this.props.arbIndex)
                } value="x" />
            </td>
            <td>
                <select value={this.props.status} onChange={(event) => {
                    this.props.changeStatus(event, this.props.index, this.props.arbIndex)
                }}>
                    <option name="pending">Pending</option>
                    <option name="won">won</option>
                    <option name="lost">lost</option>
                </select>
            </td>
        </tr>
    }
}

export default ArbEvent;