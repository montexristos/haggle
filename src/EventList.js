import React from "react";
import EventView from "./EventView";

class EventList extends React.Component {

    render() {
        let rows = [];
        const sites = this.props.sites;
        for (const eventId in this.props.events) {
            const events = this.props.events[eventId];
            if (!events) {
                continue;
            }
            if (events.length) {
                rows.push(<EventView events={events} sites={sites}/>);
            }
        }
        return <table className="table is-striped is-fullwidth  is-hoverable">
            <thead>
            <tr>
                <th>Event name</th>
                <th className="columns">
                    <span>match result</span>
                    <span>under/over</span>
                    <span>BTTS</span>
                </th>

            </tr>
            </thead>
            <tbody>
            {rows}
            </tbody>
        </table>;
    }
}

export default EventList;