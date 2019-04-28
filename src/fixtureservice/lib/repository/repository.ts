import {FixtureSchema} from "../schema/fixture";
import * as mongoose from 'mongoose';
import {Fixture} from "../model/fixture";

const DATABASE = 'fixture';
const COLLECTION = 'fixture';

const FixtureModel = mongoose.model(COLLECTION, FixtureSchema);

export class Repository{
    constructor() {
        this.mongoSetup();
    }

    private mongoSetup(): void {
        let mongoScheme = process.env.MONGO_SCHEME;
        let mongoUser = process.env.MONGO_USERNAME;
        let mongoPassword = process.env.MONGO_PASSWORD;
        let mongoHost = process.env.MONGO_HOST;
        let mongoPort = process.env.MONGO_PORT;

        let mongoUrl: string = mongoScheme + '://' + mongoUser + ':' + mongoPassword + '@' + mongoHost;

        if (mongoPort != '') {
            mongoUrl += ':' + mongoPort;
        }

        mongoUrl += '/' + DATABASE + '?retryWrites=true';

        mongoose.Promise = global.Promise;
        mongoose.connect(mongoUrl, {useNewUrlParser: true});
    }

    public getAllFixtures(): Fixture[] {
        return [];
    }

    public getFixturesById(id: string): Fixture {
        return new Fixture();
    }
}