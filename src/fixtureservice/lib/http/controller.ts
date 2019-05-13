import * as mongoose from 'mongoose';
import {FixtureSchema} from '../schema/fixture';
import {Request, Response} from 'express';
import {Service} from "../service/service";
import {Forms} from "../model/form";

const Contact = mongoose.model('Contact', FixtureSchema);

export class Controller {
    constructor(private service: Service) {
    }

    public getAllFixtures(req: Request, res: Response) {
        this.service.getAllFixtures().then(fixtures => {
            if (fixtures.length == 0) {
                res.status(404)
                    .send();
                return;
            }

            res.json(fixtures);
        }, err => {
            res.status(500)
                .send(err);
        });
    }

    public getFixtureById(req: Request, res: Response) {
        this.service.getFixtureById(req.params.id).then(fixture => {
            if (!fixture) {
                res.status(404)
                    .send();
                return;
            }

            res.json(fixture);
        }, err => {
            res.status(500)
                .send(err);
        });
    }

    public getTeamForms(req: Request, res: Response) {
        this.service.getTeamForms().then((forms: Forms) => {
            res.json(forms.teams);
        }, err => {
            res.status(500)
                .send(err);
        });
    }
}