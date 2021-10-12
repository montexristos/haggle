import React from 'react';

class EventSiteView extends React.Component {

    render = () => {
        const event = this.props.event;
        const siteId = this.props.siteId;
        const site = this.props.site;
        const matchResult = this.props.matchResult;
        const overUnder = this.props.overUnder;
        const btts =  this.props.btts;
        return <div className="columns">
            <span className="column">{site}</span>
            <span className="column">{matchResult && matchResult.Selections && matchResult.Selections[0] ? matchResult.Selections[0].Price : ""}</span>
            <span className="column">{matchResult && matchResult.Selections && matchResult.Selections[1] ? matchResult.Selections[1].Price : ""}</span>
            <span className="column">{matchResult && matchResult.Selections && matchResult.Selections[2] ? matchResult.Selections[2].Price : ""}</span>
            <span className="column">{overUnder && overUnder.Selections && overUnder.Selections[0] ? overUnder.Selections[0].Price : ""}</span>
            <span className="column">{overUnder && overUnder.Selections && overUnder.Selections[1] ? overUnder.Selections[1].Price : ""}</span>
            <span className="column">{btts && btts.Selections && btts.Selections[0] ? btts.Selections[0].Price : ""}</span>
            <span className="column">{btts && btts.Selections&& btts.Selections[1] ? btts.Selections[1].Price : ""}</span>
        </div>
    }
}

export default EventSiteView;