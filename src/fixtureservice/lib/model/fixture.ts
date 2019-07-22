import * as t from 'google-protobuf/google/protobuf/timestamp_pb.js';

export class Fixture {
    id : string;
    played : number;
    dateTime : Date;
    matchday : number;
    hTeam : string;
    aTeam : string;
    hGoals : number;
    aGoals : number;
}

export const fixtureToGrpc = (fixture: Fixture) => {
    let dateTime = t.Timestamp.fromDate(fixture.dateTime);

    return {
        id : fixture.id,
        played : fixture.played,
        dateTime : {
            seconds: dateTime.getSeconds(),
            nanos: dateTime.getNanos()
        },
        matchday : fixture.matchday,
        hTeam : fixture.hTeam,
        aTeam : fixture.aTeam,
        hGoals : fixture.hGoals,
        aGoals : fixture.aGoals,
    }
};