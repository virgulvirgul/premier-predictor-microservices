import "jasmine";
import {beforeAll, it, expect, describe} from "jasmine-ts"
import {Repository} from "../../lib/repository/repository";

let repository: Repository;

beforeAll(() => {
    process.env.MONGO_SCHEME = 'mongodb';
    process.env.MONGO_HOST = 'localhost';
    process.env.MONGO_PORT = '27017';

    repository = new Repository();
});

describe("something", () => {
    it("should work with promises", () => {
        let fixture = repository.getFixturesById("1");
        expect(fixture).toEqual(1);
    });
});