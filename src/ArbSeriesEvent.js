import React from 'react';

class ArbSeriesEvent extends React.Component {

    calculatePerc = (sel1, sel2) => {
        if (sel1 === 0 || sel2 === 0) {
            return 0;
        }
        const perc1=(1/sel1)*100;
        const perc2=  (1/sel2) * 100;
        return Math.round(100 * (perc1+perc2)) / 100;
    }

    render = () => {
        const arbPerc = this.calculatePerc(this.props.odd, this.props.arb);
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
            <td><input type="number" step="0.01" value={this.props.arb} onChange={(event) => {
                this.props.setArb(event, this.props.index, this.props.arbIndex)
            }} />({arbPerc}%)</td>
            { this.props.series }
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

export default ArbSeriesEvent;