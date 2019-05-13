import {Service} from "../../lib/service/service";
import {Repository} from "../../lib/repository/repository";
import {instance, mock, when} from "ts-mockito";
import {Fixture} from "../../lib/model/fixture";
import {FixtureFormatter} from "../../lib/component/fixture-formatter";
import {FormFormatter} from "../../lib/component/form-formatter";
import {Forms, TeamForm} from "../../lib/model/form";

let chai = require("chai");
chai.should();

let mockRepository: Repository = mock(Repository);
let repository: Repository = instance(mockRepository);

let mockFixtureFormatter: FixtureFormatter = mock(FixtureFormatter);
let fixtureFormatter: FixtureFormatter = instance(mockFixtureFormatter);

let mockFormFormatter: FormFormatter = mock(FormFormatter);
let formFormatter: FormFormatter = instance(mockFormFormatter);

let service = new Service(repository, fixtureFormatter, formFormatter);

describe("Service", () => {
    describe("getTeamForms", () => {
        it("should get past fixtures from database and map into forms object", (done) => {
            const fixtures: Fixture[] = [];
            when(mockRepository.getAllPastFixtures()).thenResolve(fixtures);

            const teamFixtures: Map<string, Fixture[]> = new Map<string, Fixture[]>();
            when(mockFixtureFormatter.groupIntoTeams(fixtures)).thenResolve(teamFixtures);

            const teamForms: Map<string, TeamForm> = new Map<string, TeamForm>();
            when(mockFormFormatter.formatLastFiveGames(teamFixtures)).thenResolve(teamForms);

            service.getTeamForms().then( (forms: Forms) => {
                forms.teams.should.equal(teamForms);
                done();
            }, err => {
                done(new Error(err));
            });
        });
    });
});