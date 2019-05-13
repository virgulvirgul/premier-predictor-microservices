import {Service} from "../service/service";
import {Fixture, fixtureToGrpc} from "../model/fixture";
import {Forms} from "../model/form";

export class Handler {
    constructor(private service: Service) {
    }

    public getMatches(call, callback) {
        this.service.getAllFixtures().then((fixtures: Fixture[]) => {
            const matches = fixtures.map(f => fixtureToGrpc(f));
            callback(null, {matches: matches});
        }, err => {
            callback(err);
        });
    }

    public getTeamForms(call, callback) {
        this.service.getTeamForms().then((forms: Forms) => {
            callback(null, forms);
        }, err => {
            callback(err);
        });
    }
}