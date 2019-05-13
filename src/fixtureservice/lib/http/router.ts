import {Request, Response, NextFunction} from "express";
import {Controller} from "./controller";
import {Middleware} from "../middleware/middleware";

export class Router {

    constructor(private controller: Controller,
                private middleware: Middleware) {
    }

    public route(app): void {
        app.get('/health', (req, res) => res.send('Service is healthy!'));

        app.get('/*', (req: Request, res: Response, next: NextFunction) =>
            this.middleware.validateHttp(req, res, next));

        app.route('/')
            .get((req: Request, res: Response, next: NextFunction) =>
                this.controller.getAllFixtures(req, res));

        app.route('/:id')
            .get((req: Request, res: Response, next: NextFunction) =>
                this.controller.getFixtureById(req, res));

        app.route('/team/form')
            .get((req: Request, res: Response, next: NextFunction) =>
                this.controller.getTeamForms(req, res));
    }
}