import React from 'react';

class EventDetails extends React.Component {

    render = () => {
        const fixture = this.props.fixture;
        // const fixture = JSON.parse("{\"date\":\"2020-11-21T15:00:00Z\",\"homeTeam\":{\"name\":\"Exeter\",\"goalsScored\":14.4,\"goalsConceided\":11.4,\"homeGoalsScored\":7.800000000000002,\"homeGoalsConceided\":3.800000000000001,\"awayGoalsScored\":6.600000000000002,\"awayGoalsConceided\":7.600000000000002,\"HomeAwayScored\":0,\"HomeAwayConceided\":0,\"AwayHomeScored\":0,\"AwayHomeConceided\":0,\"noEvents\":49,\"homeEvents\":23,\"awayEvents\":26,\"points\":86,\"corners\":553,\"homeCorners\":153,\"homeAwayCorners\":100,\"awayCorners\":163,\"awayHomeCorners\":137,\"yellowCards\":58,\"redCards\":4,\"totalCards\":155,\"fixtures\":[{\"date\":\"2019-08-03T15:00:00Z\",\"homeTeamName\":\"Exeter\",\"awayTeamName\":\"Macclesfield\",\"score\":\"1 - 0\",\"homeCorners\":7,\"awayCorners\":6,\"homeYellowCards\":1,\"awayYellowCards\":3,\"homeRedCards\":0,\"awayRedCards\":0},{\"date\":\"2019-08-10T15:00:00Z\",\"homeTeamName\":\"Stevenage\",\"awayTeamName\":\"Exeter\",\"score\":\"0 - 1\",\"homeCorners\":4,\"awayCorners\":4,\"homeYellowCards\":1,\"awayYellowCards\":2,\"homeRedCards\":0,\"awayRedCards\":0},{\"date\":\"2019-08-17T15:00:00Z\",\"homeTeamName\":\"Exeter\",\"awayTeamName\":\"Swindon\",\"score\":\"1 - 1\",\"homeCorners\":4,\"awayCorners\":10,\"homeYellowCards\":1,\"awayYellowCards\":2,\"homeRedCards\":0,\"awayRedCards\":0},{\"date\":\"2019-08-20T19:45:00Z\",\"homeTeamName\":\"Oldham\",\"awayTeamName\":\"Exeter\",\"score\":\"0 - 0\",\"homeCorners\":7,\"awayCorners\":10,\"homeYellowCards\":2,\"awayYellowCards\":0,\"homeRedCards\":0,\"awayRedCards\":0},{\"date\":\"2019-08-24T15:00:00Z\",\"homeTeamName\":\"Morecambe\",\"awayTeamName\":\"Exeter\",\"score\":\"2 - 3\",\"homeCorners\":5,\"awayCorners\":9,\"homeYellowCards\":0,\"awayYellowCards\":0,\"homeRedCards\":0,\"awayRedCards\":1},{\"date\":\"2019-08-31T15:00:00Z\",\"homeTeamName\":\"Exeter\",\"awayTeamName\":\"Mansfield\",\"score\":\"1 - 0\",\"homeCorners\":3,\"awayCorners\":6,\"homeYellowCards\":2,\"awayYellowCards\":1,\"homeRedCards\":0,\"awayRedCards\":1},{\"date\":\"2019-09-07T15:00:00Z\",\"homeTeamName\":\"Carlisle\",\"awayTeamName\":\"Exeter\",\"score\":\"1 - 3\",\"homeCorners\":4,\"awayCorners\":5,\"homeYellowCards\":2,\"awayYellowCards\":0,\"homeRedCards\":0,\"awayRedCards\":0},{\"date\":\"2019-09-14T15:00:00Z\",\"homeTeamName\":\"Exeter\",\"awayTeamName\":\"Leyton Orient\",\"score\":\"2 - 2\",\"homeCorners\":9,\"awayCorners\":2,\"homeYellowCards\":0,\"awayYellowCards\":6,\"homeRedCards\":0,\"awayRedCards\":0},{\"date\":\"2019-09-17T19:45:00Z\",\"homeTeamName\":\"Exeter\",\"awayTeamName\":\"Port Vale\",\"score\":\"2 - 0\",\"homeCorners\":9,\"awayCorners\":4,\"homeYellowCards\":0,\"awayYellowCards\":4,\"homeRedCards\":0,\"awayRedCards\":0},{\"date\":\"2019-09-21T15:00:00Z\",\"homeTeamName\":\"Newport County\",\"awayTeamName\":\"Exeter\",\"score\":\"1 - 1\",\"homeCorners\":8,\"awayCorners\":7,\"homeYellowCards\":1,\"awayYellowCards\":3,\"homeRedCards\":0,\"awayRedCards\":0},{\"date\":\"2019-09-28T15:00:00Z\",\"homeTeamName\":\"Exeter\",\"awayTeamName\":\"Grimsby\",\"score\":\"1 - 3\",\"homeCorners\":8,\"awayCorners\":1,\"homeYellowCards\":4,\"awayYellowCards\":4,\"homeRedCards\":0,\"awayRedCards\":0},{\"date\":\"2019-10-05T15:00:00Z\",\"homeTeamName\":\"Crewe\",\"awayTeamName\":\"Exeter\",\"score\":\"1 - 1\",\"homeCorners\":4,\"awayCorners\":10,\"homeYellowCards\":3,\"awayYellowCards\":1,\"homeRedCards\":0,\"awayRedCards\":0},{\"date\":\"2019-10-12T15:00:00Z\",\"homeTeamName\":\"Exeter\",\"awayTeamName\":\"Forest Green\",\"score\":\"1 - 0\",\"homeCorners\":11,\"awayCorners\":4,\"homeYellowCards\":3,\"awayYellowCards\":3,\"homeRedCards\":0,\"awayRedCards\":0},{\"date\":\"2019-10-19T15:00:00Z\",\"homeTeamName\":\"Cambridge\",\"awayTeamName\":\"Exeter\",\"score\":\"4 - 0\",\"homeCorners\":3,\"awayCorners\":3,\"homeYellowCards\":0,\"awayYellowCards\":2,\"homeRedCards\":0,\"awayRedCards\":0},{\"date\":\"2019-10-22T19:45:00Z\",\"homeTeamName\":\"Scunthorpe\",\"awayTeamName\":\"Exeter\",\"score\":\"3 - 1\",\"homeCorners\":1,\"awayCorners\":4,\"homeYellowCards\":2,\"awayYellowCards\":1,\"homeRedCards\":0,\"awayRedCards\":0},{\"date\":\"2019-10-26T13:00:00Z\",\"homeTeamName\":\"Exeter\",\"awayTeamName\":\"Plymouth\",\"score\":\"4 - 0\",\"homeCorners\":4,\"awayCorners\":4,\"homeYellowCards\":2,\"awayYellowCards\":2,\"homeRedCards\":0,\"awayRedCards\":0},{\"date\":\"2019-11-02T15:00:00Z\",\"homeTeamName\":\"Bradford\",\"awayTeamName\":\"Exeter\",\"score\":\"2 - 0\",\"homeCorners\":3,\"awayCorners\":6,\"homeYellowCards\":1,\"awayYellowCards\":0,\"homeRedCards\":1,\"awayRedCards\":2},{\"date\":\"2019-11-16T15:00:00Z\",\"homeTeamName\":\"Exeter\",\"awayTeamName\":\"Cheltenham\",\"score\":\"0 - 0\",\"homeCorners\":6,\"awayCorners\":0,\"homeYellowCards\":2,\"awayYellowCards\":2,\"homeRedCards\":0,\"awayRedCards\":0},{\"date\":\"2019-11-23T15:00:00Z\",\"homeTeamName\":\"Crawley Town\",\"awayTeamName\":\"Exeter\",\"score\":\"0 - 1\",\"homeCorners\":6,\"awayCorners\":9,\"homeYellowCards\":2,\"awayYellowCards\":1,\"homeRedCards\":0,\"awayRedCards\":0},{\"date\":\"2019-12-07T15:00:00Z\",\"homeTeamName\":\"Exeter\",\"awayTeamName\":\"Northampton\",\"score\":\"3 - 2\",\"homeCorners\":9,\"awayCorners\":4,\"homeYellowCards\":2,\"awayYellowCards\":1,\"homeRedCards\":0,\"awayRedCards\":0},{\"date\":\"2019-12-14T15:00:00Z\",\"homeTeamName\":\"Salford\",\"awayTeamName\":\"Exeter\",\"score\":\"0 - 1\",\"homeCorners\":5,\"awayCorners\":5,\"homeYellowCards\":1,\"awayYellowCards\":1,\"homeRedCards\":0,\"awayRedCards\":0},{\"date\":\"2019-12-21T15:00:00Z\",\"homeTeamName\":\"Exeter\",\"awayTeamName\":\"Walsall\",\"score\":\"3 - 3\",\"homeCorners\":8,\"awayCorners\":8,\"homeYellowCards\":1,\"awayYellowCards\":2,\"homeRedCards\":0,\"awayRedCards\":0},{\"date\":\"2019-12-26T15:00:00Z\",\"homeTeamName\":\"Exeter\",\"awayTeamName\":\"Newport County\",\"score\":\"1 - 0\",\"homeCorners\":9,\"awayCorners\":3,\"homeYellowCards\":4,\"awayYellowCards\":4,\"homeRedCards\":0,\"awayRedCards\":0},{\"date\":\"2019-12-29T15:00:00Z\",\"homeTeamName\":\"Colchester\",\"awayTeamName\":\"Exeter\",\"score\":\"2 - 2\",\"homeCorners\":5,\"awayCorners\":4,\"homeYellowCards\":2,\"awayYellowCards\":1,\"homeRedCards\":1,\"awayRedCards\":0},{\"date\":\"2020-01-01T15:00:00Z\",\"homeTeamName\":\"Forest Green\",\"awayTeamName\":\"Exeter\",\"score\":\"0 - 1\",\"homeCorners\":8,\"awayCorners\":6,\"homeYellowCards\":3,\"awayYellowCards\":2,\"homeRedCards\":0,\"awayRedCards\":1},{\"date\":\"2020-01-11T15:00:00Z\",\"homeTeamName\":\"Exeter\",\"awayTeamName\":\"Cambridge\",\"score\":\"2 - 0\",\"homeCorners\":4,\"awayCorners\":7,\"homeYellowCards\":1,\"awayYellowCards\":0,\"homeRedCards\":0,\"awayRedCards\":0},{\"date\":\"2020-01-18T15:00:00Z\",\"homeTeamName\":\"Grimsby\",\"awayTeamName\":\"Exeter\",\"score\":\"0 - 1\",\"homeCorners\":8,\"awayCorners\":8,\"homeYellowCards\":1,\"awayYellowCards\":1,\"homeRedCards\":0,\"awayRedCards\":0},{\"date\":\"2020-01-25T15:00:00Z\",\"homeTeamName\":\"Exeter\",\"awayTeamName\":\"Colchester\",\"score\":\"0 - 0\",\"homeCorners\":3,\"awayCorners\":5,\"homeYellowCards\":0,\"awayYellowCards\":2,\"homeRedCards\":0,\"awayRedCards\":0},{\"date\":\"2020-01-28T19:45:00Z\",\"homeTeamName\":\"Port Vale\",\"awayTeamName\":\"Exeter\",\"score\":\"3 - 1\",\"homeCorners\":6,\"awayCorners\":7,\"homeYellowCards\":1,\"awayYellowCards\":1,\"homeRedCards\":0,\"awayRedCards\":0},{\"date\":\"2020-02-01T15:00:00Z\",\"homeTeamName\":\"Swindon\",\"awayTeamName\":\"Exeter\",\"score\":\"2 - 1\",\"homeCorners\":8,\"awayCorners\":6,\"homeYellowCards\":1,\"awayYellowCards\":1,\"homeRedCards\":0,\"awayRedCards\":0},{\"date\":\"2020-02-08T15:00:00Z\",\"homeTeamName\":\"Exeter\",\"awayTeamName\":\"Stevenage\",\"score\":\"2 - 1\",\"homeCorners\":11,\"awayCorners\":4,\"homeYellowCards\":0,\"awayYellowCards\":4,\"homeRedCards\":0,\"awayRedCards\":0},{\"date\":\"2020-02-11T19:45:00Z\",\"homeTeamName\":\"Exeter\",\"awayTeamName\":\"Oldham\",\"score\":\"5 - 1\",\"homeCorners\":11,\"awayCorners\":3,\"homeYellowCards\":1,\"awayYellowCards\":1,\"homeRedCards\":0,\"awayRedCards\":1},{\"date\":\"2020-02-15T15:00:00Z\",\"homeTeamName\":\"Macclesfield\",\"awayTeamName\":\"Exeter\",\"score\":\"2 - 3\",\"homeCorners\":3,\"awayCorners\":3,\"homeYellowCards\":1,\"awayYellowCards\":2,\"homeRedCards\":0,\"awayRedCards\":0},{\"date\":\"2020-02-22T15:00:00Z\",\"homeTeamName\":\"Northampton\",\"awayTeamName\":\"Exeter\",\"score\":\"2 - 0\",\"homeCorners\":4,\"awayCorners\":5,\"homeYellowCards\":1,\"awayYellowCards\":1,\"homeRedCards\":0,\"awayRedCards\":0},{\"date\":\"2020-02-29T15:00:00Z\",\"homeTeamName\":\"Exeter\",\"awayTeamName\":\"Crawley Town\",\"score\":\"1 - 1\",\"homeCorners\":6,\"awayCorners\":2,\"homeYellowCards\":1,\"awayYellowCards\":0,\"homeRedCards\":0,\"awayRedCards\":0},{\"date\":\"2020-03-03T19:45:00Z\",\"homeTeamName\":\"Exeter\",\"awayTeamName\":\"Crewe\",\"score\":\"1 - 1\",\"homeCorners\":9,\"awayCorners\":2,\"homeYellowCards\":1,\"awayYellowCards\":1,\"homeRedCards\":0,\"awayRedCards\":0},{\"date\":\"2020-03-07T15:00:00Z\",\"homeTeamName\":\"Walsall\",\"awayTeamName\":\"Exeter\",\"score\":\"3 - 1\",\"homeCorners\":5,\"awayCorners\":4,\"homeYellowCards\":2,\"awayYellowCards\":2,\"homeRedCards\":0,\"awayRedCards\":0},{\"date\":\"2020-09-12T15:00:00Z\",\"homeTeamName\":\"Salford\",\"awayTeamName\":\"Exeter\",\"score\":\"2 - 2\",\"homeCorners\":4,\"awayCorners\":10,\"homeYellowCards\":1,\"awayYellowCards\":0,\"homeRedCards\":0,\"awayRedCards\":0},{\"date\":\"2020-09-19T15:00:00Z\",\"homeTeamName\":\"Exeter\",\"awayTeamName\":\"Port Vale\",\"score\":\"0 - 2\",\"homeCorners\":6,\"awayCorners\":1,\"homeYellowCards\":0,\"awayYellowCards\":0,\"homeRedCards\":0,\"awayRedCards\":0},{\"date\":\"2020-09-26T15:00:00Z\",\"homeTeamName\":\"Mansfield\",\"awayTeamName\":\"Exeter\",\"score\":\"1 - 2\",\"homeCorners\":4,\"awayCorners\":11,\"homeYellowCards\":2,\"awayYellowCards\":1,\"homeRedCards\":0,\"awayRedCards\":0},{\"date\":\"2020-10-03T15:00:00Z\",\"homeTeamName\":\"Exeter\",\"awayTeamName\":\"Cambridge\",\"score\":\"2 - 0\",\"homeCorners\":2,\"awayCorners\":7,\"homeYellowCards\":3,\"awayYellowCards\":2,\"homeRedCards\":0,\"awayRedCards\":0},{\"date\":\"2020-10-10T15:00:00Z\",\"homeTeamName\":\"Southend\",\"awayTeamName\":\"Exeter\",\"score\":\"2 - 2\",\"homeCorners\":4,\"awayCorners\":5,\"homeYellowCards\":0,\"awayYellowCards\":0,\"homeRedCards\":0,\"awayRedCards\":0},{\"date\":\"2020-10-17T15:00:00Z\",\"homeTeamName\":\"Walsall\",\"awayTeamName\":\"Exeter\",\"score\":\"0 - 0\",\"homeCorners\":10,\"awayCorners\":4,\"homeYellowCards\":0,\"awayYellowCards\":1,\"homeRedCards\":0,\"awayRedCards\":0},{\"date\":\"2020-10-20T19:00:00Z\",\"homeTeamName\":\"Exeter\",\"awayTeamName\":\"Crawley Town\",\"score\":\"2 - 1\",\"homeCorners\":4,\"awayCorners\":4,\"homeYellowCards\":0,\"awayYellowCards\":1,\"homeRedCards\":0,\"awayRedCards\":0},{\"date\":\"2020-10-24T13:00:00Z\",\"homeTeamName\":\"Exeter\",\"awayTeamName\":\"Scunthorpe\",\"score\":\"3 - 1\",\"homeCorners\":7,\"awayCorners\":7,\"homeYellowCards\":0,\"awayYellowCards\":1,\"homeRedCards\":0,\"awayRedCards\":0},{\"date\":\"2020-10-27T19:00:00Z\",\"homeTeamName\":\"Leyton Orient\",\"awayTeamName\":\"Exeter\",\"score\":\"1 - 1\",\"homeCorners\":5,\"awayCorners\":4,\"homeYellowCards\":3,\"awayYellowCards\":0,\"homeRedCards\":0,\"awayRedCards\":0},{\"date\":\"2020-10-31T13:00:00Z\",\"homeTeamName\":\"Exeter\",\"awayTeamName\":\"Carlisle\",\"score\":\"1 - 0\",\"homeCorners\":3,\"awayCorners\":6,\"homeYellowCards\":1,\"awayYellowCards\":0,\"homeRedCards\":0,\"awayRedCards\":1},{\"date\":\"2020-11-03T18:30:00Z\",\"homeTeamName\":\"Morecambe\",\"awayTeamName\":\"Exeter\",\"score\":\"2 - 2\",\"homeCorners\":7,\"awayCorners\":4,\"homeYellowCards\":0,\"awayYellowCards\":1,\"homeRedCards\":0,\"awayRedCards\":0},{\"date\":\"2020-11-14T15:00:00Z\",\"homeTeamName\":\"Bradford\",\"awayTeamName\":\"Exeter\",\"score\":\"2 - 2\",\"homeCorners\":6,\"awayCorners\":10,\"homeYellowCards\":0,\"awayYellowCards\":3,\"homeRedCards\":0,\"awayRedCards\":0}],\"over0\":45,\"over1\":35,\"over2\":23,\"over3\":19,\"over4\":5,\"under1\":4,\"under2\":14,\"under3\":26,\"under4\":30,\"homeIndex\":0,\"awayIndex\":0,\"cardIndex\":0},\"awayTeam\":{\"name\":\"Oldham\",\"goalsScored\":11.8,\"goalsConceided\":15.799999999999995,\"homeGoalsScored\":7.200000000000002,\"homeGoalsConceided\":6.800000000000001,\"awayGoalsScored\":4.600000000000001,\"awayGoalsConceided\":9,\"HomeAwayScored\":0,\"HomeAwayConceided\":0,\"AwayHomeScored\":0,\"AwayHomeConceided\":0,\"noEvents\":49,\"homeEvents\":25,\"awayEvents\":24,\"points\":52,\"corners\":527,\"homeCorners\":156,\"homeAwayCorners\":123,\"awayCorners\":114,\"awayHomeCorners\":134,\"yellowCards\":74,\"redCards\":4,\"totalCards\":155,\"fixtures\":[{\"date\":\"2019-08-03T15:00:00Z\",\"homeTeamName\":\"Forest Green\",\"awayTeamName\":\"Oldham\",\"score\":\"1 - 0\",\"homeCorners\":3,\"awayCorners\":0,\"homeYellowCards\":1,\"awayYellowCards\":2,\"homeRedCards\":0,\"awayRedCards\":0},{\"date\":\"2019-08-10T15:00:00Z\",\"homeTeamName\":\"Oldham\",\"awayTeamName\":\"Crewe\",\"score\":\"1 - 2\",\"homeCorners\":11,\"awayCorners\":5,\"homeYellowCards\":2,\"awayYellowCards\":1,\"homeRedCards\":0,\"awayRedCards\":0},{\"date\":\"2019-08-17T15:00:00Z\",\"homeTeamName\":\"Bradford\",\"awayTeamName\":\"Oldham\",\"score\":\"3 - 0\",\"homeCorners\":6,\"awayCorners\":11,\"homeYellowCards\":1,\"awayYellowCards\":1,\"homeRedCards\":0,\"awayRedCards\":0},{\"date\":\"2019-08-20T19:45:00Z\",\"homeTeamName\":\"Oldham\",\"awayTeamName\":\"Exeter\",\"score\":\"0 - 0\",\"homeCorners\":7,\"awayCorners\":10,\"homeYellowCards\":2,\"awayYellowCards\":0,\"homeRedCards\":0,\"awayRedCards\":0},{\"date\":\"2019-08-24T15:00:00Z\",\"homeTeamName\":\"Cambridge\",\"awayTeamName\":\"Oldham\",\"score\":\"1 - 2\",\"homeCorners\":7,\"awayCorners\":2,\"homeYellowCards\":0,\"awayYellowCards\":4,\"homeRedCards\":0,\"awayRedCards\":0},{\"date\":\"2019-08-31T15:00:00Z\",\"homeTeamName\":\"Oldham\",\"awayTeamName\":\"Colchester\",\"score\":\"0 - 1\",\"homeCorners\":4,\"awayCorners\":4,\"homeYellowCards\":2,\"awayYellowCards\":2,\"homeRedCards\":0,\"awayRedCards\":0},{\"date\":\"2019-09-07T15:00:00Z\",\"homeTeamName\":\"Plymouth\",\"awayTeamName\":\"Oldham\",\"score\":\"2 - 2\",\"homeCorners\":2,\"awayCorners\":5,\"homeYellowCards\":4,\"awayYellowCards\":5,\"homeRedCards\":0,\"awayRedCards\":0},{\"date\":\"2019-09-14T15:00:00Z\",\"homeTeamName\":\"Oldham\",\"awayTeamName\":\"Grimsby\",\"score\":\"2 - 2\",\"homeCorners\":5,\"awayCorners\":4,\"homeYellowCards\":1,\"awayYellowCards\":4,\"homeRedCards\":0,\"awayRedCards\":0},{\"date\":\"2019-09-17T19:45:00Z\",\"homeTeamName\":\"Scunthorpe\",\"awayTeamName\":\"Oldham\",\"score\":\"2 - 2\",\"homeCorners\":5,\"awayCorners\":5,\"homeYellowCards\":5,\"awayYellowCards\":5,\"homeRedCards\":0,\"awayRedCards\":0},{\"date\":\"2019-09-21T15:00:00Z\",\"homeTeamName\":\"Oldham\",\"awayTeamName\":\"Morecambe\",\"score\":\"3 - 1\",\"homeCorners\":4,\"awayCorners\":3,\"homeYellowCards\":0,\"awayYellowCards\":0,\"homeRedCards\":0,\"awayRedCards\":0},{\"date\":\"2019-09-28T15:00:00Z\",\"homeTeamName\":\"Carlisle\",\"awayTeamName\":\"Oldham\",\"score\":\"1 - 0\",\"homeCorners\":3,\"awayCorners\":5,\"homeYellowCards\":0,\"awayYellowCards\":1,\"homeRedCards\":0,\"awayRedCards\":0},{\"date\":\"2019-10-05T15:00:00Z\",\"homeTeamName\":\"Oldham\",\"awayTeamName\":\"Cheltenham\",\"score\":\"1 - 1\",\"homeCorners\":4,\"awayCorners\":8,\"homeYellowCards\":3,\"awayYellowCards\":4,\"homeRedCards\":0,\"awayRedCards\":0},{\"date\":\"2019-10-12T15:00:00Z\",\"homeTeamName\":\"Mansfield\",\"awayTeamName\":\"Oldham\",\"score\":\"6 - 1\",\"homeCorners\":6,\"awayCorners\":2,\"homeYellowCards\":2,\"awayYellowCards\":0,\"homeRedCards\":0,\"awayRedCards\":2},{\"date\":\"2019-10-19T15:00:00Z\",\"homeTeamName\":\"Oldham\",\"awayTeamName\":\"Macclesfield\",\"score\":\"0 - 1\",\"homeCorners\":4,\"awayCorners\":3,\"homeYellowCards\":4,\"awayYellowCards\":1,\"homeRedCards\":0,\"awayRedCards\":0},{\"date\":\"2019-10-22T19:45:00Z\",\"homeTeamName\":\"Oldham\",\"awayTeamName\":\"Walsall\",\"score\":\"2 - 0\",\"homeCorners\":6,\"awayCorners\":3,\"homeYellowCards\":1,\"awayYellowCards\":0,\"homeRedCards\":0,\"awayRedCards\":1},{\"date\":\"2019-10-26T15:00:00Z\",\"homeTeamName\":\"Port Vale\",\"awayTeamName\":\"Oldham\",\"score\":\"0 - 0\",\"homeCorners\":14,\"awayCorners\":6,\"homeYellowCards\":2,\"awayYellowCards\":0,\"homeRedCards\":0,\"awayRedCards\":0},{\"date\":\"2019-11-02T15:00:00Z\",\"homeTeamName\":\"Oldham\",\"awayTeamName\":\"Northampton\",\"score\":\"2 - 2\",\"homeCorners\":5,\"awayCorners\":7,\"homeYellowCards\":0,\"awayYellowCards\":2,\"homeRedCards\":0,\"awayRedCards\":0},{\"date\":\"2019-11-23T15:00:00Z\",\"homeTeamName\":\"Newport County\",\"awayTeamName\":\"Oldham\",\"score\":\"0 - 1\",\"homeCorners\":2,\"awayCorners\":4,\"homeYellowCards\":4,\"awayYellowCards\":1,\"homeRedCards\":0,\"awayRedCards\":0},{\"date\":\"2019-12-07T15:00:00Z\",\"homeTeamName\":\"Oldham\",\"awayTeamName\":\"Leyton Orient\",\"score\":\"1 - 1\",\"homeCorners\":4,\"awayCorners\":4,\"homeYellowCards\":3,\"awayYellowCards\":3,\"homeRedCards\":0,\"awayRedCards\":0},{\"date\":\"2019-12-14T15:00:00Z\",\"homeTeamName\":\"Swindon\",\"awayTeamName\":\"Oldham\",\"score\":\"2 - 0\",\"homeCorners\":5,\"awayCorners\":7,\"homeYellowCards\":1,\"awayYellowCards\":2,\"homeRedCards\":0,\"awayRedCards\":0},{\"date\":\"2019-12-21T15:00:00Z\",\"homeTeamName\":\"Oldham\",\"awayTeamName\":\"Crawley Town\",\"score\":\"2 - 1\",\"homeCorners\":9,\"awayCorners\":4,\"homeYellowCards\":0,\"awayYellowCards\":0,\"homeRedCards\":0,\"awayRedCards\":0},{\"date\":\"2019-12-26T15:00:00Z\",\"homeTeamName\":\"Morecambe\",\"awayTeamName\":\"Oldham\",\"score\":\"1 - 2\",\"homeCorners\":4,\"awayCorners\":2,\"homeYellowCards\":0,\"awayYellowCards\":1,\"homeRedCards\":0,\"awayRedCards\":0},{\"date\":\"2019-12-29T15:00:00Z\",\"homeTeamName\":\"Oldham\",\"awayTeamName\":\"Salford\",\"score\":\"1 - 4\",\"homeCorners\":9,\"awayCorners\":1,\"homeYellowCards\":1,\"awayYellowCards\":1,\"homeRedCards\":0,\"awayRedCards\":0},{\"date\":\"2020-01-01T15:00:00Z\",\"homeTeamName\":\"Oldham\",\"awayTeamName\":\"Scunthorpe\",\"score\":\"0 - 2\",\"homeCorners\":8,\"awayCorners\":9,\"homeYellowCards\":0,\"awayYellowCards\":0,\"homeRedCards\":0,\"awayRedCards\":0},{\"date\":\"2020-01-04T15:00:00Z\",\"homeTeamName\":\"Cheltenham\",\"awayTeamName\":\"Oldham\",\"score\":\"3 - 0\",\"homeCorners\":9,\"awayCorners\":3,\"homeYellowCards\":0,\"awayYellowCards\":1,\"homeRedCards\":0,\"awayRedCards\":0},{\"date\":\"2020-01-11T15:00:00Z\",\"homeTeamName\":\"Macclesfield\",\"awayTeamName\":\"Oldham\",\"score\":\"1 - 1\",\"homeCorners\":5,\"awayCorners\":13,\"homeYellowCards\":1,\"awayYellowCards\":1,\"homeRedCards\":0,\"awayRedCards\":1},{\"date\":\"2020-01-14T19:45:00Z\",\"homeTeamName\":\"Stevenage\",\"awayTeamName\":\"Oldham\",\"score\":\"0 - 0\",\"homeCorners\":6,\"awayCorners\":7,\"homeYellowCards\":1,\"awayYellowCards\":3,\"homeRedCards\":0,\"awayRedCards\":0},{\"date\":\"2020-01-18T15:00:00Z\",\"homeTeamName\":\"Oldham\",\"awayTeamName\":\"Carlisle\",\"score\":\"1 - 1\",\"homeCorners\":7,\"awayCorners\":6,\"homeYellowCards\":1,\"awayYellowCards\":1,\"homeRedCards\":0,\"awayRedCards\":0},{\"date\":\"2020-01-25T15:00:00Z\",\"homeTeamName\":\"Salford\",\"awayTeamName\":\"Oldham\",\"score\":\"1 - 1\",\"homeCorners\":3,\"awayCorners\":5,\"homeYellowCards\":1,\"awayYellowCards\":3,\"homeRedCards\":0,\"awayRedCards\":0},{\"date\":\"2020-01-28T19:45:00Z\",\"homeTeamName\":\"Oldham\",\"awayTeamName\":\"Mansfield\",\"score\":\"3 - 1\",\"homeCorners\":7,\"awayCorners\":5,\"homeYellowCards\":0,\"awayYellowCards\":2,\"homeRedCards\":0,\"awayRedCards\":0},{\"date\":\"2020-02-01T15:00:00Z\",\"homeTeamName\":\"Oldham\",\"awayTeamName\":\"Bradford\",\"score\":\"3 - 0\",\"homeCorners\":12,\"awayCorners\":4,\"homeYellowCards\":2,\"awayYellowCards\":3,\"homeRedCards\":0,\"awayRedCards\":0},{\"date\":\"2020-02-08T15:00:00Z\",\"homeTeamName\":\"Crewe\",\"awayTeamName\":\"Oldham\",\"score\":\"2 - 1\",\"homeCorners\":7,\"awayCorners\":1,\"homeYellowCards\":1,\"awayYellowCards\":2,\"homeRedCards\":0,\"awayRedCards\":0},{\"date\":\"2020-02-11T19:45:00Z\",\"homeTeamName\":\"Exeter\",\"awayTeamName\":\"Oldham\",\"score\":\"5 - 1\",\"homeCorners\":11,\"awayCorners\":3,\"homeYellowCards\":1,\"awayYellowCards\":1,\"homeRedCards\":0,\"awayRedCards\":1},{\"date\":\"2020-02-15T15:00:00Z\",\"homeTeamName\":\"Oldham\",\"awayTeamName\":\"Forest Green\",\"score\":\"1 - 1\",\"homeCorners\":7,\"awayCorners\":7,\"homeYellowCards\":1,\"awayYellowCards\":0,\"homeRedCards\":0,\"awayRedCards\":0},{\"date\":\"2020-02-22T15:00:00Z\",\"homeTeamName\":\"Leyton Orient\",\"awayTeamName\":\"Oldham\",\"score\":\"2 - 2\",\"homeCorners\":7,\"awayCorners\":4,\"homeYellowCards\":1,\"awayYellowCards\":3,\"homeRedCards\":0,\"awayRedCards\":0},{\"date\":\"2020-02-29T15:00:00Z\",\"homeTeamName\":\"Oldham\",\"awayTeamName\":\"Newport County\",\"score\":\"5 - 0\",\"homeCorners\":4,\"awayCorners\":4,\"homeYellowCards\":0,\"awayYellowCards\":1,\"homeRedCards\":0,\"awayRedCards\":0},{\"date\":\"2020-03-07T15:00:00Z\",\"homeTeamName\":\"Crawley Town\",\"awayTeamName\":\"Oldham\",\"score\":\"3 - 0\",\"homeCorners\":6,\"awayCorners\":1,\"homeYellowCards\":2,\"awayYellowCards\":0,\"homeRedCards\":0,\"awayRedCards\":0},{\"date\":\"2020-09-12T15:00:00Z\",\"homeTeamName\":\"Oldham\",\"awayTeamName\":\"Leyton Orient\",\"score\":\"0 - 1\",\"homeCorners\":5,\"awayCorners\":3,\"homeYellowCards\":0,\"awayYellowCards\":2,\"homeRedCards\":0,\"awayRedCards\":0},{\"date\":\"2020-09-19T15:00:00Z\",\"homeTeamName\":\"Stevenage\",\"awayTeamName\":\"Oldham\",\"score\":\"3 - 0\",\"homeCorners\":13,\"awayCorners\":3,\"homeYellowCards\":1,\"awayYellowCards\":1,\"homeRedCards\":0,\"awayRedCards\":0},{\"date\":\"2020-09-26T15:00:00Z\",\"homeTeamName\":\"Oldham\",\"awayTeamName\":\"Crawley Town\",\"score\":\"2 - 3\",\"homeCorners\":8,\"awayCorners\":2,\"homeYellowCards\":0,\"awayYellowCards\":2,\"homeRedCards\":0,\"awayRedCards\":0},{\"date\":\"2020-10-03T15:00:00Z\",\"homeTeamName\":\"Colchester\",\"awayTeamName\":\"Oldham\",\"score\":\"3 - 3\",\"homeCorners\":5,\"awayCorners\":8,\"homeYellowCards\":4,\"awayYellowCards\":1,\"homeRedCards\":0,\"awayRedCards\":0},{\"date\":\"2020-10-10T15:00:00Z\",\"homeTeamName\":\"Oldham\",\"awayTeamName\":\"Morecambe\",\"score\":\"2 - 3\",\"homeCorners\":6,\"awayCorners\":2,\"homeYellowCards\":0,\"awayYellowCards\":1,\"homeRedCards\":0,\"awayRedCards\":0},{\"date\":\"2020-10-17T15:00:00Z\",\"homeTeamName\":\"Bolton\",\"awayTeamName\":\"Oldham\",\"score\":\"1 - 2\",\"homeCorners\":2,\"awayCorners\":9,\"homeYellowCards\":1,\"awayYellowCards\":4,\"homeRedCards\":0,\"awayRedCards\":0},{\"date\":\"2020-10-20T19:00:00Z\",\"homeTeamName\":\"Oldham\",\"awayTeamName\":\"Carlisle\",\"score\":\"1 - 1\",\"homeCorners\":3,\"awayCorners\":7,\"homeYellowCards\":1,\"awayYellowCards\":0,\"homeRedCards\":0,\"awayRedCards\":0},{\"date\":\"2020-10-24T15:00:00Z\",\"homeTeamName\":\"Oldham\",\"awayTeamName\":\"Port Vale\",\"score\":\"1 - 2\",\"homeCorners\":7,\"awayCorners\":7,\"homeYellowCards\":3,\"awayYellowCards\":1,\"homeRedCards\":0,\"awayRedCards\":0},{\"date\":\"2020-10-27T19:00:00Z\",\"homeTeamName\":\"Southend\",\"awayTeamName\":\"Oldham\",\"score\":\"1 - 2\",\"homeCorners\":1,\"awayCorners\":5,\"homeYellowCards\":0,\"awayYellowCards\":2,\"homeRedCards\":0,\"awayRedCards\":0},{\"date\":\"2020-10-31T15:00:00Z\",\"homeTeamName\":\"Salford\",\"awayTeamName\":\"Oldham\",\"score\":\"2 - 0\",\"homeCorners\":2,\"awayCorners\":3,\"homeYellowCards\":2,\"awayYellowCards\":1,\"homeRedCards\":0,\"awayRedCards\":0},{\"date\":\"2020-11-03T18:00:00Z\",\"homeTeamName\":\"Oldham\",\"awayTeamName\":\"Cheltenham\",\"score\":\"2 - 1\",\"homeCorners\":4,\"awayCorners\":2,\"homeYellowCards\":1,\"awayYellowCards\":3,\"homeRedCards\":0,\"awayRedCards\":0},{\"date\":\"2020-11-14T15:00:00Z\",\"homeTeamName\":\"Oldham\",\"awayTeamName\":\"Scunthorpe\",\"score\":\"0 - 2\",\"homeCorners\":6,\"awayCorners\":9,\"homeYellowCards\":1,\"awayYellowCards\":1,\"homeRedCards\":0,\"awayRedCards\":0}],\"over0\":46,\"over1\":40,\"over2\":28,\"over3\":14,\"over4\":7,\"under1\":3,\"under2\":9,\"under3\":21,\"under4\":35,\"homeIndex\":0,\"awayIndex\":0,\"cardIndex\":0},\"homeTeamName\":\"Exeter\",\"awayTeamName\":\"Oldham\",\"score\":\"\",\"omit\":\"\",\"homeCorners\":5,\"awayCorners\":5,\"homeYellowCards\":0,\"awayYellowCards\":0,\"homeRedCards\":0,\"awayRedCards\":0,\"homeIndex\":0.8312709030100338,\"awayIndex\":0.41977777777777786,\"totalIndex\":60.04194723151247,\"predicted0\":1,\"predicted1\":0,\"predicted2\":0,\"predicted3\":0,\"predicted4\":0,\"homeOdd\":0,\"drawOdd\":0,\"awayOdd\":0,\"overOdd\":0,\"underOdd\":0,\"over0\":0,\"over1\":0,\"over2\":0,\"over3\":0,\"over4\":0,\"cornerClass\":\"cornerUnder cornerWrong\",\"cardClass\":\"cardUnder cardWrong\"}");
        const eventDate = new Date(fixture.date);
        const homeTeamName = fixture.homeTeam.name;
        const awayTeamName = fixture.awayTeam.name;
        const size = 60;
        const home = fixture.homeTeam;
        const away = fixture.awayTeam;
        const homeFixtures = home.fixtures.reverse().slice(0, size);
        const awayFixtures = away.fixtures.reverse().slice(0, size);

        const homeCalcs = this.getCalculations(homeFixtures, homeTeamName);
        const awayCalcs = this.getCalculations(awayFixtures, awayTeamName);
        return <React.Fragment>
            <table className="eventDetails">
                <thead>
                <tr>
                    <th colSpan="3">{homeTeamName}</th>
                    <th>{eventDate.toDateString()}></th>
                    <th colSpan="3">{awayTeamName}</th>
                </tr>
                <tr>
                    <th>Home</th>
                    <th>Away</th>
                    <th>Total</th>
                    <th>-</th>
                    <th>Home</th>
                    <th>Away</th>
                    <th>Total</th>
                </tr>
                </thead>
                <tbody>
                <tr>
                    <td>{home.homeEvents}</td>
                    <td>{home.awayEvents}</td>
                    <td>{home.noEvents}</td>
                    <td>Events</td>
                    <td>{away.homeEvents}</td>
                    <td>{away.awayEvents}</td>
                    <td>{away.noEvents}</td>
                </tr>
                <tr>
                    <td>{homeCalcs.homeOver0}/{homeCalcs.homeCount}</td>
                    <td>{homeCalcs.awayOver0}/{homeCalcs.awayCount}</td>
                    <td>{homeCalcs.over0}/{homeCalcs.count}</td>
                    <td>Over 0</td>
                    <td>{awayCalcs.homeOver0}/{awayCalcs.homeCount}</td>
                    <td>{awayCalcs.awayOver0}/{awayCalcs.awayCount}</td>
                    <td>{awayCalcs.over0}/{awayCalcs.count}</td>
                </tr>
                <tr>
                    <td>{homeCalcs.homeOver1}/{homeCalcs.homeCount}</td>
                    <td>{homeCalcs.awayOver1}/{homeCalcs.awayCount}</td>
                    <td>{homeCalcs.over1}/{homeCalcs.count}</td>
                    <td>Over 1</td>
                    <td>{awayCalcs.homeOver1}/{awayCalcs.homeCount}</td>
                    <td>{awayCalcs.awayOver1}/{awayCalcs.awayCount}</td>
                    <td>{awayCalcs.over1}/{awayCalcs.count}</td>
                </tr>
                <tr>
                    <td>{homeCalcs.homeOver2}/{homeCalcs.homeCount}</td>
                    <td>{homeCalcs.awayOver2}/{homeCalcs.awayCount}</td>
                    <td>{homeCalcs.over2}/{homeCalcs.count}</td>
                    <td>Over 2</td>
                    <td>{awayCalcs.homeOver2}/{awayCalcs.homeCount}</td>
                    <td>{awayCalcs.awayOver2}/{awayCalcs.awayCount}</td>
                    <td>{awayCalcs.over2}/{awayCalcs.count}</td>
                </tr>
                <tr>
                    <td>{homeCalcs.homeOver3}/{homeCalcs.homeCount}</td>
                    <td>{homeCalcs.awayOver3}/{homeCalcs.awayCount}</td>
                    <td>{homeCalcs.over3}/{homeCalcs.count}</td>
                    <td>Over 3</td>
                    <td>{awayCalcs.homeOver3}/{awayCalcs.homeCount}</td>
                    <td>{awayCalcs.awayOver3}/{awayCalcs.awayCount}</td>
                    <td>{awayCalcs.over3}/{awayCalcs.count}</td>
                </tr>
                <tr>
                    <td>{homeCalcs.homeOver4}/{homeCalcs.homeCount}</td>
                    <td>{homeCalcs.awayOver4}/{homeCalcs.awayCount}</td>
                    <td>{homeCalcs.over4}/{homeCalcs.count}</td>
                    <td>Over 4</td>
                    <td>{awayCalcs.homeOver4}/{awayCalcs.homeCount}</td>
                    <td>{awayCalcs.awayOver4}/{awayCalcs.awayCount}</td>
                    <td>{awayCalcs.over4}/{awayCalcs.count}</td>
                </tr>
                <tr>
                    <td>{homeCalcs.homeGoals}/{homeCalcs.homeCount}</td>
                    <td>{homeCalcs.awayGoals}/{homeCalcs.awayCount}</td>
                    <td></td>
                    <td>Total Goals</td>
                    <td>{awayCalcs.homeGoals}/{awayCalcs.homeCount}</td>
                    <td>{awayCalcs.awayGoals}/{awayCalcs.awayCount}</td>
                    <td></td>
                </tr>
                <tr>
                    <td>{homeCalcs.homeScored}/{homeCalcs.homeCount}</td>
                    <td>{homeCalcs.awayScored}/{homeCalcs.awayCount}</td>
                    <td></td>
                    <td>Goals Scored</td>
                    <td>{awayCalcs.homeScored}/{awayCalcs.homeCount}</td>
                    <td>{awayCalcs.awayScored}/{awayCalcs.awayCount}</td>
                    <td></td>
                </tr>
                <tr>
                    <td>{homeCalcs.homeConceided}/{homeCalcs.homeCount}</td>
                    <td>{homeCalcs.awayConceided}/{homeCalcs.awayCount}</td>
                    <td></td>
                    <td>Goals Conceided</td>
                    <td>{awayCalcs.homeConceided}/{awayCalcs.homeCount}</td>
                    <td>{awayCalcs.awayConceided}/{awayCalcs.awayCount}</td>
                    <td></td>
                </tr>
                <tr>
                    <td>{home.homeCorners}</td>
                    <td>{home.awayHomeCorners}</td>
                    <td></td>
                    <td>Corners</td>
                    <td>{away.homeCorners}</td>
                    <td>{away.awayHomeCorners}</td>
                    <td></td>
                </tr>
                <tr>
                    <td>{home.yellowCards}</td>
                    <td>{home.redCards}</td>
                    <td>{home.yellowCards}</td>
                    <td>Cards</td>
                    <td>{away.yellowCards}</td>
                    <td>{away.redCards}</td>
                    <td>{away.yellowCards}</td>
                </tr>
                </tbody>
            </table>
            <table className="eventDetails">
                <thead>
                    <tr>
                        <th colSpan="3">Home Fixtures</th>
                        <th colSpan="3">Away Fixtures</th>
                    </tr>
                </thead>
                <tbody>
                    <tr>
                        <td colSpan="3">{homeCalcs.fixtureListHome}</td>
                        <td colSpan="3">{awayCalcs.fixtureListAway}</td>
                    </tr>
                </tbody>
            </table>
        </React.Fragment>

    }

    getCalculations = (fixtures, team) => {
        let over0 = 0, over1 = 0, over2 = 0, over3 = 0, over4 = 0;
        let under1 = 0, under2 = 0, under3 = 0, under4 = 0;
        let homeOver0 = 0, homeOver1 = 0, homeOver2 = 0, homeOver3 = 0, homeOver4 = 0;
        let homeUnder1 = 0, homeUnder2 = 0, homeUnder3 = 0, homeUnder4 = 0;
        let awayOver0 = 0, awayOver1 = 0, awayOver2 = 0, awayOver3 = 0, awayOver4 = 0;
        let awayUnder1 = 0, awayUnder2 = 0, awayUnder3 = 0, awayUnder4 = 0;
        let homeScored = 0, homeConceided = 0, awayScored = 0, awayConceided = 0;
        let homeCount = 0, awayCount = 0, count = 0;
        let list = [],homeList = [],awayList = [], fixtureList, fixtureListHome, fixtureListAway;
        for (let i=0; i< fixtures.length; i++) {
            const fixture = fixtures[i];
            if (fixture.score !== "") {
                count++;
                const isHome = (team === fixture.homeTeamName);
                const isAway = (team === fixture.awayTeamName);
                const scores = fixture.score.split(" - ");
                const score = parseInt(scores[0]) + parseInt(scores[1]);
                if (score > 0) {
                    over0++
                }
                if (score > 1) {
                    over1++
                }
                if (score > 2) {
                    over2++
                }
                if (score > 3) {
                    over3++
                }
                if (score > 4) {
                    over4++
                }
                if (score < 1) {
                    under1++
                }
                if (score < 2) {
                    under2++
                }
                if (score < 3) {
                    under3++
                }
                if (score < 4) {
                    under4++
                }
                if (isHome) {
                    homeCount++;
                    homeScored += parseInt(scores[0]);
                    homeConceided += parseInt(scores[1]);
                    if (score > 0) {
                        homeOver0++
                    }
                    if (score > 1) {
                        homeOver1++
                    }
                    if (score > 2) {
                        homeOver2++
                    }
                    if (score > 3) {
                        homeOver3++
                    }
                    if (score > 4) {
                        homeOver4++
                    }
                    if (score < 1) {
                        homeUnder1++
                    }
                    if (score < 2) {
                        homeUnder2++
                    }
                    if (score < 3) {
                        homeUnder3++
                    }
                    if (score < 4) {
                        homeUnder4++
                    }
                } 
                if (isAway) {
                    awayCount++;
                    awayScored += parseInt(scores[1]);
                    awayConceided += parseInt(scores[0]);
                    if (score > 0) {
                        awayOver0++
                    }
                    if (score > 1) {
                        awayOver1++
                    }
                    if (score > 2) {
                        awayOver2++
                    }
                    if (score > 3) {
                        awayOver3++
                    }
                    if (score > 4) {
                        awayOver4++
                    }
                    if (score < 1) {
                        awayUnder1++
                    }
                    if (score < 2) {
                        awayUnder2++
                    }
                    if (score < 3) {
                        awayUnder3++
                    }
                    if (score < 4) {
                        awayUnder4++
                    }
                }
                let homePart, awayPart;
                if (isHome) {
                    homePart = `${fixture.homeTeamName} ${scores[0]}`;
                } else {
                    homePart = `${fixture.homeTeamName} ${scores[0]}`;
                }
                if (isAway) {
                    awayPart = `${scores[1]} ${fixture.awayTeamName} `;
                } else {
                    awayPart = `${scores[1]} ${fixture.awayTeamName}`;
                }
                if (isHome) {
                    homeList.push(`${homePart}:${awayPart}`);
                }
                if (isAway) {
                    awayList.push(`${homePart}:${awayPart}`);
                }
                list.push(`${homePart}:${awayPart}`);
            }
            fixtureListHome = homeList.map((item) =>
                <li>{item}</li>
            );
            fixtureListAway = awayList.map((item) =>
                <li>{item}</li>
            );
            fixtureList = list.map((item) =>
                <li>{item}</li>
            );
        }
        return {
            over0, over1, over2, over3, over4,
            under1, under2, under3, under4,
            homeOver0, homeOver1, homeOver2, homeOver3, homeOver4,
            homeUnder1, homeUnder2, homeUnder3, homeUnder4,
            awayOver0, awayOver1, awayOver2, awayOver3, awayOver4,
            awayUnder1, awayUnder2, awayUnder3, awayUnder4,
            homeScored, homeConceided, awayScored, awayConceided,
            homeCount, awayCount, count,
            fixtureListHome, fixtureListAway, fixtureList
        };
    }
}

export default EventDetails;