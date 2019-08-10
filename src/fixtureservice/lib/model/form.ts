export class Forms {
    teams: Map<string, TeamForm>
}

export class TeamForm {
    forms: TeamMatchResult[]
}

export class TeamMatchResult {
    result: Result;
    score: string;
    opponent: string;
    location: Location;
}

export enum Result {
    WIN = 0,
    DRAW = 1,
    LOSS = 2,
}

export enum Location {
    HOME = 0,
    AWAY = 1,
}