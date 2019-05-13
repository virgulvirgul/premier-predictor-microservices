import {Fixture} from "../model/fixture";
import {asSequence} from "sequency";

export class FixtureFormatter {
    public groupIntoTeams(fixtures: Fixture[]): Promise<Map<string, Fixture[]>> {
        return new Promise<Map<string, Fixture[]>>((resolve, reject) => {
            const homeTeamMatches: Map<string, Fixture[]> = asSequence(fixtures)
                .groupBy(f => f.hTeam);
            const awayTeamMatches: Map<string, Fixture[]> = asSequence(fixtures)
                .groupBy(f => f.aTeam);

            const allMatches = homeTeamMatches;

            awayTeamMatches.forEach((v: Fixture[], k: string) => {
                allMatches.set(k, this.mergeArrays(k, allMatches, awayTeamMatches));

                const sortedFixtures: Fixture[] = asSequence<Fixture>(allMatches.get(k))
                    .sortedBy((f: Fixture) => {
                        return f.dateTime
                    })
                    .toArray();

                allMatches.set(k, sortedFixtures);
            });

            resolve(allMatches);
        });
    }

    private mergeArrays(k: string, allMatches: Map<string, Fixture[]>, awayTeamMatches: Map<string, Fixture[]>): Fixture[] {
        if (!allMatches.has(k)) {
            return awayTeamMatches.get(k);
        }

        return asSequence<Fixture>(allMatches.get(k))
            .merge(awayTeamMatches.get(k), (f: Fixture) => {
                return f.id;
            })
            .toArray();
    }
}