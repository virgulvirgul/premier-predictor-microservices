import * as express from "express";
import * as bodyParser from "body-parser";
import {Router} from "./http/router";
import * as fs from "fs";
import * as https from "https";

const PORT = process.env.HTTP_PORT;

export class Http {
    private app: express.Application = express();

    constructor(private router: Router) {
    }

    public start() {
        this.configure();
        this.router.route(this.app);

        const httpsOptions = {
            key: fs.readFileSync('./certs/key.pem'),
            cert: fs.readFileSync('./certs/cert.pem')
        };

        https.createServer(httpsOptions, this.app).listen(PORT, () => {
            console.log('Express server listening on port ' + PORT);
        });
    }

    private configure(): void {
        this.app.use(bodyParser.json());
        this.app.use(bodyParser.urlencoded({extended: false}));
    }
}
