import {Fixture} from "../../lib/model/fixture";
import {FormFormatter} from "../../lib/component/form-formatter";
import {Location, Result, TeamForm} from "../../lib/model/form";

const {HOME, AWAY} = Location;
const {WIN, DRAW, LOSS} = Result;

let chai = require("chai");
chai.should();

let formFormatter = new FormFormatter();

const TEAM_1 = "Team 1";
const TEAM_2 = "Team 2";
const TEAM_3 = "Team 3";

const m1: Fixture = {id: "1", hTeam: TEAM_1, aTeam: TEAM_2, dateTime: new Date('2014-04-30'), hGoals: 1, aGoals: 0, matchday: 0, played: 1};
const m2: Fixture = {id: "2", hTeam: TEAM_1, aTeam: TEAM_2, dateTime: new Date('2014-04-29'), hGoals: 2, aGoals: 1, matchday: 0, played: 1};
const m3: Fixture = {id: "3", hTeam: TEAM_1, aTeam: TEAM_2, dateTime: new Date('2014-04-28'), hGoals: 0, aGoals: 4, matchday: 0, played: 1};
const m4: Fixture = {id: "4", hTeam: TEAM_1, aTeam: TEAM_2, dateTime: new Date('2014-04-27'), hGoals: 3, aGoals: 3, matchday: 0, played: 1};
const m5: Fixture = {id: "5", hTeam: TEAM_1, aTeam: TEAM_2, dateTime: new Date('2014-04-26'), hGoals: 1, aGoals: 2, matchday: 0, played: 1};
const m6: Fixture = {id: "6", hTeam: TEAM_1, aTeam: TEAM_2, dateTime: new Date('2014-04-25'), hGoals: 3, aGoals: 1, matchday: 0, played: 1};
const m7: Fixture = {id: "7", hTeam: TEAM_1, aTeam: TEAM_3, dateTime: new Date('2014-04-24'), hGoals: 2, aGoals: 0, matchday: 0, played: 1};
const m8: Fixture = {id: "8", hTeam: TEAM_1, aTeam: TEAM_3, dateTime: new Date('2014-04-23'), hGoals: 1, aGoals: 1, matchday: 0, played: 1};
const m9: Fixture = {id: "9", hTeam: TEAM_1, aTeam: TEAM_3, dateTime: new Date('2014-04-22'), hGoals: 0, aGoals: 1, matchday: 0, played: 1};
const m10: Fixture = {id: "10", hTeam: TEAM_1, aTeam: TEAM_3, dateTime: new Date('2014-04-21'), hGoals: 0, aGoals: 1, matchday: 0, played: 1};
const m11: Fixture = {id: "11", hTeam: TEAM_1, aTeam: TEAM_3, dateTime: new Date('2014-04-20'), hGoals: 2, aGoals: 2, matchday: 0, played: 1};
const m12: Fixture = {id: "12", hTeam: TEAM_1, aTeam: TEAM_3, dateTime: new Date('2014-04-19'), hGoals: 1, aGoals: 3, matchday: 0, played: 1};

const input: Map<string, Fixture[]> = new Map<string, Fixture[]>([
   [TEAM_1, [m1, m2, m3, m4, m5, m6, m7, m8, m9, m10, m11, m12]],
   [TEAM_2, [m1, m2, m3, m4, m5, m6]],
   [TEAM_3, [m7, m8, m9, m10, m11, m12]]
]);

const expectedOutput: Map<string, TeamForm> = new Map<string, TeamForm>([
    [
        TEAM_1,
        {
            forms:[
                {result: DRAW, score: "1-1", opponent: TEAM_3, location: HOME},
                {result: LOSS, score: "0-1", opponent: TEAM_3, location: HOME},
                {result: LOSS, score: "0-1", opponent: TEAM_3, location: HOME},
                {result: DRAW, score: "2-2", opponent: TEAM_3, location: HOME},
                {result: LOSS, score: "1-3", opponent: TEAM_3, location: HOME},
            ]
        }
    ],
    [
        TEAM_2,
        {
            forms:[
                {result: LOSS, score: "2-1", opponent: TEAM_1, location: AWAY},
                {result: WIN, score: "0-4", opponent: TEAM_1, location: AWAY},
                {result: DRAW, score: "3-3", opponent: TEAM_1, location: AWAY},
                {result: WIN, score: "1-2", opponent: TEAM_1, location: AWAY},
                {result: LOSS, score: "3-1", opponent: TEAM_1, location: AWAY},
            ]
        }
    ],
    [
        TEAM_3,
        {
            forms:[
                {result: DRAW, score: "1-1", opponent: TEAM_1, location: AWAY},
                {result: WIN, score: "0-1", opponent: TEAM_1, location: AWAY},
                {result: WIN, score: "0-1", opponent: TEAM_1, location: AWAY},
                {result: DRAW, score: "2-2", opponent: TEAM_1, location: AWAY},
                {result: WIN, score: "1-3", opponent: TEAM_1, location: AWAY},
            ]
        }
    ],
]);

describe("FormFormatter", () => {
    describe("formatLastFiveGames", () => {
        it("should take only the last 5 matches for each team and return in correct format", (done) => {
            formFormatter.formatLastFiveGames(input).then((teamForms: Map<string, TeamForm>) => {
                teamForms.should.eql(expectedOutput);
                done();
            }, e => {
                done(new Error(e));
            });
        });
    });
});