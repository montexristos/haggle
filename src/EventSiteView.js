import React from 'react';

class EventSiteView extends React.Component {

    render = () => {
        const site = this.props.site;
        const matchResult = this.props.matchResult;
        const overUnder = this.props.overUnder;
        const btts =  this.props.btts;
        return <tr>
            <td>{site}</td>
            <td className="has-text-centered">{matchResult && matchResult.Selections && matchResult.Selections[0] ? matchResult.Selections[0].Price : ""}</td>
            <td className="has-text-centered">{matchResult && matchResult.Selections && matchResult.Selections[1] ? matchResult.Selections[1].Price : ""}</td>
            <td className="has-text-centered">{matchResult && matchResult.Selections && matchResult.Selections[2] ? matchResult.Selections[2].Price : ""}</td>
            <td className="has-text-centered">{overUnder && overUnder.Selections && overUnder.Selections[0] ? overUnder.Selections[0].Price : ""}</td>
            <td className="has-text-centered">{overUnder && overUnder.Selections && overUnder.Selections[1] ? overUnder.Selections[1].Price : ""}</td>
            <td className="has-text-centered">{btts && btts.Selections && btts.Selections[0] ? btts.Selections[0].Price : ""}</td>
            <td className="has-text-centered">{btts && btts.Selections&& btts.Selections[1] ? btts.Selections[1].Price : ""}</td>
        </tr>
    }
}

export default EventSiteView;