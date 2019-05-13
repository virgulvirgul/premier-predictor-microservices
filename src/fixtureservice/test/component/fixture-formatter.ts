import {FixtureFormatter} from "../../lib/component/fixture-formatter";
import {Fixture} from "../../lib/model/fixture";

let chai = require("chai");
chai.should();

let fixtureFormatter = new FixtureFormatter();

const TEAM_1 = "Team 1";
const TEAM_2 = "Team 2";
const TEAM_3 = "Team 3";

const DATE_1 = new Date('2014-04-03');
const DATE_2 = new Date('2014-04-02');
const DATE_3 = new Date('2014-04-01');

describe("FixtureFormatter", () => {
    describe("groupIntoTeams", () => {
        it("should take a list of matches and groups them into teams and orders matches by date", (done) => {
            const m1: Fixture = {
                id: "1",
                hTeam: TEAM_1,
                aTeam: TEAM_2,
                dateTime: DATE_1,
                aGoals: 0, hGoals: 0, matchday: 0, played: 0
            };
            const m2: Fixture = {
                id: "2",
                hTeam: TEAM_1,
                aTeam: TEAM_3,
                dateTime: DATE_2,
                aGoals: 0, hGoals: 0, matchday: 0, played: 0
            };
            const m3: Fixture = {
                id: "3",
                hTeam: TEAM_2,
                aTeam: TEAM_1,
                dateTime: DATE_3,
                aGoals: 0, hGoals: 0, matchday: 0, played: 0
            };

            const expectedResult = new Map<string, Fixture[]>();
            expectedResult.set(TEAM_1, [m3, m2, m1]);
            expectedResult.set(TEAM_2, [m3, m1]);
            expectedResult.set(TEAM_3, [m2]);

            fixtureFormatter.groupIntoTeams([m1, m2, m3]).then((teamFixtures: Map<string, Fixture[]>) => {
                teamFixtures.should.eql(expectedResult);
                done();
            }, e => {
                done(new Error(e));
            });
        });
    });
});