import React from 'react';
import Popup from "reactjs-popup";
import EventDetails from "./EventDetails";

class Event extends React.Component {


    getNoEvents(homeTeam, awayTeam, over0, over1, over2, over3, over4,
                predicted0,predicted1,predicted2,predicted3,predicted4) {
        let class0 = over0 > 0 ? "classWon" : over0 < 0 ? "classLost" : "";
        let class1 = over1 > 0 ? "classWon" : over1 < 0 ? "classLost" : "";
        let class2 = over2 > 0 ? "classWon" : over2 < 0 ? "classLost" : "";
        let class3 = over3 > 0 ? "classWon" : over3 < 0 ? "classLost" : "";
        let class4 = over4 > 0 ? "classWon" : over4 < 0 ? "classLost" : "";
        if (predicted4 === 1) {
            class4 += " classPredicted";
        } else if (predicted4 === 0) {
            class4 += " classPredictedNot";
        }
        if (predicted3 === 1) {
            class3 += " classPredicted";
        } else if (predicted4 === 0) {
            class3 += " classPredictedNot";
        }
        if (predicted2 === 1) {
            class2 += " classPredicted";
        } else if (predicted4 === 0) {
            class2 += " classPredictedNot";
        }
        if (predicted1 === 1) {
            class1 += " classPredicted";
        } else if (predicted4 === 0) {
            class1 += " classPredictedNot";
        }
        if (predicted0 === 1) {
            class0 += " classPredicted";
        } else if (predicted4 === 0) {
            class0 += " classPredictedNot";
        }
        return <React.Fragment>
            <table>
                <tbody>
                    <tr>
                        <td className={class0}>o0</td>
                        <td className={class1}>o1</td>
                        <td className={class2}>o2</td>
                        <td className={class3}>o3</td>
                        <td className={class4}>o4</td>
                    </tr>
                    <tr>
                        <td>{homeTeam.over0}</td>
                        <td>{homeTeam.over1}</td>
                        <td>{homeTeam.over2}</td>
                        <td>{homeTeam.over3}</td>
                        <td>{homeTeam.over4}</td>
                    </tr>
                    <tr>
                        <td>{awayTeam.over0}</td>
                        <td>{awayTeam.over1}</td>
                        <td>{awayTeam.over2}</td>
                        <td>{awayTeam.over3}</td>
                        <td>{awayTeam.over4}</td>
                    </tr>
                </tbody>
            </table>
        </React.Fragment>;
    }
    getUnderEvents(homeTeam, awayTeam, under0, under1, under2, under3, under4,
                predicted0,predicted1,predicted2,predicted3,predicted4) {
        let class0 = under0 > 0 ? "classWon" : under0 < 0 ? "classLost" : "";
        let class1 = under1 > 0 ? "classWon" : under1 < 0 ? "classLost" : "";
        let class2 = under2 > 0 ? "classWon" : under2 < 0 ? "classLost" : "";
        let class3 = under3 > 0 ? "classWon" : under3 < 0 ? "classLost" : "";
        let class4 = under4 > 0 ? "classWon" : under4 < 0 ? "classLost" : "";
        if (predicted4 === 1) {
            class4 += " classPredicted";
        } else if (predicted4 === 0) {
            class4 += " classPredictedNot";
        }
        if (predicted3 === 1) {
            class3 += " classPredicted";
        } else if (predicted4 === 0) {
            class3 += " classPredictedNot";
        }
        if (predicted2 === 1) {
            class2 += " classPredicted";
        } else if (predicted4 === 0) {
            class2 += " classPredictedNot";
        }
        if (predicted1 === 1) {
            class1 += " classPredicted";
        } else if (predicted4 === 0) {
            class1 += " classPredictedNot";
        }
        if (predicted0 === 1) {
            class0 += " classPredicted";
        } else if (predicted4 === 0) {
            class0 += " classPredictedNot";
        }
        return <React.Fragment>
            <table>
                <tbody>
                    <tr>
                        <td className={class0}>u0</td>
                        <td className={class1}>u1</td>
                        <td className={class2}>u2</td>
                        <td className={class3}>u3</td>
                        <td className={class4}>u4</td>
                    </tr>
                    <tr>
                        <td>{homeTeam.under0}</td>
                        <td>{homeTeam.under1}</td>
                        <td>{homeTeam.under2}</td>
                        <td>{homeTeam.under3}</td>
                        <td>{homeTeam.under4}</td>
                    </tr>
                    <tr>
                        <td>{awayTeam.under0}</td>
                        <td>{awayTeam.under1}</td>
                        <td>{awayTeam.under2}</td>
                        <td>{awayTeam.under3}</td>
                        <td>{awayTeam.under4}</td>
                    </tr>
                </tbody>
            </table>
        </React.Fragment>;
    }

    render = () => {
        const fixture = this.props.fixture;
        const eventDate = new Date(fixture.date);
        const homeTeamName = fixture.homeTeam.name;
        const awayTeamName = fixture.awayTeam.name;
        const noEvents = this.getNoEvents(
            fixture.homeTeam, fixture.awayTeam,
            fixture.over0, fixture.over1, fixture.over2, fixture.over3, fixture.over4,
            fixture.predictedOver0,fixture.predictedOver1, fixture.predictedOver2,
            fixture.predictedOver3, fixture.predictedOver4
        );
        const noUnderEvents = this.getUnderEvents(
            fixture.homeTeam, fixture.awayTeam,
            fixture.under0, fixture.under1, fixture.under2, fixture.under3, fixture.under4,
            fixture.predictedUnder0, fixture.predictedUnder1, fixture.predictedUnder2,
            fixture.predictedUnder3, fixture.predictedUnder4
        );
        const totalIndex = fixture.totalIndex;
        const bg = {background: 'transparent'};
        if (eventDate < this.props.today && this.props.hideCompleted) {
            return null;
        }
        if (fixture.homeTeam.Name === '') {
            console.log(fixture);
        }
        const position = "right bottom";
        const dateOptions = { weekday: 'long', month: 'numeric', day: 'numeric', hour: 'numeric', minute: 'numeric' };
        const popup = <EventDetails fixture={fixture} />
        const name = homeTeamName + " - " + awayTeamName;
        const over = this.props.overs.includes(name) ? 'OVER' : 'UNDER';
        const gg = this.props.ggs.includes(name) ? 'GG' : 'NG';

        return <tr>
            <td>{ new Date(fixture.date).toLocaleDateString("el-GR", dateOptions) }</td>
            {/*<td>{ new Date(fixture.date).toTimeString() }</td>*/}
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
            <td>{noEvents}</td>
            <td>{noUnderEvents}</td>
            {/*<td className={cardBg}>{cards}</td>*/}
            {/*<td className={cornerBg}>{cornerIndex}</td>*/}
            {/*<td>{homeCorners} - {awayCorners}</td>*/}
            <td>{over}</td>
            <td>{gg}</td>
            <td style={bg}>{Math.round(totalIndex)}</td>
            <td>{fixture.score}</td>
        </tr>
    }

}

export default Event;