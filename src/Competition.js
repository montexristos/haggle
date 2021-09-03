import React from 'react';
import Event from "./Event";

class Competition extends React.Component {

    shouldComponentUpdate = (nextProps, nextState) => {
        return nextProps !== this.props;
      }
    render = () => {
        if (!this.props.competition.fixtures.length) {
            return null;
        }
        const events = this.props.events.map((fixture) => {
            if (this.props.hideCompleted && fixture.score !== "") {
                return null;
            }
            return <Event fixture={fixture}
                          today={this.props.today}
                          hideCompleted={this.props.hideCompleted}
                          ggs={this.props.ggs}
                          overs={this.props.overs}
                          key={fixture.homeTeam.name + "-" + fixture.awayTeam.name} />
        })
        // return <Event fixture={fixture} today={today} hideCompleted={this.state.hideCompleted}  key={fixture.homeTeam.name + "-" + fixture.awayTeam.name} />

        return  <React.Fragment>
            <h1>{this.props.competition.name}</h1>
            <table>
                {this.props.tableHeader}
                <tbody>
                {events}
                </tbody>
            </table>
        </React.Fragment>
    }
}

export default Competition;