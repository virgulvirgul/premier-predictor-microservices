import {Fixture} from "../model/fixture";
import {Result, TeamForm, TeamMatchResult, Location} from "../model/form";

const {HOME, AWAY} = Location;
const {WIN, DRAW, LOSS} = Result;

export class FormFormatter {
    public formatLastFiveGames(teamFixtures: Map<string, Fixture[]>): Promise<Map<string, TeamForm>> {
        return new Promise<Map<string, TeamForm>>((resolve, reject) => {
            const forms: Map<string, TeamForm> = new Map<string, TeamForm>();

            teamFixtures.forEach((v: Fixture[], k: string) => {
                forms.set(k, this.createTeamForm(k, v.slice(-5)));
            });

            resolve(forms);
        });
    }

    private createTeamForm(teamName: String, fixtures: Fixture[]): TeamForm {
        return {
            forms: fixtures.map(f => this.convertMatchToResultSummary(teamName, f))
        }
    }

    private convertMatchToResultSummary(teamName: String, fixture: Fixture): TeamMatchResult {
        return {
            result: this.getResult(teamName, fixture),
            score: this.getScore(fixture),
            opponent: this.getOpponent(teamName, fixture),
            location: this.getLocation(teamName, fixture)
        }
    }

    private getResult(teamName: String, fixture: Fixture): Result {
        if (fixture.hGoals == fixture.aGoals) {
            return DRAW
        }

        if (this.getLocation(teamName, fixture) == HOME) {
            if (fixture.hGoals > fixture.aGoals) {
                return WIN
            }
        } else {
            if (fixture.aGoals > fixture.hGoals) {
                return WIN
            }
        }

        return LOSS
    }

    private getScore = (fixture: Fixture) => `${fixture.hGoals}-${fixture.aGoals}`;

    private getOpponent(teamName: String, fixture: Fixture): string {
        if (fixture.hTeam === teamName) {
            return fixture.aTeam;
        }

        return fixture.hTeam;
    }

    private getLocation(teamName: String, fixture: Fixture): Location {
        if (fixture.hTeam === teamName) {
            return HOME;
        }

        return AWAY;
    }
}