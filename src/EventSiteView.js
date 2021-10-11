import React from 'react';

class EventSiteView extends React.Component {

    render = () => {
        const event = this.props.event;
        const siteId = this.props.siteId;
        const site = this.props.site;
        const matchResult = {};
        const overUnder = {};
        const btts = {};
        matchResult[siteId] =event.Markets[0].MarketType ? event.Markets[0] : "";
        overUnder[siteId] = event.Markets[1].MarketType ? event.Markets[1] : "";
        btts[siteId] = event.Markets[2].MarketType ? event.Markets[2] : "";

        return <div className="columns">
            <span className="column">{site}</span>
            <span className="column">{matchResult[siteId] && matchResult[siteId].Selections ? matchResult[siteId].Selections[0].Price : ""}</span>
            <span className="column">{matchResult[siteId] && matchResult[siteId].Selections ? matchResult[siteId].Selections[1].Price : ""}</span>
            <span className="column">{matchResult[siteId] && matchResult[siteId].Selections ? matchResult[siteId].Selections[2].Price : ""}</span>
            <span className="column">{overUnder[siteId] && overUnder[siteId].Selections ? overUnder[siteId].Selections[0].Price : ""}</span>
            <span className="column">{overUnder[siteId] && overUnder[siteId].Selections ? overUnder[siteId].Selections[1].Price : ""}</span>
            <span className="column">{btts[siteId] && btts[siteId].Selections ? btts[siteId].Selections[0].Price : ""}</span>
            <span className="column">{btts[siteId] && btts[siteId].Selections ? btts[siteId].Selections[1].Price : ""}</span>
        </div>
    }
}

export default EventSiteView;