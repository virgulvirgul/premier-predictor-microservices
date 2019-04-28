import * as express from "express";
import * as bodyParser from "body-parser";
import {Router} from "./router/router";
import * as mongoose from "mongoose";

class App {

    app: express.Application = express();
    private router: Router = new Router();

    constructor() {
        this.configure();
        this.router.route(this.app);
    }

    private configure(): void {
        this.app.use(bodyParser.json());
        this.app.use(bodyParser.urlencoded({extended: false}));
        // // serving static files
        // this.app.use(express.static('public'));
    }

}

export default new App().app;
