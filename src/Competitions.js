import React from 'react';
import Competition from "./Competition";
import Event from "./Event";

class Competitions extends React.Component {

    constructor(props) {
        super(props);
        this.state = {
            competitions: [],
            selectedDate: "all",
            correct: 0,
            wrong: 0
        };    
    }

    // Returns an array of dates between the two dates
    getDays = (startDate, endDate) => {
        var current = new Date(startDate);
        var end = new Date(endDate);
        var dates = [],
            currentDate = current,
            addDays = function(days) {
                var date = new Date(this.valueOf());
                date.setDate(date.getDate() + days);
                return date;
            };
        while (currentDate <= end) {
            dates.push(currentDate);
            currentDate = addDays.call(currentDate, 1);
        }
        return dates;
    };

    render = () => {
        let evs = [];

        if (!this.props.tournaments || !this.props.weeks) {
            return null;
        }
        const weeks =  this.props.weeks.map((week) => {
            return <option key={week.start} value={week.end}>{week.start} - {week.end}</option>;
        });
        const weekSelect = <select onChange={this.props.changeWeek}>
            <option value="all">all</option>
            { weeks }
        </select>;
        const days = this.getDays(this.props.start, this.props.end);
        const dayList = days.map((day) => {
            return <option value={day} key={day.toLocaleDateString()}>{ day.toLocaleDateString() }</option>
        });
        const daySelect = <select onChange={this.props.filterDay}>
            <option value="all">All</option>
            { dayList }
        </select>;

        const tableHeader = this.getTableHeader();
        const comps = this.props.tournaments.map((comp) => {
            const events = comp.fixtures.filter((fixture) =>
                (fixture.homeTeam && fixture.awayTeam) && ("all" === this.props.date || new Date(fixture.date).getDay() === this.props.selectedDate)
            ).map((fixture) => {
                if (
                    (this.props.showO0 && fixture.predictedOver0 < 1) ||
                    (this.props.showO1 && fixture.predictedOver1 < 1) ||
                    (this.props.showO2 && fixture.predictedOver2 < 1) ||
                    (this.props.showO3 && fixture.predictedOver3 < 1) ||
                    (this.props.showO4 && fixture.predictedOver4 < 1) ||
                    (this.props.showU0 && fixture.predictedOver0 > 0) ||
                    (this.props.showU1 && fixture.predictedOver1 > 0) ||
                    (this.props.showU2 && fixture.predictedOver2 > 0) ||
                    (this.props.showU3 && fixture.predictedOver3 > 0) ||
                    (this.props.showU4 && fixture.predictedOver4 > 0)
                ) {
                    return null;
                }
                return fixture;
                // return <Event fixture={fixture} today={today} hideCompleted={this.state.hideCompleted}  key={fixture.homeTeam.name + "-" + fixture.awayTeam.name} />
            }).filter(function (el) {
                return el != null;
              });
            if (events.length === 0) {
                return null
            }
            evs = [...evs, ...events];
            return <div key={comp.name}>
                <Competition competition={comp}
                             today={this.props.today}
                             hideCompleted={this.props.hideCompleted}
                             date={this.props.filterDay}
                             events={events}
                             tableHeader={tableHeader}
                             ggs={this.props.ggs}
                             overs={this.props.overs}
                />
            </div>
        });

        let competitions = [];
        if (this.props.sort === 'date') {
            evs.sort(function(a,b){
                // Turn your strings into dates, and then subtract them
                // to get a value that is either negative, positive, or zero.
                return new Date(a.date) - new Date(b.date);
            });
            const events = evs.map((fixture) => {
                console.log(this.props.showO1);
                if (
                    (this.props.showO0 && fixture.predicted0 < 1) ||
                    (this.props.showO1 && fixture.predicted1 < 1) ||
                    (this.props.showO2 && fixture.predicted2 < 1) ||
                    (this.props.showO3 && fixture.predicted3 < 1) ||
                    (this.props.showO4 && fixture.predicted4 < 1) ||
                    (this.props.hideCompleted && fixture.score !== "")
                ) {
                    return null;
                }
                return <Event fixture={fixture}
                              today={this.props.today}
                              hideCompleted={this.props.hideCompleted}
                              ggs={this.props.ggs}
                              overs={this.props.overs}
                              key={fixture.homeTeam.name + "-" + fixture.awayTeam.name}
                />
            })

            competitions = <table id="events">
                {/*<button onClick={this.sortTable}>sort</button>*/}
                {tableHeader}
                <tbody>
                {events}
                </tbody>
            </table>
        } else {
            competitions = comps;
        }


        return <div>
            {weekSelect}
            {daySelect}
            {competitions}
        </div>
    }

    getTableHeader = () => {
        return <thead>
                <tr>
                    <th className="eventName">Date</th>
                    <th className="eventName">Name</th>
                    <th className="eventName">Overs</th>
                    <th className="eventName">Unders</th>
                    {/*<th>Cards</th>*/}
                    {/*<th>CornerIndex</th>*/}
                    {/*<th>Corners</th>*/}
                    <th>over</th>
                    <th>gg</th>
                    <th>Index</th>
                    <th>score</th>
                </tr>
            </thead>;
    }

    sortTable = () => {
        var table, rows, switching, i, x, y, shouldSwitch;
        table = document.getElementById("events");
        switching = true;
        /*Make a loop that will continue until
        no switching has been done:*/
        while (switching) {
            //start by saying: no switching is done:
            switching = false;
            rows = table.rows;
            /*Loop through all table rows (except the
            first, which contains table headers):*/
            for (i = 1; i < (rows.length - 1); i++) {
                //start by saying there should be no switching:
                shouldSwitch = false;
                /*Get the two elements you want to compare,
                one from current row and one from the next:*/
                x = rows[i].getElementsByTagName("TD")[0];
                y = rows[i + 1].getElementsByTagName("TD")[0];
                //check if the two rows should switch place:
                const xdate = Date.parse(x.innerHTML);
                const ydate = Date.parse(y.innerHTML);
                if (xdate > ydate) {
                    //if so, mark as a switch and break the loop:
                    shouldSwitch = true;
                    break;
                }
            }
            if (shouldSwitch) {
                /*If a switch has been marked, make the switch
                and mark that a switch has been done:*/
                rows[i].parentNode.insertBefore(rows[i + 1], rows[i]);
                switching = true;
            }
        }
    }
};



export default Competitions;