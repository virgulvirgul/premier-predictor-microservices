import {Repository} from "../repository/repository";
import {Fixture} from "../model/fixture";
import {Forms, TeamForm} from "../model/form";
import {FixtureFormatter} from "../component/fixture-formatter";
import {FormFormatter} from "../component/form-formatter";

export class Service {
    constructor(private repository: Repository,
                private fixtureFormatter: FixtureFormatter,
                private formFormatter: FormFormatter) {
    }

    public getAllFixtures(): Promise<Fixture[]> {
        return new Promise<Fixture[]>((resolve, reject) => {
            this.repository.getAllFixtures().then((fixtures: Fixture[]) => {
                resolve(fixtures);
            }, err => {
                reject(err);
            });
        });
    }

    public getFixtureById(id: string): Promise<Fixture> {
        return new Promise<Fixture>((resolve, reject) => {
            this.repository.getFixtureById(id).then((fixture: Fixture) => {
                resolve(fixture);
            }, err => {
                reject(err);
            });
        });
    }

    public getTeamForms(): Promise<Forms> {
        return new Promise<Forms>((resolve, reject) => {
            this.repository.getAllPastFixtures().then((fixtures: Fixture[]) => {
                this.fixtureFormatter.groupIntoTeams(fixtures).then((teamFixtures: Map<string, Fixture[]>) => {
                    this.formFormatter.formatLastFiveGames(teamFixtures).then((teamForms: Map<string, TeamForm>) => {
                        resolve({teams: teamForms});
                    }, err => {
                        reject(err);
                    });
                }, err => {
                    reject(err);
                });
            }, err => {
                reject(err);
            });
        });
    }
}