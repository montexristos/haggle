import React from 'react';
import Popup from "reactjs-popup";
import EventDetails from "./EventDetails";
import EventSiteView from "./EventSiteView";

class Event extends React.Component {

    render = () => {
        const fixture = this.props.fixture;
        const homeTeamName = fixture.homeTeam.name;
        const awayTeamName = fixture.awayTeam.name;
        const position = "right bottom";
        const dateOptions = { weekday: 'long', month: 'numeric', day: 'numeric', hour: 'numeric', minute: 'numeric' };
        const popup = <EventDetails fixture={fixture} />
        const events = fixture.odds;
        const sites = this.props.sites;
        const eventSites = [];
        let date = "";
        let marketMap = {
            "MRES": {},
            "OU": {},
            "BTTS": {},
        };
        let marketTable = null;
        if (this.props.sites && events && events.length) {
            for (let siteId=0; siteId<events.length;siteId++) {
                if (date === "" && events[siteId].Date) {
                    date = events[siteId].Date;
                }
                const site = sites[events[siteId].SiteID];
                const event = events[siteId];
                event.Markets.forEach((market) => {
                    if (!marketMap["MRES"][siteId] && market.MarketType === "MRES") {
                        marketMap["MRES"][siteId] = market;
                    }

                    if (!marketMap["OU"][siteId] && market.MarketType === "OU" && market.Line === 2.5) {
                        marketMap["OU"][siteId] = market;
                    }
                    if (!marketMap["BTTS"][siteId] && market.MarketType === "BTTS") {
                        marketMap["BTTS"][siteId] = market;
                    }

                })
                eventSites.push(
                    <EventSiteView site={site}
                                   siteId={siteId}
                                   event={events[siteId]}
                                   marketMap={marketMap}
                                   matchResult={marketMap["MRES"][siteId]}
                                   overUnder={marketMap["OU"][siteId]}
                                   btts={marketMap["BTTS"][siteId]}
                                   key={events[0].Name + site}
                    />
                );
            }
            marketTable = <table className="table is-striped is-fullwidth  is-hoverable">
                <thead>
                <tr>
                    <th>Site</th>
                    <th colSpan="3" className="has-text-centered">
                        Match Result
                    </th>
                    <th colSpan="2" className="has-text-centered">
                        Under/Over
                    </th>
                    <th colSpan="2" className="has-text-centered">
                        GG/NG
                    </th>
                </tr>
                </thead>
                <tbody>
                { eventSites }
                </tbody>
            </table>;
        }
        return <tr>
            <td>{ new Date(fixture.date).toLocaleDateString("el-GR", dateOptions) }</td>
            <td>
                <Popup
                    key={homeTeamName}
                    trigger={
                        <span>{homeTeamName} - {awayTeamName}</span>
                    }
                    position={position}
                    on={['hover']}
                    arrow={position !== 'center center'}
                >
                    {popup}
                </Popup>
            </td>
            <td>{marketTable}</td>
            <td>{fixture.score}</td>
        </tr>
    }

}

export default Event;