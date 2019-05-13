import * as mongoose from 'mongoose';
import {ERR_FIXTURES_NOT_FOUND, Repository} from "../../lib/repository/repository";
import {FixtureSchema} from "../../lib/schema/fixture";
import {Fixture} from "../../lib/model/fixture";

let chai = require("chai");
chai.should();

process.env.MONGO_SCHEME = 'mongodb';
process.env.MONGO_HOST = 'localhost';
process.env.MONGO_USERNAME = '';
process.env.MONGO_PASSWORD = '';
process.env.MONGO_PORT = '27017';

let repository = new Repository();

const FixtureModel = mongoose.model('fixture', FixtureSchema);

const HOME_GOALS = 2;
const AWAY_GOALS = 3;
const PLAYED = 1;
const HOME_TEAM = "hTeam";
const AWAY_TEAM = "aTeam";
const FUTURE_DATE_TIME = new Date('2024-04-03');
const PAST_DATE_TIME = new Date('2014-04-03');
const MATCHDAY = 5;

const addNewFixture = function (id: string): Promise<Fixture> {
    let fixture = new FixtureModel({
        _id: id,
        hTeam: HOME_TEAM,
        aTeam: AWAY_TEAM,
        dateTime: FUTURE_DATE_TIME,
        matchday: MATCHDAY,
    });

    return addFixture(fixture, id)
};

const addPastFixture = function (id: string): Promise<Fixture> {
    let fixture = new FixtureModel({
        _id: id,
        played: PLAYED,
        hTeam: HOME_TEAM,
        aTeam: AWAY_TEAM,
        hGoals: HOME_GOALS,
        aGoals: AWAY_GOALS,
        dateTime: PAST_DATE_TIME,
        matchday: MATCHDAY,
    });

    return addFixture(fixture, id)
};

const addFixture = function (fixture, id: string): Promise<Fixture> {
    return new Promise((resolve, reject) => {
        fixture.save((err) => {
            if (err) {
                reject(err)
            } else {
                resolve();
            }
        });
    });
};

const cleanupDb = function () {
    FixtureModel.deleteMany({}, (e, _) => {
        if (e) {
            console.log(e);
        }
    });
};

describe("Repository", () => {
    describe("getFixtureById", () => {
        const ID = "1";

        before(function (done) {
            addNewFixture(ID).then(() => {
                done();
            }, e => {
                cleanupDb();
                throw new Error(e);
            });
        });

        after(function (done) {
            cleanupDb();
            done();
        });

        it("should get fixture from database with id", (done) => {
            repository.getFixtureById(ID).then(f => {
                f.id.should.equal(ID);
                f.hTeam.should.equal(HOME_TEAM);
                f.aTeam.should.equal(AWAY_TEAM);
                f.dateTime.toISOString().should.equal(FUTURE_DATE_TIME.toISOString());
                f.matchday.should.equal(MATCHDAY);

                console.log(f);
                done();
            }, e => {
                done(new Error(e));
            });
        });

        it("should return error if no fixtures are found", (done) => {
            repository.getFixtureById("not real id").then(f => {
                done(new Error('was not supposed to succeed'));
            }, e => {
                e.should.equal(ERR_FIXTURES_NOT_FOUND);
                done();
            });
        });
    });

    describe("getAllFixtures", () => {
        const ID = "1";
        const ID2 = "2";

        before(function (done) {
            addNewFixture(ID).then(() => {
                addNewFixture(ID2).then(() => {
                    done();
                }, e => {
                    cleanupDb();
                    throw new Error(e);
                });
            }, e => {
                cleanupDb();
                done(new Error(e));
            });
        });

        afterEach(function (done) {
            cleanupDb();
            done();
        });

        it("should get all fixtures from database", (done) => {
            repository.getAllFixtures().then(fixtures => {
                fixtures[0].id.should.equal(ID);
                fixtures[0].hTeam.should.equal(HOME_TEAM);
                fixtures[0].aTeam.should.equal(AWAY_TEAM);
                fixtures[0].dateTime.toISOString().should.equal(FUTURE_DATE_TIME.toISOString());
                fixtures[0].matchday.should.equal(MATCHDAY);
                fixtures[1].id.should.equal(ID2);
                fixtures[1].hTeam.should.equal(HOME_TEAM);
                fixtures[1].aTeam.should.equal(AWAY_TEAM);
                fixtures[1].dateTime.toISOString().should.equal(FUTURE_DATE_TIME.toISOString());
                fixtures[1].matchday.should.equal(MATCHDAY);

                console.log(fixtures);
                done();
            }, e => {
                done(new Error(e));
            });
        });

        it("should return empty array if no fixtures are found", async () => {
            await cleanupDb();
            await repository.getAllFixtures().then(fixtures => {
                fixtures.length.should.equal(0);
            }, e => {
                throw new Error('was not supposed to fail');
            });
        });
    });

    describe("getAllPastFixtures", () => {
        const ID = "1";
        const ID2 = "2";

        before(function (done) {
            addNewFixture(ID).then(() => {
                addPastFixture(ID2).then(() => {
                    done();
                }, e => {
                    cleanupDb();
                    done();
                });
            }, e => {
                cleanupDb();
                done(new Error(e));
            });
        });

        it("should get all fixtures from database that are in past", (done) => {
            repository.getAllPastFixtures().then(fixtures => {
                fixtures[0].id.should.equal(ID2);
                fixtures[0].hTeam.should.equal(HOME_TEAM);
                fixtures[0].aTeam.should.equal(AWAY_TEAM);
                fixtures[0].dateTime.toISOString().should.equal(PAST_DATE_TIME.toISOString());
                fixtures[0].matchday.should.equal(MATCHDAY);
                fixtures[0].played.should.equal(PLAYED);
                fixtures[0].hGoals.should.equal(HOME_GOALS);
                fixtures[0].aGoals.should.equal(AWAY_GOALS);

                console.log(fixtures);
                done();
            }, e => {
                done(new Error(e));
            });
        });

        afterEach(function (done) {
            cleanupDb();
            done();
        });
    });
});