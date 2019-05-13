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
    return {
        id : fixture.id,
        played : fixture.played,
        dateTime : {
            seconds: fixture.dateTime.getTime(),
            nanos: (fixture.dateTime.getTime() % 1000) * 1e6,
        },
        matchday : fixture.matchday,
        hTeam : fixture.hTeam,
        aTeam : fixture.aTeam,
        hGoals : fixture.hGoals,
        aGoals : fixture.aGoals,
    }
};