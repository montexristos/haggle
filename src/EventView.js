import React from 'react';
import EventSiteView from "./EventSiteView";

class EventView extends React.Component {

    render = () => {
        const events = this.props.events;
        const sites = this.props.sites;
        const eventSites = [];
        let nameTag = "";
        let date = "";
        let marketMap = {
            "MRES": {},
            "OU": {},
            "BTTS": {},
        };
        for (let siteId=0; siteId<events.length;siteId++) {
            if (siteId === 0) {
                nameTag = events[0].CanonicalName;
            }
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
        let dt;
        if (date !== "") {
            dt = Date.parse(date);
            dt.toString();
        }
        return <tr key={events[0].Name}>
            <td>
                {nameTag} <br/>
                {date}
            </td>
            <td>
                {eventSites}
            </td>
        </tr>
    }
}

export default EventView;