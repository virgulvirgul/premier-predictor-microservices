import {Fixture} from "../model/fixture";

export class FixtureEntity {
    _id: string;
    played: number;
    dateTime: Date;
    matchday: number;
    hTeam: string;
    aTeam: string;
    hGoals: number;
    aGoals: number;

    constructor(id: string, played: number, dateTime: Date, matchday: number, hTeam: string, aTeam: string, hGoals: number, aGoals: number) {
        this._id = id;
        this.played = played;
        this.dateTime = dateTime;
        this.matchday = matchday;
        this.hTeam = hTeam;
        this.aTeam = aTeam;
        this.hGoals = hGoals;
        this.aGoals = aGoals;
    }

    public toFixture(): Fixture {
        return {
            id: this._id,
            played: this.played,
            dateTime: this.dateTime,
            matchday: this.matchday,
            hTeam: this.hTeam,
            aTeam: this.aTeam,
            hGoals: this.hGoals,
            aGoals: this.aGoals,
        };
    }

    static fromFixture(fixture: Fixture): FixtureEntity {
        return new FixtureEntity(
            fixture.id,
            fixture.played,
            fixture.dateTime,
            fixture.matchday,
            fixture.hTeam,
            fixture.aTeam,
            fixture.hGoals,
            fixture.aGoals,
        );
    }
}