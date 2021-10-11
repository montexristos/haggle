import React from 'react';
import EventSiteView from "./EventSiteView";

class EventView extends React.Component {

    render = () => {
        const events = this.props.events;
        const sites = this.props.sites;
        const eventSites = [];
        let nameTag = "";
        let date = "";
        for (let siteId=0; siteId<events.length;siteId++) {
            if (siteId === 0) {
                nameTag = events[0].Name;
            }
            if (date === "" && events[siteId].Date) {
                date = events[siteId].Date;
            }
            const site = sites[events[siteId].SiteID];
            eventSites.push(
                <EventSiteView site={site} siteId={siteId} event={events[siteId]}/>
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